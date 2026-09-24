import { ColumnMeta, ColumnModel, ColumnPayloadModel } from '../model/column.model';

export interface ColumnModelDTO {
    id: number;
    title: number;
    meta: ColumnMeta;
    board_id: number;
    position: number;
    created_at: Date;
    update_at: Date;
}

export function toColumnModel(c: ColumnModelDTO): ColumnModel {
    return {
        id: c.id,
        title: c.title,
        meta: c.meta,
        boardId: c.board_id,
        position: c.position,
        createdAt: c.created_at,
        updateAt: c.update_at,
    }
}

export function toColumnModels(c: ColumnModelDTO[]): ColumnModel[] {
    return c.map((column) => toColumnModel(column));
}

export interface ColumnPayloadModelDTO {
    title: string;
    order: number;
    board_id: number;
}

export function toColumnPayloadModelDTO(p: ColumnPayloadModel): ColumnPayloadModelDTO {
    return {
        title: p.title,
        order: p.order,
        board_id: p.boardId,
    }
}
