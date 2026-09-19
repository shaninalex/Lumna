import { Component, inject } from '@angular/core';
import { Dialog, DialogRef } from '@angular/cdk/dialog';
import { Router, RouterLink } from "@angular/router";
import { Store } from '@ngrx/store';
import { selectWorkspaces } from '@entities/workspace/model';
import { AsyncPipe, NgClass } from '@angular/common';
import { filter, map, switchMap } from 'rxjs';


@Component({
    selector: 'lu-workspace-switcher',
    imports: [AsyncPipe, RouterLink],
    template: `
        @if (workspace$ | async; as workspace) {
            <button
                class="btn btn-sm btn-outline-secondary d-block w-100 text-left bg-body text-body"
                (click)="openDialog()"
            >
                <span class="d-block text-start">
                    {{ workspace.title }}
                </span>
                <small class="d-block text-start" style="font-size: 0.7rem">
                    Switch workspace
                </small>
            </button>
        } @else if (workspaces$ | async; as workspaces) {
            @if (workspaces.length > 0) {
                <button
                    class="btn btn-sm btn-outline-secondary d-block w-100 text-left bg-body text-body"
                    (click)="openDialog()"
                >
                    <span class="d-block text-start">
                        Select workspace
                    </span>
                </button>
            } @else {
                <a 
                    class="btn btn-sm btn-outline-secondary d-block bg-body text-body"
                    routerLink="/app/workspaces/create">
                    Create Workspace
                </a>
            }
        }
    `,
})
export class WorkspaceSwitcherComponent {
    private store = inject(Store);
    dialog = inject(Dialog);
    workspace$ = this.store.select(selectWorkspaces.currentWorkspace);
    workspaces$ = this.store.select(selectWorkspaces.all);

    openDialog() {
        this.dialog.open(SwitchWorkspaceModal, {
            minWidth: '300px',
        });
    }
}

@Component({
    selector: 'lu-switch-workspace-modal',
    imports: [AsyncPipe, NgClass],
    template: `
        <div class="card">
            @if (data$ | async; as data) {
                <div class="list-group overflow-y-auto" style="max-height: 20rem">
                    @for (item of data.workspaces; track $index) {
                        <button 
                            (click)="handleLink(['/app/w', item.id.toString()])"
                            [ngClass]="{'active': data.currentWorkspace!.id === item.id}"
                            class="list-group-item list-group-item-action">
                        {{ item.title }}
                    </button>
                    }
                </div>
            }
            <button (click)="handleLink(['/app/workspaces'])" class="btn btn-link">See all</button>
        </div>
        
    `,
})
export class SwitchWorkspaceModal {
    private dialogRef = inject(DialogRef);
    private router = inject(Router);
    private store = inject(Store);

    data$ = this.store.select(selectWorkspaces.all).pipe(
        filter(workspace => workspace !== null),
        switchMap((workspaces) => 
            this.store.select(selectWorkspaces.currentWorkspace).pipe(
                map((currentWorkspace) => ({workspaces, currentWorkspace}))
            )
        ),
    );

    handleLink(link: Array<string>) {
        this.dialogRef.close();
        this.router.navigate(link);
    }
}