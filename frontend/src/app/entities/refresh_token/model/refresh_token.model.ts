export interface RefreshToken {
    id: number
    identityId: number
    clientId: string
    scopes: string
    expiresAt: Date
    revoked: boolean
    createdAt: Date
}
