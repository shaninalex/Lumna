import type { OnInit } from '@angular/core';
import { Component, inject } from '@angular/core';
import { UiService } from '@shared/ui';
import { MainLayout } from '@core/layout';

@Component({
    selector: 'lu-inbox-page',
    imports: [MainLayout],
    template: `
        <lu-main-layout>
            <div class="container-fluid py-4">
                <h4>Few latest tasks</h4>
                <h4>Latest comments</h4>
                <h4>Latest assignments, attachments and other activity</h4>
            </div>
        </lu-main-layout>
    `,
})
export class InboxPage implements OnInit {
    private ui = inject(UiService);

    ngOnInit(): void {
        this.ui.setPageTitle('Inbox');
    }
}
