import { Component, inject } from '@angular/core';
import { MatButton } from '@angular/material/button';
import { MatDialog, MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatDivider } from '@angular/material/divider';
import { MatIcon } from '@angular/material/icon';
import { MatListModule } from '@angular/material/list';
import { MatMenuItem } from '@angular/material/menu';
import { Router, RouterLink } from '@angular/router';
import { Store } from '@ngrx/store';
import { selectWorkspaces } from '@entities/workspace/model';
import { AsyncPipe } from '@angular/common';
import { filter, map, switchMap } from 'rxjs';

@Component({
    selector: 'lu-workspace-switcher',
    imports: [AsyncPipe, RouterLink, MatDivider, MatIcon, MatMenuItem],

    template: `
        @if (workspace$ | async; as workspace) {
            <mat-divider />
            <button mat-menu-item type="button" class="switcher-item" (click)="openDialog()">
                <mat-icon>swap_horiz</mat-icon>
                <span>
                    <span class="d-block">{{ workspace.title }}</span>
                    <small class="d-block text-muted">Switch workspace</small>
                </span>
            </button>
        } @else if (workspaces$ | async; as workspaces) {
            <mat-divider />
            @if (workspaces.length > 0) {
                <button mat-menu-item type="button" (click)="openDialog()">
                    <mat-icon>workspaces</mat-icon>
                    <span>Select workspace</span>
                </button>
            } @else {
                <a mat-menu-item routerLink="/app/workspaces/create">
                    <mat-icon>add</mat-icon>
                    <span>Create Workspace</span>
                </a>
            }
        }
    `,
    styles: `
        .switcher-item {
            height: auto;
            min-height: 56px;
        }
    `,
})
export class WorkspaceSwitcherComponent {
    private store = inject(Store);
    dialog = inject(MatDialog);
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
    imports: [AsyncPipe, MatButton, MatDialogModule, MatListModule],

    template: `
        <h2 mat-dialog-title>Switch workspace</h2>
        <mat-dialog-content class="p-0">
            @if (data$ | async; as data) {
                <mat-action-list class="overflow-y-auto" style="max-height: 20rem">
                    @for (item of data.workspaces; track $index) {
                        <button
                            mat-list-item
                            (click)="handleLink(['/app/w', item.id.toString()])"
                            [activated]="data.currentWorkspace?.id === item.id"
                        >
                            <span matListItemAvatar class="lu-avatar lu-avatar--tertiary">{{ item.title.charAt(0) }}</span>
                            <span matListItemTitle>{{ item.title }}</span>
                        </button>
                    }
                </mat-action-list>
            }
        </mat-dialog-content>
        <mat-dialog-actions align="end">
            <button matButton type="button" (click)="handleLink(['/app/workspaces'])">See all</button>
        </mat-dialog-actions>
    `,
})
export class SwitchWorkspaceModal {
    private dialogRef = inject(MatDialogRef);
    private router = inject(Router);
    private store = inject(Store);

    data$ = this.store.select(selectWorkspaces.all).pipe(
        filter((workspace) => workspace !== null),
        switchMap((workspaces) =>
            this.store
                .select(selectWorkspaces.currentWorkspace)
                .pipe(map((currentWorkspace) => ({ workspaces, currentWorkspace }))),
        ),
    );

    handleLink(link: Array<string>) {
        this.dialogRef.close();
        this.router.navigate(link);
    }
}
