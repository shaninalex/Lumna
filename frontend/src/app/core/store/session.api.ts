import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import type { Observable } from 'rxjs';
import { map } from 'rxjs';
import type { UserModel } from '@entities/user';
import type { APIResponse } from '@shared/models';

@Injectable({
    providedIn: 'root',
})
export class SessionApi {
    http = inject(HttpClient);

    login(email: string, password: string): Observable<UserModel> {
        return this.http
            .post<
                APIResponse<UserModel>
            >(`/api/v1/auth/login`, { email, password }, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    logout(): Observable<void> {
        return this.http
            .post<APIResponse<void>>(`/api/v1/auth/logout`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    refresh(): Observable<void> {
        return this.http
            .post<APIResponse<void>>(`/api/v1/auth/refresh`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
