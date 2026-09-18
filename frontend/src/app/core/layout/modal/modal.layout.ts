import { Component, output } from '@angular/core';

@Component({
    selector: 'lu-modal-layout',
    styleUrl: './modal.layout.css',
    template: `
        <div class="modal-layout">
            <div class="modal-layout-backdrop" (click)="closed.emit()"></div>
            <div class="modal-layout-content container card p-0">
                <div class="card-body modal-content-scroller">
                    <button class="btn btn-sm btn-outline-secondary close-btn" (click)="closed.emit()">x</button>
                    <ng-content/>
                </div>
            </div>
        </div>
    `
})
export class ModalLayout {
    readonly closed = output<void>();
}
