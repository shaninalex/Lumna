import { inject } from '@angular/core';
import type { CanActivateFn} from '@angular/router';
import { Router } from '@angular/router';
import { Store } from '@ngrx/store';
import { map, switchMap, take } from 'rxjs/operators';
import { of } from 'rxjs';
import type {
    WorkspaceModel} from '@entities/workspace';
import {
    actionWorkspace,
    selectWorkspaces,
    WorkspaceApi
} from '@entities/workspace';
import { parseRouteId } from '@shared/utils';

function getOrFetchWorkspaces() {
    const store = inject(Store);
    const workspaceApi = inject(WorkspaceApi);

    return store.select(selectWorkspaces.all).pipe(
        take(1),
        switchMap((workspaces) => {
            if (workspaces && workspaces.length > 0) {
                return of(workspaces);
            }
            return workspaceApi.List().pipe(
                map((list: WorkspaceModel[]) => {
                    store.dispatch(actionWorkspace.setList({ list }));
                    return list;
                }),
            );
        }),
    );
}

export const activeWorkspaceGuard: CanActivateFn = (route) => {
    const store = inject(Store);
    const router = inject(Router);
    const workspaceId = parseRouteId(route.paramMap.get('workspaceId'));

    return getOrFetchWorkspaces().pipe(
        map((workspaces: WorkspaceModel[]) => {
            if (!workspaces || workspaces.length === 0) {
                return router.createUrlTree(['/app/workspaces/create']);
            }

            const targetWorkspace =
                workspaceId === null
                    ? undefined
                    : workspaces.find((w: WorkspaceModel) => w.id === workspaceId);

            if (!targetWorkspace) {
                // Malformed or unknown workspace ID in URL
                return router.createUrlTree(['/app/workspaces']);
            }

            store.dispatch(actionWorkspace.setCurrent({ id: targetWorkspace.id }));
            return true;
        }),
    );
};
