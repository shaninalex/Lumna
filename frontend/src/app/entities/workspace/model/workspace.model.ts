export interface WorkspaceModel {
    id: number;
    title: string;
    active: string;
    owner_email: string;
    created_at: Date;
    updated_at?: Date;
}

export interface WorkspaceCreateModel {
    title: string;
    email: string;
}
