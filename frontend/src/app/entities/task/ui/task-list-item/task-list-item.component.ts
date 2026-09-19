import { Component, inject, Input, OnInit } from '@angular/core';
import { TaskModel } from '../../model';
import { AppRoutes } from '@core';
import { standardTimeFormat } from '@shared/utils';
import { Store } from '@ngrx/store';
import { ColumnModel, selectColumns } from '@entities/column'; // bad (?)
import { filter, Observable } from 'rxjs';
import { RouterLink } from '@angular/router';
import { AsyncPipe, DatePipe } from '@angular/common';

@Component({
    selector: 'lu-task-list-item',
    imports: [RouterLink, DatePipe, AsyncPipe],
    template: `
        <a
            [routerLink]="appRoutes.task(task.board_id, task.id)"
            class="d-flex justify-content-between align-items-start text-decoration-none"
        >
            <div class="flex-grow-1">
                <div class="d-flex align-items-center gap-2 mb-1">
                    @if (column$ | async; as column) {
                        <span class="badge text-bg-primary">{{ column.title }}</span>
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
        class: 'list-group-item list-group-item-action',
    },
})
export class TaskListItemComponent implements OnInit {
    private store = inject(Store);
    @Input() task: TaskModel;
    readonly appRoutes = inject(AppRoutes);

    standardTime = standardTimeFormat;
    column$: Observable<ColumnModel>;

    ngOnInit() {
        this.column$ = this.store.select(selectColumns.byId(this.task.column_id)).pipe(
            filter(c => c !== undefined),
        );
    }
}
