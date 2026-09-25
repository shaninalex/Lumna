import { Component, inject, ChangeDetectionStrategy } from '@angular/core';
import { CdkMenu, CdkMenuItem, CdkMenuTrigger } from '@angular/cdk/menu';
import { AsyncPipe } from '@angular/common';
import { UiService } from '../ui.service';
import type { Theme } from '../theme.manager';
import type { Observable } from 'rxjs';

@Component({
    selector: 'lu-theme-switcher',
    imports: [CdkMenu, CdkMenuItem, CdkMenuTrigger, AsyncPipe],
    changeDetection: ChangeDetectionStrategy.Eager,
    template: `
        <button class="btn btn-sm" [cdkMenuTriggerFor]="theme">
            @if (themeMode | async; as themeMode) {
                @switch (themeMode) {
                    @case ('dark') {
                        <i class="fa-solid fa-moon"></i>
                    }
                    @case ('light') {
                        <i class="fa-solid fa-sun"></i>
                    }
                    @case ('auto') {
                        <i class="fa-solid fa-circle-half-stroke"></i>
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
