import { Component, inject } from '@angular/core';
import { BoardCreateFeature } from '@features';
import { MainLayout } from '@core/layout';
import { UiService } from '@shared/ui';

@Component({
    selector: 'lu-board-create-page',
    imports: [MainLayout, BoardCreateFeature],
    template: `
        <lu-main-layout>
            <div class="container-fluid">
                <h1>Create Board</h1>

                <lu-board-create-feature />
            </div>
        </lu-main-layout>
    `,
})
export class BoardCreatePage {
    private ui = inject(UiService);

    constructor() {
        this.ui.setPageTitle("Create Board");
    }
}
