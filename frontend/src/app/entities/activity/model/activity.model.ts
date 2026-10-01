export interface ActivityModel {
    id: number;
    message: string;
    eventType: string;
    entityType: string
    entityId: number;
    createdAt: Date;
}
