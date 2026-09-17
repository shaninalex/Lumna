import type { Routes } from '@angular/router';
import {
    WorkspacesPage,
    WorkspaceCreateComponent,
    WorkspaceEntryPage,
    InboxPage,
    BoardsPage,
    BacklogPage,
    ProjectsPage,
    ProjectsCreatePage,
    TaskCreatePage,
    TaskDetailPage,
    BoardCreatePage,
    BoardDetailPage,
} from '@pages';
import { MODAL_OUTLET, PROJECT_ROUTE_PATH } from '@core';
import { authGuard } from './auth.guard';
import { activeProjectGuard } from './project.guard';
import { activeWorkspaceGuard } from './workspace.guard';
import { lastRouteRedirect } from './lastRouteRedirect';
import { paramMatches, paramMatchesDigitsOnly } from '@shared/utils';

export const routes: Routes = [
    {
        path: '',
        canMatch: [authGuard],
        children: [
            {
                path: 'workspaces',
                component: WorkspacesPage,
            },
            {
                path: 'workspaces/create',
                component: WorkspaceCreateComponent,
            },
            {
                path: 'w/:workspaceId',
                canActivate: [activeWorkspaceGuard],
                children: [
                    {
                        path: '',
                        component: WorkspaceEntryPage,
                    },
                    {
                        path: PROJECT_ROUTE_PATH,
                        canActivate: [activeProjectGuard],
                        children: [
                            {
                                path: 'inbox',
                                component: InboxPage,
                            },
                            {
                                path: 'boards',
                                component: BoardsPage,
                            },
                            {
                                path: 'board/create',
                                component: BoardCreatePage,
                            },
                            {
                                path: 'board/:boardId',
                                component: BoardDetailPage,
                            },
                            {
                                path: 'backlog',
                                component: BacklogPage,
                            },
                            {
                                path: 'task/create',
                                component: TaskCreatePage,
                            },
                            {
                                path: 'task/:taskId',
                                component: TaskDetailPage,
                                outlet: MODAL_OUTLET,
                                canMatch: [paramMatches("taskId", paramMatchesDigitsOnly)]
                            },
                            {
                                path: '',
                                pathMatch: 'full',
                                redirectTo: 'inbox',
                            },
                        ],
                    },
                    {
                        path: 'projects',
                        component: ProjectsPage,
                    },
                    {
                        path: 'projects/create',
                        component: ProjectsCreatePage,
                    },
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
