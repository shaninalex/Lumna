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

    logging_out_completed$ = createEffect(() => {
        return this.actions$.pipe(
            ofType(actionSession.loggedOut),
            map(() => actionUser.clear()),
        );
    });
    logging_out_completed_redirect$ = createEffect(() => {
            return this.actions$.pipe(
                ofType(actionSession.loggedOut),
                map(() => window.location.href = '/auth/login'),
            );
        },
        {dispatch: false},
    );
    private sessionApi = inject(SessionApi);
    authenticate_start$ = createEffect(() => {
        return this.actions$.pipe(
            ofType(actionSession.startAuthenticate),
            exhaustMap((action) =>
                this.sessionApi.login(action.email, action.password).pipe(
                    map(() => actionSession.authenticatedSuccessful()),
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

    private router = inject(Router);
    authenticated_successfull_redirect$ = createEffect(
        () => {
            return this.actions$.pipe(
                ofType(actionSession.authenticatedSuccessful),
                map(() => this.router.navigateByUrl('/app')),
            );
        },
        {dispatch: false},
    );
}
