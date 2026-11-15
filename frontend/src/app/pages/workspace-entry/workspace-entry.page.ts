import { Component, effect, inject, signal, Signal } from '@angular/core';
import { MainLayout } from '@core/layout';
import { UiService } from '@shared/ui';
import { Store } from '@ngrx/store';
import { selectWorkspaces } from '@entities/workspace';
import { AppRoutes } from '@core';
import { RouterLink } from '@angular/router';
import { NgIcon, provideIcons } from '@ng-icons/core';
import { reiconPlus, reiconSettings2 } from '@ng-icons/reicon';
import { ProjectModel, selectProjects } from '@entities/project';

@Component({
    selector: 'lu-workspace-entry-page',
    imports: [MainLayout, RouterLink, NgIcon],
    templateUrl: "./workspace-entry.page.html",
    providers: [provideIcons({reiconPlus, reiconSettings2})],
})
export class WorkspaceEntryPage {
    readonly appRoutes = inject(AppRoutes);
    projects: Signal<ProjectModel[]> = signal<ProjectModel[]>([]);
    private store = inject(Store);
    workspace = this.store.selectSignal(selectWorkspaces.currentWorkspace)
    private ui = inject(UiService);

    constructor() {
        effect(() => {
            const w = this.workspace();
            if (!w) return;
            this.ui.setPageTitle(`Workspace: ${ w.title }`);
            this.projects = this.store.selectSignal(selectProjects.byWorkspaceId(w.id));
        });
    }
}
