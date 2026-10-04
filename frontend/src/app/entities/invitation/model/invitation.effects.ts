import { inject, Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { exhaustMap, of } from 'rxjs';
import { switchMap } from 'rxjs/operators';
import { actionInvitation } from "./invitation.actions";
import { InvitationApi } from "../api";


@Injectable()
export class InvitationEffects {
    private actions$ = inject(Actions);
    private api = inject(InvitationApi);

    list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionInvitation.getList),
            exhaustMap((action) => this.api.list(action.workspaceId).pipe(
                switchMap((invitations) => of(actionInvitation.setList({ invitations })))
            )),
        )
    );

    create$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionInvitation.create),
            exhaustMap((action) => this.api.create(action.workspaceId, action.data).pipe(
                switchMap((invitation) => of(actionInvitation.set({ invitation })))
            )),
        )
    );
}
