import { createFeatureSelector, createSelector } from '@ngrx/store';
import { invitationAdapter, InvitationState } from './invitation.store';

const feature = createFeatureSelector<InvitationState>("invitation");
const entitySelectors = invitationAdapter.getSelectors();

const selectAll = createSelector(feature, entitySelectors.selectAll);

export const selectInvitation = {
    all: selectAll,
    byWorkspaceId: (workspaceId: number) => createSelector(
        selectAll,
        (invitations) => invitations.filter((invitation) => invitation.workspaceId === workspaceId),
    )
}
