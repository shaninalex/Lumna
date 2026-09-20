export interface BoardModel {
    id: number;
    title: string;
    projectId: number;
    createdAt: Date;
    updatedAt: Date;
}

export interface BoardPayloadModel {
    title: string;
    projectId: number;
}
