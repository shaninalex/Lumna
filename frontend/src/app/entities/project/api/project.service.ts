import { inject, Injectable } from '@angular/core';
import type { Observable } from 'rxjs';
import { map } from 'rxjs';
import type { ProjectModel, ProjectCreateModel } from '../model/project.model';
import type { APIResponse } from '@shared/models';
import { HttpClient, HttpParams } from '@angular/common/http';
import {
    ProjectModelDTO,
    toProjectCreateModelDTO,
    toProjectModel,
    toProjectsModel
} from '@entities/project/api/project.dto';

@Injectable()
export class ProjectApi {
    http = inject(HttpClient);

    GetProjects(workspaceId: number): Observable<ProjectModel[]> {
        return this.http
            .get<APIResponse<ProjectModelDTO[]>>(`/api/v1/projects`, {
                params: new HttpParams().set('workspace_id', workspaceId),
                withCredentials: true,
            })
            .pipe(map((response) => toProjectsModel(response.data)));
    }

    CreateProject(payload: ProjectCreateModel): Observable<ProjectModel> {
        return this.http
            .post<APIResponse<ProjectModelDTO>>(`/api/v1/projects`, toProjectCreateModelDTO(payload), { withCredentials: true })
            .pipe(map((response) => toProjectModel(response.data)));
    }

    DeleteProject(projectId: number): Observable<void> {
        return this.http
            .delete<APIResponse<void>>(`/api/v1/projects/${projectId}`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Patch(projectId: number, payload: ProjectCreateModel): Observable<ProjectModel> {
        return this.http
            .patch<
                APIResponse<ProjectModelDTO>
            >(`/api/v1/projects/${projectId}`, toProjectCreateModelDTO(payload), { withCredentials: true })
            .pipe(map((response) => toProjectModel(response.data)));
    }
}
