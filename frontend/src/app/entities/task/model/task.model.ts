export interface TaskModel {
    id: number;
    title: string;
    body: string;
    completed: boolean;
    meta: string;
    project_id: number;
    board_id: number;
    column_id: number;
    position: number;
    owner_id: number;
    assignees_ids: number[];
    due_to?: Date;
    created_at: Date;
    updated_at: Date;
    task_events: EntityEvent[];
}

export interface EntityEvent {
    id: number;
    identity_id?: number;
    entity_id?: number;
    entity_type?: string;
    event_type: string;
    data: string;
    created_at: Date;
}

export interface TaskCreateModel {
    title: string;
    body: string;
    project_id: number;
    position: number;
    column_id: number;
    board_id: number;
    due_to?: Date;
}

export interface TaskListQueryModel {
    board_id: number;
}
