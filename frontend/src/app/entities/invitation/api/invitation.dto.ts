import { Invitation } from '../model/invitation.model';

export interface invitationDto {
    id: number;
    workspace_id: number;
    email: string;
    role: string;
    invited_by?: number;
    expires_at: Date;
    accepted_at?: Date;
    revoked_at?: Date;
    created_at: Date;
}

export function toInvitationModel(i: invitationDto): Invitation {
    return {
        id: i.id,
        workspaceId: i.workspace_id,
        email: i.email,
        role: i.role,
        invitedBy: i.invited_by,
        expiresAt: i.expires_at,
        acceptedAt: i.accepted_at,
        revokedAt: i.revoked_at,
        createdAt: i.created_at,
    }
}

export function toInvitationDTO(i: Invitation): invitationDto {
    return {
        id: i.id,
        workspace_id: i.workspaceId,
        email: i.email,
        role: i.role,
        invited_by: i.invitedBy,
        expires_at: i.expiresAt,
        accepted_at: i.acceptedAt,
        revoked_at: i.revokedAt,
        created_at: i.createdAt,
    }
}

export function toInvitationModels(m: invitationDto[]): Invitation[] {
    return m.map(s => toInvitationModel(s));
}
