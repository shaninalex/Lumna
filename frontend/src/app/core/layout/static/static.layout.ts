import { Component, inject, ChangeDetectionStrategy } from '@angular/core';
import { UiService } from '@shared/ui';

@Component({
    selector: 'lu-static-layout',
    imports: [],
    changeDetection: ChangeDetectionStrategy.Eager,
    template: `
        <div class="p-5 d-flex align-items-center justify-content-center">
            <div class="w-100">
                <div style="container">
                    <div class="card">
                        <div class="card-body">
                            <ng-content />
                        </div>
                    </div>
                </div>
            </div>
        </div>
    `,
})
export class StaticLayout {
    private _ = inject(UiService);
}
