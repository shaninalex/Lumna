import { Component, ChangeDetectionStrategy } from '@angular/core';
import { CdkMenu, CdkMenuItem, CdkMenuTrigger } from '@angular/cdk/menu';

@Component({
    selector: 'lu-notifications-dropdown',
    imports: [CdkMenu, CdkMenuItem, CdkMenuTrigger],
    changeDetection: ChangeDetectionStrategy.Eager,
    templateUrl: './notifications-dropdown.component.html',
})
export class NotificationsDropdownComponent {}
