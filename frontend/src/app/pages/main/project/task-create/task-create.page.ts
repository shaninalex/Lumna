import { Component, inject } from '@angular/core';
import { ModalLayout } from '@core/layout';
import { UiService } from '@shared/ui';

@Component({
    selector: 'lu-task-create-page',
    imports: [ModalLayout],
    template: `
        <lu-modal-layout>
            <div class="container-fluid py-4">
                <h3>Create task</h3>
            </div>
        </lu-modal-layout>
    `
})
export class TaskCreatePage {
    private ui = inject(UiService);

    constructor() {
        this.ui.setPageTitle("Create Task")
    }
}
