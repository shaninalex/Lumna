import type { Routes } from '@angular/router';
import { SettingsRoot } from './settings.root';
import { SessionsPage } from './sessions';
import { AccountSecurityPage } from './account-security';
import { AccountPage } from './account';
import { AccountRemovePage } from './account-remove';

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
                component: AccountSecurityPage
            },
            {
                path: 'remove',
                component: AccountRemovePage,
            }
        ]
    }
]
