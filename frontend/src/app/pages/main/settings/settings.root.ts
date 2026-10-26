import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { GlobalLayout } from '@core/layout';

@Component({
    imports: [
        RouterOutlet,
        GlobalLayout,
    ],
    selector: 'lu-settings_root',
    template: `
        <lu-global-layout>
            <div class="container">
                <h4>Settings</h4>
                <router-outlet/>
            </div>
        </lu-global-layout>
    `
})
export class SettingsRoot {}
