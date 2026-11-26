import { Component, Input } from '@angular/core';
import { MatIconButton } from '@angular/material/button';
import { MatIcon } from '@angular/material/icon';
import { MatMenuModule } from '@angular/material/menu';

@Component({
    selector: 'lu-column-menu-dropdown',
    imports: [MatIconButton, MatIcon, MatMenuModule],
    template: `
        <button matIconButton [matMenuTriggerFor]="taskMenu" aria-label="Column menu">
            <mat-icon>menu</mat-icon>
        </button>
        <mat-menu #taskMenu="matMenu">
            <button mat-menu-item type="button" (click)="delete()">Close</button>
            <ng-content select="[columnMenuSlot]" />
        </mat-menu>
    `,
})
export class ColumnMenuDropdownComponent {
    @Input() stageId: number;

    delete() {}
}
