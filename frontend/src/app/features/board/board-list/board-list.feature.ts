import { AsyncPipe } from '@angular/common';
import { Component, inject } from '@angular/core';
import { BoardCardComponent, selectBoard } from '@entities/board';
import { selectProjects } from '@entities/project';
import { Store } from '@ngrx/store';
import { filter, switchMap } from 'rxjs';

@Component({
    selector: 'lu-board-list-feature',
    imports: [AsyncPipe, BoardCardComponent],
    template: `
        @if (boards$ | async; as boards) {
            <div class="row g-3">
                @for (board of boards; track board.id) {
                    <lu-board-card [board]="board" />
                }
            </div>
        }
    `,
})
export class BoardListFeature {
    private store = inject(Store);
    boards$ = this.store.select(selectProjects.currentProjectId).pipe(
        filter((projectId) => projectId !== null),
        switchMap((projectId) => this.store.select(selectBoard.byProjectId(projectId))),
    );
}
