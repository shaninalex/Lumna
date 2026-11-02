import { Component, inject } from '@angular/core';
import { Dialog, DIALOG_DATA, DialogRef } from "@angular/cdk/dialog";
import { Store } from "@ngrx/store";
import { actionUser } from "@entities/user";

@Component({
    imports: [],
    selector: 'lu-account-remove-feature',
    template: `
        <div class="alert alert-danger">
            <h4>Danger!</h4>
            <p>Warning text here...</p>
            <hr>
            <button class="btn btn-danger" (click)="openConfirmationWindow()">Remove</button>
        </div>
    `,
})
export class AccountRemoveFeature {
    dialog = inject(Dialog);
    private store = inject(Store);

    openConfirmationWindow() {
        const dialogRef = this.dialog.open<dialogData>(AccountRemoveFeatureDialog, {
            width: '250px',
            data: {},
        });

        dialogRef.closed.subscribe((data) => {
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
    template: `
        <div class="modal d-block">
            <div class="modal-dialog">
                <div class="modal-content">
                    <div class="modal-header">
                        <h5 class="modal-title">Warning</h5>
                        <button type="button" class="btn-close" (click)="dialogRef.close({ confirm: false })"></button>
                    </div>
                    <div class="modal-body">
                        <p>Are you sure want to delete account?</p>
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
export class AccountRemoveFeatureDialog {
    dialogRef = inject<DialogRef<dialogData>>(DialogRef<dialogData>);
    data = inject(DIALOG_DATA);
}
