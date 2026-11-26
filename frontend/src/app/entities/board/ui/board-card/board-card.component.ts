import { Component, inject, input } from '@angular/core';
import { RouterLink } from '@angular/router';
import { MatCardModule } from '@angular/material/card';
import { MatIcon } from '@angular/material/icon';
import { MatProgressBar } from '@angular/material/progress-bar';

import { BoardModel } from '../../model';
import { AppRoutes } from '@core';


@Component({
    imports: [RouterLink, MatCardModule, MatIcon, MatProgressBar],
    selector: 'lu-board-card',
    styleUrl: './board-card.component.scss',
    template: `
        <a [routerLink]="appRoutes.board(board().id)" class="lu-plain-link d-block h-100">
            <mat-card class="board-card h-100">
                <mat-card-content>
                    <div class="d-flex justify-content-between align-items-start mb-3">
                        <span class="lu-icon-tile lu-icon-tile--primary">
                            <mat-icon>view_kanban</mat-icon>
                        </span>
                        <span class="lu-tag lu-tag--success">Active</span>
                    </div>

                    <h6 class="mb-1">{{ board().title }}</h6>

                    @if (board().description !== '') {
                        <p class="text-muted small mb-3">{{ board().description }}</p>
                    }

                    <div class="d-flex justify-content-between small mb-2">
                        <span class="text-muted">Progress</span>
                        <span class="fw-medium">75%</span>
                    </div>
                    <mat-progress-bar class="mb-3" mode="determinate" [value]="75" />

                    <div class="lu-meta d-flex justify-content-between text-muted small">
                        <span><mat-icon class="me-1">checklist</mat-icon>24 tasks</span>
                        <span><mat-icon class="me-1">schedule</mat-icon>Updated 20m ago</span>
                    </div>
                </mat-card-content>
            </mat-card>
        </a>
    `,

    host: {
        class: "col-12 col-md-6",
    }
})
export class BoardCardComponent {
    readonly appRoutes = inject(AppRoutes);
    board = input.required<BoardModel>();
}
