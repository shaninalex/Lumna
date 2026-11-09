import { Component, inject } from '@angular/core';
import { UiService } from '@shared/ui';
import { MainLayout } from '@core/layout';
import { Router, RouterLink } from '@angular/router';
import { AppRoutes } from '@core';
import { Store } from '@ngrx/store';
import { selectProjects } from '@entities/project';

@Component({
    selector: 'lu-inbox-page',
    imports: [MainLayout, RouterLink],
    templateUrl: "./index.page.html"
})
export class InboxPage {
    router = inject(Router);
    readonly appRoutes = inject(AppRoutes);
    private ui = inject(UiService);
    private store = inject(Store);
    project = this.store.selectSignal(selectProjects.currentProject);

    constructor() {
        this.ui.setPageTitle('Inbox');
    }
}
