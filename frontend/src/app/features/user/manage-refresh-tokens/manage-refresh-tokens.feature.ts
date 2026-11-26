import { Component, inject, OnInit } from '@angular/core';
import { DatePipe } from '@angular/common';
import { MatButton } from '@angular/material/button';
import { MatTableModule } from '@angular/material/table';
import { Store } from '@ngrx/store';
import { actionRefreshToken, selectRefreshToken } from '@entities/refresh_token';
import { standardTimeFormat } from '@shared/utils';

@Component({
    imports: [DatePipe, MatButton, MatTableModule],
    selector: 'lu-manage-refresh-tokens-feature',
    template: `
        <h2>Tokens</h2>
        <table mat-table [dataSource]="refreshTokens()" class="w-100">
            <ng-container matColumnDef="id">
                <th mat-header-cell *matHeaderCellDef>id</th>
                <td mat-cell *matCellDef="let token">{{ token.id }}</td>
            </ng-container>
            <ng-container matColumnDef="created">
                <th mat-header-cell *matHeaderCellDef>created</th>
                <td mat-cell *matCellDef="let token">
                    {{ token.createdAt | date: standardTimeFormat }}
                </td>
            </ng-container>
            <ng-container matColumnDef="expired">
                <th mat-header-cell *matHeaderCellDef>expired</th>
                <td mat-cell *matCellDef="let token">
                    {{ token.expiresAt | date: standardTimeFormat }}
                </td>
            </ng-container>
            <ng-container matColumnDef="revoked">
                <th mat-header-cell *matHeaderCellDef>revoked</th>
                <td mat-cell *matCellDef="let token">{{ token.revoked }}</td>
            </ng-container>
            <ng-container matColumnDef="actions">
                <th mat-header-cell *matHeaderCellDef></th>
                <td mat-cell *matCellDef="let token" class="text-end">
                    <div class="d-flex gap-2 justify-content-end">
                        <button
                            matButton="tonal"
                            [disabled]="token.revoked"
                            (click)="revokeToken(token.id)"
                        >
                            Revoke
                        </button>
                        <button matButton="tonal" class="lu-danger" (click)="deleteToken(token.id)">
                            Delete
                        </button>
                    </div>
                </td>
            </ng-container>

            <tr mat-header-row *matHeaderRowDef="columns"></tr>
            <tr mat-row *matRowDef="let row; columns: columns"></tr>
        </table>
    `,
})
export class ManageRefreshTokensFeature implements OnInit {
    protected readonly standardTimeFormat = standardTimeFormat;
    private store = inject(Store);
    refreshTokens = this.store.selectSignal(selectRefreshToken.all);
    readonly columns = ['id', 'created', 'expired', 'revoked', 'actions'];

    ngOnInit() {
        this.store.dispatch(actionRefreshToken.getList());
    }

    revokeToken(tokenId: number): void {
        this.store.dispatch(actionRefreshToken.revoke({ tokenId }));
    }

    deleteToken(tokenId: number): void {
        this.store.dispatch(actionRefreshToken.delete({ tokenId }));
    }
}
