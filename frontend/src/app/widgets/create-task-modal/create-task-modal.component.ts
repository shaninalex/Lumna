import {
    Component,
    inject,
    Input,
    signal,
    WritableSignal,
    ChangeDetectionStrategy,
} from '@angular/core';
import { FormsModule } from '@angular/forms';
import { Dialog, DIALOG_DATA, DialogRef } from '@angular/cdk/dialog';
import { Store } from '@ngrx/store';
import { actionTask, TaskCreateModel } from '@entities/task';
import { form, FormField } from '@angular/forms/signals';
import { RouterLink } from '@angular/router';

@Component({
    selector: 'lu-create-task-modal',
    imports: [],
    changeDetection: ChangeDetectionStrategy.Eager,
    template: `
        <button class="button" type="button" (click)="openDialog()">
            <i class="fa-solid fa-plus"></i>
        </button>
    `,
})
export class CreateTaskModalComponent {
    @Input() columnId: number;
    @Input() boardId: number;
    @Input() taskCount: number;
    @Input() projectId: number;
    @Input() nextPosition: number;

    dialog = inject(Dialog);
    private store = inject(Store);

    openDialog(): void {
        const dialogRef = this.dialog.open<TaskCreateModel>(CreateTaskModalDialog, {
            width: '650px',
            data: {
                title: '',
                body: '',
                projectId: this.projectId,
                position: this.nextPosition,
                columnId: this.columnId,
                boardId: this.boardId,
            },
        });
        dialogRef.closed.subscribe((data) => {
            if (!data) {
                return;
            }
            this.store.dispatch(actionTask.create({ data }));
        });
    }
}

@Component({
    selector: 'lu-column-delete-prompt-dialog',
    templateUrl: 'column-delete-prompt.dialog.html',
    changeDetection: ChangeDetectionStrategy.Eager,
    imports: [FormsModule, RouterLink, FormField],
})
export class CreateTaskModalDialog {
    dialogRef = inject<DialogRef<TaskCreateModel>>(DialogRef<TaskCreateModel>);
    data = inject(DIALOG_DATA);

    taskEditFormModel: WritableSignal<TaskCreateModel> = signal<TaskCreateModel>({
        title: this.data.title,
        body: this.data.body,
        projectId: this.data.projectId,
        position: this.data.position,
        columnId: this.data.columnId,
        boardId: this.data.boardId,
        dueTo: this.data.dueTo,
    });
    taskEditForm = form(this.taskEditFormModel);

    submit(event: Event) {
        event.preventDefault();
        this.dialogRef.close(this.taskEditFormModel() as TaskCreateModel);
    }
}
