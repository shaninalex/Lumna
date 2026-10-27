import { Component, inject, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { actionRefreshToken, selectRefreshToken } from '@entities/refresh_token';
import { DatePipe } from '@angular/common';
import { standardTimeFormat } from '@shared/utils';

@Component({
    imports: [DatePipe],
    selector: 'lu-sessions-page',
    template: `
        <h2>Sessions</h2>
        <div class="alert alert-secondary">not implemented</div>
        <hr class="my-4">
        <h2>Tokens</h2>

        <table class="table">
            <thead>
            <tr>
                <th>id</th>
                <th>created</th>
                <th>expired</th>
                <th>revoked</th>
                <th></th>
            </tr>
            </thead>
            <tbody>
                @for (token of refreshTokens(); track token.id) {
                    <tr>
                        <td>{{ token.id }}</td>
                        <td>{{ token.createdAt | date: standardTimeFormat }}</td>
                        <td>{{ token.expiresAt | date: standardTimeFormat }}</td>
                        <td>{{ token.revoked }}</td>
                        <td>
                            <div class="btn-group">
                                <button class="btn btn-sm btn-secondary" [disabled]="token.revoked" (click)="revokeToken(token.id)">Revoke</button>
                                <button class="btn btn-sm btn-danger" (click)="deleteToken(token.id)">
                                    Delete
                                </button>
                            </div>
                        </td>
                    </tr>
                }
            </tbody>
        </table>
    `,
})
export class SessionsPage implements OnInit {
    protected readonly standardTimeFormat = standardTimeFormat;
    private store = inject(Store);
    refreshTokens = this.store.selectSignal(selectRefreshToken.all)

    ngOnInit() {
        this.store.dispatch(actionRefreshToken.getList())
    }

    revokeToken(tokenId: number): void {
        this.store.dispatch(actionRefreshToken.revoke({ tokenId }));
    }

    deleteToken(tokenId: number): void {
        this.store.dispatch(actionRefreshToken.delete({ tokenId }));
    }
}
