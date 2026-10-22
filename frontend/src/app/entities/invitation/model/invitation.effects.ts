import { inject, Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { catchError, exhaustMap, of } from 'rxjs';
import { switchMap } from 'rxjs/operators';
import { actionInvitation } from "./invitation.actions";
import { InvitationApi } from "../api";
import type { HttpErrorResponse } from '@angular/common/http';
import { fromErrorResponse } from '@shared/models';


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

    delete$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionInvitation.delete),
            exhaustMap((action) => this.api.delete(action.invitationId).pipe(
                switchMap(() => of(actionInvitation.deleteSuccess({ invitationId: action.invitationId }))),
                catchError((err: HttpErrorResponse) => of(
                    actionInvitation.deleteFailed({ errors: fromErrorResponse(err) })
                ))
            ))
        )
    )

    revoke$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionInvitation.revoke),
            exhaustMap((action) => this.api.revoke(action.invitationId).pipe(
                switchMap((invitation) => of(actionInvitation.revokeSuccess({ invitation }))),
                catchError((err: HttpErrorResponse) => of(
                    actionInvitation.revokeFailed({ errors: fromErrorResponse(err) })
                ))
            ))
        )
    )
}
