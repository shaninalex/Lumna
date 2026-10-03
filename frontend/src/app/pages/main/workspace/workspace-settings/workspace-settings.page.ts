import { Component, inject } from '@angular/core';
import { MainLayout } from '@core/layout';
import { RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';
import { AppRoutes } from '@core';

@Component({
    selector: 'lu-workspace-settings',
    imports: [MainLayout, RouterOutlet, RouterLink, RouterLinkActive],
    styleUrl: './workspace-settings.page.css',
    template: `
        <lu-main-layout>
            <ul class="workspace-settings-menu">
                <li>
                    <a [routerLink]="appRoutes.workspaceRoot(['settings'])"
                        routerLinkActive="active" [routerLinkActiveOptions]="{exact: true}">
                        Main
                    </a>
                </li>
                <li>
                    <a [routerLink]="appRoutes.workspaceRoot(['settings', 'members'])"
                       routerLinkActive="active">
                        Members
                    </a>
                </li>
            </ul>
            <router-outlet />
        </lu-main-layout>
    `
})
export class WorkspaceSettingsPage {
    readonly appRoutes = inject(AppRoutes);

}
