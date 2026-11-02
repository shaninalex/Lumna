import { inject, Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { catchError, exhaustMap, of, switchMap, tap } from 'rxjs';
import { MemberApi } from '../api/member.api';
import { actionMember } from './member.actions';
import type { HttpErrorResponse } from '@angular/common/http';
import { fromErrorResponse } from '@shared/models';

@Injectable()
export class MemberEffects {
    private actions$ = inject(Actions);
    private listApi = inject(MemberApi);

    list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionMember.getList),
            tap(action => console.log(action)),
            exhaustMap((action) =>
                this.listApi
                    .list(action.workspaceId)
                    .pipe(switchMap((members) => of(actionMember.setList({ members })))),
            ),
        ),
    );

    remove$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionMember.removeMember),
            tap(action => console.log(action)),
            exhaustMap((action) =>
                this.listApi.removeMember(action.workspaceId, action.identityId).pipe(
                    switchMap(() => of(actionMember.removeMemberSuccess({ workspaceId: action.workspaceId, identityId: action.identityId }))),
                    catchError((err: HttpErrorResponse) =>
                        of(actionMember.removeMemberFailed({errors: fromErrorResponse(err)}))
                    )
                ),
            ),
        ),
    );
}
