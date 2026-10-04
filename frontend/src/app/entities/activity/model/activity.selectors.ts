import { createFeatureSelector, createSelector } from '@ngrx/store';
import { activityAdapter, ActivityState } from './activity.store';

const feature = createFeatureSelector<ActivityState>('activity');
const entitySelectors = activityAdapter.getSelectors();

const selectAll = createSelector(feature, entitySelectors.selectAll);

export const selectActivity = {
    all: selectAll,
    entities: createSelector(feature, entitySelectors.selectEntities),
    byIdAndType: (entityId: number, entityType: string) =>
        createSelector(selectAll, (activities) => activities.filter((b) => b.entityType === entityType && b.entityId === entityId)),
};
