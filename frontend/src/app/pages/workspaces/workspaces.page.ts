import type { OnInit } from '@angular/core';
import { Component, inject } from '@angular/core';
import { UiService } from '@shared/ui';
import { selectWorkspaces, WorkspaceCardComponent } from '@entities/workspace';
import { RouterLink } from '@angular/router';
import { Store } from '@ngrx/store';
import { AsyncPipe } from '@angular/common';
import { GlobalLayout } from '@core/layout';
import { MatButton } from '@angular/material/button';
import { MatIcon } from '@angular/material/icon';

@Component({
    selector: 'lu-workspaces-page',
    imports: [GlobalLayout, WorkspaceCardComponent, RouterLink, AsyncPipe, MatButton, MatIcon],

    templateUrl: 'workspaces.page.html',
})
export class WorkspacesPage implements OnInit {
    private ui = inject(UiService);
    private store = inject(Store);

    workspaces$ = this.store.select(selectWorkspaces.all);

    ngOnInit(): void {
        this.ui.setPageTitle('Workspaces');
    }
}
