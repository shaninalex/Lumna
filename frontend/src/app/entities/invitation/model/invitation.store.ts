import { EntityState, createEntityAdapter } from '@ngrx/entity';
import { createReducer, on } from '@ngrx/store';
import { Invitation } from './invitation.model';
import { actionInvitation } from './invitation.actions';

export type InvitationState = EntityState<Invitation>;
export const invitationAdapter = createEntityAdapter<Invitation>({
    sortComparer: (a, b) => new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime(),
});
const initialState = invitationAdapter.getInitialState();

export const invitationReducer = createReducer(
    initialState,
    on(actionInvitation.setList, (state, { invitations }) => invitationAdapter.addMany(invitations, state)),
    on(actionInvitation.set, (state, { invitation }) => invitationAdapter.addOne(invitation, state)),
);
