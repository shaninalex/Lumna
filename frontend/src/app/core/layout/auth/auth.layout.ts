import { Component, inject } from '@angular/core';
import { UiService } from '@shared/ui';

@Component({
    selector: 'lu-auth-layout',

    template: `
        <div class="vh-100 vw-100 d-flex align-items-center justify-content-center">
            <div style="width: 18rem">
                <div class="card">
                    <div class="card-body">
                        <ng-content />
                    </div>
                </div>
            </div>
        </div>
    `,
})
export class AuthLayout {
    private _ = inject(UiService);
}
