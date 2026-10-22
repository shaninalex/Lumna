import { Component, inject } from '@angular/core';
import { Store } from '@ngrx/store';
import { filter, Observable, tap } from 'rxjs';
import { selectWorkspaces } from '@entities/workspace';
import { switchMap } from 'rxjs/operators';
import { actionInvitation, Invitation, selectInvitation } from '@entities/invitation';
import { AsyncPipe, DatePipe, NgClass } from '@angular/common';
import { AppRoutes } from '@core';
import { InviteData, MembersInvitationsDialogComponent } from './members-invitations.dialog';
import { Dialog } from '@angular/cdk/dialog';
import { isAfter } from 'date-fns';

@Component({
    selector: 'lu-members-invitations',
    imports: [AsyncPipe, DatePipe, NgClass],
    styleUrl: './members-invitations.page.scss',
    template: `
        <button class="btn" (click)="openInviteModal()">
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
                    <th></th>
                </tr>
                </thead>
                <tbody>
                    @for (invitation of invitations; track invitation.id) {
                        <tr [ngClass]="{
                            'disabled-invitation': isExpired(invitation),
                        }">
                            <td>{{ invitation.email }}</td>
                            <td>{{ invitation.role }}</td>
                            <td><img alt="" src="images/7.png"
                                     style="width:1.5rem; height: 1.5rem; border-radius: 1.5rem"/></td>
                            <td>{{ invitation.expiresAt | date }}</td>
                            <td>
                                @if (invitation.acceptedAt) {
                                    {{ invitation.acceptedAt | date }}
                                } @else {
                                    ---
                                }
                            </td>
                            <td>
                                @if (invitation.revokedAt) {
                                    {{ invitation.revokedAt | date }}
                                } @else {
                                    ---
                                }
                            </td>
                            <td>{{ invitation.createdAt | date }}</td>
                            <td>
                                <div class="btn-group" role="group" aria-label="Invitation actions">
                                    <button type="button" class="btn btn-sm btn-outline-secondary"
                                            (click)="invitationRevoke(invitation.id)">Revoke
                                    </button>
                                    <button type="button" class="btn btn-sm btn-outline-secondary"
                                            (click)="invitationDelete(invitation.id)">Delete
                                    </button>
                                </div>
                            </td>
                        </tr>
                    }
                </tbody>
            </table>
        }
    `,
})
export class MembersInvitationsPage {
    readonly appRoutes = inject(AppRoutes);
    private store = inject(Store);
    private workspaceId: number;

    dialog = inject(Dialog);
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
                role: 'member',
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

    isExpired(invitation: Invitation): boolean {
        return isAfter(new Date(), new Date(invitation.expiresAt))
    }

    invitationDelete(invitationId: number): void {
        this.store.dispatch(actionInvitation.delete({ invitationId }))
    }

    invitationRevoke(invitationId: number): void {
        this.store.dispatch(actionInvitation.revoke({ invitationId }))
    }
}
