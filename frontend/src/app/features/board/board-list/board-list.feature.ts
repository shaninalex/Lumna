import { AsyncPipe } from '@angular/common';
import { Component, inject, ChangeDetectionStrategy } from '@angular/core';
import { BoardListItemComponent, selectBoard } from '@entities/board';
import { selectProjects } from '@entities/project';
import { Store } from '@ngrx/store';
import { filter, switchMap } from 'rxjs';

@Component({
    selector: 'lu-board-list-feature',
    imports: [AsyncPipe, BoardListItemComponent],
    changeDetection: ChangeDetectionStrategy.Eager,
    template: `
        @if (boards$ | async; as boards) {
            @for (board of boards; track $index) {
                <lu-board-list-item class="mb-3 d-block" [board]="board" />
            }
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
