import { ActivityModel } from '../model/activity.model';

export interface activityDTO {
    id: number;
    message: string;
    event_type: string;
    entity_type: string;
    entity_id: number;
    created_at: Date;
}

export function toActivityModel(dto: activityDTO): ActivityModel {
    return {
        id: dto.id,
        message: dto.message,
        eventType: dto.event_type,
        entityType: dto.entity_type,
        entityId: dto.entity_id,
        createdAt: dto.created_at,
    }
}

export function toActivityModels(activityDTOs: activityDTO[]): ActivityModel[] {
    return activityDTOs.map(m => toActivityModel(m));
}
