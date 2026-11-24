import { inject, Injectable } from "@angular/core";
import type { Observable } from "rxjs";
import { map } from "rxjs";
import type { APIResponse } from "@shared/models";
import { HttpClient } from "@angular/common/http";
import type { WorkspaceCreateModel, WorkspaceModel } from "../model/workspace.model";
import { toWorkspaceModel, toWorkspaceModels, workspaceModelDTO } from "./workspace.dto";

@Injectable()
export class WorkspaceApi {
    http = inject(HttpClient);

    List(): Observable<WorkspaceModel[]> {
        return this.http
            .get<APIResponse<workspaceModelDTO[]>>(`/api/v1/workspaces`, {withCredentials: true})
            .pipe(map((response) => toWorkspaceModels(response.data)));
    }

    Create(data: WorkspaceCreateModel): Observable<WorkspaceModel> {
        return this.http
            .post<APIResponse<workspaceModelDTO>>(`/api/v1/workspaces`, data, {withCredentials: true})
            .pipe(map((response) => toWorkspaceModel(response.data)));
    }

    remove(id: number): Observable<void> {
        return this.http
            .delete<APIResponse<void>>(`/api/v1/workspaces/${id}`, {withCredentials: true})
            .pipe(map((response) => response.data));
    }
}
