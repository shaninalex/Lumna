import { Component, DestroyRef, inject, signal } from '@angular/core';
import { MatButton } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInput } from '@angular/material/input';
import { debounce, form, FormField, minLength, required, validate } from '@angular/forms/signals';
import { Store } from '@ngrx/store';
import { actionUser } from '@entities/user';
import { Actions, ofType } from '@ngrx/effects';
import { tap } from 'rxjs';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { SpinnerComponent } from '@shared/ui';

const defaultValue = {
    password: '',
    passwordConfirm: '',
}

@Component({
    imports: [FormField, SpinnerComponent, MatButton, MatFormFieldModule, MatInput],
    selector: 'lu-change-password-feature',
    template: `
        <form (submit)="onSubmit($event)">
            <div class="mb-4">
                <mat-form-field class="w-100 mb-3">
                    <mat-label>New password</mat-label>
                    <input matInput type="password" id="password" placeholder="***" [formField]="changePasswordForm.password">
                    @if (changePasswordForm().dirty() && changePasswordForm.password().invalid()) {
                        @for (error of changePasswordForm.password().errors(); track error.kind) {
                            <mat-error>{{ error.message }}</mat-error>
                        }
                    }
                </mat-form-field>
                <mat-form-field class="w-100 mb-3">
                    <mat-label>Password Confirm</mat-label>
                    <input matInput type="password" id="password-confirm" placeholder="***" [formField]="changePasswordForm.passwordConfirm">
                    @if (changePasswordForm().dirty() && changePasswordForm.passwordConfirm().invalid()) {
                        @for (error of changePasswordForm.passwordConfirm().errors(); track error.kind) {
                            <mat-error>{{ error.message }}</mat-error>
                        }
                    }
                </mat-form-field>

                <button matButton="filled" type="submit" [disabled]="!changePasswordForm().valid()">
                    Save
                    <lu-spinner [loading]="loading()"/>
                </button>
            </div>
        </form>
    `,

})
export class ChangePasswordFeature {
    private store = inject(Store);
    private readonly actions$ = inject(Actions);
    private destroyRef = inject(DestroyRef);
    loading = signal(false);

    changePassword = signal<{ password: string; passwordConfirm: string }>(defaultValue);
    changePasswordForm = form(this.changePassword, (schemaPath) => {
        debounce(schemaPath.password, 500);
        debounce(schemaPath.passwordConfirm, 500);
        required(schemaPath.password, { message: 'Password is required' });
        required(schemaPath.passwordConfirm, { message: 'Password confirmation is required' });
        minLength(schemaPath.password, 8, {message: 'Password must be at least 8 characters'});
        validate(schemaPath.passwordConfirm, ({value, valueOf}) => {
            const confirmPassword = value();
            const password = valueOf(schemaPath.password);
            if (confirmPassword !== password) {
                return {
                    kind: 'passwordMismatch',
                    message: 'Passwords do not match',
                };
            }
            return null;
        });
    });

    constructor() {
        this.actions$.pipe(
            ofType(actionUser.changePasswordSuccess, actionUser.changePasswordFailed),
            tap(() => {
                this.changePassword.set({...defaultValue})
                this.changePasswordForm().reset()
                this.loading.set(false)
            }),
            takeUntilDestroyed(this.destroyRef),
        ).subscribe()

        this.actions$.pipe(
            ofType(actionUser.changePassword),
            takeUntilDestroyed(this.destroyRef),
            tap(() => this.loading.set(true)),
        ).subscribe()
    }

    onSubmit(event: Event): void {
        event.preventDefault();
        if (this.changePasswordForm().errors().length) {
            return
        }
        this.store.dispatch(actionUser.changePassword({
            password: this.changePassword().password,
            passwordConfirm: this.changePassword().passwordConfirm,
        }))
    }
}
