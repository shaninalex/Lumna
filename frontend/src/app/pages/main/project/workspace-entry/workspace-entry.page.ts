import { Component, inject, ChangeDetectionStrategy } from '@angular/core';
import { MainLayout } from '@core/layout';
import { UiService } from '@shared/ui';
import { Store } from '@ngrx/store';
import { selectWorkspaces, WorkspaceModel } from '@entities/workspace';
import { Observable, tap } from 'rxjs';
import { AsyncPipe } from '@angular/common';

@Component({
    selector: 'lu-workspace-entry-page',
    imports: [MainLayout, AsyncPipe],
    changeDetection: ChangeDetectionStrategy.Eager,
    template: `
        <lu-main-layout>
            @if (workspace$ | async; as workspace) {
                <h1>{{ workspace.title }} workspace</h1>
            }
        </lu-main-layout>
    `,
})
export class WorkspaceEntryPage {
    private store = inject(Store);
    private ui = inject(UiService);

    workspace$: Observable<WorkspaceModel | null> = this.store
        .select(selectWorkspaces.currentWorkspace)
        .pipe(tap((workspace) => this.ui.setPageTitle(`Workspace: ${workspace?.title}`)));
}
