import { Component, DestroyRef, inject } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialog, MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatButton } from '@angular/material/button';
import { MatIcon } from '@angular/material/icon';
import { Store } from '@ngrx/store';
import { actionProject, selectProjects } from '@entities/project';
import { Actions, ofType } from '@ngrx/effects';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { Router } from '@angular/router';
import { AppRoutes } from '@core';

@Component({
    imports: [MatButton, MatIcon],
    selector: 'lu-project-remove-feature',
    template: `
        <div class="lu-banner lu-banner--danger" role="alert">
            <mat-icon>error</mat-icon>
            <div class="flex-grow-1">
                <h4 class="mt-0">Danger!</h4>
                <p>Warning text here...</p>
                <button matButton="filled" class="lu-danger" (click)="openConfirmationWindow()">Remove</button>
            </div>
        </div>
    `,
})
export class ProjectRemoveFeature {
    dialog = inject(MatDialog);
    private store = inject(Store);
    private actions = inject(Actions);
    private destroyRef = inject(DestroyRef);
    private projectId = this.store.selectSignal(selectProjects.currentProjectId)
    private router = inject(Router);
    private appRoutes = inject(AppRoutes);

    constructor() {
        this.actions.pipe(
            ofType(actionProject.deleteSuccessful),
            takeUntilDestroyed(this.destroyRef)
        ).subscribe(() => this.router.navigate(
            this.appRoutes.workspaceRoot(["projects"])
        ))
    }

    openConfirmationWindow() {
        const dialogRef = this.dialog.open<ProjectRemoveFeatureDialog, object, dialogData>(ProjectRemoveFeatureDialog, {
            width: '400px',
            data: {},
        });

        dialogRef.afterClosed().subscribe((data) => {
            const id = this.projectId()
            if (!data || !data.confirm || !id) {
                return;
            }
            this.store.dispatch(actionProject.delete({projectId: id}));
        });
    }
}

interface dialogData {
    confirm: boolean;
}

@Component({
    selector: 'lu-account-remove-feature-dialog',
    imports: [MatDialogModule, MatButton],
    template: `
        <h2 mat-dialog-title>Warning</h2>
        <mat-dialog-content>
            <p>Are you sure want to delete project?</p>
        </mat-dialog-content>
        <mat-dialog-actions align="end">
            <button type="button" matButton (click)="dialogRef.close({ confirm: false })">Cancel</button>
            <button type="button" matButton="filled" class="lu-danger" (click)="dialogRef.close({ confirm: true })">Yes</button>
        </mat-dialog-actions>
    `,
})
export class ProjectRemoveFeatureDialog {
    dialogRef = inject<MatDialogRef<ProjectRemoveFeatureDialog, dialogData>>(MatDialogRef);
    data = inject(MAT_DIALOG_DATA);
}
