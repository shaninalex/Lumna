import { inject, Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { exhaustMap, of, switchMap, tap } from 'rxjs';
import { MemberApi } from '../api/member.api';
import { actionMember } from './member.actions';

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
}
