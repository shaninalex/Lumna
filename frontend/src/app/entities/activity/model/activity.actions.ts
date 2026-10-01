import { createActionGroup, props } from '@ngrx/store';
import { ActivityModel } from './activity.model';

export const actionActivity = createActionGroup({
    source: 'Activity',
    events: {
        'get list': props<{ entityId: number, entityType: string }>(),
        'set list': props<{ activities: ActivityModel[] }>(),
    },
});
