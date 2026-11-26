import { Component, inject, input } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialog, MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatButton, MatIconButton } from '@angular/material/button';
import { MatIcon } from '@angular/material/icon';
import { Store } from '@ngrx/store';
import { actionMember } from '@entities/member/model';

@Component({
    selector: 'lu-remove-member-button',
    imports: [MatIconButton, MatIcon],
    template: `
        <button matIconButton class="lu-danger" aria-label="Remove member" (click)="openConfirmationDialog()">
            <mat-icon>person_remove</mat-icon>
        </button>
    `,
})
export class RemoveMemberButtonFeature {
    private store = inject(Store);
    dialog = inject(MatDialog);

    identityId = input.required<number>()
    workspaceId = input.required<number>()

    openConfirmationDialog(): void {
        const dialogRef = this.dialog.open<RemoveMemberConfirmDialog, object, dialogData>(RemoveMemberConfirmDialog, {
            width: '400px',
            data: {},
        });

        dialogRef.afterClosed().subscribe((data) => {
            if (!data || !data.confirm) { return }
            this.store.dispatch(actionMember.removeMember({
                workspaceId: this.workspaceId(),
                identityId: this.identityId(),
            }))
        });
    }
}

interface dialogData {
    confirm: boolean;
}

@Component({
    selector: 'lu-remove-member-confirm-dialog',
    imports: [MatDialogModule, MatButton],
    template: `
        <h2 mat-dialog-title>Warning!</h2>
        <mat-dialog-content>
            <p>Are you sure want to remove member from workspace?</p>
        </mat-dialog-content>
        <mat-dialog-actions align="end">
            <button type="button" matButton (click)="dialogRef.close({ confirm: false })">Cancel</button>
            <button type="button" matButton="filled" class="lu-danger" (click)="dialogRef.close({ confirm: true })">Yes</button>
        </mat-dialog-actions>
    `,
})
export class RemoveMemberConfirmDialog {
    dialogRef = inject<MatDialogRef<RemoveMemberConfirmDialog, dialogData>>(MatDialogRef);
    data = inject(MAT_DIALOG_DATA);
}
