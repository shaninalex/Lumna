import { createFeatureSelector, createSelector } from '@ngrx/store';
import type { UserState } from './user.store';

const feature = createFeatureSelector<UserState>('user');

export const selectUser = {
    user: createSelector(feature, (state: UserState) => state.user),
    userState: createSelector(feature, (state: UserState) => state),
};
