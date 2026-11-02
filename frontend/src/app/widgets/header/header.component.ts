import { Component, inject } from '@angular/core';
import { Store } from '@ngrx/store';
import { UserMenuComponent } from '@entities/user';
import { NotificationsDropdownComponent } from '@entities/notification';
import { ThemeSwitcherComponent } from '@shared/ui';
import { ProjectDropdownComponent } from '@features/project';
import { WorkspaceSwitcherComponent } from '@features';
import { actionUI, selectUI } from '@core/store/ui';
import { NgClass } from '@angular/common';
import { NgIcon, provideIcons } from '@ng-icons/core';
import { reiconChevronLeft, reiconMenu } from '@ng-icons/reicon';

@Component({
    selector: 'lu-header',
    imports: [
        UserMenuComponent,
        ThemeSwitcherComponent,
        NotificationsDropdownComponent,
        ProjectDropdownComponent,
        WorkspaceSwitcherComponent,
        NgClass,
        NgIcon,
    ],
    providers: [provideIcons({ reiconChevronLeft, reiconMenu })],
    styleUrl: './header.component.scss',
    template: `
        <nav class="navbar navbar-expand-lg bg-body-tertiary">
            <div class="container-fluid">
                <div class="d-flex align-items-center gap-2">
                    <button class="button logo-button"
                            [ngClass]="{'logo-button-closed': !sidebarOpen()}"
                            (click)="toggleSidebar()">
                        @if (!sidebarOpen()) {
                            <ng-icon name="reiconChevronLeft" />
                        } @else {
                            <ng-icon name="reiconMenu" />
                            <img src="images/logo-icon.svg">
                        }
                    </button>
                    <lu-project-dropdown />
                </div>

                <div class="flex align-items-center">
                    <lu-theme-switcher />
                    <lu-notifications-dropdown />
                    <lu-user-menu>
                        <lu-workspace-switcher bottomMenuItems />
                    </lu-user-menu>
                </div>
            </div>
        </nav>
    `,
})
export class HeaderComponent {
    private store = inject(Store);
    sidebarOpen = this.store.selectSignal(selectUI.sidebarOpen);

    toggleSidebar(): void {
        this.store.dispatch(actionUI.sidebarState({ state: !this.sidebarOpen() }));
    }
}
