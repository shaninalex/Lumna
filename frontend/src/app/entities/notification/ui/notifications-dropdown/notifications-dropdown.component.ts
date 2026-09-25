import { Component } from '@angular/core';
import { CdkMenu, CdkMenuItem, CdkMenuTrigger } from '@angular/cdk/menu';

@Component({
    selector: 'lu-notifications-dropdown',
    imports: [CdkMenu, CdkMenuItem, CdkMenuTrigger],

    templateUrl: './notifications-dropdown.component.html',
})
export class NotificationsDropdownComponent {}
