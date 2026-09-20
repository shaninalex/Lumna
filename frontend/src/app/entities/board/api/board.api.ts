import { inject, Injectable } from '@angular/core';
import type { Observable } from 'rxjs';
import { map } from 'rxjs';
import type { APIResponse } from '@shared/models';
import { HttpClient, HttpParams } from '@angular/common/http';
import { BoardModel, BoardPayloadModel } from '../model/board.model';
import { BoardModelDTO, toBoardCreateDTO, toBoardModel, toBoardModels } from './board.dto';

@Injectable()
export class BoardApi {
    private http = inject(HttpClient);

    list(projectId: number): Observable<BoardModel[]> {
        const params = new HttpParams().set('project_id', projectId);
        return this.http
            .get<APIResponse<BoardModelDTO[]>>('/api/v1/boards', {params, withCredentials: true})
            .pipe(map((res) => toBoardModels(res.data)));
    }

    create(payload: BoardPayloadModel): Observable<BoardModel> {
        return this.http
            .post<APIResponse<BoardModelDTO>>('/api/v1/boards', toBoardCreateDTO(payload), {withCredentials: true})
            .pipe(map((res) => toBoardModel(res.data)));
    }

    get(boardId: number): Observable<BoardModel> {
        return this.http
            .get<APIResponse<BoardModelDTO>>(`/api/v1/boards/${boardId}`, {withCredentials: true})
            .pipe(map((res) => toBoardModel(res.data)));
    }

    patch(boardId: number, payload: BoardPayloadModel): Observable<BoardModel> {
        return this.http
            .patch<APIResponse<BoardModelDTO>>(`/api/v1/boards/${boardId}`, toBoardCreateDTO(payload), {withCredentials: true})
            .pipe(map((res) => toBoardModel(res.data)));
    }

    delete(boardId: number): Observable<void> {
        return this.http
            .delete<APIResponse<void>>(`/api/v1/boards/${boardId}`, {withCredentials: true})
            .pipe(map((res) => res.data));
    }
}

