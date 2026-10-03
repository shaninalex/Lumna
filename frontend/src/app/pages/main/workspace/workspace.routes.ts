import type { Routes } from '@angular/router';
import { WorkspaceEntryPage } from './workspace-entry';
import { WorkspaceSettingsPage } from './workspace-settings';
import { WorkspaceSettingsMainPage } from './workspace-settings-main';
import { WorkspaceSettingsMembersPage } from './workspace-settings-members';
import { ProjectListPage } from './project-list';
import { ProjectCreatePage } from './project-create';

import { routes as workspaceMembersRoutes } from './workspace-settings-members/workspace-members.routes'

export const routes: Routes = [
    {
        path: '',
        component: WorkspaceEntryPage,
    },
    {
        path: 'settings',
        component: WorkspaceSettingsPage,
        children: [
            {
                path: "",
                component: WorkspaceSettingsMainPage,
            },
            {
                path: "members",
                component: WorkspaceSettingsMembersPage,
                children: workspaceMembersRoutes,
            }
        ]
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
