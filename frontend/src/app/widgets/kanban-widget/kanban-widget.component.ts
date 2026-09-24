import { AsyncPipe } from '@angular/common';
import type { OnInit } from '@angular/core';
import { Component, DestroyRef, effect, inject, input } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { Actions, ofType } from '@ngrx/effects';
import { Store } from '@ngrx/store';
import { actionsColumns, ColumnMenuDropdownComponent } from '@entities/column';
import { filter, tap, type Observable } from 'rxjs';
import { TimeAgoPipe } from '@shared/utils';
import { selectBoard, type BoardModel } from '@entities/board';
import { TaskCardComponent, actionTask } from '@entities/task';

import type { CdkDragDrop } from '@angular/cdk/drag-drop';
import {
    CdkDrag,
    CdkDragHandle,
    CdkDropList,
    CdkDropListGroup,
} from '@angular/cdk/drag-drop';
import type { KanbanCard, KanbanColumn } from './model/kanban.models';
import { KanbanService } from './service';
import { AppRoutes } from '@core';
import { AssignmentDropdown } from '@features/task';
import { ColumnDeletePromptComponent, NewColumnFormComponent } from '@features/stage';
import { CreateTaskModalComponent } from '../create-task-modal';
import { selectProjects } from '@entities/project';

@Component({
    selector: 'lu-kanban-board-feature',
    imports: [
        CdkDropListGroup,
        CdkDropList,
        CdkDrag,
        CdkDragHandle,
        AsyncPipe,
        NewColumnFormComponent,
        TimeAgoPipe,
        TaskCardComponent,
        AssignmentDropdown,
        ColumnMenuDropdownComponent,
        ColumnDeletePromptComponent,
        CreateTaskModalComponent,
    ],
    templateUrl: './kanban-widget.component.html',
    styleUrl: './kanban-widget.component.css',
    providers: [KanbanService],
})
export class KanbanBoardWidget implements OnInit {
    readonly appRoutes = inject(AppRoutes);
    boardId = input.required<number>();
    board$: Observable<BoardModel>;
    private store = inject(Store);
    projectId$ = this.store.select(selectProjects.currentProjectId);
    private actions$ = inject(Actions);
    private destroyRef = inject(DestroyRef);
    private kanban = inject(KanbanService);
    kolumns$: Observable<KanbanColumn[]> = this.kanban.boardData();

    constructor() {
        effect(() => {
            this.kanban.setBoardId(this.boardId());
        });
    }

    ngOnInit() {
        const _q = {boardId: this.boardId()};
        this.store.dispatch(actionTask.getList({query: _q}));
        this.store.dispatch(actionsColumns.loadByBoardId(_q));
        this.board$ = this.store
            .select(selectBoard.byId(_q.boardId))
            .pipe(filter((board) => board !== null));

        this.actions$
            .pipe(
                ofType(actionsColumns.reorderFailed),
                takeUntilDestroyed(this.destroyRef),
                tap(() => this.store.dispatch(actionsColumns.loadByBoardId(_q))),
            )
            .subscribe();
    }

    dropColumn(event: CdkDragDrop<KanbanColumn[]>): void {
        this.kanban.dropColumn(event, this.boardId());
    }

    dropTask(event: CdkDragDrop<KanbanCard[]>, column: KanbanColumn): void {
        const isSameList = event.previousContainer === event.container;

        if (isSameList) {
            this.kanban.moveTask(event, column, this.boardId());
        } else {
            this.kanban.transferTask(event, column, this.boardId());
        }
    }

    public columnsAmount(): number {
        return this.kanban.getColumnsLength();
    }

    public cardsAmount(): number {
        return this.kanban.cardsCount;
    }
}
