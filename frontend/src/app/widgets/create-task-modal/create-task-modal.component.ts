import { Component, inject, Input, signal, WritableSignal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { Dialog, DIALOG_DATA, DialogRef } from '@angular/cdk/dialog';
import { Store } from '@ngrx/store';
import { actionTask, TaskCreateModel } from '@entities/task';
import { form, FormField } from '@angular/forms/signals';
import { RouterLink } from '@angular/router';

@Component({
    selector: 'lu-create-task-modal',
    imports: [],
    template: `
        <button class="button" type="button" (click)="openDialog()">
            <i class="fa-solid fa-plus"></i>
        </button>
    `,
})
export class CreateTaskModalComponent {
    @Input() column_id: number;
    @Input() board_id: number;
    @Input() task_count: number;
    @Input() project_id: number;

    dialog = inject(Dialog);
    private store = inject(Store)


    openDialog(): void {
        const dialogRef = this.dialog.open<TaskCreateModel>(CreateTaskModalDialog, {
            width: '650px',
            data: {
                title: '',
                body: '',
                project_id: this.project_id,
                position: 0,
                column_id: this.column_id,
                board_id: this.board_id,
            },
        });
        dialogRef.closed.subscribe(data => {
            if (!data) { return; }
            this.store.dispatch(actionTask.create({ data }));
        })
    }
}

@Component({
    selector: 'lu-column-delete-prompt-dialog',
    templateUrl: 'column-delete-prompt.dialog.html',
    imports: [FormsModule, RouterLink, FormField],
})
export class CreateTaskModalDialog {
    dialogRef = inject<DialogRef<TaskCreateModel>>(DialogRef<TaskCreateModel>);
    data = inject(DIALOG_DATA);

    taskEditFormModel: WritableSignal<TaskCreateModel> = signal<TaskCreateModel>({
        title: this.data.title,
        body: this.data.body,
        project_id: this.data.project_id,
        position: this.data.position,
        column_id: this.data.column_id,
        board_id: this.data.board_id,
        due_to: this.data.due_to,
    })
    taskEditForm = form(this.taskEditFormModel)

    submit(event: Event) {
        event.preventDefault();
        this.dialogRef.close(this.taskEditFormModel() as TaskCreateModel);
    }
}
