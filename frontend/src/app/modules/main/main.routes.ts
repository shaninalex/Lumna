import type { Routes } from '@angular/router';
import { authGuard } from './auth.guard';
import { activeWorkspaceGuard } from './workspace.guard';
import { lastRouteRedirect } from './lastRouteRedirect';
import { workspaceRoutes, projectRoutes } from "@pages";

export const routes: Routes = [
    {
        path: '',
        canMatch: [authGuard],
        children: [
            ...workspaceRoutes,
            {
                path: 'w/:workspaceId',
                canActivate: [activeWorkspaceGuard],
                children: projectRoutes,
            },
            {
                path: '',
                pathMatch: 'full',
                canActivate: [lastRouteRedirect],
                children: [],
            },
        ],
    },
];
