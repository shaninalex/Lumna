import { Component, DestroyRef, inject } from '@angular/core';
import { Dialog, DIALOG_DATA, DialogRef } from '@angular/cdk/dialog';
import { Store } from '@ngrx/store';
import { actionProject, selectProjects } from '@entities/project';
import { Actions, ofType } from '@ngrx/effects';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { Router } from '@angular/router';
import { AppRoutes } from '@core';

@Component({
    imports: [],
    selector: 'lu-project-remove-feature',
    template: `
        <div class="alert alert-danger">
            <h4>Danger!</h4>
            <p>Warning text here...</p>
            <hr>
            <button class="btn btn-danger" (click)="openConfirmationWindow()">Remove</button>
        </div>
    `,
})
export class ProjectRemoveFeature {
    dialog = inject(Dialog);
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
        const dialogRef = this.dialog.open<dialogData>(ProjectRemoveFeatureDialog, {
            width: '250px',
            data: {},
        });

        dialogRef.closed.subscribe((data) => {
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
    template: `
        <div class="modal d-block">
            <div class="modal-dialog">
                <div class="modal-content">
                    <div class="modal-header">
                        <h5 class="modal-title">Warning</h5>
                        <button type="button" class="btn-close" (click)="dialogRef.close({ confirm: false })"></button>
                    </div>
                    <div class="modal-body">
                        <p>Are you sure want to delete project?</p>
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-danger" (click)="dialogRef.close({ confirm: true })">
                            Yes
                        </button>
                        <button type="button" class="btn btn-secondary" (click)="dialogRef.close({ confirm: false })">
                            Cancel
                        </button>
                    </div>
                </div>
            </div>
        </div>
    `,
})
export class ProjectRemoveFeatureDialog {
    dialogRef = inject<DialogRef<dialogData>>(DialogRef<dialogData>);
    data = inject(DIALOG_DATA);
}
