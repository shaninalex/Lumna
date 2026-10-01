import { inject, Injectable } from '@angular/core';
import { Observable, map } from 'rxjs';
import { APIResponse } from '@shared/models';
import { HttpClient, HttpParams } from '@angular/common/http';
import { ActivityModel } from '../model/activity.model';
import { activityDTO, toActivityModels } from './activity.dto';

@Injectable()
export class ActivityApi {
    private http = inject(HttpClient);

    list(entityId: number, entityType: string): Observable<ActivityModel[]> {
        let params = new HttpParams();
        params = params.set('entity_id', entityId)
        params = params.set('entity_type', entityType);
        return this.http
            .get<APIResponse<activityDTO[]>>(`/api/v1/tasks/${entityId}/activities`, {params, withCredentials: true})
            .pipe(map((res) => toActivityModels(res.data)));
    }
}

