import { Component, inject } from '@angular/core';
import { UserMenuComponent } from '@entities/user';
import { RouterLink } from '@angular/router';
import { Store } from '@ngrx/store';
import { selectWorkspaces } from '@entities/workspace';
import { ThemeSwitcherComponent } from '@shared/ui';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatIconButton } from '@angular/material/button';
import { MatIcon } from '@angular/material/icon';

@Component({
    selector: 'lu-global-layout',
    imports: [ThemeSwitcherComponent, UserMenuComponent, RouterLink, MatToolbarModule, MatIconButton, MatIcon],
    template: `
        <mat-toolbar class="border-bottom">
            <a [routerLink]="['/app/w', currentWorkspaceId() || '']" class="d-flex">
                <img src="images/logo-h.svg" alt="Lumna" style="width: 160px" />
            </a>
            <span class="flex-grow-1"></span>
            <lu-theme-switcher />
            <button matIconButton aria-label="Notifications">
                <mat-icon>notifications</mat-icon>
            </button>
            <lu-user-menu />
        </mat-toolbar>

        <ng-content />
    `,
})
export class GlobalLayout {
    private store = inject(Store);
    currentWorkspaceId = this.store.selectSignal(selectWorkspaces.currentWorkspaceId);
}
