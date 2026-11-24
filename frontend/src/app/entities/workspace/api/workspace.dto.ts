import {WorkspaceModel} from '../model/workspace.model';

export interface workspaceModelDTO {
    id: number;
    title: string;
    active: boolean;
    owner_email: string;
    created_at: Date;
}


export function toWorkspaceModel(u: workspaceModelDTO): WorkspaceModel {
    return {
        id: u.id,
        title: u.title,
        active: u.active,
        ownerEmail: u.owner_email,
        createdAt: u.created_at,
    }
}

export function toWorkspaceModels(u: workspaceModelDTO[]): WorkspaceModel[] {
    return u.map(w => toWorkspaceModel(w))
}
