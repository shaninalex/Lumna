import { BoardModel, BoardPayloadModel } from '../model/board.model';


export interface BoardModelDTO {
    id: number;
    title: string;
    project_id: number;
    stage_count: number;
    issue_count: number;
    created_at: Date;
    updated_at: Date;
}

export function toBoardModel(boardModel: BoardModelDTO): BoardModel {
    return {
        id: boardModel.id,
        title: boardModel.title,
        description: '',
        projectId: boardModel.project_id,
        stageCount: boardModel.stage_count,
        issueCount: boardModel.issue_count,
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
