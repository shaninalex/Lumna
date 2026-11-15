import { Component } from '@angular/core';
import { MainLayout } from '@core/layout';

@Component({
    imports: [MainLayout],
    selector: 'lu-recent',
    template: `
        <lu-main-layout>
            <p>recent works!</p>
        </lu-main-layout>
    `,
})
export class RecentPage {
}
