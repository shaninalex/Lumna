import { routerReducer } from '@ngrx/router-store';
import { SessionEffects } from './session';
import { uiReducer } from './ui';

export const rootEffects = [
    SessionEffects,
];

export const rootReducers = {
    router: routerReducer,
    ui: uiReducer,
};
