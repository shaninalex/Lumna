import type { Routes } from '@angular/router';
import { SettingsRoot } from './settings.root';
import { SessionsPage } from './sessions';
import { SecurityPage } from './security';
import { AccountPage } from './account';
import { RemovePage } from './remove';

export const settingsRoutes: Routes = [
    {
        path: '',
        component: SettingsRoot,
        children: [
            {
                path: '',
                component: AccountPage
            },
            {
                path: 'sessions',
                component: SessionsPage
            },
            {
                path: 'security',
                component: SecurityPage
            },
            {
                path: 'remove',
                component: RemovePage,
            }
        ]
    }
]
