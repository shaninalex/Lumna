import { createFeatureSelector, createSelector } from '@ngrx/store';
import { refreshTokenAdapter, RefreshTokenState } from './refresh_token.store';

const feature = createFeatureSelector<RefreshTokenState>('refresh_token');
const entitySelectors = refreshTokenAdapter.getSelectors();

const selectAll = createSelector(feature, entitySelectors.selectAll);

export const selectRefreshToken = {
    all: selectAll,
    entities: createSelector(feature, entitySelectors.selectEntities),
};
