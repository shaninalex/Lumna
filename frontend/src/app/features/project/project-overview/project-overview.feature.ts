import { Component, inject } from '@angular/core';
import { Router, RouterLink } from '@angular/router';
import { AppRoutes } from '@core';
import { UiService } from '@shared/ui';
import { Store } from '@ngrx/store';
import { selectProjects } from '@entities/project';

@Component({
    imports: [RouterLink],
    selector: 'lu-project-overview-feature',
    templateUrl: './project-overview.feature.html',
})
export class ProjectOverviewFeature {
    router = inject(Router);
    readonly appRoutes = inject(AppRoutes);
    private store = inject(Store);
    project = this.store.selectSignal(selectProjects.currentProject);
}
