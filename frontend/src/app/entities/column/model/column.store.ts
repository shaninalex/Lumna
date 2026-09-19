import { createReducer, on } from '@ngrx/store';
import { createEntityAdapter, type EntityState } from '@ngrx/entity';
import type { Error } from '@shared/models';
import type { ColumnModel } from './column.model';
import { actionsColumns } from './column.actions';

export interface ColumnState extends EntityState<ColumnModel> {
    loading: boolean;
    errors: Error[];
}

export const stageAdapter = createEntityAdapter<ColumnModel>({
    sortComparer: (a, b) => a.position - b.position,
});

const initialState: ColumnState = stageAdapter.getInitialState({
    loading: false,
    errors: [],
});

export const statusReducer = createReducer(
    initialState,
    on(actionsColumns.createSuccess, (state, action) => stageAdapter.addOne(action.column, state)),
    on(actionsColumns.loadByBoardIdSuccess, (state, action) =>
        stageAdapter.upsertMany(action.columns, state),
    ),
    on(
        actionsColumns.loadByBoardIdFailed,
        actionsColumns.createFailed,
        actionsColumns.reorderFailed,
        actionsColumns.deleteStageFailed,
        (state, action) => ({
            ...state,
            errors: action.errors,
        }),
    ),
    on(actionsColumns.deleteStageSuccess, (state, { data }) => stageAdapter.removeOne(data.id, state)),
);
