import { routerReducer } from '@ngrx/router-store';
import { SessionEffects } from './session/session.effects';
import { sessionReducer } from './session/session.store';

export const rootEffects = [
    SessionEffects,
];

export const rootReducers = {
    session: sessionReducer,
    router: routerReducer,
};
