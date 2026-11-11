import { Component, inject } from '@angular/core';
import { UiService } from '@shared/ui';
import { MainLayout } from '@core/layout';
import { Router } from '@angular/router';
import { AppRoutes } from '@core';
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
    readonly appRoutes = inject(AppRoutes);
    private ui = inject(UiService);
    private store = inject(Store);
    project = this.store.selectSignal(selectProjects.currentProject);

    constructor() {
        const project = this.project()
        if (!project) { return }

        this.ui.setPageTitle(project.title);
    }
}
