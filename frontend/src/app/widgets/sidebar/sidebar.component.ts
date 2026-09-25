import { Component, inject } from '@angular/core';
import { AsyncPipe, NgClass } from '@angular/common';
import { RouterLink } from '@angular/router';
import { Store } from '@ngrx/store';
import { selectWorkspaces } from '@entities/workspace';
import { filter, map, switchMap } from 'rxjs';
import { selectProjects } from '@entities/project';
import { selectUI } from "@core/store/ui";

@Component({
    selector: 'lu-sidebar',
    imports: [NgClass, RouterLink, AsyncPipe],
    styleUrl: './sidebar.component.css',
    templateUrl: './sidebar.component.html',
})
export class SidebarComponent {
    private store = inject(Store);

    hideSidebar = this.store.selectSignal(selectUI.sidebarOpen);
    currentProject = this.store.selectSignal(selectProjects.currentProject);

    workspace$ = this.store.select(selectWorkspaces.currentWorkspaceId).pipe(
        filter((workspaceId) => workspaceId !== null),
        switchMap((workspaceId) =>
            this.store
                .select(selectProjects.byWorkspaceId(workspaceId))
                .pipe(map((projects) => ({workspaceId, projects}))),
        ),
    );

}
