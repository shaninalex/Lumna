import { createActionGroup, emptyProps, props } from '@ngrx/store';
import type { UserModel } from './user.model';
import type { Error } from '@shared/models';

export const actionUser = createActionGroup({
    source: 'User',
    events: {
        get: emptyProps(),
        clear: emptyProps(),
        set: props<{ user: UserModel }>(),
        changePassword: props<{ password: string, passwordConfirm: string }>(),
        changePasswordSuccess: emptyProps(),
        changePasswordFailed: props<{ errors: Error[] }>(),

        update: props<{ email: string, fullName: string }>(),
        updateSuccess: props<{ user: UserModel }>(),
        updateFailed: props<{ errors: Error[] }>(),
    },
});
