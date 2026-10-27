import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import { exhaustMap, of, switchMap } from 'rxjs';
import { actionRefreshToken } from './refresh_token.actions';
import { RefreshTokenService } from '../api';

@Injectable()
export class RefreshTokenEffects {
    private actions$ = inject(Actions);
    private refreshTokenApi = inject(RefreshTokenService);

    list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionRefreshToken.getList),
            exhaustMap(() =>
                this.refreshTokenApi.List().pipe(
                    switchMap((list) => of(actionRefreshToken.setList({list})))
                ),
            ),
        ),
    );

    delete$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionRefreshToken.delete),
            exhaustMap((action) =>
                this.refreshTokenApi.Delete(action.tokenId).pipe(
                    switchMap(() => of(actionRefreshToken.removeOne({tokenId: action.tokenId})))
                ),
            ),
        ),
    );

    revoke$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionRefreshToken.revoke),
            exhaustMap((action) =>
                this.refreshTokenApi.Revoke(action.tokenId).pipe(
                    switchMap((token) => of(actionRefreshToken.updateOne({
                        tokenId: action.tokenId,
                        updates: token,
                    })))
                )
            ),
        ),
    );
}
