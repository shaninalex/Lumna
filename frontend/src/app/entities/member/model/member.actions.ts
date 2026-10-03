import { createActionGroup, props } from '@ngrx/store';
import { Member } from './member.model';

export const actionMember = createActionGroup({
    source: 'Member',
    events: {
        'get list': props<{ workspaceId: number }>(),
        'set list': props<{ members: Member[] }>(),
    },
});
