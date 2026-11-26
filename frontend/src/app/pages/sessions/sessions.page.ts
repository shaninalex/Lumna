import { Component } from '@angular/core';
import { ManageRefreshTokensFeature } from '@features';
import { MatDivider } from '@angular/material/divider';
import { MatIcon } from '@angular/material/icon';

@Component({
    imports: [ManageRefreshTokensFeature, MatDivider, MatIcon],
    selector: 'lu-sessions-page',
    template: `
        <h2 class="lu-page-title mb-3">Sessions</h2>
        <div class="lu-banner lu-banner--neutral" role="status">
            <mat-icon>info</mat-icon>
            <span>not implemented</span>
        </div>
        <mat-divider class="my-4" />
        <lu-manage-refresh-tokens-feature />
    `,
})
export class SessionsPage {}
