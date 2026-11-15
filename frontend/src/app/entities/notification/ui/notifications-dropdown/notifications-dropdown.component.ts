import { Component } from '@angular/core';
import { CdkMenu, CdkMenuItem, CdkMenuTrigger } from '@angular/cdk/menu';
import { NgIcon, provideIcons } from '@ng-icons/core';
import { reiconBell, reiconPlus, reiconCommentDots} from '@ng-icons/reicon';
import { reiconArrowToTopLeftFill, reiconAtFill } from '@ng-icons/reicon/fill';
import { RouterLink } from '@angular/router';


@Component({
    selector: 'lu-notifications-dropdown',
    imports: [CdkMenu, CdkMenuItem, CdkMenuTrigger, NgIcon, RouterLink],
    providers: [provideIcons({ reiconBell, reiconPlus, reiconArrowToTopLeftFill, reiconCommentDots, reiconAtFill })],
    templateUrl: './notifications-dropdown.component.html',
})
export class NotificationsDropdownComponent {}
