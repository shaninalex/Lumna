import { Component, output } from '@angular/core';
import { MatIconButton } from '@angular/material/button';
import { MatIcon } from '@angular/material/icon';

@Component({
    selector: 'lu-modal-layout',
    styleUrl: './modal.layout.scss',
    imports: [MatIconButton, MatIcon],
    template: `
        <div class="modal-layout">
            <div class="modal-layout-backdrop" (click)="closed.emit()"></div>
            <div class="modal-layout-content" role="dialog" aria-modal="true">
                <div class="modal-content-scroller">
                    <button matIconButton class="modal-close-button" aria-label="Close" (click)="closed.emit()">
                        <mat-icon>close</mat-icon>
                    </button>

                    <ng-content />
                </div>
            </div>
        </div>
    `,
})
export class ModalLayout {
    readonly closed = output<void>();
}
