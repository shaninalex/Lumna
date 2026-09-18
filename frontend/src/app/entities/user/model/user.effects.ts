import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import { tap } from 'rxjs';
import { actionUser } from './user.actions';

@Injectable()
export class UserEffects {
    private actions$ = inject(Actions);

    // set_user$ = createEffect(() =>
    //     this.actions$.pipe(
    //         ofType(actionUser.set),
    //         // dispatch event to reconnect websocket for example
    //         tap((action) => console.log(`user ${action.user.full_name} added to store`)),
    //     ),
    //     { dispatch: false },
    // );

    clear_user$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionUser.clear),
            // dispatch event to close websocket connection for example
            tap(() => console.log('user store cleaned')),
        ),
        { dispatch: false },
    );
}
