import { Component, inject, input } from '@angular/core';
import { Store } from '@ngrx/store';
import { actionUI } from '@core/store/ui';

@Component({
    selector: 'lu-add-task-btn',
    template: `
        <button class="button button-pill w-100 text-start" (click)="openForm($event)">
            <i class="fa-solid fa-plus"></i> Add Task
        </button>
    `
})
export class AddTaskBtnComponent {
    columnId = input.required<number>();
    private store = inject(Store);

    openForm(event: MouseEvent): void {
        event.stopPropagation();
        this.store.dispatch(actionUI.taskColumnForm({columnId: this.columnId(), position: "bottom"}));
    }
}
