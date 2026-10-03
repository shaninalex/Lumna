import { Component, inject, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { actionMember, Member, selectMember } from '@entities/member/model';
import { filter, Observable, tap } from 'rxjs';
import { selectWorkspaces } from '@entities/workspace';
import { AsyncPipe, DatePipe } from '@angular/common';
import { switchMap } from 'rxjs/operators';

@Component({
    imports: [AsyncPipe, DatePipe],
    selector: 'lu-members-list',
    template: `
        @if (members$ | async; as members) {
            <table class="table">
                <thead>
                <tr>
                    <th></th>
                    <th>email</th>
                    <th>full name</th>
                    <th>role</th>
                    <th>date joined</th>
                </tr>
                </thead>
                <tbody>
                    @for (member of members; track member.id) {
                        <tr>
                            <td><img src="images/7.png" style="width:2rem; height: 2rem; border-radius: 2rem"></td>
                            <td>{{ member.email }}</td>
                            <td>{{ member.fullName }}</td>
                            <td>{{ member.role }}</td>
                            <td>{{ member.dateJoined | date }}</td>
                        </tr>
                    }
                </tbody>
            </table>
        }
    `,
})
export class MembersListPage {
    private store = inject(Store);
    members$: Observable<Member[]> = this.store.select(selectWorkspaces.currentWorkspaceId).pipe(
        filter(workspaceId => workspaceId !== null),
        tap((workspaceId) => this.store.dispatch(actionMember.getList({workspaceId}))),
        switchMap(workspaceId => this.store.select(selectMember.byWorkspaceId(workspaceId)))
    );
}
