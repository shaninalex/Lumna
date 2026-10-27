import { createReducer, on } from '@ngrx/store';
import { createEntityAdapter, EntityState } from '@ngrx/entity';
import { RefreshToken } from './refresh_token.model';
import { actionRefreshToken } from './refresh_token.actions';

export type RefreshTokenState = EntityState<RefreshToken>;
export const refreshTokenAdapter = createEntityAdapter<RefreshToken>({
    sortComparer: (a, b) => new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime(),
})
const initialState = refreshTokenAdapter.getInitialState();
export const refreshTokenReducer = createReducer(
    initialState,
    on(actionRefreshToken.setList, (state, {list}) => refreshTokenAdapter.addMany(list, state)),
    on(actionRefreshToken.removeOne, (state, {tokenId}) => refreshTokenAdapter.removeOne(tokenId, state)),
    on(actionRefreshToken.updateOne, (state, {tokenId, updates}) => refreshTokenAdapter.updateOne({
        id: tokenId,
        changes: updates
    }, state))
)

