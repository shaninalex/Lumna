import { Component, output, ChangeDetectionStrategy } from '@angular/core';

@Component({
    selector: 'lu-modal-layout',
    styleUrl: './modal.layout.css',
    changeDetection: ChangeDetectionStrategy.Eager,
    template: `
        <div class="modal-layout">
            <div class="modal-layout-backdrop" (click)="closed.emit()"></div>
            <div class="modal-layout-content container card p-0">
                <div class="card-body modal-content-scroller">
                    <div class="d-flex align-content-end">
                        <button class="button ms-auto" (click)="closed.emit()">
                            <i class="fa-solid fa-x"></i>
                        </button>
                    </div>

                    <ng-content />
                </div>
            </div>
        </div>
    `,
})
export class ModalLayout {
    readonly closed = output<void>();
}
