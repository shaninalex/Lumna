import { Component, computed, inject } from '@angular/core';
import { Router, RouterLink } from '@angular/router';
import { AppRoutes } from '@core';
import { Store } from '@ngrx/store';
import { selectProjects } from '@entities/project';
import { BoardCardComponent, selectBoard } from '@entities/board';
import { MatButton, MatIconButton } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatIcon } from '@angular/material/icon';
import { MatListModule } from '@angular/material/list';
import { MatProgressBar } from '@angular/material/progress-bar';
import { MatTabsModule } from '@angular/material/tabs';
import { selectWorkspaces } from '@entities/workspace';

@Component({
    imports: [
        RouterLink,
        BoardCardComponent,
        MatButton,
        MatIconButton,
        MatCardModule,
        MatIcon,
        MatListModule,
        MatProgressBar,
        MatTabsModule,
    ],
    selector: 'lu-project-overview-feature',
    templateUrl: './project-overview.feature.html',
})
export class ProjectOverviewFeature {
    router = inject(Router);
    readonly appRoutes = inject(AppRoutes);
    private store = inject(Store);

    workspace = this.store.selectSignal(selectWorkspaces.currentWorkspace);
    project = this.store.selectSignal(selectProjects.currentProject);
    boards = computed(() => {
        const project = this.project();

        if (!project) {
            return [];
        }

        return this.store.selectSignal(selectBoard.byProjectId(project.id))();
    });

    workspaceLink: unknown[] = [];
    projectsLink: unknown[] = [];

    constructor() {
        const project = this.project();
        if (!project) {
            return;
        }

        const wp = this.workspace();
        if (!wp) {
            return;
        }

        this.workspaceLink = this.appRoutes.workspaceDetail(wp.id);
        this.projectsLink = this.appRoutes.workspaceRoot(['projects']);
    }
}
