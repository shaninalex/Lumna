export interface ColumnModel {
    id: number;
    title: number;
    meta: ColumnMeta;
    boardId: number;
    position: number;
    createdAt: Date;
    updateAt: Date;
}

export interface ColumnPayloadModel {
    title: string;
    order: number;
    boardId: number;
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
