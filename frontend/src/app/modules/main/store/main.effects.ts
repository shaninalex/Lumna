import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import { filter, map, tap } from 'rxjs';
import { ROUTER_NAVIGATED } from '@ngrx/router-store';
import { LocalStorageService } from '@shared/services';
import { actionBoard } from '@entities/board';
import { actionSession } from '@core/store';
import { actionWorkspace } from '@entities/workspace';
import { actionProject } from '@entities/project';

@Injectable()
export class MainEffects {
    private actions$ = inject(Actions);
    private localStorageService = inject(LocalStorageService);

    routeChanged$ = createEffect(
        () =>
            this.actions$.pipe(
                ofType(ROUTER_NAVIGATED),
                map((action) => action.payload.event.url),
                filter((url: string) => url.startsWith('/app')),
                tap((url) => this.localStorageService.set('last_url', url)),
            ),
        { dispatch: false },
    );

    authenticated_successfull$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionSession.authenticated),
            map(() => actionWorkspace.getList()),
        ),
    );

    setCurrentWorkspace$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionWorkspace.setCurrent),
            map((action) => action.id),
            filter((id) => id !== null),
            map((id) => actionProject.getList({ workspaceId: id })),
        ),
    );

    // loadProjectTasks$ = createEffect(() =>
    //     this.actions$.pipe(
    //         ofType(actionProject.setCurrent),
    //         map((action) => action.id),
    //         filter((id) => id !== null),
    //         map((id) => actionTask.getList({ query: { project_id: id } })),
    //     ),
    // );

    onSetCurrentProject$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionProject.setCurrent),
            map((action) => action.id),
            filter((id) => id !== null),
            map((id) => actionBoard.getList({ projectId: id })),
        ),
    );
}
