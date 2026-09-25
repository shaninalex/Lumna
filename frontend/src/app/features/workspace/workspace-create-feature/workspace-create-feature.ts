import { Component, DestroyRef, inject, signal, ChangeDetectionStrategy } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { email, form, FormField, required } from '@angular/forms/signals';
import { actionWorkspace, type WorkspaceCreateModel } from '@entities/workspace';
import { Actions, ofType } from '@ngrx/effects';
import { Store } from '@ngrx/store';
import { selectUser } from '@entities/user';
import { filter } from 'rxjs';

@Component({
    selector: 'lu-workspace-create-feature',
    imports: [FormField],
    changeDetection: ChangeDetectionStrategy.Eager,
    template: `
        <form (submit)="onSubmit($event)">
            <div class="mb-4">
                <label for="workspace_title" class="form-label">Workspaces title</label>
                <input
                    type="text"
                    class="form-control"
                    id="workspace_title"
                    [formField]="wspForm.title"
                />
            </div>

            <div class="mb-4">
                <label for="workspace_email" class="form-label">Workspaces email Owner</label>
                <input
                    type="email"
                    class="form-control"
                    id="workspace_email"
                    [formField]="wspForm.email"
                />
            </div>
            <div>
                <button class="btn btn-primary" type="submit">Create</button>
            </div>
        </form>
    `,
})
export class WorkspaceCreateFeature {
    private store = inject(Store);
    private actions$ = inject(Actions);
    private destroyRef = inject(DestroyRef);

    wspFormModel = signal<WorkspaceCreateModel>({
        title: '',
        email: '',
    });
    wspForm = form(this.wspFormModel, (schemaPath) => {
        required(schemaPath.title);
        required(schemaPath.email);
        email(schemaPath.email);
    });

    constructor() {
        this.store
            .select(selectUser.user)
            .pipe(
                takeUntilDestroyed(this.destroyRef),
                filter((user) => user !== null),
            )
            .subscribe((user) => {
                this.wspFormModel().email = user.email;
            });
        this.actions$
            .pipe(ofType(actionWorkspace.createFailed), takeUntilDestroyed(this.destroyRef))
            .subscribe((action) => console.log(action));
    }

    onSubmit(event: Event): void {
        event.preventDefault();
        this.store.dispatch(actionWorkspace.create({ data: this.wspFormModel() }));
    }
}
