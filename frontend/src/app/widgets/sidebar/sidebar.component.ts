import { Component, effect, inject, signal, Signal } from '@angular/core';
import { NgClass } from '@angular/common';
import { RouterLink } from '@angular/router';
import { Store } from '@ngrx/store';
import { selectWorkspaces } from '@entities/workspace';
import { ProjectModel, selectProjects } from '@entities/project';
import { selectUI } from "@core/store/ui";
import { AppRoutes } from '@core';
import {
    reiconInbox,
    reiconPage,
    reiconSettings2,
    reiconBell,
    reiconClockSquare,
    reiconPlug2,
    reiconNotebookSquare
} from '@ng-icons/reicon';
import { NgIcon, provideIcons } from '@ng-icons/core';

@Component({
    selector: 'lu-sidebar',
    imports: [NgClass, RouterLink, NgIcon],
    styleUrl: './sidebar.component.scss',
    templateUrl: './sidebar.component.html',
    providers: [provideIcons({
        reiconSettings2,
        reiconInbox,
        reiconPage,
        reiconBell,
        reiconClockSquare,
        reiconPlug2,
        reiconNotebookSquare
    })]
})
export class SidebarComponent {
    readonly appRoutes = inject(AppRoutes);
    projects: Signal<ProjectModel[]> = signal<ProjectModel[]>([])
    private store = inject(Store);
    hideSidebar = this.store.selectSignal(selectUI.sidebarOpen);
    workspaces = this.store.selectSignal(selectWorkspaces.all)
    currentWorkspace = this.store.selectSignal(selectWorkspaces.currentWorkspace)
    currentProject = this.store.selectSignal(selectProjects.currentProject);

    constructor() {
        effect(() => {
            const wp = this.currentWorkspace()
            if (!wp) {
                return
            }
            this.projects = this.store.selectSignal(selectProjects.byWorkspaceId(wp.id))
        });
    }
}
