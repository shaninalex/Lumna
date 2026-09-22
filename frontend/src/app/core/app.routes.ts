import type { Routes } from '@angular/router';

export const routes: Routes = [
    {
        path: 'auth',
        loadChildren: () => import('../pages/auth/auth.routes').then((m) => m.routes),
    },
    {
        path: 'app',
        loadChildren: () => import('../modules/main/main.routes').then((m) => m.routes),
    },
    {
        path: '',
        pathMatch: 'full',
        redirectTo: 'app',
    },
    {
        path: '404',
        loadComponent: () => import('../pages/system/page-404').then((m) => m.Page404),
    },
    {
        path: '**',
        redirectTo: '404',
    },
];
