import type { ColumnModel } from '@entities/column';
import type { TaskModel } from '@entities/task';

export type KanbanColumn = ColumnModel & { tasks: KanbanCard[] };

export type KanbanCard = TaskModel & { column: number; position: number };

export interface KanbanMoveTask {
    task_id: number
    position: number
    board_id: number
}

export interface KanbanTransferTask {
    task_id: number
    board_id: number
    column_id: number
    position: number
}

export interface KanbanMoveColumn {
    board_id: number
    column_id: number
    position: number
}
