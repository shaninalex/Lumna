import { createFeatureSelector, createSelector } from '@ngrx/store';
import { memberAdapter, MemberState } from './member.store';

const feature = createFeatureSelector<MemberState>('member');
const entitySelectors = memberAdapter.getSelectors();

const selectAll = createSelector(feature, entitySelectors.selectAll);

export const selectMember = {
    all: selectAll,
    entities: createSelector(feature, entitySelectors.selectEntities),
    byWorkspaceId: (workspaceId: number) =>
        createSelector(selectAll, (members) => members.filter((b) => b.workspaceId === workspaceId)),
};
