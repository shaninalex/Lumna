import { Component, inject, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { actionSession } from '@core/store/session.actions';
import { Store } from '@ngrx/store';
import { form, required, email, FormField } from '@angular/forms/signals';

interface LoginFormPayload {
    email: string;
    password: string;
}

@Component({
    selector: 'lu-auth-login-feature',
    imports: [FormsModule, FormField],

    templateUrl: './auth-login.component.html',
})
export class AuthLoginFeature {
    private store = inject(Store);

    loginFormModel = signal<LoginFormPayload>({
        email: '',
        password: '',
    });

    loginForm = form(this.loginFormModel, (schemaPath) => {
        required(schemaPath.email, { message: 'Email is required' });
        required(schemaPath.password, { message: 'Password is required' });
        email(schemaPath.email, { message: 'invalid email format' });
    });

    onSubmit(): void {
        if (!this.loginForm.email().errors().length && !this.loginForm.password().errors().length) {
            this.store.dispatch(
                actionSession.startAuthenticate({
                    email: this.loginForm.email().value(),
                    password: this.loginForm.password().value(),
                }),
            );
        }
    }
}
