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
import {
    TaskAssignModelDTO, TaskModelDTO,
    toTaskAssignModel,
    toTaskAssignModelDTO,
    toTaskCreateModelDTO,
    toTaskEditModelDTO, toTaskModel, toTaskModels
} from './task.dto';

@Injectable()
export class TaskApi {
    private http = inject(HttpClient);

    list(q: TaskListQueryModel): Observable<TaskModel[]> {
        let params = new HttpParams()
        params = params.append("board_id", q.boardId)

        return this.http
            .get<
                APIResponse<TaskModelDTO[]>
            >(`/api/v1/tasks`, {params, withCredentials: true})
            .pipe(map((response) => toTaskModels(response.data)));
    }

    create(data: TaskCreateModel): Observable<TaskModel> {
        return this.http
            .post<
                APIResponse<TaskModelDTO>
            >(`/api/v1/tasks`, toTaskCreateModelDTO(data), {withCredentials: true})
            .pipe(map((response) => toTaskModel(response.data)));
    }

    update(data: TaskEditModel): Observable<TaskModel> {
        return this.http
            .patch<
                APIResponse<TaskModelDTO>
            >(`/api/v1/tasks/${data.taskId}`, toTaskEditModelDTO(data), {withCredentials: true})
            .pipe(map((response) => toTaskModel(response.data)));
    }

    assign(data: TaskAssignModel): Observable<TaskAssignModel> {
        return this.http
            .patch<
                APIResponse<TaskAssignModelDTO>
            >(`/api/v1/tasks/${data.taskId}/assign`, toTaskAssignModelDTO(data), {withCredentials: true})
            .pipe(map((response) => toTaskAssignModel(response.data)));
    }

    delete(taskId: number): Observable<unknown> {
        return this.http
            .delete<
                APIResponse<unknown>
            >(`/api/v1/tasks/${taskId}`, {withCredentials: true})
            .pipe(map((response) => response.data));
    }
}
