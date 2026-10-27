import { createActionGroup, props, emptyProps } from '@ngrx/store';
import { RefreshToken } from './refresh_token.model';

export const actionRefreshToken = createActionGroup({
    source: 'RefreshToken',
    events: {
        getList: emptyProps(),
        setList: props<{ list: RefreshToken[] }>(),
        delete: props<{ tokenId: number }>(),
        removeOne: props<{ tokenId: number }>(),

        revoke: props<{ tokenId: number }>(),
        updateOne: props<{ tokenId: number, updates: Partial<RefreshToken> }>(),
    },
});
