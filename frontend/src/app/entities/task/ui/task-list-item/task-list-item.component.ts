import { Component, inject, Input, OnInit } from '@angular/core';
import type { TaskModel } from '@entities/task/model';
import { AppRoutes } from '@core';
import { standardTimeFormat } from '@shared/utils'
import { Store } from '@ngrx/store';
import {ColumnModel, selectColumns } from '@entities/column';
import { Observable } from 'rxjs';
import { RouterLink } from '@angular/router';
import { AsyncPipe, DatePipe } from '@angular/common';
import { BoardModel, selectBoard } from '@entities/board';

@Component({
    selector: 'lu-task-list-item',
    imports: [RouterLink, DatePipe, AsyncPipe],
    template: `
        <a [routerLink]="[...appRoutes.task(task.id)]" class="d-flex justify-content-between align-items-start text-decoration-none">
            <div class="flex-grow-1">
                <div class="d-flex align-items-center gap-2 mb-1">
                    @if (columns$ | async; as columns) {
                        @for (column of columns; track column.id) {
                            <span class="badge text-bg-primary">{{ column.title }}</span>
                        }
                    }
                </div>
                <h6 class="mb-1 text-body">
                    {{ task.title }}
                </h6>
                <div class="small text-muted">
                    #Feat-123 • Created by Alex • {{ task.created_at | date: standardTime }}
                </div>
            </div>
            <i class="fa-solid fa-chevron-right text-muted mt-1"></i>
        </a>
    `,
    host: {
        class: "list-group-item list-group-item-action",
    }
})
export class TaskListItemComponent implements OnInit {
    private store = inject(Store);
    @Input() task: TaskModel;
    readonly appRoutes = inject(AppRoutes)

    standardTime = standardTimeFormat;
    columns$: Observable<ColumnModel[]>;

    ngOnInit() {
        this.columns$ = this.store.select(selectColumns.byIds(this.task.boards.map(t => t.column_id)))
    }
}
