import { Component, DestroyRef, inject, signal } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { form, FormField, required } from '@angular/forms/signals';
import { actionProject, type ProjectCreateModel } from '@entities/project';
import { selectWorkspaces } from '@entities/workspace';
import { Actions, ofType } from '@ngrx/effects';
import { Store } from '@ngrx/store';
import { filter, tap } from 'rxjs';

@Component({
    selector: 'lu-project-create-feature',
    imports: [FormField],

    template: `
        <form (submit)="onSubmit($event)">
            <div class="mb-4">
                <label for="project_title" class="form-label">Project title</label>
                <input
                    type="text"
                    class="form-control"
                    id="project_title"
                    [formField]="pForm.title"
                />
            </div>

            <div>
                <button class="btn btn-primary" type="submit">Create</button>
            </div>
        </form>
    `,
})
export class ProjectCreateFeature {
    private store = inject(Store);
    private actions$ = inject(Actions);
    private destroyRef = inject(DestroyRef);
    private workspace$ = this.store.select(selectWorkspaces.currentWorkspace);

    pFormModel = signal<ProjectCreateModel>({ title: '', workspaceId: 0 });
    pForm = form(this.pFormModel, (schemaPath) => required(schemaPath.title));

    constructor() {
        this.workspace$
            .pipe(
                takeUntilDestroyed(this.destroyRef),
                filter((wp) => wp !== null),
                tap((wp) => (this.pFormModel().workspaceId = wp.id)),
            )
            .subscribe();

        this.actions$
            .pipe(ofType(actionProject.createFailed), takeUntilDestroyed(this.destroyRef))
            .subscribe((action) => console.log(action));
    }

    onSubmit(event: Event): void {
        event.preventDefault();
        this.store.dispatch(actionProject.create({ payload: this.pFormModel() }));
    }
}
