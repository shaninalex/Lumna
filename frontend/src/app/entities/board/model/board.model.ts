export interface BoardModel {
    id: number;
    title: string;
    description: string;
    projectId: number;
    stageCount: number;
    issueCount: number;
    createdAt: Date;
    updatedAt?: Date;
}

export interface BoardPayloadModel {
    title: string;
    projectId: number;
}
