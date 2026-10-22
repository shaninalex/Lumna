import { createActionGroup, props } from '@ngrx/store';
import { Invitation } from './invitation.model';
import { Error } from '@shared/models';

export const actionInvitation = createActionGroup({
    source: 'Invitation',
    events: {
        'get list': props<{ workspaceId: number }>(),
        'set list': props<{ invitations: Invitation[] }>(),

        'create': props<{ workspaceId: number, data: { email: string; role: string } }>(),
        'set': props<{ invitation: Invitation }>(),

        'delete': props<{ invitationId: number }>(),
        'delete success': props<{ invitationId: number }>(),
        'delete failed': props<{ errors: Error[] }>(),

        'revoke': props<{ invitationId: number }>(),
        'revoke success': props<{ invitation: Invitation }>(),
        'revoke failed': props<{ errors: Error[] }>(),
    },
});
