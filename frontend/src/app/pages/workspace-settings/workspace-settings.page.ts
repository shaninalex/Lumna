import { Component, inject } from '@angular/core';
import { MainLayout } from '@core/layout';
import { RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';
import { MatTabsModule } from '@angular/material/tabs';
import { AppRoutes } from '@core';

@Component({
    selector: 'lu-workspace-settings',
    imports: [MainLayout, RouterOutlet, RouterLink, RouterLinkActive, MatTabsModule],
    template: `
        <lu-main-layout>
            <nav mat-tab-nav-bar [tabPanel]="settingsPanel" mat-stretch-tabs="false" mat-align-tabs="start">
                <a mat-tab-link
                   [routerLink]="appRoutes.workspaceRoot(['settings'])"
                   routerLinkActive
                   #mainRla="routerLinkActive"
                   [routerLinkActiveOptions]="{exact: true}"
                   [active]="mainRla.isActive">
                    Main
                </a>
                <a mat-tab-link
                   [routerLink]="appRoutes.workspaceRoot(['settings', 'members'])"
                   routerLinkActive
                   #membersRla="routerLinkActive"
                   [active]="membersRla.isActive">
                    Members
                </a>
                <a mat-tab-link
                   [routerLink]="appRoutes.workspaceRoot(['settings', 'remove'])"
                   routerLinkActive
                   #removeRla="routerLinkActive"
                   [active]="removeRla.isActive">
                    Remove
                </a>
            </nav>

            <mat-tab-nav-panel #settingsPanel class="d-block pt-4">
                <router-outlet />
            </mat-tab-nav-panel>
        </lu-main-layout>
    `
})
export class WorkspaceSettingsPage {
    readonly appRoutes = inject(AppRoutes);

}
