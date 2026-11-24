import { createActionGroup, emptyProps, props } from '@ngrx/store';

export const actionSession = createActionGroup({
    source: 'Task',
    events: {
        startAuthenticate: props<{ email: string; password: string }>(),
        authenticatedSuccessful: emptyProps(),
        authenticated: emptyProps(),
        loggingOut: emptyProps(),
        loggedOut: emptyProps(),
    },
});
