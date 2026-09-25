import { Component, ChangeDetectionStrategy } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { AuthLayout } from '@core/layout';

@Component({
    selector: 'lu-auth-root',
    imports: [RouterOutlet, AuthLayout],
    changeDetection: ChangeDetectionStrategy.Eager,
    template: `
        <lu-auth-layout>
            <router-outlet />
        </lu-auth-layout>
    `,
})
export class AuthRoot {}
