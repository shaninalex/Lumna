import { Component, inject, ChangeDetectionStrategy } from '@angular/core';
import { UiService } from '@shared/ui';
import { AuthLoginFeature } from '@features';

@Component({
    selector: 'lu-login-page',
    imports: [AuthLoginFeature],
    changeDetection: ChangeDetectionStrategy.Eager,
    templateUrl: './login.page.html',
})
export class LoginPage {
    private ui = inject(UiService);

    constructor() {
        this.ui.setPageTitle('Login');
    }
}
