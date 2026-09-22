import type { Routes } from '@angular/router';
import { provideState } from "@ngrx/store";
import { authGuard } from './auth.guard';
import { provideEffects } from '@ngrx/effects';

import { AppRoutes } from '@core';
import { activeWorkspaceGuard } from './workspace.guard';
import { lastRouteRedirect } from './lastRouteRedirect';
import { projectRoutes, workspaceRoutes } from "@pages";
import { MainComponent } from './main.component';
import { mainEffects } from './store';
import { WorkspaceApi, workspaceFeature } from '@entities/workspace';
import { ProjectApi, projectFeature } from '@entities/project';
import { UserApi, userFeature } from '@entities/user';
import { TaskApi, taskFeature } from '@entities/task';
import { ColumnApi, columnFeature } from '@entities/column';
import { BoardApi, boardFeature } from '@entities/board';
import { KanbanApi } from '@widgets/kanban-widget/api';
import { WebSocketService } from './websocket.service';


export const routes: Routes = [
    {
        path: '',
        canMatch: [authGuard],
        component: MainComponent,
        providers: [
            AppRoutes,
            WorkspaceApi,
            ProjectApi,
            UserApi,
            TaskApi,
            BoardApi,
            ColumnApi,
            KanbanApi,
            WebSocketService,

            provideEffects(mainEffects),

            provideState(workspaceFeature),
            provideState(projectFeature),
            provideState(userFeature),
            provideState(taskFeature),
            provideState(boardFeature),
            provideState(columnFeature),
        ],
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
