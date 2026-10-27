import { inject, Injectable } from '@angular/core';
import type { Observable } from 'rxjs';
import { map } from 'rxjs';
import type { APIResponse } from '@shared/models';
import { HttpClient } from '@angular/common/http';
import { RefreshToken } from '../model/refresh_token.model';
import { RefreshTokenDTO, toRefreshToken, toRefreshTokens } from './refresh_token.dto';

@Injectable()
export class RefreshTokenService {
    http = inject(HttpClient);

    List(): Observable<RefreshToken[]> {
        return this.http
            .get<APIResponse<RefreshTokenDTO[]>>(`/api/v1/sessions/refresh-tokens/`, {
                withCredentials: true,
            })
            .pipe(map((response) => toRefreshTokens(response.data)));
    }

    Delete(refreshTokenId: number): Observable<void> {
        return this.http
            .delete<APIResponse<void>>(`/api/v1/sessions/refresh-tokens/${ refreshTokenId }`, {withCredentials: true})
            .pipe(map((response) => response.data));
    }

    Revoke(refreshTokenId: number): Observable<RefreshToken> {
        return this.http
            .patch<
                APIResponse<RefreshTokenDTO>
            >(`/api/v1/sessions/refresh-tokens/${ refreshTokenId }/revoke`, null, {withCredentials: true})
            .pipe(map((response) => toRefreshToken(response.data)));
    }
}
