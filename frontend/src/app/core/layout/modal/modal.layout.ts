import { Component, output } from '@angular/core';

@Component({
    selector: 'lu-modal-layout',
    styleUrl: './modal.layout.css',
    template: `
        <div class="modal-layout">
            <div class="modal-layout-backdrop" (click)="closed.emit()"></div>
            <div class="modal-layout-content container ">
                <ng-content/>
            </div>
        </div>
    `
})
export class ModalLayout {
    readonly closed = output<void>();
}
