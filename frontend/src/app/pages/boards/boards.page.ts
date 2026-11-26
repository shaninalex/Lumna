import { Component, inject } from '@angular/core';
import { UiService } from '@shared/ui';
import { RouterLink } from '@angular/router';
import { AppRoutes } from '@core';
import { BoardListFeature } from '@features';
import { MainLayout } from '@core/layout';
import { MatButton } from '@angular/material/button';
import { MatIcon } from '@angular/material/icon';

@Component({
    selector: 'lu-boards-page',
    imports: [MainLayout, RouterLink, BoardListFeature, MatButton, MatIcon],
    template: `
        <lu-main-layout>
            <div class="container-fluid py-4">
                <nav class="lu-breadcrumb mb-2" aria-label="Breadcrumb">
                    <a [routerLink]="">Workspace</a>
                    <mat-icon>chevron_right</mat-icon>
                    <a [routerLink]="">Project</a>
                </nav>

                <!-- Header -->
                <div class="d-flex justify-content-between align-items-center mb-4">
                    <div>
                        <h2 class="lu-page-title mb-1">Boards</h2>
                        <p class="text-muted mb-0">
                            Organize your project's work into multiple boards.
                        </p>
                    </div>

                    <a matButton="filled" [routerLink]="appRoutes.boardsCreate()">
                        <mat-icon>add</mat-icon>
                        New Board
                    </a>
                </div>

                <lu-board-list-feature />
            </div>
        </lu-main-layout>
    `,
})
export class BoardsPage {
    readonly appRoutes = inject(AppRoutes);
    private ui = inject(UiService);

    constructor() {
        this.ui.setPageTitle('Boards');
    }
}
