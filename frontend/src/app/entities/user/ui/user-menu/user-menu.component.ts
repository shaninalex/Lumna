import { Component, inject } from '@angular/core';
import { MatIconButton } from '@angular/material/button';
import { MatIcon } from '@angular/material/icon';
import { MatMenuModule } from '@angular/material/menu';
import { Store } from '@ngrx/store';
import { actionSession } from '@core/store/session/session.actions';
import { RouterLink } from '@angular/router';

@Component({
    selector: 'lu-user-menu',
    imports: [MatIconButton, MatIcon, MatMenuModule, RouterLink],
    template: `
        <button matIconButton [matMenuTriggerFor]="menu" aria-label="User menu">
            <img src="images/7.png" alt="" class="rounded-circle avatar" />
        </button>

        <mat-menu #menu="matMenu" xPosition="before">
            <a mat-menu-item [routerLink]="['/app/settings']">
                <mat-icon>settings</mat-icon>
                <span>Settings</span>
            </a>
            <button mat-menu-item type="button" (click)="logout()">
                <mat-icon>logout</mat-icon>
                <span>Logout</span>
            </button>
            <ng-content select="[bottomMenuItems]" />
        </mat-menu>
    `,
    styles: `
        .avatar {
            display: block;
            width: 24px;
            height: 24px;
        }
    `,
})
export class UserMenuComponent {
    readonly store = inject(Store);

    logout(): void {
        this.store.dispatch(actionSession.loggingOut());
    }
}
