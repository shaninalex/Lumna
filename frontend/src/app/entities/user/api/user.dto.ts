import { UserModel } from '../model/user.model';

export interface UserModelDTO {
    id: number;
    full_name: string;
    email: string;
    active: boolean;
    created_at: Date;
    updated_at: Date;
}

export function toUserModel(u: UserModelDTO): UserModel {
    return {
        id: u.id,
        fullName: u.full_name,
        email: u.email,
        active: u.active,
        createdAt: u.created_at,
        updatedAt: u.updated_at,
    }
}
