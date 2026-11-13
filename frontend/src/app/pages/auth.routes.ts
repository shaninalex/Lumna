import type { Routes } from '@angular/router';
import { AuthRoot } from './auth.root';
import { LoginPage } from './login';

export const routes: Routes = [
    {
        path: '',
        component: AuthRoot,
        children: [
            {
                path: 'login',
                component: LoginPage,
            },
            {
                path: '',
                pathMatch: 'full',
                redirectTo: '/auth/login'
            }
        ]
    }
];
