import { FormsModule } from '@angular/forms';
import { Component, inject } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatButton } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select';

export interface InviteData {
    email: string;
    role: string;
}

@Component({
    selector: 'lu-members-invitations-dialog',
    templateUrl: 'members-invitations.dialog.html',
    imports: [FormsModule, MatDialogModule, MatButton, MatFormFieldModule, MatInputModule, MatSelectModule],
})
export class MembersInvitationsDialogComponent {
    dialogRef = inject<MatDialogRef<MembersInvitationsDialogComponent, InviteData>>(MatDialogRef);
    data = inject<InviteData>(MAT_DIALOG_DATA);

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
