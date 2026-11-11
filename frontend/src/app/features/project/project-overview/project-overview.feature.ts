import { Component, inject, signal, Signal } from '@angular/core';
import { Router, RouterLink } from '@angular/router';
import { AppRoutes } from '@core';
import { Store } from '@ngrx/store';
import { selectProjects } from '@entities/project';
import { BoardCardComponent, BoardModel, selectBoard } from '@entities/board';

@Component({
    imports: [RouterLink, BoardCardComponent],
    selector: 'lu-project-overview-feature',
    templateUrl: './project-overview.feature.html',
})
export class ProjectOverviewFeature {
    router = inject(Router);
    readonly appRoutes = inject(AppRoutes);
    private store = inject(Store);
    project = this.store.selectSignal(selectProjects.currentProject);
    boards: Signal<BoardModel[]> = signal<BoardModel[]>([]);

    constructor() {
        const project = this.project()
        if (!project) { return }

        this.boards = this.store.selectSignal(selectBoard.byProjectId(project.id))
    }
}
