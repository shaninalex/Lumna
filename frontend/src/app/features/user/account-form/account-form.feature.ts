import { Component, DestroyRef, inject, signal } from '@angular/core';
import { MatButton } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInput } from '@angular/material/input';
import { email, form, FormField, required } from '@angular/forms/signals';
import { Store } from '@ngrx/store';
import { actionUser, selectUser } from '@entities/user';
import { SpinnerComponent } from '@shared/ui';
import { Actions, ofType } from '@ngrx/effects';
import { tap } from 'rxjs';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';

@Component({
    imports: [
        FormField,
        SpinnerComponent,
        MatButton,
        MatFormFieldModule,
        MatInput,
    ],
    selector: 'lu-account-form-feature',
    template: `
        <form (submit)="onSubmit($event)">
            <mat-form-field class="w-100 mb-3">
                <mat-label>Email</mat-label>
                <input matInput type="email" id="email" placeholder="example@email.com" [formField]="accountForm.email" />
                @if (accountForm().dirty() && accountForm.email().invalid()) {
                    @for (error of accountForm.email().errors(); track error.kind) {
                        <mat-error>{{ error.message }}</mat-error>
                    }
                }
            </mat-form-field>
            <mat-form-field class="w-100 mb-3">
                <mat-label>Full Name</mat-label>
                <input matInput type="text" id="fullName" placeholder="Full Name" [formField]="accountForm.fullName" />
                @if (accountForm().dirty() && accountForm.fullName().invalid()) {
                    @for (error of accountForm.fullName().errors(); track error.kind) {
                        <mat-error>{{ error.message }}</mat-error>
                    }
                }
            </mat-form-field>

            <button matButton="filled" type="submit" [disabled]="!accountForm().valid()">
                Save
                <lu-spinner [loading]="loading()"/>
            </button>
        </form>
    `,

})
export class AccountFormFeature {
    private store = inject(Store);
    private readonly actions$ = inject(Actions);
    private destroyRef = inject(DestroyRef);

    loading = signal(false);
    accountFormModel = signal<{ email: string; fullName: string }>({
        email: '',
        fullName: '',
    });
    accountForm = form(this.accountFormModel, (schemaPath) => {
        required(schemaPath.email, {message: 'Email is required'});
        email(schemaPath.email, {message: 'Invalid email format'});
        required(schemaPath.fullName, {message: 'Full name is required'});
    });

    constructor() {
        const user = this.store.selectSignal(selectUser.user)()
        if (!user) return

        this.accountFormModel.set({email: user.email, fullName: user.fullName})

        this.actions$.pipe(
            ofType(actionUser.update),
            takeUntilDestroyed(this.destroyRef),
            tap(() => this.loading.set(true)),
        ).subscribe()

        this.actions$.pipe(
            ofType(actionUser.updateSuccess, actionUser.updateFailed),
            takeUntilDestroyed(this.destroyRef),
            tap(() => this.loading.set(false)),
        ).subscribe()
    }

    onSubmit(event: Event): void {
        event.preventDefault();
        if (!this.accountForm().valid()) {
            return;
        }
        this.store.dispatch(actionUser.update(this.accountFormModel()))
    }
}
