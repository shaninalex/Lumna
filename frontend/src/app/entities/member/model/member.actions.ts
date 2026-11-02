import { createActionGroup, props } from '@ngrx/store';
import { Member } from './member.model';
import type { Error } from '@shared/models';

export const actionMember = createActionGroup({
    source: 'Member',
    events: {
        getList: props<{ workspaceId: number }>(),
        setList: props<{ members: Member[] }>(),

        removeMember: props<{ workspaceId: number, identityId: number }>(),
        removeMemberFailed: props<{ errors: Error[] }>(),
        removeMemberSuccess: props<{ workspaceId: number, identityId: number }>(),
    },
});
