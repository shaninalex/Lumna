import { Component, inject } from '@angular/core';
import { UserMenuComponent } from '@entities/user';
import { RouterLink } from '@angular/router';
import { Store } from '@ngrx/store';
import { selectWorkspaces } from '@entities/workspace';
import { ThemeSwitcherComponent } from '@shared/ui';
import { NgIcon, provideIcons } from '@ng-icons/core';
import { reiconBell } from '@ng-icons/reicon';

@Component({
    selector: 'lu-global-layout',
    imports: [ThemeSwitcherComponent, UserMenuComponent, RouterLink, NgIcon],
    providers: [provideIcons({ reiconBell })],
    template: `
        <nav class="navbar navbar-expand-lg border-bottom bg-body">
            <div class="container-fluid">
                <a [routerLink]="['/app/w', currentWorkspaceId() || '']">
                    <img src="images/logo-h.svg" alt="" style="width: 160px" />
                </a>
                <div class="flex align-items-center">
                    <lu-theme-switcher />
                    <button class="btn btn-sm">
                        <ng-icon name="reiconBell" />
                    </button>
                    <lu-user-menu />
                </div>
            </div>
        </nav>

        <ng-content />
    `,
})
export class GlobalLayout {
    private store = inject(Store);
    currentWorkspaceId = this.store.selectSignal(selectWorkspaces.currentWorkspaceId);
}
