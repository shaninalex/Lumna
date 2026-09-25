import { createFeature, createReducer, on } from '@ngrx/store';
import { actionUI } from './ui.effects';

export interface UIState {
    sidebar: boolean;
}

const initialState: UIState = {
    sidebar: true,
};

export const uiReducer = createReducer(
    initialState,
    on(actionUI.sidebarState, (state, action) => ({sidebar: action.state})),
);

export const uiFeature = createFeature({
    name: 'sidebar',
    reducer: uiReducer,
});
