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
                        const kolumns: KanbanColumn[] = [];
                        columns.forEach((col) => {
                            const kolumn: KanbanColumn = {...col, tasks: []};
                            kolumn.tasks = tasks.filter((t) => t.columnId === col.id);
                            kolumn.tasks.sort((a, b) => a.position - b.position);
                            kolumns.push(kolumn);
                        });
                        this.data.next(kolumns);
                    }),
                )
                .subscribe(),
        );
    }

    public get cardsCount(): number {
        let tasks = 0;
        this.data.getValue().forEach((value) => tasks += value.tasks.length)
        return tasks
    }

    ngOnDestroy(): void {
        this.sub.unsubscribe();
    }

    // Moving columns
    public dropColumn(event: CdkDragDrop<KanbanColumn[]>, boardId: number): void {
        const data = [...this.data.getValue()];
        moveItemInArray(data, event.previousIndex, event.currentIndex);
        this.data.next(data);
        const position = this.calculatePosition(data, event.currentIndex);
        this.store.dispatch(actionKanban.dropColumn({
            event: {
                board_id: boardId,
                column_id: event.item.data.id,
                position: position,
            }
        }));
    }

    public moveTask(event: CdkDragDrop<KanbanCard[]>, column: KanbanColumn, boardId: number): void {
        moveItemInArray(event.container.data, event.previousIndex, event.currentIndex);
        const items = event.container.data;
        const index = event.currentIndex;

        const prev = items[index - 1]?.position;
        const next = items[index + 1]?.position;

        let position: number;

        if (prev === undefined && next === undefined) {
            position = 1;
        } else if (prev === undefined) {
            position = next! / 2;
        } else if (next === undefined) {
            position = prev + 1;
        } else {
            position = (prev + next) / 2;
        }
        this.store.dispatch(actionKanban.moveTask({
            event: {
                board_id: boardId,
                task_id: event.item.data.id,
                position: position,
            },
        }));
    }

    // moving tasks between columns
    public transferTask(event: CdkDragDrop<KanbanCard[]>, column: KanbanColumn, boardId: number): void {
        const card: KanbanCard = event.item.data;
        transferArrayItem(event.previousContainer.data, event.container.data, event.previousIndex, event.currentIndex,);
        // card.column_id = column.id;
        const position = this.calculatePosition(event.container.data, event.currentIndex);
        this.store.dispatch(actionKanban.transferTask({
            event: {
                task_id: card.id,
                board_id: boardId,
                column_id: column.id,
                position: position,
            }
        }));
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

    private calculatePosition<T extends { position: number }>(items: T[], index: number): number {
        const prev = items[index - 1]?.position;
        const next = items[index + 1]?.position;

        if (prev === undefined && next === undefined) {
            return 1;
        }

        if (prev === undefined) {
            return next! / 2;
        }

        if (next === undefined) {
            return prev + 1;
        }

        return (prev + next) / 2;
    }
}
