import { inject, Injectable } from "@angular/core";
import type { Observable } from "rxjs";
import { map } from "rxjs";
import type { APIResponse } from "@shared/models";
import { HttpClient, HttpParams } from "@angular/common/http";
import { ColumnDeleteModel, ColumnDeleteResponseModel, ColumnModel, ColumnPayloadModel } from "../model/column.model";

@Injectable()
export class ColumnApi {
    private http = inject(HttpClient);

    list(listId: number): Observable<ColumnModel[]> {
        let params = new HttpParams()
        params = params.append('board_id', listId);
        return this.http
            .get<
                APIResponse<ColumnModel[]>
            >(`/api/v1/columns`, { params, withCredentials: true })
            .pipe(map((response) => response.data));
    }

    create(data: ColumnPayloadModel): Observable<ColumnModel> {
        return this.http
            .post<
                APIResponse<ColumnModel>
            >(`/api/v1/columns`, data, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    delete(data: ColumnDeleteModel): Observable<ColumnDeleteResponseModel> {
        let params = new HttpParams()
        params = params.append('with_tasks', data.withTasks);
        return this.http
            .delete<
                APIResponse<ColumnDeleteResponseModel>
            >(`/api/v1/columns/${data.id}`, { params, withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
