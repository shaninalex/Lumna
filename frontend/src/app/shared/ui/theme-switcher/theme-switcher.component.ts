import { Component, inject } from '@angular/core';
import { AsyncPipe, NgClass } from '@angular/common';
import { MatIconButton } from '@angular/material/button';
import { MatIcon } from '@angular/material/icon';
import { MatMenuModule } from '@angular/material/menu';
import { MatTooltip } from '@angular/material/tooltip';
import { UiService } from '../ui.service';
import type { Theme } from '../theme.manager';
import type { Observable } from 'rxjs';

@Component({
    selector: 'lu-theme-switcher',
    imports: [AsyncPipe, MatIconButton, MatIcon, MatMenuModule, MatTooltip, NgClass],
    template: `
        @let current = themeMode | async;
        <button
            matIconButton
            [matMenuTriggerFor]="theme"
            matTooltip="Theme"
            aria-label="Change theme"
        >
            <mat-icon>{{ icons[current ?? 'auto'] }}</mat-icon>
        </button>
        <mat-menu #theme="matMenu">
            @for (option of options; track option.value) {
                <button
                    mat-menu-item
                    (click)="changeTheme(option.value)"
                    class="theme-option"
                    [ngClass]="{ selected: current === option.value }"
                >
                    <mat-icon>{{ icons[option.value] }}</mat-icon>
                    <span>{{ option.label }}</span>
                </button>
            }
        </mat-menu>
    `,
    styles: `
        .theme-option.selected {
            background-color: var(--mat-sys-surface-dim);
        }
    `,
})
export class ThemeSwitcherComponent {
    private ui = inject(UiService);
    themeMode: Observable<Theme> = this.ui.getTheme();

    readonly icons: Record<Theme, string> = {
        light: 'light_mode',
        dark: 'dark_mode',
        auto: 'contrast',
    };
    readonly options: { value: Theme; label: string }[] = [
        { value: 'light', label: 'Light' },
        { value: 'dark', label: 'Dark' },
        { value: 'auto', label: 'System' },
    ];

    changeTheme(v: Theme) {
        this.ui.setTheme(v);
    }
}
