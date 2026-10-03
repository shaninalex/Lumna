import { EntityState, createEntityAdapter } from '@ngrx/entity';
import { createReducer, on } from '@ngrx/store';
import { Member } from './member.model';
import { actionMember } from './member.actions';

export type MemberState = EntityState<Member>;
export const memberAdapter = createEntityAdapter<Member>({
    sortComparer: (a, b) => new Date(b.dateJoined).getTime() - new Date(a.dateJoined).getTime(),
});
const initialState = memberAdapter.getInitialState();

export const memberReducer = createReducer(
    initialState,
    on(actionMember.setList, (state, { members }) => memberAdapter.addMany(members, state)),
);
