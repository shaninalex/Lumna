import { createFeatureSelector, createSelector } from '@ngrx/store';
import { UIState } from './ui.store';

const feature = createFeatureSelector<UIState>('ui');

export const selectUI = {
    sidebarOpen: createSelector(feature, (state: UIState) => state.sidebarOpen),
};
