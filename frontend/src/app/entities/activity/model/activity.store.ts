import { EntityState, createEntityAdapter } from '@ngrx/entity';
import { createReducer, on } from '@ngrx/store';
import { ActivityModel } from './activity.model';
import { actionActivity } from './activity.actions';

export type ActivityState = EntityState<ActivityModel>;
export const activityAdapter = createEntityAdapter<ActivityModel>({
    sortComparer: (a, b) => new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime(),
});
const initialState = activityAdapter.getInitialState();

export const activityReducer = createReducer(
    initialState,
    on(actionActivity.setList, (state, { activities }) => activityAdapter.addMany(activities, state)),
);
