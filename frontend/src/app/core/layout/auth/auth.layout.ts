import { Component, inject } from '@angular/core';
import { MatCardModule } from '@angular/material/card';
import { UiService } from '@shared/ui';

@Component({
    selector: 'lu-auth-layout',
    imports: [MatCardModule],
    template: `
        <div class="vh-100 vw-100 d-flex align-items-center justify-content-center">
            <mat-card appearance="raised" style="width: 22rem">
                <mat-card-content class="p-4">
                    <ng-content />
                </mat-card-content>
            </mat-card>
        </div>
    `,
})
export class AuthLayout {
    private _ = inject(UiService);
}
