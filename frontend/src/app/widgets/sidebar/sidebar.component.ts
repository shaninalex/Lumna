import { Component, effect, inject, signal, Signal } from '@angular/core';
import { RouterLink, RouterLinkActive } from '@angular/router';
import { Store } from '@ngrx/store';
import { selectWorkspaces } from '@entities/workspace';
import { ProjectModel, selectProjects } from '@entities/project';
import { selectUI } from "@core/store/ui";
import { AppRoutes } from '@core';
import { MatListModule } from '@angular/material/list';
import { MatIcon } from '@angular/material/icon';
import { MatDivider } from '@angular/material/divider';
import { MatTooltip } from '@angular/material/tooltip';

@Component({
    selector: 'lu-sidebar',
    imports: [RouterLink, RouterLinkActive, MatListModule, MatIcon, MatDivider, MatTooltip],
    styleUrl: './sidebar.component.scss',
    templateUrl: './sidebar.component.html',
})
export class SidebarComponent {
    readonly appRoutes = inject(AppRoutes);
    readonly mainLinks = [
        { link: '/app/overview', label: 'Overview', icon: 'inbox' },
        { link: '/app/recent', label: 'Recent', icon: 'schedule' },
        { link: '/app/notifications', label: 'Notifications', icon: 'notifications' },
        { link: '/app/plugins', label: 'Plugins', icon: 'extension' },
    ];
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
