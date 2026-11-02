import { Component, inject, input } from '@angular/core';
import { Store } from '@ngrx/store';
import { actionUI } from '@core/store/ui';
import { NgIcon, provideIcons } from '@ng-icons/core';
import { reiconPlus } from '@ng-icons/reicon';

@Component({
    selector: 'lu-add-task-plus-btn',
    imports: [
        NgIcon
    ],
    providers: [provideIcons({ reiconPlus })],
    template: `
        <button class="button" (click)="openForm($event)">
            <ng-icon name="reiconPlus"/>
        </button>
    `
})
export class AddTaskPlusBtnComponent {
    columnId = input.required<number>();

    private store = inject(Store);

    openForm(event: MouseEvent): void {
        event.stopPropagation();
        this.store.dispatch(actionUI.taskColumnForm({columnId: this.columnId(), position: 'top'}));
    }
}
