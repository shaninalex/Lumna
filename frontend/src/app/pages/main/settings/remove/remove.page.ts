import { Component } from '@angular/core';
import { AccountRemoveFeature } from "@features";

@Component({
    imports: [AccountRemoveFeature],
    selector: 'lu-remove',
    template: `<lu-account-remove-feature />`,
})
export class RemovePage {
}
