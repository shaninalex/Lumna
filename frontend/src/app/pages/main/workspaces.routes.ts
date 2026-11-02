import type { Routes } from '@angular/router';
import { WorkspacesPage } from './workspaces';
import { WorkspaceCreateComponent } from './workspace-create';


export const workspacesRoutes: Routes = [
    {
        path: 'workspaces',
        component: WorkspacesPage,
    },
    {
        path: 'workspaces/create',
        component: WorkspaceCreateComponent,
    },
]
