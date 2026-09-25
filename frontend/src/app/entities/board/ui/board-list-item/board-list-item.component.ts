import { Component, inject, input } from '@angular/core';
import type { BoardModel } from '../../model/board.model';
import { RouterLink } from '@angular/router';
import { AppRoutes } from '@core';
import { TimeAgoPipe } from '@shared/utils';

@Component({
    selector: 'lu-board-list-item',
    imports: [RouterLink, TimeAgoPipe],

    template: `
        <a [routerLink]="appRoutes.board(board().id)" class="text-decoration-none">
            <h5 class="mb-1">{{ board().title }}</h5>
        </a>
        @if (board().issueCount) {
            <p class="mb-2 text-muted">{{ board().description }}</p>
        }
        <div class="d-flex gap-3 small text-muted">
            @if (board().issueCount) {
                <span>{{ board().issueCount }} cards</span>
            }
            @if (board().stageCount) {
                <span>{{ board().stageCount }} columns</span>
            }
            @if (board().updatedAt) {
                <span>Updated {{ board()?.updatedAt | timeAgo }}</span>
            }
        </div>
    `,
})
export class BoardListItemComponent {
    board = input.required<BoardModel>();
    readonly appRoutes = inject(AppRoutes);
}
