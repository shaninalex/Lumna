import { Component, inject } from '@angular/core';
import { RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';
import { MatCardModule } from '@angular/material/card';
import { MatTabsModule } from '@angular/material/tabs';
import { AppRoutes } from '@core';

@Component({
    selector: 'lu-workspace-settings-members',
    imports: [RouterOutlet, RouterLink, RouterLinkActive, MatCardModule, MatTabsModule],
    template: `
        <mat-card>
            <nav mat-tab-nav-bar [tabPanel]="membersPanel" mat-stretch-tabs="false" mat-align-tabs="start">
                <a mat-tab-link
                   [routerLink]="appRoutes.workspaceRoot(['settings', 'members'])"
                   routerLinkActive
                   #listRla="routerLinkActive"
                   [routerLinkActiveOptions]="{exact: true}"
                   [active]="listRla.isActive">
                    List
                </a>
                <a mat-tab-link
                   [routerLink]="appRoutes.workspaceRoot(['settings', 'members', 'invitations'])"
                   routerLinkActive
                   #invitationsRla="routerLinkActive"
                   [active]="invitationsRla.isActive">
                    Invitations
                </a>
            </nav>
            <mat-card-content class="pt-3">
                <mat-tab-nav-panel #membersPanel>
                    <router-outlet/>
                </mat-tab-nav-panel>
            </mat-card-content>
        </mat-card>
    `,
})
export class WorkspaceSettingsMembersPage {
    readonly appRoutes = inject(AppRoutes);
}
