import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import { catchError, exhaustMap, of, switchMap } from 'rxjs';
import { ProjectApi } from '../api/project.service';
import type { HttpErrorResponse } from '@angular/common/http';
import { fromErrorResponse } from '@shared/models';
import { actionProject } from './project.actions';

@Injectable()
export class ProjectEffects {
    private actions$ = inject(Actions);
    private projectsApi = inject(ProjectApi);

    get_projects_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionProject.getList),
            exhaustMap((action) =>
                this.projectsApi.GetProjects(action.workspaceId).pipe(
                    switchMap((data) => of(actionProject.setList({ projects: data })))
                ),
            ),
        ),
    );

    create_project$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionProject.create),
            exhaustMap((action) =>
                this.projectsApi.CreateProject(action.payload).pipe(
                    switchMap((data) => of(actionProject.createSuccefull({ project: data }))),
                    catchError((err: HttpErrorResponse) => of(actionProject.createFailed({ errors: fromErrorResponse(err) })))
                ),
            ),
        ),
    );

    update_project$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionProject.patch),
            exhaustMap((action) =>
                this.projectsApi.Patch(action.id, action.data).pipe(
                    switchMap((data) => of(actionProject.patchSuccessfull({ data })))
                ),
            ),
        ),
    );

    delete_project$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionProject.delete),
            exhaustMap((action) =>
                this.projectsApi.DeleteProject(action.project_id).pipe(
                    switchMap(() => of(actionProject.deleteSuccefull({ projectId: action.project_id }))))
                )
        ),
    );
}
