import { createFeatureSelector, createSelector } from '@ngrx/store';
import type { ColumnState } from './column.store';
import { statusAdapter } from './column.store';

const feature = createFeatureSelector<ColumnState>('column');
const entitySelectors = statusAdapter.getSelectors();

const selectAll = createSelector(feature, entitySelectors.selectAll);

export const selectColumns = {
    all: selectAll,
    entities: createSelector(feature, entitySelectors.selectEntities),
    total: createSelector(feature, entitySelectors.selectTotal),
    byId: (id: number) => createSelector(selectAll, (list) => list.find((a) => a.id === id)),
    byListId: (listId: number) =>
        createSelector(selectAll, (list) => list.filter((a) => a.board_id === listId)),
    loading: createSelector(feature, (state) => state?.loading ?? false),
    error: createSelector(feature, (state) => state?.errors ?? []),
};
