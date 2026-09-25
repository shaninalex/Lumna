import { Component, inject } from '@angular/core';
import { CdkMenu, CdkMenuItem, CdkMenuTrigger } from '@angular/cdk/menu';
import { Store } from '@ngrx/store';
import { actionSession } from '@core/store/session.actions';

@Component({
    selector: 'lu-user-menu',
    imports: [CdkMenu, CdkMenuItem, CdkMenuTrigger],

    template: `
        <button [cdkMenuTriggerFor]="menu" class="btn btn-sm">
            <img src="images/7.png" alt="" class="rounded-circle" style="width: 24px" />
        </button>

        <ng-template #menu>
            <div class="d-flex flex-column gap-2" cdkMenu>
                <div class="list-group">
                    <button
                        cdkMenuItem
                        type="button"
                        class="list-group-item list-group-item-action"
                    >
                        Account
                    </button>
                    <button
                        cdkMenuItem
                        type="button"
                        class="list-group-item list-group-item-action"
                    >
                        Settings
                    </button>
                    <button
                        cdkMenuItem
                        type="button"
                        class="list-group-item list-group-item-action"
                        (click)="logout()"
                    >
                        Logout
                    </button>
                </div>
                <ng-content select="[bottomMenuItems]" />
            </div>
        </ng-template>
    `,
})
export class UserMenuComponent {
    readonly store = inject(Store);

    logout(): void {
        this.store.dispatch(actionSession.loggingOut());
    }
}
