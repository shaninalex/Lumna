import { ProjectCreateModel, ProjectModel } from '../model/project.model';

export interface ProjectModelDTO {
    id: number
    title: string
    key: string
    workspace_id: number
    owner_id: number
    meta: string
    created_at: Date;
    updated_at: Date;
}

export function toProjectModel(p: ProjectModelDTO): ProjectModel {
    return {
        id: p.id,
        title: p.title,
        key: p.key,
        workspaceId: p.workspace_id,
        ownerId: p.owner_id,
        meta: p.meta,
        createdAt: p.created_at,
        updatedAt: p.updated_at,
    }
}

export function toProjectsModel(projects: ProjectModelDTO[]): ProjectModel[] {
    return projects.map(p => toProjectModel(p));
}

// Used to create/patch projects
export interface ProjectCreateModelDTO {
    title: string;
    workspace_id: number;
}

export function toProjectCreateModelDTO(p: ProjectCreateModel): ProjectCreateModelDTO {
    return {
        title: p.title,
        workspace_id: p.workspaceId,
    }
}
