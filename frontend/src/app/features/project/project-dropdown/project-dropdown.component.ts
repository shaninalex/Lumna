import { AsyncPipe } from '@angular/common';
import { RouterLink } from '@angular/router';
import { Component, inject } from '@angular/core';
import { CdkMenu, CdkMenuItem, CdkMenuTrigger } from '@angular/cdk/menu';
import { Store } from '@ngrx/store';
import { filter, switchMap, map } from 'rxjs';
import { selectWorkspaces } from '@entities/workspace';
import { ProjectListItemComponent, selectProjects } from '@entities/project';

@Component({
    selector: 'lu-project-dropdown',
    imports: [
        CdkMenu,
        CdkMenuItem,
        CdkMenuTrigger,
        AsyncPipe,
        ProjectListItemComponent,
        RouterLink,
    ],

    templateUrl: './project-dropdown.component.html',
})
export class ProjectDropdownComponent {
    private store = inject(Store);

    open = false;
    project$ = this.store.select(selectProjects.currentProject);
    workspace$ = this.store.select(selectWorkspaces.currentWorkspace).pipe(
        filter((workspace) => workspace !== null),
        switchMap((workspace) =>
            this.store
                .select(selectProjects.byWorkspaceId(workspace.id))
                .pipe(map((projects) => ({ workspace, projects }))),
        ),
    );
}
