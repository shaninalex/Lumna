import { Component } from '@angular/core';
import { CdkMenu, CdkMenuItem, CdkMenuTrigger } from '@angular/cdk/menu';
import { NgIcon, provideIcons } from '@ng-icons/core';
import { reiconBell, reiconPlus, reiconCommentDots} from '@ng-icons/reicon';
import { reiconArrowToTopLeftFill } from '@ng-icons/reicon/fill';

@Component({
    selector: 'lu-notifications-dropdown',
    imports: [CdkMenu, CdkMenuItem, CdkMenuTrigger, NgIcon],
    providers: [provideIcons({ reiconBell, reiconPlus, reiconArrowToTopLeftFill, reiconCommentDots })],
    templateUrl: './notifications-dropdown.component.html',
})
export class NotificationsDropdownComponent {}
