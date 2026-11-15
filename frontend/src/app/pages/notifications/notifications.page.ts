import { Component } from '@angular/core';
import { MainLayout } from '@core/layout';

@Component({
    imports: [MainLayout],
    selector: 'lu-notifications',
    template: `
        <lu-main-layout>
            <p>notifications works!</p>
        </lu-main-layout>
    `,
})
export class NotificationsPage {
}
