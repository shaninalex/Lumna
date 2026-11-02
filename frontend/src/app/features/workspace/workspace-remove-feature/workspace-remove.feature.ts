import { Component, inject } from '@angular/core';
import { Store } from '@ngrx/store';
import { actionWorkspace, selectWorkspaces } from '@entities/workspace';
import { Dialog, DIALOG_DATA, DialogRef } from '@angular/cdk/dialog';

@Component({
    imports: [],
    selector: 'lu-workspace-remove-feature',
    template: `
        <div class="alert alert-danger">
            <h4>Remove Workspace</h4>
            <p>Warning! This is destructive operation that can't be undone.</p>
            <hr>
            <button class="btn btn-danger" (click)="openConfirmationWindow()">Remove</button>
        </div>
    `,
})
export class WorkspaceRemoveFeature {
    private store = inject(Store);
    dialog = inject(Dialog);
    private workspaceID = this.store.selectSignal(selectWorkspaces.currentWorkspaceId)

    openConfirmationWindow(): void {
        const dialogRef = this.dialog.open<dialogData>(WorkspaceRemoveFeatureDialog, {
            width: '250px',
            data: {},
        });

        dialogRef.closed.subscribe((data) => {
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
    template: `
        <div class="modal d-block">
            <div class="modal-dialog">
                <div class="modal-content">
                    <div class="modal-header">
                        <h5 class="modal-title">Warning!</h5>
                        <button type="button" class="btn-close" (click)="dialogRef.close({ confirm: false })"></button>
                    </div>
                    <div class="modal-body">
                        <p>Are you sure want to delete workspace?</p>
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-danger" (click)="dialogRef.close({ confirm: true })">Yes</button>
                        <button type="button" class="btn btn-secondary" (click)="dialogRef.close({ confirm: false })">Cancel</button>
                    </div>
                </div>
            </div>
        </div>
    `,
})
export class WorkspaceRemoveFeatureDialog {
    dialogRef = inject<DialogRef<dialogData>>(DialogRef<dialogData>);
    data = inject(DIALOG_DATA);
}
