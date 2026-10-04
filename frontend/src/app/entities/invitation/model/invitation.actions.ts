import { createActionGroup, props } from '@ngrx/store';
import { Invitation } from './invitation.model';

export const actionInvitation = createActionGroup({
    source: 'Invitation',
    events: {
        'get list': props<{ workspaceId: number }>(),
        'set list': props<{ invitations: Invitation[] }>(),

        'create': props<{ workspaceId: number, data: { email: string; role: string } }>(),
        'set': props<{ invitation: Invitation }>(),
    },
});
