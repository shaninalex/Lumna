export interface TaskModel {
    id: number;
    title: string;
    body?: string;
    completed: boolean;
    meta: string;
    projectId: number;
    boardId: number;
    columnId: number;
    position: number;
    ownerId: number;
    assignees: number[];
    dueTo?: Date;
    createdAt: Date;
    updatedAt?: Date;
    taskEvents: EntityEvent[];
}

export interface EntityEvent {
    id: number;
    identityId?: number;
    entityId?: number;
    entityType?: string;
    eventType: string;
    data: string;
    createdAt: Date;
}

export interface TaskCreateModel {
    title: string;
    body: string;
    projectId: number;
    position: number;
    columnId: number;
    boardId: number;
    dueTo?: Date;
}

export const taskCreateModelDefault: TaskCreateModel = {
    title: '',
    body: '',
    projectId: 0,
    position: 0,
    columnId: 0,
    boardId: 0,
}

export interface TaskListQueryModel {
    boardId: number;
}

export interface TaskEditModel {
    taskId: number;
    title: string;
    body: string;
}

export interface TaskAssignModel {
    taskId: number;
    identityId: number;
}
