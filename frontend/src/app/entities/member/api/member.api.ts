import { inject, Injectable } from '@angular/core';
import { Observable, map } from 'rxjs';
import { APIResponse } from '@shared/models';
import { HttpClient } from '@angular/common/http';
import { Member } from '../model/member.model';
import { memberDTO, toMemberModels } from './member.dto';


@Injectable()
export class MemberApi {
    private http = inject(HttpClient);

    list(workspaceId: number): Observable<Member[]> {
        return this.http
            .get<APIResponse<memberDTO[]>>(`/api/v1/workspaces/${workspaceId}/members`, {withCredentials: true})
            .pipe(map((res) => toMemberModels(res.data)));
    }
}

