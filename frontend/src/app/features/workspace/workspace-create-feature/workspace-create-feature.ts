import { Component, DestroyRef, inject, signal } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { email, form, FormField, required } from '@angular/forms/signals';
import { actionWorkspace, type WorkspaceCreateModel } from '@entities/workspace';
import { Store } from '@ngrx/store';
import { MatButton } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInput } from '@angular/material/input';
import { selectUser } from '@entities/user';
import { filter } from 'rxjs';

@Component({
    selector: 'lu-workspace-create-feature',
    imports: [FormField, MatButton, MatFormFieldModule, MatInput],

    template: `
        <form (submit)="onSubmit($event)">
            <mat-form-field class="w-100 mb-3">
                <mat-label>Workspaces title</mat-label>
                <input matInput type="text" id="workspace_title" [formField]="wspForm.title" />
            </mat-form-field>

            <mat-form-field class="w-100 mb-3">
                <mat-label>Workspaces email Owner</mat-label>
                <input matInput type="email" id="workspace_email" [formField]="wspForm.email" />
            </mat-form-field>
            <div>
                <button matButton="filled" type="submit">Create</button>
            </div>
        </form>
    `,

})
export class WorkspaceCreateFeature {
    private store = inject(Store);
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
    }

    onSubmit(event: Event): void {
        event.preventDefault();
        this.store.dispatch(actionWorkspace.create({ data: this.wspFormModel() }));
    }
}
