import { inject, Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { BoardApi } from '../api/board.api';
import { catchError, exhaustMap, of, switchMap } from 'rxjs';
import type { HttpErrorResponse } from '@angular/common/http';
import { fromErrorResponse } from '@shared/models';
import { actionBoard } from './board.actions';

@Injectable()
export class BoardEffects {
    private actions$ = inject(Actions);
    private listApi = inject(BoardApi);

    list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionBoard.getList),
            exhaustMap((action) =>
                this.listApi
                    .list(action.projectId)
                    .pipe(switchMap((boards) => of(actionBoard.setList({boards})))),
            ),
        ),
    );

    create$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionBoard.create),
            exhaustMap((action) =>
                this.listApi.create(action.data).pipe(
                    switchMap((board) => of(actionBoard.createSuccess({board}))),
                    catchError((err: HttpErrorResponse) =>
                        of(actionBoard.createFailed({errors: fromErrorResponse(err)})),
                    ),
                ),
            ),
        ),
    );

    patch_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionBoard.patch),
            exhaustMap((action) =>
                this.listApi.patch(action.boardId, action.data).pipe(
                    switchMap((board) => of(actionBoard.patchSuccess({board}))),
                    catchError((err: HttpErrorResponse) =>
                        of(actionBoard.patchFailed({errors: fromErrorResponse(err)})),
                    ),
                ),
            ),
        ),
    );

    delete_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionBoard.delete),
            exhaustMap((action) =>
                this.listApi.delete(action.boardId).pipe(
                    switchMap(() => of(actionBoard.deleteSuccess({boardId: action.boardId}))),
                    catchError((err: HttpErrorResponse) =>
                        of(actionBoard.deleteFailed({errors: fromErrorResponse(err)})),
                    ),
                ),
            ),
        ),
    );

    get_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionBoard.getById),
            exhaustMap((action) =>
                this.listApi.get(action.boardId).pipe(
                    switchMap((board) => of(actionBoard.set({board})),),
                    catchError((err: HttpErrorResponse) =>
                        of(actionBoard.getFailed({errors: fromErrorResponse(err)})),
                    ),
                ),
            ),
        ),
    );
}
