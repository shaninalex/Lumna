import { Component, inject, input, OnDestroy, OnInit, signal, WritableSignal } from '@angular/core';
import { Store } from '@ngrx/store';
import { TimeAgoPipe } from '@shared/utils';
import { AppRoutes } from '@core';
import { filter, Observable, tap } from 'rxjs';
import { form, FormField, required } from '@angular/forms/signals';
import { AsyncPipe, DatePipe } from '@angular/common';
import { RouterLink } from '@angular/router';
import { actionTask, selectTasks, TaskEditModel, TaskModel } from '@entities/task/model';
import { NgxEditorComponent, NgxEditorMenuComponent, Editor, Toolbar, toDoc, toHTML } from 'ngx-editor';
import { FormsModule } from '@angular/forms';
import { defaultToolbar } from '@shared/ui';

@Component({
    selector: 'lu-task-detail-view',
    imports: [TimeAgoPipe, DatePipe, AsyncPipe, RouterLink, FormField, NgxEditorComponent, NgxEditorMenuComponent, FormsModule],
    templateUrl: './task-detail-view.html',
})
export class TaskDetailView implements OnInit, OnDestroy {
    readonly appRoutes = inject(AppRoutes);
    taskId = input.required<number>();
    task$: Observable<TaskModel>;
    taskEditFormModel: WritableSignal<TaskEditModel> = signal<TaskEditModel>({taskId: 0, title: '', body: ''})
    taskEditForm = form(this.taskEditFormModel, (schemaPath) => {
        required(schemaPath.taskId, {message: "Task ID is required"});
        required(schemaPath.title, {message: "Title is required"});
    })
    editor: Editor;
    toolbar: Toolbar = defaultToolbar;
    html: Record<string, unknown> | string = '';
    private store = inject(Store);

    ngOnInit(): void {
        this.editor = new Editor();
        this.task$ = this.store.select(selectTasks.byId(this.taskId())).pipe(
            filter(task => !!task),
            tap(task => {
                this.taskEditForm.taskId().value.set(task.id);
                this.taskEditForm.title().value.set(task.title);
                if (task.body) {
                    this.taskEditForm.body().value.set(task.body);
                }
                this.html = toDoc(task.body || '', this.editor.schema);
            })
        );
    }

    ngOnDestroy(): void {
        this.editor.destroy();
    }

    onBodyChange(content: Record<string, unknown> | string): void {
        const renderedHtml = typeof content === 'string' ? content : toHTML(content, this.editor.schema);
        this.taskEditForm.body().value.set(renderedHtml);
    }

    submit(event: Event): void {
        event.preventDefault();
        const _data = this.taskEditFormModel();
        this.store.dispatch(actionTask.updateTask({
            data: {
                taskId: _data.taskId,
                title: _data.title.trim(),
                body: _data.body.trim(),
            }
        }))
    }
}
