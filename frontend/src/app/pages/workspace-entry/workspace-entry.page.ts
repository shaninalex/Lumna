import { Component, effect, inject, signal, Signal } from '@angular/core';
import { MainLayout } from '@core/layout';
import { UiService } from '@shared/ui';
import { Store } from '@ngrx/store';
import { selectWorkspaces } from '@entities/workspace';
import { AppRoutes } from '@core';
import { RouterLink } from '@angular/router';
import { MatButton } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatCheckbox } from '@angular/material/checkbox';
import { MatIcon } from '@angular/material/icon';
import { MatListModule } from '@angular/material/list';
import { ProjectModel, selectProjects } from '@entities/project';

@Component({
    selector: 'lu-workspace-entry-page',
    imports: [MainLayout, RouterLink, MatButton, MatCardModule, MatCheckbox, MatIcon, MatListModule],
    templateUrl: "./workspace-entry.page.html",
    styles: `
        .task-check {
            width: auto !important;
            height: auto !important;
            margin-block: 0 !important;
            align-self: center !important;
        }

        .projects-list {
            --mat-list-list-item-leading-avatar-shape: var(--mat-sys-corner-medium);
        }
    `,
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
