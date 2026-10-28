import { Component, inject, signal } from '@angular/core';
import { email, form, FormField, required } from '@angular/forms/signals';
import { Store } from '@ngrx/store';
import { actionUser, selectUser } from '@entities/user';

@Component({
    imports: [
        FormField,
    ],
    selector: 'lu-account-form-feature',
    template: `
        <form (submit)="onSubmit($event)">
            <div class="mb-4">
                <label for="email" class="form-label">Email</label>
                <input
                    class="form-control"
                    type="email"
                    id="email"
                    placeholder="example@email.com"
                    [class.is-invalid]="accountForm().dirty() && accountForm.email().invalid()"
                    [formField]="accountForm.email"
                />
                @if (accountForm().dirty() && accountForm.email().invalid()) {
                    @for (error of accountForm.email().errors(); track error.kind) {
                        <div class="invalid-feedback">{{ error.message }}</div>
                    }
                }
            </div>
            <div class="mb-4">
                <label for="fullName" class="form-label">Full Name</label>
                <input
                    class="form-control"
                    type="text"
                    id="fullName"
                    placeholder="Full Name"
                    [class.is-invalid]="accountForm().dirty() && accountForm.fullName().invalid()"
                    [formField]="accountForm.fullName"
                />
                @if (accountForm().dirty() && accountForm.fullName().invalid()) {
                    @for (error of accountForm.fullName().errors(); track error.kind) {
                        <div class="invalid-feedback">{{ error.message }}</div>
                    }
                }
            </div>

            <button class="btn btn-primary" type="submit" [disabled]="!accountForm().valid()">Save</button>
        </form>
    `,
})
export class AccountFormFeature {
    private store = inject(Store);

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
    }

    onSubmit(event: Event): void {
        event.preventDefault();
        if (!this.accountForm().valid()) {
            return;
        }
        this.store.dispatch(actionUser.update(this.accountFormModel()))
    }
}
