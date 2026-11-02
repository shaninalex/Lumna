import { Component, Input } from '@angular/core';
import { CdkMenu, CdkMenuItem, CdkMenuTrigger } from '@angular/cdk/menu';
import { NgIcon, provideIcons } from '@ng-icons/core';
import { reiconMenu2 } from '@ng-icons/reicon';

@Component({
    selector: 'lu-column-menu-dropdown',
    imports: [CdkMenu, CdkMenuItem, CdkMenuTrigger, NgIcon],
    providers: [provideIcons({ reiconMenu2 })],
    template: `
        <button [cdkMenuTriggerFor]="taskMenu" class="btn btn-sm p-1 lh-1">
            <ng-icon name="reiconMenu2" />
        </button>
        <ng-template #taskMenu>
            <div class="list-group" cdkMenu>
                <button
                    cdkMenuItem
                    type="button"
                    class="list-group-item list-group-item-action"
                    (click)="delete()"
                >
                    Close
                </button>
                <ng-content select="[columnMenuSlot]" />
            </div>
        </ng-template>
    `,
})
export class ColumnMenuDropdownComponent {
    @Input() stageId: number;

    delete() {}
}
