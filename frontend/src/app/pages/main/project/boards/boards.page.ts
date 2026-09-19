import { Component, inject } from '@angular/core';
import { UiService } from '@shared/ui';
import { RouterLink } from "@angular/router";
import { AppRoutes } from '@core';
import { BoardListFeature } from '@features';
import { MainLayout } from '@core/layout';

@Component({
    selector: 'lu-boards-page',
    imports: [MainLayout, RouterLink, BoardListFeature],
    template: `
        <lu-main-layout>
            <div class="container-fluid py-4">
                <!-- Header -->
                <div class="d-flex justify-content-between align-items-center mb-4">
                    <div>
                        <h2 class="mb-1">Boards</h2>
                        <p class="text-muted mb-0">Organize your project's work into multiple boards.</p>
                    </div>

                    <a [routerLink]="appRoutes.boardsCreate()" class="btn btn-primary btn-sm">
                        <i class="fa-solid fa-plus me-2"></i>
                        New Board
                    </a>
                </div>

                <lu-board-list-feature/>
            </div>
        </lu-main-layout>

    `,
})
export class BoardsPage {
    readonly appRoutes = inject(AppRoutes);
    private ui = inject(UiService);

    constructor() {
        this.ui.setPageTitle("Boards")
    }
}
