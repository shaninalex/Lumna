import { Component, inject } from '@angular/core';
import { RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';
import { AppRoutes } from '@core';

@Component({
    selector: 'lu-workspace-settings-members',
    imports: [RouterOutlet, RouterLink, RouterLinkActive],
    template: `
        <div class="card">
            <div class="card-header">
                <ul class="nav nav-tabs card-header-tabs">
                    <li class="nav-item">
                        <a class="nav-link" [routerLink]="appRoutes.workspaceRoot(['settings', 'members'])"
                           routerLinkActive="active" [routerLinkActiveOptions]="{exact: true}">
                            List
                        </a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" aria-current="page"
                           [routerLink]="appRoutes.workspaceRoot(['settings', 'members', 'invitations'])"
                           routerLinkActive="active">
                            Invitations
                        </a>
                    </li>
                </ul>
            </div>
            <div class="card-body">
                <router-outlet/>
            </div>
        </div>

    `,
})
export class WorkspaceSettingsMembersPage {
    readonly appRoutes = inject(AppRoutes);
}
