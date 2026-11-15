import { Component, effect, inject, signal, Signal } from '@angular/core';
import { NgClass } from '@angular/common';
import { RouterLink } from '@angular/router';
import { Store } from '@ngrx/store';
import { selectWorkspaces, WorkspaceModel } from '@entities/workspace';
import { ProjectModel, selectProjects } from '@entities/project';
import { selectUI } from "@core/store/ui";
import { AppRoutes } from '@core';
import { reiconInbox, reiconPage, reiconSettings2, reiconBell, reiconClockSquare } from '@ng-icons/reicon';
import { NgIcon, provideIcons } from '@ng-icons/core';

@Component({
    selector: 'lu-sidebar',
    imports: [NgClass, RouterLink, NgIcon],
    styleUrl: './sidebar.component.scss',
    templateUrl: './sidebar.component.html',
    providers: [provideIcons({reiconSettings2, reiconInbox, reiconPage, reiconBell, reiconClockSquare})]
})
export class SidebarComponent {
    readonly appRoutes = inject(AppRoutes);
    private store = inject(Store);

    hideSidebar = this.store.selectSignal(selectUI.sidebarOpen);
    workspaces = this.store.selectSignal(selectWorkspaces.all)
    projects: Signal<ProjectModel[]> = signal<ProjectModel[]>([])

    currentWorkspace = this.store.selectSignal(selectWorkspaces.currentWorkspace)
    currentProject = this.store.selectSignal(selectProjects.currentProject);

    constructor() {
        effect(() => {
            const wp = this.currentWorkspace()
            if (!wp) { return }
            this.projects = this.store.selectSignal(selectProjects.byWorkspaceId(wp.id))
        });
    }
}
