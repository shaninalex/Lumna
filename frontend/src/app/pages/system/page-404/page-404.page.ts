import { Component, inject, ChangeDetectionStrategy } from '@angular/core';
import { StaticLayout } from '@core/layout';
import { UiService } from '@shared/ui';

@Component({
    selector: 'lu-page-404-page',
    imports: [StaticLayout],
    changeDetection: ChangeDetectionStrategy.Eager,
    template: `
        <lu-static-layout>
            <h1>Page Not Found</h1>
        </lu-static-layout>
    `,
})
export class Page404 {
    private ui = inject(UiService);

    constructor() {
        this.ui.setPageTitle('Page not found');
    }
}
