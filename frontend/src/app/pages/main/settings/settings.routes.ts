import type { Routes } from '@angular/router';
import { SettingsRoot } from './settings.root';
import { SessionsPage } from './sessions';

export const routes: Routes = [
    {
        path: '',
        component: SettingsRoot,
        children: [
            {
                path: 'sessions',
                component: SessionsPage
            }
        ]
    }
]
