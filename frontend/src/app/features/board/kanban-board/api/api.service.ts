import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { ColumnModel } from '@entities/column';
import { TaskModel } from '@entities/task';
import { APIResponse } from '@shared/models';
import { map, Observable } from 'rxjs';
import type { KanbanMoveColumn, KanbanMoveTask, KanbanTransferTask } from '@features/board/kanban-board/model';

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
            .post<APIResponse<ColumnModel>>(`/api/v1/tasks/move`, payload, {withCredentials: true})
            .pipe(map((response) => response.data));
    }

    TaskAction(action: "move_task" |"change_stage", data: KanbanMoveTask | KanbanTransferTask): Observable<TaskModel> {
        const payload: BoardAction = {
            action: action,
            data: data
        }
        return this.http
            .post<APIResponse<TaskModel>>(`/api/v1/tasks/move`, payload, {withCredentials: true})
            .pipe(map((response) => response.data));
    }
}
