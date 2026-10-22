import { Component, inject } from '@angular/core';
import { MainLayout } from '@core/layout';
import { RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';
import { AppRoutes } from '@core';

@Component({
    selector: 'lu-workspace-settings',
    imports: [MainLayout, RouterOutlet, RouterLink, RouterLinkActive],
    styleUrl: './workspace-settings.page.scss',
    template: `
        <lu-main-layout>
            <nav class="navbar navbar-expand-lg bg-body-tertiary rounded mb-4">
                <div class="container-fluid">
                    <div class="collapse navbar-collapse">
                        <ul class="navbar-nav">
                            <li class="nav-item">
                                <a class="nav-link active" [routerLink]="appRoutes.workspaceRoot(['settings'])"
                                   routerLinkActive="active" [routerLinkActiveOptions]="{exact: true}">
                                    Main
                                </a>
                            </li>
                            <li class="nav-item">
                                <a class="nav-link" [routerLink]="appRoutes.workspaceRoot(['settings', 'members'])"
                                   routerLinkActive="active">
                                    Members
                                </a>
                            </li>
                        </ul>
                    </div>
                </div>
            </nav>

            <router-outlet />
        </lu-main-layout>
    `
})
export class WorkspaceSettingsPage {
    readonly appRoutes = inject(AppRoutes);

}
