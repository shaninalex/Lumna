import { Component, effect, inject, signal, Signal } from '@angular/core';
import { MainLayout } from '@core/layout';
import { UiService } from '@shared/ui';
import { Store } from '@ngrx/store';
import { selectWorkspaces} from '@entities/workspace';
import { AsyncPipe } from '@angular/common';
import { AppRoutes } from '@core';
import { RouterLink } from '@angular/router';
import { NgIcon, provideIcons } from '@ng-icons/core';
import { reiconPlus } from '@ng-icons/reicon';
import { ProjectModel, selectProjects } from '@entities/project';

@Component({
    selector: 'lu-workspace-entry-page',
    imports: [MainLayout, AsyncPipe, RouterLink, NgIcon],
    templateUrl: "./workspace-entry.page.html",
    providers: [provideIcons({reiconPlus})],
})
export class WorkspaceEntryPage {
    private store = inject(Store);
    private ui = inject(UiService);
    readonly appRoutes = inject(AppRoutes);

    workspace = this.store.selectSignal(selectWorkspaces.currentWorkspace)
    projects: Signal<ProjectModel[]> = signal<ProjectModel[]>([]);

    constructor() {
        effect(() => {
            const w = this.workspace();
            if (!w) return;
            this.ui.setPageTitle(`Workspace: ${w.title}`);
            this.projects = this.store.selectSignal(selectProjects.byWorkspaceId(w.id));
        });
    }
}
