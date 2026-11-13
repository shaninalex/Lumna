import type { Routes } from '@angular/router';
import { WorkspaceEntryPage } from './workspace-entry';
import { WorkspaceSettingsPage } from './workspace-settings';
import { WorkspaceSettingsMainPage } from './workspace-settings-main';
import { WorkspaceSettingsMembersPage } from './workspace-settings-members';
import { ProjectListPage } from './project-list';
import { ProjectCreatePage } from './project-create';

import { MembersListPage } from './members-list';
import { MembersInvitationsPage } from './members-invitations';
import { WorkspaceRemovePage } from './workspace-remove';

export const workspaceRoutes: Routes = [
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
                children: [
                    {
                        path: "",
                        component: MembersListPage,
                    },
                    {
                        path: "invitations",
                        component: MembersInvitationsPage,
                    }
                ],
            },
            {
                path: "remove",
                component: WorkspaceRemovePage,
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
