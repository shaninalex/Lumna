import type { Routes } from '@angular/router';
import { provideState } from "@ngrx/store";
import { authGuard } from './auth.guard';
import { provideEffects } from '@ngrx/effects';

import { AppRoutes } from '@core';
import { activeWorkspaceGuard } from './workspace.guard';
import { lastRouteRedirect } from './lastRouteRedirect';
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
import { ActivityApi } from '@entities/activity/api';
import { activityFeature } from '@entities/activity';
import { routes as workspacesRoutes } from '@pages/main/workspaces';
import { routes as workspaceRoutes } from '@pages/main/workspace';
import { routes as projectRoutes } from '@pages/main/project';
import { MemberApi } from '@entities/member/api/member.api';
import { memberFeature } from '@entities/member/provider';
import { InvitationApi } from '@entities/invitation/api';
import { invitationFeature } from '@entities/invitation';

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
            ActivityApi,
            MemberApi,
            InvitationApi,

            provideEffects(mainEffects),

            provideState(workspaceFeature),
            provideState(projectFeature),
            provideState(userFeature),
            provideState(taskFeature),
            provideState(boardFeature),
            provideState(columnFeature),
            provideState(activityFeature),
            provideState(memberFeature),
            provideState(invitationFeature),
        ],
        children: [
            ...workspacesRoutes,
            {
                path: 'w/:workspaceId',
                canActivate: [activeWorkspaceGuard],
                children: [
                    ...workspaceRoutes,
                    ...projectRoutes,
                ],
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
