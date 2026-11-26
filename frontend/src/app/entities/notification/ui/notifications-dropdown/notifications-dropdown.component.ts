import { Component, ViewEncapsulation } from '@angular/core';
import { RouterLink } from '@angular/router';
import { MatBadge } from '@angular/material/badge';
import { MatButton, MatIconButton } from '@angular/material/button';
import { MatDivider } from '@angular/material/divider';
import { MatIcon } from '@angular/material/icon';
import { MatMenuModule } from '@angular/material/menu';


@Component({
    selector: 'lu-notifications-dropdown',
    imports: [RouterLink, MatBadge, MatButton, MatIconButton, MatDivider, MatIcon, MatMenuModule],
    templateUrl: './notifications-dropdown.component.html',
    // The menu panel is rendered in an overlay, outside this component's host.
    encapsulation: ViewEncapsulation.None,
    styles: `
        .mat-mdc-menu-panel.lu-notifications-menu {
            width: 380px;
            max-width: 380px;

            .notification {
                height: auto;
                white-space: normal;

                .mat-mdc-menu-item-text {
                    width: 100%;
                }
            }

            .notification--unread {
                background-color: var(--mat-sys-surface-container-high);
            }

            .lu-avatar .mat-icon {
                margin: 0;
                color: inherit;
            }

            .unread-dot {
                flex-shrink: 0;
                width: 7px;
                height: 7px;
                margin-top: 6px;
                border-radius: 50%;
                background-color: var(--mat-sys-primary);
            }
        }
    `,
})
export class NotificationsDropdownComponent {}
