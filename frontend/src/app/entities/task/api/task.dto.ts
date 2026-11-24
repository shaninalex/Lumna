import { TaskAssignModel, TaskCreateModel, TaskEditModel, TaskModel } from '../model/task.model';

export interface TaskAssignModelDTO {
    task_id: number;
    identity_id: number;
}

export function toTaskAssignModelDTO(t: TaskAssignModel): TaskAssignModelDTO {
    return {
        task_id: t.taskId,
        identity_id: t.identityId,
    }
}

export interface TaskEditModelDTO {
    task_id: number;
    title: string;
    body: string;
}

export function toTaskEditModelDTO(t: TaskEditModel): TaskEditModelDTO {
    return {
        task_id: t.taskId,
        title: t.title,
        body: t.body,
    }
}

export interface TaskCreateModelDTO {
    title: string;
    body: string;
    project_id: number;
    position: number;
    column_id: number;
    board_id: number;
    due_to?: Date;
}

export function toTaskCreateModelDTO(t: TaskCreateModel): TaskCreateModelDTO {
    return {
        title: t.title,
        body: t.body,
        project_id: t.projectId,
        position: t.position,
        column_id: t.columnId,
        board_id: t.boardId,
        due_to: t.dueTo,
    }
}

export interface TaskModelDTO {
    id: number;
    title: string;
    body?: string;
    completed: boolean;
    meta: string;
    project_id: number;
    board_id: number;
    column_id: number;
    position: number;
    owner_id: number;
    assignees: number[];
    due_to?: Date;
    created_at: Date;
    updated_at?: Date;
}

export function toTaskModel(t: TaskModelDTO): TaskModel {
    return {
        id: t.id,
        title: t.title,
        body: t.body,
        completed: t.completed,
        meta: t.meta,
        projectId: t.project_id,
        boardId: t.board_id,
        columnId: t.column_id,
        position: t.position,
        ownerId: t.owner_id,
        assignees: t.assignees,
        dueTo: t.due_to,
        createdAt: t.created_at,
        updatedAt: t.updated_at,
    }
}

export function toTaskModels(tasks: TaskModelDTO[]): TaskModel[] {
    return tasks.map(t => toTaskModel(t));
}
