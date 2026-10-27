import { RefreshToken } from '../model/refresh_token.model';

export interface RefreshTokenDTO {
    id: number
    identity_id: number
    client_id: string
    scopes: string
    expires_at: Date
    revoked: boolean
    created_at: Date
}

export function toRefreshToken(dto: RefreshTokenDTO): RefreshToken {
    return {
        id: dto.id,
        identityId: dto.identity_id,
        clientId: dto.client_id,
        scopes: dto.scopes,
        expiresAt: dto.expires_at,
        revoked: dto.revoked,
        createdAt: dto.created_at,
    }
}

export function toRefreshTokens(dtos: RefreshTokenDTO[]): RefreshToken[] {
    return dtos.map(d => toRefreshToken(d))
}
