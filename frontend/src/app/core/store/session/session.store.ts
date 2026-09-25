import { createReducer, on } from '@ngrx/store';
import { actionSession } from './session.actions';

type SessionState = {
    authenticated: boolean;
};

const initialState: SessionState = {
    authenticated: false,
};

export const sessionReducer = createReducer(
    initialState,
    on(
        actionSession.authenticatedSuccessfull,
        actionSession.authenticated,
        () => ({ authenticated: true }),
    ),
    on(actionSession.loggedOut, () => ({ authenticated: false })),
);
