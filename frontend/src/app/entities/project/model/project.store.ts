import { createReducer, on } from '@ngrx/store';
import type { ProjectModel } from './project.model';
import type { EntityState } from '@ngrx/entity';
import { createEntityAdapter } from '@ngrx/entity';
import { actionProject } from './project.actions';
import { actionWorkspace } from '@entities/workspace';

export interface ProjectState extends EntityState<ProjectModel> {
    currentId: number | null;
}
export const projectsAdapter = createEntityAdapter<ProjectModel>();
export const initialState: ProjectState = projectsAdapter.getInitialState({
    currentId: null,
});

export const projectReducer = createReducer(
    initialState,
    on(actionProject.setList, (state, action) => projectsAdapter.addMany(action.projects, state)),
    on(actionProject.createSuccefull, (state, action) => projectsAdapter.upsertOne(action.project, state)),
    on(actionProject.deleteSuccefull, (state, action) =>
        projectsAdapter.removeOne(action.projectId, {
            ...state,
            currentId: state.currentId === action.projectId ? null : state.currentId,
        }),
    ),
    on(actionProject.setCurrent, (state, { id }) => ({ ...state, currentId: id })),
    // Drop the current project as soon as it no longer belongs to the active workspace
    on(actionWorkspace.setCurrent, (state, { id }) => {
        if (state.currentId === null) {
            return state;
        }
        const current = state.entities[state.currentId];
        return current && current.workspaceId === id ? state : { ...state, currentId: null };
    }),
);
