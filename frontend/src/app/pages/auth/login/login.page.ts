import { Component, inject } from '@angular/core';
import { UiService } from '@shared/ui';
import { AuthLoginFeature } from '@features';

@Component({
    selector: 'lu-login-page',
    imports: [AuthLoginFeature],

    templateUrl: './login.page.html',
})
export class LoginPage {
    private ui = inject(UiService);

    constructor() {
        this.ui.setPageTitle('Login');
    }
}
