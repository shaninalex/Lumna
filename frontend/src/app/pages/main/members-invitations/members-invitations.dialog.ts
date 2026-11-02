import { FormsModule } from '@angular/forms';
import { Component, inject } from '@angular/core';
import { DialogRef, DIALOG_DATA } from '@angular/cdk/dialog';

export interface InviteData {
    email: string;
    role: string;
}

@Component({
    selector: 'lu-members-invitations-dialog',
    templateUrl: 'members-invitations.dialog.html',
    imports: [FormsModule],
})
export class MembersInvitationsDialogComponent {
    dialogRef = inject<DialogRef<InviteData>>(DialogRef<InviteData>);
    data = inject(DIALOG_DATA);

    close(): void {
        this.dialogRef.close()
    }

    submit(): void {
        this.dialogRef.close({
            email: this.data.email,
            role: this.data.role,
        })
    }
}
