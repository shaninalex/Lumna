export interface WorkspaceModel {
    id: number;
    title: string;
    active: boolean;
    ownerEmail: string;
    createdAt: Date;
}

export interface WorkspaceCreateModel {
    title: string;
    email: string;
}
