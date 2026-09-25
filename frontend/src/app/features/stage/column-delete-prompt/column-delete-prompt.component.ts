import { Component, inject, Input } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { Dialog, DIALOG_DATA, DialogRef } from '@angular/cdk/dialog';
import { Store } from '@ngrx/store';
import { actionsColumns } from '@entities/column';

export interface dialogData {
    withTasks: boolean;
    cancel: boolean;
}

@Component({
    selector: 'lu-column-delete-prompt',
    imports: [],

    template: ` <button type="button" class="btn" (click)="openDialog()">Delete</button> `,
})
export class ColumnDeletePromptComponent {
    @Input() stageId: number;
    dialog = inject(Dialog);
    private store = inject(Store);

    openDialog(): void {
        const dialogRef = this.dialog.open<dialogData>(ColumnDeletePromptComponentDialog, {
            width: '250px',
            data: {},
        });

        dialogRef.closed.subscribe((data) => {
            if (!data || data.cancel) {
                return;
            }
            this.store.dispatch(
                actionsColumns.deleteStage({
                    data: {
                        id: this.stageId,
                        withTasks: data.withTasks,
                    },
                }),
            );
        });
    }
}

@Component({
    selector: 'lu-column-delete-prompt-dialog',
    templateUrl: 'column-delete-prompt.dialog.html',

    imports: [FormsModule],
})
export class ColumnDeletePromptComponentDialog {
    dialogRef = inject<DialogRef<dialogData>>(DialogRef<dialogData>);
    data = inject(DIALOG_DATA);
}
