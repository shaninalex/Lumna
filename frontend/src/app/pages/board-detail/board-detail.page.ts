import { AsyncPipe } from '@angular/common';
import { Component, DestroyRef, inject } from '@angular/core';
import { ActivatedRoute, RouterLink, RouterOutlet } from '@angular/router';
import { Store } from '@ngrx/store';
import { filter, map, type Observable } from 'rxjs';
import { MainLayout } from '@core/layout';
import { KanbanBoardWidget } from '@root/src/app/widgets';
import { UiService } from '@shared/ui';
import { switchMap } from 'rxjs/operators';
import { selectBoard } from '@entities/board';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { selectWorkspaces } from '@entities/workspace';
import { selectProjects } from '@entities/project';
import { AppRoutes } from '@core';

@Component({
    selector: 'lu-board-detail-page',
    imports: [MainLayout, AsyncPipe, KanbanBoardWidget, RouterOutlet, RouterLink],
    template: `
        <lu-main-layout>
            @if (boardId$ | async; as boardId) {
                <div class="container-fluid pt-0 h-100 d-flex flex-column">
                    <div class="d-flex align-items-center gap-2">
                        <a [routerLink]="appRoutes.workspaceDetail(workspace()?.id)"
                           class="text-muted text-decoration-none small">
                            {{ workspace()?.title }}
                        </a>
                        <i class="text-muted small">></i>
                        <a [routerLink]="appRoutes.projectDetail(project()?.id)"
                           class="text-muted text-decoration-none small">
                            {{ project()?.title }}
                        </a>
                    </div>
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
    private appRoutes = inject(AppRoutes)

    workspace = this.store.selectSignal(selectWorkspaces.currentWorkspace)
    project = this.store.selectSignal(selectProjects.currentProject)

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
