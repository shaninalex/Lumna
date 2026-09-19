export interface ColumnModel {
    id: number;
    title: number;
    meta: ColumnMeta;
    board_id: number;
    position: number;
    created_at: Date;
    update_at: Date;
}

export interface ColumnPayloadModel {
    title: string;
    order: number;
    board_id: number;
}

export interface ColumnMeta {
    order: number;
    color: string;
    icon: string;
    expanded: boolean;
}

export interface ColumnDeleteModel {
    id: number;
    withTasks: boolean;
}

export interface ColumnDeleteResponseModel {
    id: number;
    tasks: number[];
}
