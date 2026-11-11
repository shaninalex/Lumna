import { Component, inject, input } from '@angular/core';
import { BoardModel } from '@entities/board';
import { AppRoutes } from '@core';
import { RouterLink } from '@angular/router';

@Component({
    imports: [RouterLink],
    selector: 'lu-board-card',
    template: `
        <a [routerLink]="appRoutes.board(board().id)" class="card h-100 text-decoration-none border shadow-none">
            <div class="card-body">
                <div class="d-flex justify-content-between align-items-start mb-3">
                    <div class="rounded bg-primary bg-opacity-10 p-2">
                        <i class="bi bi-kanban text-primary fs-5"></i>
                    </div>
                    <span class="badge text-bg-success">Active</span>
                </div>

                <h6 class="mb-1">{{ board().title }}</h6>

                @if (board().description !== '') {
                    <p class="text-muted small mb-3">{{ board().description }}</p>
                }

                <div class="d-flex justify-content-between small mb-2">
                    <span class="text-muted">Progress</span>
                    <span class="fw-medium">75%</span>
                </div>
                <div class="progress mb-3" style="height:5px">
                    <div class="progress-bar" style="width:75%"></div>
                </div>

                <div class="d-flex justify-content-between text-muted small">
                    <span><i class="bi bi-list-task me-1"></i>24 tasks</span>
                    <span><i class="bi bi-clock me-1"></i>Updated 20m ago</span>
                </div>
            </div>
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
