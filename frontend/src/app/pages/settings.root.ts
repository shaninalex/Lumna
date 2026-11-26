import { Component } from '@angular/core';
import { RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';
import { MatTabsModule } from '@angular/material/tabs';
import { GlobalLayout } from '@core/layout';

@Component({
    imports: [
        RouterOutlet,
        GlobalLayout,
        RouterLink,
        RouterLinkActive,
        MatTabsModule,
    ],
    selector: 'lu-settings_root',
    template: `
        <lu-global-layout>
            <div class="container pt-4">
                <h1 class="lu-page-title mb-3">Settings</h1>
                <nav mat-tab-nav-bar [tabPanel]="settingsPanel" mat-stretch-tabs="false" mat-align-tabs="start">
                    @for (link of links; track link.path) {
                        <a mat-tab-link
                           [routerLink]="link.path"
                           routerLinkActive
                           #rla="routerLinkActive"
                           [routerLinkActiveOptions]="{ exact: link.exact }"
                           [active]="rla.isActive">
                            {{ link.label }}
                        </a>
                    }
                </nav>
                <mat-tab-nav-panel #settingsPanel class="d-block pt-4">
                    <router-outlet/>
                </mat-tab-nav-panel>
            </div>
        </lu-global-layout>
    `
})
export class SettingsRoot {
    readonly links: { path: string[]; label: string; exact: boolean }[] = [
        { path: ['/app/settings'], label: 'Home', exact: true },
        { path: ['/app/settings/sessions'], label: 'Sessions', exact: false },
        { path: ['/app/settings/security'], label: 'Security', exact: false },
        { path: ['/app/settings/remove'], label: 'Remove', exact: false },
    ];
}
