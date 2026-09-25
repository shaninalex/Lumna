import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import { catchError, EMPTY, exhaustMap, map } from 'rxjs';
import { SessionApi } from './session.api';
import { Router } from '@angular/router';
import { actionSession } from './session.actions';
import { actionUser } from '@entities/user';

@Injectable()
export class SessionEffects {
    private actions$ = inject(Actions);
    private sessionApi = inject(SessionApi);
    private router = inject(Router);

    authenticate_start$ = createEffect(() => {
        return this.actions$.pipe(
            ofType(actionSession.startAuthenticate),
            exhaustMap((action) =>
                this.sessionApi.login(action.email, action.password).pipe(
                    map((user) => actionSession.authenticatedSuccessfull({ user: user })),
                    catchError(() => EMPTY), // TODO: handle auth errors
                ),
            ),
        );
    });

    logging_out_init$ = createEffect(() => {
        return this.actions$.pipe(
            ofType(actionSession.loggingOut),
            exhaustMap(() =>
                this.sessionApi.logout().pipe(
                    map(() => actionSession.loggedOut()),
                    catchError(() => EMPTY),
                ),
            ),
        );
    });

    authenticated_successfull$ = createEffect(() => {
        return this.actions$.pipe(
            ofType(actionSession.authenticatedSuccessfull),
            map((action) => actionUser.set({ user: action.user })),
        );
    });

    authenticated_successfull_redirect$ = createEffect(
        () => {
            return this.actions$.pipe(
                ofType(actionSession.authenticatedSuccessfull),
                map(() => this.router.navigateByUrl('/app')),
            );
        },
        { dispatch: false },
    );

    logging_out_completed$ = createEffect(() => {
        return this.actions$.pipe(
            ofType(actionSession.loggedOut),
            map(() => actionUser.clear()),
        );
    });

    logging_out_completed_redirect$ = createEffect(
        () => {
            return this.actions$.pipe(
                ofType(actionSession.loggedOut),
                map(() => this.router.navigateByUrl('/auth/login')),
            );
        },
        { dispatch: false },
    );
}
