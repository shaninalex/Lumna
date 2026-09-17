import type { ColumnModel } from '@entities/column';
import type { TaskModel } from '@entities/task';

export type KanbanColumn = ColumnModel & { tasks: KanbanCard[] };

export type KanbanCard = TaskModel & { column: number; position: number };

export interface RearangeTasks {
    task_id: number;
    rank: number;
}

export interface KanbanMoveTask extends RearangeTasks {
    board_id: number
}

export interface KanbanTransferTask {
    board_id: number;
    from: RearangeTasks;
    to: RearangeTasks;
}

export interface KanbanMoveColumn {
    id: number;
    previous_index: number;
    current_index: number;
    board_id: number;
    columns_order: number[];
}
