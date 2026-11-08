import { Component, inject, input } from '@angular/core';
import { NgIcon, provideIcons } from '@ng-icons/core';
import { reiconUserX } from '@ng-icons/reicon';
import { Dialog, DIALOG_DATA, DialogRef } from '@angular/cdk/dialog';
import { Store } from '@ngrx/store';
import { actionMember } from '@entities/member/model';

@Component({
    selector: 'lu-remove-member-button',
    imports: [NgIcon],
    providers: [provideIcons({ reiconUserX })],
    template: `
        <button class="btn btn-danger btn-sm" (click)="openConfirmationDialog()">
            <ng-icon name="reiconUserX" />
        </button>
    `,
})
export class RemoveMemberButtonFeature {
    private store = inject(Store);
    dialog = inject(Dialog);

    identityId = input.required<number>()
    workspaceId = input.required<number>()

    openConfirmationDialog(): void {
        const dialogRef = this.dialog.open<dialogData>(RemoveMemberConfirmDialog, {
            width: '250px',
            data: {},
        });

        dialogRef.closed.subscribe((data) => {
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
    template: `
        <div class="modal d-block">
            <div class="modal-dialog">
                <div class="modal-content">
                    <div class="modal-header">
                        <h5 class="modal-title">Warning!</h5>
                        <button type="button" class="btn-close" (click)="dialogRef.close({ confirm: false })"></button>
                    </div>
                    <div class="modal-body">
                        <p>Are you sure want to remove member from workspace?</p>
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
export class RemoveMemberConfirmDialog {
    dialogRef = inject<DialogRef<dialogData>>(DialogRef<dialogData>);
    data = inject(DIALOG_DATA);
}
