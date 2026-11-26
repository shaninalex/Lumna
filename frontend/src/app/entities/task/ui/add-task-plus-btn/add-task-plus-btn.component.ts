import { Component, inject, input } from '@angular/core';
import { Store } from '@ngrx/store';
import { actionUI } from '@core/store/ui';
import { MatIconButton } from '@angular/material/button';
import { MatIcon } from '@angular/material/icon';

@Component({
    selector: 'lu-add-task-plus-btn',
    imports: [MatIconButton, MatIcon],
    template: `
        <button matIconButton type="button" aria-label="Add task" (click)="openForm($event)">
            <mat-icon>add</mat-icon>
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
