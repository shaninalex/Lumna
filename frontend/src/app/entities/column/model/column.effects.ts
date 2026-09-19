import { inject, Injectable } from "@angular/core";
import { Actions, createEffect, ofType } from "@ngrx/effects";
import { catchError, of } from "rxjs";
import { switchMap } from "rxjs/operators";
import type { HttpErrorResponse } from "@angular/common/http";
import { fromErrorResponse } from "@shared/models";
import { actionsColumns } from "./column.actions";
import { ColumnApi } from "../api";

@Injectable()
export class ColumnEffects {
    private actions$ = inject(Actions);
    private api = inject(ColumnApi);

    list_columns$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionsColumns.loadByBoardId),
            switchMap((action) =>
                this.api.list(action.board_id).pipe(
                    switchMap((statuses) => of(actionsColumns.loadByBoardIdSuccess({ columns: statuses }))),
                    catchError((err: HttpErrorResponse) =>
                        of(
                            actionsColumns.loadByBoardIdFailed({
                                errors: fromErrorResponse(err)
                            })
                        )
                    )
                )
            )
        )
    );

    column_create$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionsColumns.create),
            switchMap((action) =>
                this.api.create(action.payload).pipe(
                    switchMap((column) => of(actionsColumns.createSuccess({ column }))),
                    catchError((err: HttpErrorResponse) =>
                        of(
                            actionsColumns.createFailed({
                                errors: fromErrorResponse(err)
                            })
                        )
                    )
                )
            )
        )
    );

    column_delete$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionsColumns.deleteStage),
            switchMap((action) =>
                this.api.delete(action.data).pipe(
                    switchMap((data) => of(actionsColumns.deleteStageSuccess({ data }))),
                    // TODO: call tasks batch delete if data.tasks not empty!
                    catchError((err: HttpErrorResponse) =>
                        of(
                            actionsColumns.createFailed({
                                errors: fromErrorResponse(err)
                            })
                        )
                    )
                )
            )
        )
    )
}
