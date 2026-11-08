import { Component, inject } from '@angular/core';
import { UiService } from '@shared/ui';
import { MainLayout } from '@core/layout';
import { Router, RouterLink } from '@angular/router';
import { AppRoutes } from '@core';

@Component({
    selector: 'lu-inbox-page',
    imports: [MainLayout, RouterLink],
    template: `
        <lu-main-layout>
            <div class="container-fluid py-4">
                <h4>Few latest tasks</h4>
                <h4>Latest comments</h4>
                <h4>Latest assignments, attachments and other activity</h4>
                <hr />

                <a [routerLink]="[...appRoutes.projectRoute(), 'settings']" class="btn btn-primary">Settings</a>
            </div>
        </lu-main-layout>
    `,
})
export class InboxPage {
    private ui = inject(UiService);
    router = inject(Router);
    readonly appRoutes = inject(AppRoutes);

    constructor() {
        this.ui.setPageTitle('Inbox');
    }
}
