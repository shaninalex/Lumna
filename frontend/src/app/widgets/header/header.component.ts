import { Component, inject } from '@angular/core';
import { Store } from '@ngrx/store';
import { UserMenuComponent } from '@entities/user';
import { NotificationsDropdownComponent } from '@entities/notification';
import { ThemeSwitcherComponent } from '@shared/ui';
import { WorkspaceSwitcherComponent } from '@features';
import { actionUI, selectUI } from '@core/store/ui';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatIconButton } from '@angular/material/button';
import { MatIcon } from '@angular/material/icon';
import { MatTooltip } from '@angular/material/tooltip';

@Component({
    selector: 'lu-header',
    imports: [
        UserMenuComponent,
        ThemeSwitcherComponent,
        NotificationsDropdownComponent,
        WorkspaceSwitcherComponent,
        MatToolbarModule,
        MatIconButton,
        MatIcon,
        MatTooltip,
    ],
    styleUrl: './header.component.scss',
    template: `
        <mat-toolbar>
            <div class="d-flex align-items-center gap-2">
                <button matIconButton
                        class="logo-button"
                        [class.logo-button-closed]="!sidebarOpen()"
                        [matTooltip]="sidebarOpen() ? 'Collapse menu' : 'Expand menu'"
                        (click)="toggleSidebar()">
                    @if (!sidebarOpen()) {
                        <mat-icon>menu</mat-icon>
                    } @else {
                        <mat-icon class="hover-icon">menu_open</mat-icon>
                        <img src="images/logo-icon.svg" alt="">
                    }
                </button>
                <span class="lu-page-title">Lumna</span>
            </div>

            <span class="flex-grow-1"></span>

            <div class="d-flex align-items-center gap-1">
                <lu-theme-switcher />
                <lu-notifications-dropdown />
                <lu-user-menu>
                    <lu-workspace-switcher bottomMenuItems />
                </lu-user-menu>
            </div>
        </mat-toolbar>
    `,
})
export class HeaderComponent {
    private store = inject(Store);
    sidebarOpen = this.store.selectSignal(selectUI.sidebarOpen);

    toggleSidebar(): void {
        this.store.dispatch(actionUI.sidebarState({ state: !this.sidebarOpen() }));
    }
}
