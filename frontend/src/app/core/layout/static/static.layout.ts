import { Component, inject } from '@angular/core';
import { MatCardModule } from '@angular/material/card';
import { UiService } from '@shared/ui';

@Component({
    selector: 'lu-static-layout',
    imports: [MatCardModule],
    template: `
        <div class="container p-5">
            <mat-card>
                <mat-card-content class="p-4">
                    <ng-content />
                </mat-card-content>
            </mat-card>
        </div>
    `,
})
export class StaticLayout {
    private _ = inject(UiService);
}
