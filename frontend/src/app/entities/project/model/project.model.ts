export interface ProjectModel {
    id: number
    title: string
    key: string
    workspaceId: number
    ownerId: number
    meta: string

    createdAt: Date;
    updatedAt: Date;
}

// Used to create/patch projects
export interface ProjectCreateModel {
    title: string;
    workspaceId: number;
}
