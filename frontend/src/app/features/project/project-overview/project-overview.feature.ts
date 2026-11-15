import { Component, inject, signal, Signal } from '@angular/core';
import { Router, RouterLink } from '@angular/router';
import { AppRoutes } from '@core';
import { Store } from '@ngrx/store';
import { ProjectModel, selectProjects } from '@entities/project';
import { BoardCardComponent, BoardModel, selectBoard } from '@entities/board';
import { NgIcon, provideIcons } from '@ng-icons/core';
import { reiconChevronRight } from '@ng-icons/reicon';
import { selectWorkspaces } from '@entities/workspace';

@Component({
    imports: [RouterLink, BoardCardComponent, NgIcon, RouterLink],
    selector: 'lu-project-overview-feature',
    providers: [provideIcons({reiconChevronRight})],
    templateUrl: './project-overview.feature.html',
})
export class ProjectOverviewFeature {
    router = inject(Router);
    readonly appRoutes = inject(AppRoutes);
    private store = inject(Store);

    workspace =  this.store.selectSignal(selectWorkspaces.currentWorkspace)
    project = this.store.selectSignal(selectProjects.currentProject);
    boards: Signal<BoardModel[]> = signal<BoardModel[]>([]);

    workspaceLink: unknown[] = [];
    projectsLink: unknown[] = [];

    constructor() {
        const project = this.project()
        if (!project) { return }

        this.boards = this.store.selectSignal(selectBoard.byProjectId(project.id))
        const wp = this.workspace()
        if (!wp) { return }

        this.workspaceLink = this.appRoutes.workspaceDetail(wp.id)
        this.projectsLink = this.appRoutes.workspaceRoot(['projects'])
    }
}
