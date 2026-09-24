import { createFeatureSelector, createSelector } from '@ngrx/store';
import type { TaskState } from './task.store';
import { taskAdapter } from './task.store';

const feature = createFeatureSelector<TaskState>('task');
const entitySelectors = taskAdapter.getSelectors();

const selectAll = createSelector(feature, entitySelectors.selectAll);

export const selectTasks = {
    all: selectAll,
    entities: createSelector(feature, entitySelectors.selectEntities),
    total: createSelector(feature, entitySelectors.selectTotal),
    byId: (id: number) =>
        createSelector(selectAll, (list) => list.find((a) => a.id === id)),
    byProject: (projectId: number) =>
        createSelector(selectAll, (list) => list.filter((a) => a.projectId === projectId)),
    countByProjectId: (projectId: number) =>
        createSelector(selectAll, (list) => list.filter((a) => a.projectId === projectId).length),
    byBoardId: (boardId: number) =>
        createSelector(selectAll, (tasks) => tasks.filter((tasks) => tasks.boardId === boardId)),
};
