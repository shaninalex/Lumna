import { Component, inject } from '@angular/core';
import { Store } from '@ngrx/store';
import { actionMember, Member, selectMember } from '@entities/member/model';
import { filter, Observable, tap } from 'rxjs';
import { selectWorkspaces } from '@entities/workspace';
import { AsyncPipe, DatePipe } from '@angular/common';
import { switchMap } from 'rxjs/operators';
import { MatTableModule } from '@angular/material/table';
import { RemoveMemberButtonFeature } from '@features/user';

@Component({
    imports: [AsyncPipe, DatePipe, MatTableModule, RemoveMemberButtonFeature],
    selector: 'lu-members-list',
    styles: `
        .member-avatar {
            width: 2rem;
            height: 2rem;
            border-radius: 50%;
            display: block;
        }
    `,
    template: `
        @if (members$ | async; as members) {
            <table mat-table [dataSource]="members" class="w-100">
                <ng-container matColumnDef="avatar">
                    <th mat-header-cell *matHeaderCellDef></th>
                    <td mat-cell *matCellDef="let member">
                        <img class="member-avatar" src="images/7.png" alt="">
                    </td>
                </ng-container>
                <ng-container matColumnDef="email">
                    <th mat-header-cell *matHeaderCellDef>Email</th>
                    <td mat-cell *matCellDef="let member">{{ member.email }}</td>
                </ng-container>
                <ng-container matColumnDef="fullName">
                    <th mat-header-cell *matHeaderCellDef>Full name</th>
                    <td mat-cell *matCellDef="let member">{{ member.fullName }}</td>
                </ng-container>
                <ng-container matColumnDef="role">
                    <th mat-header-cell *matHeaderCellDef>Role</th>
                    <td mat-cell *matCellDef="let member">{{ member.role }}</td>
                </ng-container>
                <ng-container matColumnDef="dateJoined">
                    <th mat-header-cell *matHeaderCellDef>Date joined</th>
                    <td mat-cell *matCellDef="let member">{{ member.dateJoined | date }}</td>
                </ng-container>
                <ng-container matColumnDef="actions">
                    <th mat-header-cell *matHeaderCellDef></th>
                    <td mat-cell *matCellDef="let member" class="text-end">
                        <lu-remove-member-button
                            [identityId]="member.id"
                            [workspaceId]="member.workspaceId"
                        />
                    </td>
                </ng-container>

                <tr mat-header-row *matHeaderRowDef="columns"></tr>
                <tr mat-row *matRowDef="let row; columns: columns"></tr>
            </table>
        }
    `,
})
export class MembersListPage {
    private store = inject(Store);
    readonly columns = ['avatar', 'email', 'fullName', 'role', 'dateJoined', 'actions'];
    members$: Observable<Member[]> = this.store.select(selectWorkspaces.currentWorkspaceId).pipe(
        filter(workspaceId => workspaceId !== null),
        tap((workspaceId) => this.store.dispatch(actionMember.getList({workspaceId}))),
        switchMap(workspaceId => this.store.select(selectMember.byWorkspaceId(workspaceId)))
    );
}
