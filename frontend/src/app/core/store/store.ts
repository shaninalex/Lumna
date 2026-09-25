import { routerReducer } from '@ngrx/router-store';
import { sessionReducer, SessionEffects } from './session';
import { uiReducer } from './ui';

export const rootEffects = [
    SessionEffects,
];

export const rootReducers = {
    session: sessionReducer,
    router: routerReducer,
    ui: uiReducer,
};
