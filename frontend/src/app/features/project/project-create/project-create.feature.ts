import { Component, DestroyRef, inject, signal } from '@angular/core';
import { MatButton } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInput } from '@angular/material/input';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { form, FormField, required } from '@angular/forms/signals';
import { actionProject, ProjectCreateModel, ProjectModel } from '@entities/project';
import { selectWorkspaces } from '@entities/workspace';
import { Actions, ofType } from '@ngrx/effects';
import { Store } from '@ngrx/store';
import { filter, map, tap } from 'rxjs';
import { SpinnerComponent } from '@shared/ui';
import { Router } from '@angular/router';
import { AppRoutes } from '@core';

@Component({
    selector: 'lu-project-create-feature',
    imports: [FormField, SpinnerComponent, MatButton, MatFormFieldModule, MatInput],
    template: `
        <form (submit)="onSubmit($event)">
            <mat-form-field class="w-100 mb-3">
                <mat-label>Project title</mat-label>
                <input matInput type="text" id="project_title" [formField]="pForm.title" />
            </mat-form-field>
            <div>
                <button matButton="filled" type="submit">
                    Create
                    <lu-spinner [loading]="loading()"/>
                </button>
            </div>
        </form>
    `,

})
export class ProjectCreateFeature {
    loading = signal(false);
    pFormModel = signal<ProjectCreateModel>({title: '', workspaceId: 0});
    pForm = form(this.pFormModel, (schemaPath) => required(schemaPath.title));
    private store = inject(Store);
    private actions$ = inject(Actions);
    private destroyRef = inject(DestroyRef);
    private workspace$ = this.store.select(selectWorkspaces.currentWorkspace);
    private router = inject(Router);
    private appRoutes = inject(AppRoutes);

    constructor() {
        this.workspace$
            .pipe(
                takeUntilDestroyed(this.destroyRef),
                filter((wp) => wp !== null),
                tap((wp) => (this.pFormModel().workspaceId = wp.id)),
            )
            .subscribe();

        this.actions$.pipe(
            ofType(actionProject.create),
            takeUntilDestroyed(this.destroyRef),
        ).subscribe(() => this.loading.set(true))

        this.actions$.pipe(
            ofType(actionProject.createSuccessful, actionProject.createFailed),
            takeUntilDestroyed(this.destroyRef),
        ).subscribe(() => this.loading.set(false))

        this.actions$.pipe(
            ofType(actionProject.createSuccessful),
            takeUntilDestroyed(this.destroyRef),
            map(project => project.project),
        ).subscribe((project: ProjectModel) => this.router.navigate(
            this.appRoutes.workspaceRoot(['p', project.id, 'inbox'])
        ));
    }

    onSubmit(event: Event): void {
        event.preventDefault();
        this.store.dispatch(actionProject.create({payload: this.pFormModel()}));
    }
}
