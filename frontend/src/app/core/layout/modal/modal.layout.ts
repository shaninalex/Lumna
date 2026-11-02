import { Component, output } from '@angular/core';
import { NgIcon, provideIcons } from '@ng-icons/core';
import { reiconXmark } from '@ng-icons/reicon';

@Component({
    selector: 'lu-modal-layout',
    styleUrl: './modal.layout.scss',
    imports: [NgIcon],
    providers: [provideIcons({ reiconXmark })],
    template: `
        <div class="modal-layout">
            <div class="modal-layout-backdrop" (click)="closed.emit()"></div>
            <div class="modal-layout-content container card p-0">
                <div class="card-body modal-content-scroller">
                    <button class="button modal-close-button" (click)="closed.emit()">
                        <ng-icon name="reiconXmark" />
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
