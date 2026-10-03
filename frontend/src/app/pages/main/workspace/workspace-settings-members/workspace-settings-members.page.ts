import { Component, inject } from '@angular/core';
import { RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';
import { AppRoutes } from '@core';

@Component({
    selector: 'lu-workspace-settings-members',
    imports: [RouterOutlet, RouterLink, RouterLinkActive],
    template: `
        <ul class="list-menu">
            <li>
                <a [routerLink]="appRoutes.workspaceRoot(['settings', 'members'])"
                   routerLinkActive="active" [routerLinkActiveOptions]="{exact: true}">
                    List
                </a>
            </li>
            <li>
                <a [routerLink]="appRoutes.workspaceRoot(['settings', 'members', 'invitations'])"
                   routerLinkActive="active">
                    Invitations
                </a>
            </li>
        </ul>
        <router-outlet />
    `,
})
export class WorkspaceSettingsMembersPage {
    readonly appRoutes = inject(AppRoutes);
}
