import { createFeatureSelector, createSelector } from '@ngrx/store';
import type { ProjectState } from './project.store';
import { projectsAdapter } from './project.store';

const feature = createFeatureSelector<ProjectState>('project');
const entitySelectors = projectsAdapter.getSelectors();

const selectAll = createSelector(feature, entitySelectors.selectAll);
const selectCurrentProjectId = createSelector(feature, (state) => state?.currentId ?? null);

const selectCurrentProject = createSelector(
    selectAll,
    selectCurrentProjectId,
    (projects, currentId) => projects.find((p) => p.id === currentId) ?? null,
);

export const selectProjects = {
    all: selectAll,
    entities: createSelector(feature, entitySelectors.selectEntities),
    total: createSelector(feature, entitySelectors.selectTotal),
    byId: (id: number) => createSelector(selectAll, (list) => list.find((p) => p.id === id)),
    byWorkspaceId: (workspaceId: number) =>
        createSelector(selectAll, (list) => list.filter((p) => p.workspaceId === workspaceId)),
    currentProjectId: selectCurrentProjectId,
    currentProject: selectCurrentProject,
};
