import { UserModel } from '../model/user.model';

export interface UserModelDTO {
    id: number;
    full_name: string;
    email: string;
    active: boolean;
}

export function toUserModel(u: UserModelDTO): UserModel {
    return {
        id: u.id,
        fullName: u.full_name,
        email: u.email,
        active: u.active,
    }
}
