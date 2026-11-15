import type { Routes } from '@angular/router';
import { OverviewPage } from '@pages/overview';
import { NotificationsPage } from '@pages/notifications';
import { RecentPage } from '@pages/recent';
import { PluginsPage } from '@pages/plugins';

export const workRoutes: Routes = [
    {
        path: "overview",
        component: OverviewPage
    },
    {
        path: "recent",
        component: RecentPage
    },
    {
        path: "notifications",
        component: NotificationsPage
    },
    {
        path: "plugins",
        component: PluginsPage
    }
]
