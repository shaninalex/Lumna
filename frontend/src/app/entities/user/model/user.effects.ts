import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import { catchError, of, tap } from 'rxjs';
import { actionUser } from './user.actions';
import { UserApi } from '../api/user.service';
import { switchMap } from 'rxjs/operators';
import type { HttpErrorResponse } from '@angular/common/http';
import { fromErrorResponse } from '@shared/models';
import { Router } from "@angular/router";
import { actionSession } from "@core/store/session/session.actions";

@Injectable()
export class UserEffects {
    private actions$ = inject(Actions);
    private router = inject(Router);

    clearUser$ = createEffect(() =>
            this.actions$.pipe(
                ofType(actionUser.clear),
                tap(() => console.log('user store cleaned')),
            ),
        {dispatch: false},
    );
    private api = inject(UserApi);
    changePassword$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionUser.changePassword),
            switchMap((action) =>
                this.api.changePassword({password: action.password, password_confirm: action.passwordConfirm}).pipe(
                    switchMap(() => of(actionUser.changePasswordSuccess())),
                    catchError((err: HttpErrorResponse) =>
                        of(actionUser.changePasswordFailed({errors: fromErrorResponse(err)}))
                    )
                )
            )
        )
    );

    update$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionUser.update),
            switchMap((action) =>
                this.api.update({email: action.email, full_name: action.fullName}).pipe(
                    switchMap((user) => of(actionUser.updateSuccess({user}))),
                    catchError((err: HttpErrorResponse) =>
                        of(actionUser.updateFailed({errors: fromErrorResponse(err)}))
                    )
                )
            )
        )
    );

    delete$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionUser.delete),
            switchMap(() =>
                this.api.delete().pipe(
                    switchMap(() => of(actionSession.loggedOut())),
                    catchError((err: HttpErrorResponse) => {
                        console.log(err)
                        return of(actionUser.deleteFailed({errors: fromErrorResponse(err)}))
                    })
                )
            )
        )
    );
}
