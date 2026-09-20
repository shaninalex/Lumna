import { createFeatureSelector, createSelector } from '@ngrx/store';
import { boardAdapter, type BoardState } from './board.store';

const feature = createFeatureSelector<BoardState>('board');
const entitySelectors = boardAdapter.getSelectors();

const selectAll = createSelector(feature, entitySelectors.selectAll);

export const selectBoard = {
    all: selectAll,
    entities: createSelector(feature, entitySelectors.selectEntities),
    total: createSelector(feature, entitySelectors.selectTotal),
    byProjectId: (projectId: number) =>
        createSelector(selectAll, (boards) => boards.filter((b) => b.projectId === projectId)),
    byId: (listId: number) =>
        createSelector(feature, (state) => entitySelectors.selectEntities(state)[listId] ?? null),
};
