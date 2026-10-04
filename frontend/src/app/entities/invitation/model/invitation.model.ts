export interface Invitation {
    id: number;
    workspaceId: number;
    email: string;
    role: string;
    invitedBy?: number;
    expiresAt: Date;
    acceptedAt?: Date;
    revokedAt?: Date;
    createdAt: Date;
}
