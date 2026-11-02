import { Component } from '@angular/core';
import { ChangePasswordFeature } from '@features';

@Component({
    imports: [ChangePasswordFeature],
    selector: 'lu-security-page',
    template: `<lu-change-password-feature/>`,
})
export class SecurityPage {}
