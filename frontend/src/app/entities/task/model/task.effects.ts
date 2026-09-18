import { inject, Injectable } from "@angular/core";
import { Actions, createEffect, ofType } from "@ngrx/effects";
import { catchError, of } from "rxjs";
import { TaskApi } from "@entities/task/api";
import { switchMap } from "rxjs/operators";
import type { HttpErrorResponse } from "@angular/common/http";
import { fromErrorResponse } from "@shared/models";
import { actionTask } from "./task.actions";

@Injectable()
export class TaskEffects {
    private actions$ = inject(Actions);
    private api = inject(TaskApi);

    task_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionTask.getList),
            switchMap((action) =>
                this.api.list(action.query).pipe(
                    switchMap((tasks) => of(actionTask.getListSuccess({ tasks }))),
                    catchError((err: HttpErrorResponse) =>
                        of(
                            actionTask.getListFailed({
                                errors: fromErrorResponse(err)
                            })
                        )
                    )
                )
            )
        )
    );

    task_create$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionTask.create),
            switchMap((action) =>
                this.api.create(action.data).pipe(
                    switchMap((task) => of(actionTask.createSuccess({ task }))),
                    catchError((err: HttpErrorResponse) =>
                        of(
                            actionTask.createFailed({
                                errors: fromErrorResponse(err)
                            })
                        )
                    )
                )
            )
        )
    );

    task_update$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionTask.updateTask),
            switchMap((action) =>
                this.api.update(action.data).pipe(
                    switchMap((task) => of(actionTask.setTask({ task }))),
                    catchError((err: HttpErrorResponse) =>
                        of(actionTask.updateFailed({ errors: fromErrorResponse(err) }))
                    )
                )
            )
        )
    )
}
