import type { EntityState } from "@ngrx/entity";
import { createEntityAdapter } from "@ngrx/entity";
import { createFeature, createReducer, on } from "@ngrx/store";
import type { WorkspaceModel } from "./workspace.model";
import { actionWorkspace } from "./workspace.actions";

export interface WorkspaceState extends EntityState<WorkspaceModel> {
    currentId: number | null;
}

export const workspaceAdapter = createEntityAdapter<WorkspaceModel>();
const initialState: WorkspaceState = workspaceAdapter.getInitialState({
    currentId: null
});

export const workspaceReducer = createReducer(
    initialState,
    on(actionWorkspace.setList, (state, {list}) => workspaceAdapter.setAll(list, state)),
    on(actionWorkspace.created, (state, {data}) => workspaceAdapter.setOne(data, state)),
    on(actionWorkspace.setCurrent, (state, {id}) => ({...state, currentId: id})),
    on(actionWorkspace.removeSuccess, (state, {id}) => workspaceAdapter.removeOne(id, state)),
);

export const workspaceFeature = createFeature({
    name: 'workspace',
    reducer: workspaceReducer,
});
