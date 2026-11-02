import { Component, inject } from '@angular/core';
import { CdkMenu, CdkMenuItem, CdkMenuTrigger } from '@angular/cdk/menu';
import { AsyncPipe } from '@angular/common';
import { UiService } from '../ui.service';
import type { Theme } from '../theme.manager';
import type { Observable } from 'rxjs';
import { NgIcon, provideIcons } from '@ng-icons/core';
import { reiconDarkLight, reiconMoon, reiconSun } from '@ng-icons/reicon';

@Component({
    selector: 'lu-theme-switcher',
    imports: [CdkMenu, CdkMenuItem, CdkMenuTrigger, AsyncPipe, NgIcon],
    providers: [provideIcons({reiconMoon, reiconSun, reiconDarkLight})],
    template: `
        <button class="btn btn-sm" [cdkMenuTriggerFor]="theme">
            @if (themeMode | async; as themeMode) {
                @switch (themeMode) {
                    @case ('dark') {
                        <ng-icon name="reiconMoon"/>
                    }
                    @case ('light') {
                        <ng-icon name="reiconSun"/>
                    }
                    @case ('auto') {
                        <ng-icon name="reiconDarkLight"/>
                    }
                }
            }
        </button>
        <ng-template #theme>
            <div class="list-group" cdkMenu>
                @if (themeMode | async; as currentTheme) {
                    <button
                        type="button"
                        cdkMenuItem
                        (click)="changeTheme('light')"
                        class="list-group-item list-group-item-action"
                        [class.active]="currentTheme === 'light'"
                    >
                        Light
                    </button>

                    <button
                        type="button"
                        cdkMenuItem
                        (click)="changeTheme('dark')"
                        class="list-group-item list-group-item-action"
                        [class.active]="currentTheme === 'dark'"
                    >
                        Dark
                    </button>

                    <button
                        type="button"
                        cdkMenuItem
                        (click)="changeTheme('auto')"
                        class="list-group-item list-group-item-action"
                        [class.active]="currentTheme === 'auto'"
                    >
                        Auto
                    </button>
                }
            </div>
        </ng-template>
    `,
})
export class ThemeSwitcherComponent {
    private ui = inject(UiService);
    themeMode: Observable<Theme> = this.ui.getTheme();

    changeTheme(v: Theme) {
        this.ui.setTheme(v);
    }
}
