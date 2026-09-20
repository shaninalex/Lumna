import { BoardModel, BoardPayloadModel } from '../model/board.model';


export interface BoardModelDTO {
    id: number;
    title: string;
    project_id: number;
    created_at: Date;
    updated_at: Date;
}

export function toBoardModel(boardModel: BoardModelDTO): BoardModel {
    return {
        id: boardModel.id,
        title: boardModel.title,
        projectId: boardModel.project_id,
        createdAt: boardModel.created_at,
        updatedAt: boardModel.updated_at,
    }
}

export function toBoardModels(boardModels: BoardModelDTO[]): BoardModel[] {
    return boardModels.map(m => toBoardModel(m));
}


export interface BoardCreateDTO {
    title: string;
    project_id: number;
}


export function toBoardCreateDTO(b: BoardPayloadModel): BoardCreateDTO {
    return {
        title: b.title,
        project_id: b.projectId
    }
}
