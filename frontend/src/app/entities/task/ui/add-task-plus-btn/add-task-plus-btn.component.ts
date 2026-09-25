import { Component, inject, input } from '@angular/core';
import { Store } from '@ngrx/store';
import { actionUI } from '@core/store/ui';

@Component({
    selector: 'lu-add-task-plus-btn',
    template: `
        <button class="button" (click)="openForm($event)">
            <i class="fa-solid fa-plus"></i>
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
