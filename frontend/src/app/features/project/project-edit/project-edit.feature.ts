import { Component, DestroyRef, effect, inject, signal } from '@angular/core';
import { MatButton } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInput } from '@angular/material/input';
import { Store } from '@ngrx/store';
import { actionProject, selectProjects } from '@entities/project';
import { form, FormField, required } from '@angular/forms/signals';
import { SpinnerComponent } from '@shared/ui';
import { Actions, ofType } from '@ngrx/effects';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { tap } from 'rxjs';

@Component({
    imports: [FormField, SpinnerComponent, MatButton, MatFormFieldModule, MatInput],
    selector: 'lu-project-edit-feature',
    template: `
        <form (submit)="onSubmit($event)">
            <mat-form-field class="w-100 mb-3">
                <mat-label>Project name</mat-label>
                <input matInput type="text" id="email" [formField]="projectForm.projectName" />
                @if (projectForm().dirty() && projectForm.projectName().invalid()) {
                    @for (error of projectForm.projectName().errors(); track error.kind) {
                        <mat-error>{{ error.message }}</mat-error>
                    }
                }
            </mat-form-field>

            <button matButton="filled" type="submit" [disabled]="!projectForm().valid()">
                Save
                <lu-spinner [loading]="loading()"/>
            </button>
        </form>
    `,

})
export class ProjectEditFeature {
    loading = signal(false)
    projectFormModel = signal<{ projectName: string }>({projectName: ''})
    projectForm = form(this.projectFormModel, (schema) => {
        required(schema.projectName, {message: 'Title is required'})
    })
    private store = inject(Store);
    private project = this.store.selectSignal(selectProjects.currentProject);
    private readonly actions$ = inject(Actions);
    private destroyRef = inject(DestroyRef);

    constructor() {
        effect(() => {
            const p = this.project()
            if (p) {
                this.projectFormModel.set({projectName: p.title})
            }
        })

        this.actions$.pipe(
            ofType(actionProject.patch),
            takeUntilDestroyed(this.destroyRef),
            tap(() => this.loading.set(true)),
        ).subscribe()

        this.actions$.pipe(
            ofType(actionProject.patchSuccessful, actionProject.patchFailed),
            takeUntilDestroyed(this.destroyRef),
            tap(() => this.loading.set(false)),
        ).subscribe()
    }

    onSubmit(event: Event): void {
        event.preventDefault();
        const p = this.project()
        if (!this.projectForm().valid() || !p) {
            return;
        }

        // TODO: we do not need workspace id for updating the project
        this.store.dispatch(actionProject.patch({
            id: p.id,
            data: {
                title: this.projectFormModel().projectName,
                workspaceId: p.workspaceId,
            }
        }))
    }
}

