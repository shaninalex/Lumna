import { Member } from '@entities/member/model/member.model';

export interface memberDTO {
    id: number;
    workspace_id: number;
    email: string;
    full_name: string;
    image?: string;
    role: string;
    date_joined: Date;
}

export function toMemberModel(m: memberDTO): Member {
    return {
        id: m.id,
        workspaceId: m.workspace_id,
        email: m.email,
        fullName: m.full_name,
        image: m.image,
        role: m.role,
        dateJoined: m.date_joined,
    }
}

export function toMemberModels(m: memberDTO[]): Member[] {
    return m.map(s => toMemberModel(s));
}
