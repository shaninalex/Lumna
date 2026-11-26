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

    workspaceLink = computed<unknown[]>(() => {
        const wp = this.workspace();
        return this.project() && wp ? this.appRoutes.workspaceDetail(wp.id) : [];
    });
    projectsLink = computed<unknown[]>(() =>
        this.project() && this.workspace() ? this.appRoutes.workspaceRoot(['projects']) : [],
    );
}
