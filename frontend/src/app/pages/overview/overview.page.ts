import { Component } from '@angular/core';
import { MainLayout } from '@core/layout';

@Component({
    imports: [MainLayout],
    selector: 'lu-overview',
    template: `
        <lu-main-layout>
            <p>overview works!</p>
        </lu-main-layout>
    `,
})
export class OverviewPage {
}
