import { Component, inject } from '@angular/core';
import { Store } from '@ngrx/store';
import { filter, Observable, tap } from 'rxjs';
import { selectWorkspaces } from '@entities/workspace';
import { switchMap } from 'rxjs/operators';
import { actionInvitation, Invitation, selectInvitation } from '@entities/invitation';
import { AsyncPipe, DatePipe } from '@angular/common';
import { AppRoutes } from '@core';
import { InviteData, MembersInvitationsDialogComponent } from './members-invitations.dialog';
import { MatDialog } from '@angular/material/dialog';
import { MatTableModule } from '@angular/material/table';
import { MatButton } from '@angular/material/button';
import { MatIcon } from '@angular/material/icon';
import { MatTooltip } from '@angular/material/tooltip';
import { isAfter } from 'date-fns';

@Component({
    selector: 'lu-members-invitations',
    imports: [AsyncPipe, DatePipe, MatTableModule, MatButton, MatIcon, MatTooltip],
    styleUrl: './members-invitations.page.scss',
    template: `
        @if (invitations$ | async; as invitations) {
            <table mat-table [dataSource]="invitations" class="w-100 mb-4">
                <ng-container matColumnDef="status">
                    <th mat-header-cell *matHeaderCellDef></th>
                    <td mat-cell *matCellDef="let invitation">
                        @if (isAccepted(invitation)) {
                            <mat-icon class="text-success" matTooltip="Accepted">check</mat-icon>
                        } @else {
                            <mat-icon class="text-danger" matTooltip="Not accepted">close</mat-icon>
                        }
                    </td>
                </ng-container>
                <ng-container matColumnDef="email">
                    <th mat-header-cell *matHeaderCellDef>Email</th>
                    <td mat-cell *matCellDef="let invitation">{{ invitation.email }}</td>
                </ng-container>
                <ng-container matColumnDef="role">
                    <th mat-header-cell *matHeaderCellDef>Role</th>
                    <td mat-cell *matCellDef="let invitation">{{ invitation.role }}</td>
                </ng-container>
                <ng-container matColumnDef="invitedBy">
                    <th mat-header-cell *matHeaderCellDef>Invited By</th>
                    <td mat-cell *matCellDef="let invitation">
                        <img class="invited-by-avatar" alt="" src="images/7.png"/>
                    </td>
                </ng-container>
                <ng-container matColumnDef="expiresAt">
                    <th mat-header-cell *matHeaderCellDef>Expires At</th>
                    <td mat-cell *matCellDef="let invitation">{{ invitation.expiresAt | date }}</td>
                </ng-container>
                <ng-container matColumnDef="acceptedAt">
                    <th mat-header-cell *matHeaderCellDef>Accepted At</th>
                    <td mat-cell *matCellDef="let invitation">
                        @if (invitation.acceptedAt) {
                            {{ invitation.acceptedAt | date }}
                        } @else {
                            ---
                        }
                    </td>
                </ng-container>
                <ng-container matColumnDef="revokedAt">
                    <th mat-header-cell *matHeaderCellDef>Revoked At</th>
                    <td mat-cell *matCellDef="let invitation">
                        @if (invitation.revokedAt) {
                            {{ invitation.revokedAt | date }}
                        } @else {
                            ---
                        }
                    </td>
                </ng-container>
                <ng-container matColumnDef="createdAt">
                    <th mat-header-cell *matHeaderCellDef>Created At</th>
                    <td mat-cell *matCellDef="let invitation">{{ invitation.createdAt | date }}</td>
                </ng-container>
                <ng-container matColumnDef="actions">
                    <th mat-header-cell *matHeaderCellDef></th>
                    <td mat-cell *matCellDef="let invitation">
                        <div class="d-flex gap-2 justify-content-end" role="group" aria-label="Invitation actions">
                            <button type="button" matButton="tonal"
                                    [disabled]="isAccepted(invitation)"
                                    (click)="invitationRevoke(invitation.id)">Revoke
                            </button>
                            <button type="button" matButton="filled" class="lu-danger"
                                    (click)="invitationDelete(invitation.id)">Delete
                            </button>
                        </div>
                    </td>
                </ng-container>

                <tr mat-header-row *matHeaderRowDef="columns"></tr>
                <tr mat-row *matRowDef="let invitation; columns: columns"
                    [class.disabled-invitation]="isExpired(invitation)"></tr>
            </table>
        }
        <div class="d-flex justify-content-end">
            <button type="button" matButton="filled" (click)="openInviteModal()">
                <mat-icon>person_add</mat-icon>
                Invite
            </button>
        </div>
    `,
})
export class MembersInvitationsPage {
    readonly appRoutes = inject(AppRoutes);
    private store = inject(Store);
    private workspaceId: number;

    dialog = inject(MatDialog);
    readonly columns = ['status', 'email', 'role', 'invitedBy', 'expiresAt', 'acceptedAt', 'revokedAt', 'createdAt', 'actions'];
    invitations$: Observable<Invitation[]> = this.store.select(selectWorkspaces.currentWorkspaceId).pipe(
        filter(workspaceId => workspaceId !== null),
        tap((workspaceId) => this.store.dispatch(actionInvitation.getList({workspaceId}))),
        tap((workspaceId) => this.workspaceId = workspaceId),
        switchMap(workspaceId => this.store.select(selectInvitation.byWorkspaceId(workspaceId)))
    )

    openInviteModal(): void {
        const dialogRef = this.dialog.open<MembersInvitationsDialogComponent, InviteData, InviteData>(MembersInvitationsDialogComponent, {
            width: '400px',
            data: {
                email: '',
                role: 'member',
            },
        });

        dialogRef.afterClosed().subscribe(result => {
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

    isAccepted(invitation: Invitation): boolean {
        return invitation.acceptedAt !== undefined
    }

    invitationDelete(invitationId: number): void {
        this.store.dispatch(actionInvitation.delete({ invitationId }))
    }

    invitationRevoke(invitationId: number): void {
        this.store.dispatch(actionInvitation.revoke({ invitationId }))
    }
}
