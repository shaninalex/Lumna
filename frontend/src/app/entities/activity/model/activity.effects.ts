import { inject, Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { exhaustMap, of, switchMap } from 'rxjs';
import { ActivityApi } from '../api/activity.api';
import { actionActivity } from './activity.actions';

@Injectable()
export class ActivityEffects {
    private actions$ = inject(Actions);
    private listApi = inject(ActivityApi);

    list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionActivity.getList),
            exhaustMap((action) =>
                this.listApi
                    .list(action.entityId, action.entityType)
                    .pipe(switchMap((activities) => of(actionActivity.setList({ activities })))),
            ),
        ),
    );
}
