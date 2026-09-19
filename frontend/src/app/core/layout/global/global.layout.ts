import { Component, inject } from '@angular/core';
import { UserMenuComponent } from '@entities/user'
import { RouterLink } from "@angular/router";
import { Store } from '@ngrx/store';
import { selectWorkspaces } from '@entities/workspace';
import { ThemeSwitcherComponent } from '@shared/ui';

@Component({
    selector: 'lu-global-layout',
    imports: [ThemeSwitcherComponent, UserMenuComponent, RouterLink],
    template: `
        <nav class="navbar navbar-expand-lg border-bottom bg-body">
            <div class="container-fluid">
                <a [routerLink]="['/app/w', currentWorkspaceId() || '']">
                    <img src="images/logo-h.svg" alt="" style="width: 160px">
                </a>
                <div class="flex align-items-center">
                    <lu-theme-switcher/>
                    <button class="btn btn-sm">
                        <i class="fa-solid fa-bell"></i>
                    </button>
                    <lu-user-menu/>
                </div>
            </div>
        </nav>

        <ng-content/>
    `,
})
export class GlobalLayout {
    private store = inject(Store);
    currentWorkspaceId = this.store.selectSignal(selectWorkspaces.currentWorkspaceId);
}
