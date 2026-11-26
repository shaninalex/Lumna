import { Component, inject, Input } from '@angular/core';
import { MatButton } from '@angular/material/button';
import { MAT_DIALOG_DATA, MatDialog, MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatIcon } from '@angular/material/icon';
import { MatMenuItem } from '@angular/material/menu';
import { Store } from '@ngrx/store';
import { actionsColumns } from '@entities/column';

export interface dialogData {
    withTasks: boolean;
    cancel: boolean;
}

@Component({
    selector: 'lu-column-delete-prompt',
    imports: [MatMenuItem, MatIcon],
    template: `
        <button type="button" mat-menu-item class="delete-item" (click)="openDialog()">
            <mat-icon>delete</mat-icon>
            <span>Delete</span>
        </button>
    `,
    styles: `
        .delete-item {
            --mat-menu-item-label-text-color: var(--mat-sys-error);
            --mat-menu-item-icon-color: var(--mat-sys-error);
        }
    `,
})
export class ColumnDeletePromptComponent {
    @Input() stageId: number;
    dialog = inject(MatDialog);
    private store = inject(Store);

    openDialog(): void {
        const dialogRef = this.dialog.open<ColumnDeletePromptComponentDialog, object, dialogData>(
            ColumnDeletePromptComponentDialog,
            {
                width: '400px',
                data: {},
            },
        );

        dialogRef.afterClosed().subscribe((data) => {
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

    imports: [MatDialogModule, MatButton],
})
export class ColumnDeletePromptComponentDialog {
    dialogRef = inject<MatDialogRef<ColumnDeletePromptComponentDialog, dialogData>>(MatDialogRef);
    data = inject(MAT_DIALOG_DATA);
}
