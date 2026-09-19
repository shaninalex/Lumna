import { inject, Injectable } from "@angular/core";
import type { Observable } from "rxjs";
import { map } from "rxjs";
import type { APIResponse } from "@shared/models";
import { HttpClient, HttpParams } from "@angular/common/http";
import type {
    TaskAssignModel,
    TaskCreateModel,
    TaskEditModel,
    TaskListQueryModel,
    TaskModel
} from "../model/task.model";

@Injectable()
export class TaskApi {
    private http = inject(HttpClient);

    list(q: TaskListQueryModel): Observable<TaskModel[]> {
        let params = new HttpParams()
        params = params.append("board_id", q.board_id)

        return this.http
            .get<
                APIResponse<TaskModel[]>
            >(`/api/v1/tasks`, {params, withCredentials: true})
            .pipe(map((response) => response.data));
    }

    create(data: TaskCreateModel): Observable<TaskModel> {
        return this.http
            .post<
                APIResponse<TaskModel>
            >(`/api/v1/tasks`, data, {withCredentials: true})
            .pipe(map((response) => response.data));
    }

    update(data: TaskEditModel): Observable<TaskModel> {
        return this.http
            .patch<
                APIResponse<TaskModel>
            >(`/api/v1/tasks/${data.task_id}`, data, {withCredentials: true})
            .pipe(map((response) => response.data));
    }

    assign(data: TaskAssignModel): Observable<TaskAssignModel> {
        return this.http
            .patch<
                APIResponse<TaskAssignModel>
            >(`/api/v1/tasks/${data.task_id}/assign`, data, {withCredentials: true})
            .pipe(map((response) => response.data));
    }

    delete(taskId: number): Observable<unknown> {
        return this.http
            .delete<
                APIResponse<unknown>
            >(`/api/v1/tasks/${taskId}`, {withCredentials: true})
            .pipe(map((response) => response.data));
    }
}
