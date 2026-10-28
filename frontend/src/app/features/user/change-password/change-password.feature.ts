import { Component, inject, signal } from '@angular/core';
import { debounce, form, FormField, minLength, required, validate } from '@angular/forms/signals';
import { Store } from '@ngrx/store';
import { actionUser } from '@entities/user';
import { Actions, ofType } from '@ngrx/effects';
import { tap } from 'rxjs';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';

const defaultValue = {
    password: '',
    passwordConfirm: '',
}

@Component({
    imports: [FormField],
    selector: 'lu-change-password-feature',
    template: `
        <form (submit)="onSubmit($event)">
            <div class="mb-4">
                <div class="mb-3">
                    <label for="password" class="form-label">New password</label>
                    <input type="password"
                           class="form-control"
                           id="password"
                           placeholder="***"
                           [class.is-invalid]="changePasswordForm().dirty() && changePasswordForm.password().invalid()"
                           [formField]="changePasswordForm.password">
                    @if (changePasswordForm().dirty() && changePasswordForm.password().invalid()) {
                        @for (error of changePasswordForm.password().errors(); track error.kind) {
                            <div class="invalid-feedback">{{ error.message }}</div>
                        }
                    }
                </div>
                <div class="mb-3">
                    <label for="password-confirm" class="form-label">Password Confirm</label>
                    <input type="password"
                           class="form-control"
                           id="password-confirm"
                           placeholder="***"
                           [class.is-invalid]="changePasswordForm().dirty() && changePasswordForm.passwordConfirm().invalid()"
                           [formField]="changePasswordForm.passwordConfirm">
                    @if (changePasswordForm().dirty() && changePasswordForm.passwordConfirm().invalid()) {
                        @for (error of changePasswordForm.passwordConfirm().errors(); track error.kind) {
                            <div class="invalid-feedback">{{ error.message }}</div>
                        }
                    }
                </div>

                <button type="submit" class="btn btn-primary" [disabled]="!changePasswordForm().valid()">Save</button>
            </div>
        </form>
    `,
})
export class ChangePasswordFeature {
    private store = inject(Store);
    private actions$ = inject(Actions);

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
            }),
            takeUntilDestroyed(),
        ).subscribe()
    }

    onSubmit(event: Event): void {
        event.preventDefault();
        if (this.changePasswordForm().errors().length) {
            console.log(this.changePasswordForm().errors())
            return
        }
        this.store.dispatch(actionUser.changePassword({
            password: this.changePassword().password,
            passwordConfirm: this.changePassword().passwordConfirm,
        }))
    }
}
