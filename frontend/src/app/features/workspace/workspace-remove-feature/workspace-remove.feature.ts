import { Component, inject } from '@angular/core';
import { Store } from '@ngrx/store';
import { actionWorkspace, selectWorkspaces } from '@entities/workspace';
import { MAT_DIALOG_DATA, MatDialog, MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatButton } from '@angular/material/button';
import { MatIcon } from '@angular/material/icon';

@Component({
    imports: [MatButton, MatIcon],
    selector: 'lu-workspace-remove-feature',
    template: `
        <div class="lu-banner lu-banner--danger" role="alert">
            <mat-icon>error</mat-icon>
            <div class="flex-grow-1">
                <h4 class="mt-0">Remove Workspace</h4>
                <p>Warning! This is destructive operation that can't be undone.</p>
                <button matButton="filled" class="lu-danger" (click)="openConfirmationWindow()">Remove</button>
            </div>
        </div>
    `,
})
export class WorkspaceRemoveFeature {
    private store = inject(Store);
    dialog = inject(MatDialog);
    private workspaceID = this.store.selectSignal(selectWorkspaces.currentWorkspaceId)

    openConfirmationWindow(): void {
        const dialogRef = this.dialog.open<WorkspaceRemoveFeatureDialog, object, dialogData>(WorkspaceRemoveFeatureDialog, {
            width: '400px',
            data: {},
        });

        dialogRef.afterClosed().subscribe((data) => {
            if (!data || !data.confirm) { return }

            const id = this.workspaceID()
            if (id === null) { return }
            this.store.dispatch(actionWorkspace.remove({ id }))
        });
    }
}

interface dialogData {
    confirm: boolean;
}

@Component({
    selector: 'lu-workspace-remove-feature-dialog',
    imports: [MatDialogModule, MatButton],
    template: `
        <h2 mat-dialog-title>Warning!</h2>
        <mat-dialog-content>
            <p>Are you sure want to delete workspace?</p>
        </mat-dialog-content>
        <mat-dialog-actions align="end">
            <button type="button" matButton (click)="dialogRef.close({ confirm: false })">Cancel</button>
            <button type="button" matButton="filled" class="lu-danger" (click)="dialogRef.close({ confirm: true })">Yes</button>
        </mat-dialog-actions>
    `,
})
export class WorkspaceRemoveFeatureDialog {
    dialogRef = inject<MatDialogRef<WorkspaceRemoveFeatureDialog, dialogData>>(MatDialogRef);
    data = inject(MAT_DIALOG_DATA);
}
