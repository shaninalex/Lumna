import { inject, Injectable } from '@angular/core';
import { Observable, map } from 'rxjs';
import { APIResponse } from '@shared/models';
import { HttpClient } from '@angular/common/http';
import { invitationDto, toInvitationModel, toInvitationModels } from './invitation.dto';
import { Invitation } from '../model/invitation.model';


@Injectable()
export class InvitationApi {
    private http = inject(HttpClient);

    list(workspaceId: number): Observable<Invitation[]> {
        return this.http
            .get<APIResponse<invitationDto[]>>(`/api/v1/workspaces/${workspaceId}/invitations`, {withCredentials: true})
            .pipe(map((res) => toInvitationModels(res.data)));
    }

    create(workspaceId: number, data: {email: string; role: string}): Observable<Invitation> {
        const payload: Partial<invitationDto> = {
            workspace_id: workspaceId,
            email: data.email,
            role: data.role,
        }
        return this.http
            .post<APIResponse<invitationDto>>(
                `/api/v1/workspaces/${workspaceId}/invitations`,
                payload,
                {withCredentials: true},
            ).pipe(map((res) => toInvitationModel(res.data)));
    }
}

