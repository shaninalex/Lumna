import { Component, DestroyRef, effect, inject, signal } from '@angular/core';
import { Store } from '@ngrx/store';
import { actionProject, selectProjects } from '@entities/project';
import { form, FormField, required } from '@angular/forms/signals';
import { SpinnerComponent } from '@shared/ui';
import { Actions, ofType } from '@ngrx/effects';
import { actionUser } from '@entities/user';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { tap } from 'rxjs';

@Component({
    imports: [FormField, SpinnerComponent],
    selector: 'lu-project-edit-feature',
    template: `
        <form (submit)="onSubmit($event)">
            <div class="mb-4">
                <label for="email" class="form-label">Project name</label>
                <input
                    class="form-control"
                    type="text"
                    id="email"
                    placeholder="example@email.com"
                    [class.is-invalid]="projectForm().dirty() && projectForm.projectName().invalid()"
                    [formField]="projectForm.projectName"
                />
                @if (projectForm().dirty() && projectForm.projectName().invalid()) {
                    @for (error of projectForm.projectName().errors(); track error.kind) {
                        <div class="invalid-feedback">{{ error.message }}</div>
                    }
                }
            </div>

            <button class="btn btn-primary" type="submit" [disabled]="!projectForm().valid()">
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

