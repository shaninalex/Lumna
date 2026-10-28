import { Component } from '@angular/core';
import { AccountFormFeature } from '@features';

@Component({
    imports: [AccountFormFeature],
    selector: 'lu-account-page',
    template: `<lu-account-form-feature />`,
})
export class AccountPage {}
