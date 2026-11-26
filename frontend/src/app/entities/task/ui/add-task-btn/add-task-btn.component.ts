import { Component, inject, input } from '@angular/core';
import { Store } from '@ngrx/store';
import { actionUI } from '@core/store/ui';
import { MatButton } from '@angular/material/button';
import { MatIcon } from '@angular/material/icon';

@Component({
    selector: 'lu-add-task-btn',
    imports: [MatButton, MatIcon],
    template: `
        <button matButton type="button" class="w-100 add-task" (click)="openForm($event)">
            <mat-icon>add</mat-icon> Add Task
        </button>
    `,
    styles: `
        .add-task {
            justify-content: flex-start;
        }
    `,
})
export class AddTaskBtnComponent {
    columnId = input.required<number>();
    private store = inject(Store);

    openForm(event: MouseEvent): void {
        event.stopPropagation();
        this.store.dispatch(actionUI.taskColumnForm({columnId: this.columnId(), position: "bottom"}));
    }
}
