import { AsyncPipe } from '@angular/common';
import { Component, DestroyRef, inject } from '@angular/core';
import { ActivatedRoute, RouterOutlet } from '@angular/router';
import { Store } from '@ngrx/store';
import { filter, map, type Observable } from 'rxjs';
import { MainLayout } from '@core/layout';
import { KanbanBoardWidget } from '@root/src/app/widgets';
import { UiService } from '@shared/ui';
import { switchMap } from 'rxjs/operators';
import { selectBoard } from '@entities/board';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';

@Component({
    selector: 'lu-board-detail-page',
    imports: [MainLayout, AsyncPipe, KanbanBoardWidget, RouterOutlet],

    template: `
        <lu-main-layout>
            @if (boardId$ | async; as boardId) {
                <div class="container-fluid py-4 h-100 d-flex flex-column">
                    <lu-kanban-board-feature [boardId]="boardId" />
                </div>
            }
            <router-outlet />
        </lu-main-layout>
    `,
})
export class BoardDetailPage {
    private store = inject(Store);
    private activeRoute = inject(ActivatedRoute);
    private ui = inject(UiService);
    private destroyRef = inject(DestroyRef);

    boardId$: Observable<number> = this.activeRoute.paramMap.pipe(
        map((params) => params.get('boardId')),
        filter((boardId) => boardId !== null),
        map((boardId) => Number(boardId)),
    );

    constructor() {
        this.boardId$
            .pipe(
                takeUntilDestroyed(this.destroyRef),
                switchMap((boardId) => this.store.select(selectBoard.byId(boardId))),
                filter((board) => board !== null),
            )
            .subscribe((board) => this.ui.setPageTitle(`Board: ${board.title}`));
    }
}
