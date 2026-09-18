import type { Routes } from '@angular/router';
import { ProjectListPage } from './project-list'
import { ProjectCreatePage } from './project-create'
import { paramMatches, paramMatchesDigitsOnly } from '@shared/utils';
import { WorkspaceEntryPage } from './workspace-entry'
import { activeProjectGuard } from "../project.guard";
import { InboxPage } from "./inbox";
import { BoardsPage } from './boards';
import { BoardCreatePage } from './board-create';
import { BoardDetailPage } from './board-detail';
import { BacklogPage } from './backlog';
import { TaskCreatePage } from './task-create';
import { TaskDetailPage } from './task-detail';


export const routes: Routes = [
    {
        path: '',
        component: WorkspaceEntryPage,
    },
    {
        path: 'p/:projectId',
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
                children: [
                    {
                        path: 'task/create',
                        component: TaskCreatePage,
                    },
                    {
                        path: 'task/:taskId',
                        component: TaskDetailPage,
                        canMatch: [paramMatches("taskId", paramMatchesDigitsOnly)]
                    },
                ]
            },
            {
                path: 'backlog',
                component: BacklogPage,
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
        component: ProjectListPage,
    },
    {
        path: 'projects/create',
        component: ProjectCreatePage,
    },
]
