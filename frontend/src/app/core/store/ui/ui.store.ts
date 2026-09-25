import { createReducer, on } from '@ngrx/store';
import { actionUI } from './ui.actions';

export interface UIState {
    sidebarOpen: boolean;
}

const initialState: UIState = {
    sidebarOpen: true,
};

export const uiReducer = createReducer(
    initialState,
    on(actionUI.sidebarState, (state, action) => ({sidebarOpen: action.state})),
);
