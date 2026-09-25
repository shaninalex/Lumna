import { Component, inject, Input, ChangeDetectionStrategy } from '@angular/core';
import type { BoardModel } from '../../model/board.model';
import { RouterLink } from '@angular/router';
import { AppRoutes } from '@core';

@Component({
    selector: 'lu-board-list-item',
    imports: [RouterLink],
    changeDetection: ChangeDetectionStrategy.Eager,
    template: `
        <a [routerLink]="appRoutes.board(board.id)" class="text-decoration-none">
            <h5 class="mb-1">{{ board.title }}</h5>
        </a>

        <p class="mb-2 text-muted">Main board for feature development, bugs and improvements.</p>

        <div class="d-flex gap-3 small text-muted">
            <span>42 cards</span>
            <span>5 columns</span>
            <span>Updated 2 hours ago</span>
        </div>
    `,
})
export class BoardListItemComponent {
    @Input() board: BoardModel;

    readonly appRoutes = inject(AppRoutes);
}
