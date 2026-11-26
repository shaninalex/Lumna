import { Component, inject } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialog, MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatButton } from '@angular/material/button';
import { MatIcon } from '@angular/material/icon';
import { Store } from "@ngrx/store";
import { actionUser } from "@entities/user";

@Component({
    imports: [MatButton, MatIcon],
    selector: 'lu-account-remove-feature',
    template: `
        <div class="lu-banner lu-banner--danger" role="alert">
            <mat-icon>error</mat-icon>
            <div class="flex-grow-1">
                <h4 class="mt-0">Danger!</h4>
                <p>Warning text here...</p>
                <button matButton="filled" class="lu-danger" (click)="openConfirmationWindow()">Remove</button>
            </div>
        </div>
    `,
})
export class AccountRemoveFeature {
    dialog = inject(MatDialog);
    private store = inject(Store);

    openConfirmationWindow() {
        const dialogRef = this.dialog.open<AccountRemoveFeatureDialog, object, dialogData>(AccountRemoveFeatureDialog, {
            width: '400px',
            data: {},
        });

        dialogRef.afterClosed().subscribe((data) => {
            if (!data || !data.confirm) {
                return;
            }
            this.store.dispatch(actionUser.delete());
        });
    }
}

interface dialogData {
    confirm: boolean;
}

@Component({
    selector: 'lu-account-remove-feature-dialog',
    imports: [MatDialogModule, MatButton],
    template: `
        <h2 mat-dialog-title>Warning</h2>
        <mat-dialog-content>
            <p>Are you sure want to delete account?</p>
        </mat-dialog-content>
        <mat-dialog-actions align="end">
            <button type="button" matButton (click)="dialogRef.close({ confirm: false })">Cancel</button>
            <button type="button" matButton="filled" class="lu-danger" (click)="dialogRef.close({ confirm: true })">Yes</button>
        </mat-dialog-actions>
    `,
})
export class AccountRemoveFeatureDialog {
    dialogRef = inject<MatDialogRef<AccountRemoveFeatureDialog, dialogData>>(MatDialogRef);
    data = inject(MAT_DIALOG_DATA);
}
