import { createActionGroup, emptyProps, props } from '@ngrx/store';
import type { UserModel } from '@entities/user';

export const actionSession = createActionGroup({
    source: 'Task',
    events: {
        'start authenticate': props<{ email: string; password: string }>(),
        'authenticated successfull': props<{ user: UserModel }>(),
        authenticated: emptyProps(),
        'logging out': emptyProps(),
        'logged out': emptyProps(),
    },
});
