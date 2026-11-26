import { Component, effect, inject } from '@angular/core';
import { UiService } from '@shared/ui';
import { MainLayout } from '@core/layout';
import { Router } from '@angular/router';
import { Store } from '@ngrx/store';
import { selectProjects } from '@entities/project';
import { ProjectOverviewFeature } from '@features';

@Component({
    selector: 'lu-inbox-page',
    imports: [MainLayout, ProjectOverviewFeature],
    templateUrl: "./project.page.html"
})
export class ProjectPage {
    router = inject(Router);
    private ui = inject(UiService);
    private store = inject(Store);
    project = this.store.selectSignal(selectProjects.currentProject);

    constructor() {
        effect(() => {
            const project = this.project();
            if (project) this.ui.setPageTitle(project.title);
        });
    }
}
