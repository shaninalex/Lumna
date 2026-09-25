import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { ColumnModel } from '@entities/column';
import { TaskModel } from '@entities/task';
import { APIResponse } from '@shared/models';
import { map, Observable } from 'rxjs';
import type { KanbanMoveColumn, KanbanMoveTask, KanbanTransferTask } from '../model';
import { ColumnModelDTO, toColumnModel } from '@entities/column/api/column.dto';
import { TaskModelDTO, toTaskModel } from '@entities/task/api/task.dto';

interface BoardAction {
    action: "move_column" | "move_task" | "change_stage";
    data: KanbanMoveColumn | KanbanMoveTask | KanbanTransferTask
}

@Injectable()
export class KanbanApi {
    private http = inject(HttpClient);

    SetColumns(data: KanbanMoveColumn): Observable<ColumnModel> {
        const payload: BoardAction = {
            action: "move_column",
            data: data
        }
        return this.http
            .post<APIResponse<ColumnModelDTO>>(`/api/v1/tasks/move`, payload, {withCredentials: true})
            .pipe(map((response) => toColumnModel(response.data)));
    }

    TaskAction(action: "move_task" | "change_stage", data: KanbanMoveTask | KanbanTransferTask): Observable<TaskModel> {
        const payload: BoardAction = {
            action: action,
            data: data
        }
        return this.http
            .post<APIResponse<TaskModelDTO>>(`/api/v1/tasks/move`, payload, {withCredentials: true})
            .pipe(map((response) => toTaskModel(response.data)));
    }
}
