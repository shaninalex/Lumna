import type { EntityState } from '@ngrx/entity';
import { createEntityAdapter } from '@ngrx/entity';
import { createReducer, on } from '@ngrx/store';
import type { BoardModel } from './board.model';
import { actionBoard } from './board.actions';

export type BoardState = EntityState<BoardModel>;
export const boardAdapter = createEntityAdapter<BoardModel>();
const initialState = boardAdapter.getInitialState();

export const boardReducer = createReducer(
    initialState,
    on(actionBoard.setList, (state, { boards }) => boardAdapter.addMany(boards, state)),
    on(actionBoard.createSuccess, (state, { board }) => boardAdapter.upsertOne(board, state)),
);
