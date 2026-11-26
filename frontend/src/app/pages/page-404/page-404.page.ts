import { Component, inject } from '@angular/core';
import { RouterLink } from '@angular/router';
import { MatButton } from '@angular/material/button';
import { MatIcon } from '@angular/material/icon';
import { StaticLayout } from '@core/layout';
import { UiService } from '@shared/ui';

@Component({
    selector: 'lu-page-404-page',
    imports: [StaticLayout, RouterLink, MatButton, MatIcon],
    styles: `
        .not-found-icon {
            width: 48px;
            height: 48px;
            font-size: 48px;
        }
    `,
    template: `
        <lu-static-layout>
            <div class="d-flex flex-column align-items-center gap-3 text-center py-4">
                <mat-icon class="not-found-icon text-muted">explore_off</mat-icon>
                <h1 class="m-0">Page Not Found</h1>
                <p class="text-muted m-0">The page you are looking for does not exist or has been moved.</p>
                <a matButton="filled" routerLink="/app">
                    <mat-icon>home</mat-icon>
                    Go home
                </a>
            </div>
        </lu-static-layout>
    `,
})
export class Page404 {
    private ui = inject(UiService);

    constructor() {
        this.ui.setPageTitle('Page not found');
    }
}
