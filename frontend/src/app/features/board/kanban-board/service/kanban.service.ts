import type { OnDestroy } from '@angular/core';
import { inject, Injectable, signal } from '@angular/core';
import { selectColumns } from '@entities/column';
import { selectTasks } from '@entities/task';
import { Store } from '@ngrx/store';
import type { Observable } from 'rxjs';
import { BehaviorSubject, combineLatest, filter, Subscription, switchMap, tap } from 'rxjs';
import type { KanbanCard, KanbanColumn } from '../model/kanban.models';
import { toObservable } from '@angular/core/rxjs-interop';
import { CdkDragDrop, moveItemInArray, transferArrayItem } from '@angular/cdk/drag-drop';
import { actionKanban } from '../model';

@Injectable()
export class KanbanService implements OnDestroy {
    private store = inject(Store);
    private boardId = signal<number | undefined>(undefined);
    private sub = new Subscription();
    private data: BehaviorSubject<KanbanColumn[]> = new BehaviorSubject<KanbanColumn[]>([]);
    private boardId$ = toObservable(this.boardId).pipe(
        filter((id): id is number => id !== undefined),
    );
    private columns$ = this.boardId$.pipe(switchMap((id) => this.store.select(selectColumns.byListId(id))));
    private tasks$ = this.boardId$.pipe(switchMap((id) => this.store.select(selectTasks.byBoardId(id))));

    constructor() {
        this.sub.add(
            combineLatest([this.columns$, this.tasks$])
                .pipe(
                    tap(([columns, tasks]) => {
                        const boardId = this.boardId();
                        const kolumns: KanbanColumn[] = [];
                        columns.forEach((col) => {
                            const kolumn: KanbanColumn = { ...col, tasks: [] };
                            tasks.forEach((t) => {
                                const boardIndex = t.boards.findIndex(
                                    (b) => b.column_id === col.id && b.board_id === boardId,
                                );
                                if (boardIndex >= 0) {
                                    kolumn.tasks.push({
                                        ...t,
                                        column: col.id,
                                        position: t.boards[boardIndex].position,
                                    });
                                }
                            });
                            kolumn.tasks.sort((a, b) => a.position - b.position);
                            kolumns.push(kolumn);
                        });
                        this.data.next(kolumns);
                    }),
                )
                .subscribe(),
        );
    }

    ngOnDestroy(): void {
        this.sub.unsubscribe();
    }

    public dropColumn(event: CdkDragDrop<KanbanColumn[]>, boardId: number): void {
        const data = [...this.data.getValue()];
        moveItemInArray(data, event.previousIndex, event.currentIndex);
        this.data.next(data);
        this.store.dispatch(
            actionKanban.dropColumn({
                event: {
                    id: event.item.data.id,
                    previous_index: event.previousIndex,
                    current_index: event.currentIndex,
                    board_id: boardId,
                    columns_order: data.map((c) => c.id),
                },
            }),
        );
    }

    public moveTask(event: CdkDragDrop<KanbanCard[]>, column: KanbanColumn, boardId: number): void {
        moveItemInArray(event.container.data, event.previousIndex, event.currentIndex);
        const items = event.container.data;
        const index = event.currentIndex;

        const prev = items[index - 1]?.position;
        const next = items[index + 1]?.position;

        let position: number;

        if (prev === undefined && next === undefined) {
            // Only item in the column
            position = 1;
        } else if (prev === undefined) {
            // Moved to the beginning
            position = next! / 2;
        } else if (next === undefined) {
            // Moved to the end
            position = prev + 1;
        } else {
            // Moved between two items
            position = (prev + next) / 2;
        }
        this.store.dispatch(
            actionKanban.moveTask({
                event: {
                    board_id: boardId,
                    task_id: event.item.data.id,
                    rank: position,
                },
            }),
        );
    }

    public transferTask(event: CdkDragDrop<KanbanCard[]>, column: KanbanColumn, boardId: number): void {
        const card: KanbanCard = event.item.data;
        const fromColumnId = card.column;
        transferArrayItem(
            event.previousContainer.data,
            event.container.data,
            event.previousIndex,
            event.currentIndex,
        );
        card.column = column.id;
        // this.store.dispatch(
        //     actionKanban.transferTask({
        //         event: {
        //             board_id: boardId,
        //             from: {
        //                 column_id: fromColumnId,
        //                 tasks: event.previousContainer.data.map((t) => t.id),
        //             },
        //             to: {
        //                 column_id: column.id,
        //                 tasks: event.container.data.map((t) => t.id),
        //             },
        //         },
        //     }),
        // );
    }

    public setBoardId(boardId: number) {
        this.boardId.set(boardId);
    }

    public boardData(): Observable<KanbanColumn[]> {
        return this.data.asObservable();
    }

    public getColumnsLength(): number {
        return this.data.getValue().length;
    }
}
