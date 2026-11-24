import { createActionGroup, props } from '@ngrx/store';
import type { BoardModel, BoardPayloadModel } from './board.model';
import type { Error } from '@shared/models';

export const actionBoard = createActionGroup({
    source: 'Board',
    events: {
        getList: props<{ projectId: number }>(),
        setList: props<{ boards: BoardModel[] }>(),

        create: props<{ data: BoardPayloadModel }>(),
        createSuccess: props<{ board: BoardModel }>(),
        createFailed: props<{ errors: Error[] }>(),
    },
});
