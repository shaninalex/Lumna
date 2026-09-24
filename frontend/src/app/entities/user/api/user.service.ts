import { inject, Injectable } from '@angular/core';
import type { Observable } from 'rxjs';
import { map } from 'rxjs';
import type { UserModel } from '../model/user.model';
import type { APIResponse } from '@shared/models';
import { HttpClient } from '@angular/common/http';
import { toUserModel, UserModelDTO } from './user.dto';

@Injectable()
export class UserApi {
    http = inject(HttpClient);

    me(): Observable<UserModel> {
        return this.http
            .get<APIResponse<UserModelDTO>>(`/api/v1/user/me`, { withCredentials: true })
            .pipe(map((response) => toUserModel(response.data)));
    }
}
