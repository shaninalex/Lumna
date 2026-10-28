import { Component } from '@angular/core';
import { RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';
import { GlobalLayout } from '@core/layout';

@Component({
    imports: [
        RouterOutlet,
        GlobalLayout,
        RouterLink,
        RouterLinkActive
    ],
    selector: 'lu-settings_root',
    template: `
        <lu-global-layout>
            <nav class="navbar navbar-expand-lg bg-body-tertiary">
                <div class="container">
                    <a class="navbar-brand" href="#">Settings</a>
                    <div class="collapse navbar-collapse" id="navbarNav">
                        <ul class="navbar-nav">
                            <li class="nav-item">
                                <a class="nav-link" routerLinkActive="active"
                                   [routerLinkActiveOptions]="{ exact: true }" [routerLink]="['/app/settings']">Home</a>
                            </li>
                            <li class="nav-item">
                                <a class="nav-link" routerLinkActive="active" [routerLink]="['/app/settings/sessions']">Sessions</a>
                            </li>
                            <li class="nav-item">
                                <a class="nav-link" routerLinkActive="active" [routerLink]="['/app/settings/security']">Security</a>
                            </li>
                        </ul>
                    </div>
                </div>
            </nav>

            <div class="container pt-4">
                <router-outlet/>
            </div>
        </lu-global-layout>
    `
})
export class SettingsRoot {}
