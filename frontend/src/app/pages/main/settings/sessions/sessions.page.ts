import { Component } from '@angular/core';
import { ManageRefreshTokensFeature } from '@features';

@Component({
    imports: [ManageRefreshTokensFeature],
    selector: 'lu-sessions-page',
    template: `
        <h2>Sessions</h2>
        <div class="alert alert-secondary">not implemented</div>
        <hr class="my-4">
        <lu-manage-refresh-tokens-feature />
    `,
})
export class SessionsPage {}
