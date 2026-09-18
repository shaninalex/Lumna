import { Component, inject, input, OnInit, signal, WritableSignal } from '@angular/core';
import { actionTask, selectTasks, TaskEditModel, TaskModel } from '@entities/task';
import { Store } from '@ngrx/store';
import { TimeAgoPipe } from '@shared/utils';
import { AppRoutes } from '@core';
import { filter, Observable, tap } from 'rxjs';
import { form, FormField, required } from '@angular/forms/signals';
import { AsyncPipe, DatePipe } from '@angular/common';
import { RouterLink } from '@angular/router';

@Component({
    selector: 'lu-task-detail-view',
    imports: [TimeAgoPipe, DatePipe, AsyncPipe, RouterLink, FormField],
    templateUrl: './task-detail-view.view.html',
})
export class TaskDetailViewView implements OnInit {
    private store = inject(Store);
    readonly appRoutes = inject(AppRoutes);
    taskId = input.required<number>();
    task$: Observable<TaskModel>;
    taskEditFormModel: WritableSignal<TaskEditModel> = signal<TaskEditModel>({task_id: 0, title: '', body: ''})

    ngOnInit(): void {
        this.task$ = this.store.select(selectTasks.byId(this.taskId())).pipe(
            filter(task => !!task),
            tap(task => {
                this.taskEditForm.task_id().value.set(task.id);
                this.taskEditForm.title().value.set(task.title);
                this.taskEditForm.body().value.set(task.body);
            })
        )
    }

    taskEditForm = form(this.taskEditFormModel, (schemaPath) => {
        required(schemaPath.task_id, {message: "Email is required"});
        required(schemaPath.title, {message: "Email is required"});
    })

    submit(event: Event): void {
        event.preventDefault();
        const _data = this.taskEditFormModel();
        this.store.dispatch(actionTask.updateTask({
            data: {
                task_id: _data.task_id,
                title: _data.title.trim(),
                body: _data.body.trim(),
            }
        }))
    }
}
