import { Component, inject } from '@angular/core';
import { Store } from '@ngrx/store';
import { filter, Observable, tap } from 'rxjs';
import { selectWorkspaces } from '@entities/workspace';
import { switchMap } from 'rxjs/operators';
import { actionInvitation, Invitation, selectInvitation } from '@entities/invitation';
import { AsyncPipe, DatePipe } from '@angular/common';
import { AppRoutes } from '@core';
import { InviteData, MembersInvitationsDialogComponent } from './members-invitations.dialog';
import { Dialog } from '@angular/cdk/dialog';

@Component({
    selector: 'lu-members-invitations',
    imports: [AsyncPipe, DatePipe],
    template: `
        <button class="button button-pill" (click)="openInviteModal()">
            Invite
        </button>
        @if (invitations$ | async; as invitations) {
            <table class="table">
                <thead>
                <tr>
                    <th>Email</th>
                    <th>Role</th>
                    <th>Invited By</th>
                    <th>Expires At</th>
                    <th>Accepted At</th>
                    <th>Revoked At</th>
                    <th>Created At</th>
                </tr>
                </thead>
                <tbody>
                    @for (invitation of invitations; track invitation.id) {
                        <tr>
                            <td>{{ invitation.email }}</td>
                            <td>{{ invitation.role }}</td>
                            <td><img src="images/7.png" style="width:1.5rem; height: 1.5rem; border-radius: 1.5rem"></td>
                            <td>{{ invitation.expiresAt | date }}</td>
                            <td>
                                @if(invitation.acceptedAt){
                                    {{ invitation.acceptedAt | date }}
                                } @else {
                                    ---
                                }
                            </td>
                            <td>

                                @if(invitation.revokedAt){
                                    {{ invitation.revokedAt | date }}
                                } @else {
                                    ---
                                }
                            </td>
                            <td>{{ invitation.createdAt | date }}</td>
                        </tr>
                    }
                </tbody>
            </table>
        }
    `,
})
export class MembersInvitationsPage {
    readonly appRoutes = inject(AppRoutes);
    dialog = inject(Dialog);
    private store = inject(Store);
    private workspaceId: number;
    invitations$: Observable<Invitation[]> = this.store.select(selectWorkspaces.currentWorkspaceId).pipe(
        filter(workspaceId => workspaceId !== null),
        tap((workspaceId) => this.store.dispatch(actionInvitation.getList({workspaceId}))),
        tap((workspaceId) => this.workspaceId = workspaceId),
        switchMap(workspaceId => this.store.select(selectInvitation.byWorkspaceId(workspaceId)))
    )

    openInviteModal(): void {
        const dialogRef = this.dialog.open<InviteData>(MembersInvitationsDialogComponent, {
            data: {
                email: '',
                role: '',
            },
        });

        dialogRef.closed.subscribe(result => {
            if (!result) return
            this.store.dispatch(actionInvitation.create({
                workspaceId: this.workspaceId,
                data: {
                    email: result?.email,
                    role: result?.role,
                }
            }))
        });
    }
}
