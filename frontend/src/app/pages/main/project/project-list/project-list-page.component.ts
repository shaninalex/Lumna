import { Component, inject } from '@angular/core';
import { RouterLink } from "@angular/router";
import { UiService } from '@shared/ui';
import { ProjectCardComponent, selectProjects } from '@entities/project';
import { selectWorkspaces } from '@entities/workspace';
import { Store } from '@ngrx/store';
import { filter, switchMap, map } from 'rxjs';
import { AsyncPipe } from '@angular/common';
import { AppRoutes } from '@core';
import { GlobalLayout } from '@core/layout';

@Component({
    selector: 'lu-project-list-page',
    imports: [GlobalLayout, RouterLink, AsyncPipe, ProjectCardComponent],
    templateUrl: './project-list-page.component.html',
})
export class ProjectListPage {
    private ui = inject(UiService);
    private store = inject(Store);
    readonly appRoutes = inject(AppRoutes);

    workspace$ = this.store.select(selectWorkspaces.currentWorkspace).pipe(
        filter(workspace => workspace !== null),
        switchMap((workspace) =>
            this.store.select(selectProjects.byWorkspaceId(workspace.id)).pipe(
                map((projects) => ({workspace, projects}))
            )
        ),
    );

    constructor() {
        this.ui.setPageTitle("Projects")
    }
}
