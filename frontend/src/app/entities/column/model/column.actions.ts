import { createActionGroup, props } from '@ngrx/store';
import { ColumnPayloadModel, ColumnModel, ColumnDeleteModel, ColumnDeleteResponseModel } from './column.model';
import type { Error } from '@shared/models';

export const actionsColumns = createActionGroup({
    source: 'Column',
    events: {
        'load by board id': props<{ board_id: number }>(),
        'load by board id success': props<{ columns: ColumnModel[] }>(),
        'load by board id failed': props<{ errors: Error[] }>(),
        create: props<{ payload: ColumnPayloadModel }>(),
        'create success': props<{ column: ColumnModel }>(),
        'create failed': props<{ errors: Error[] }>(),
        'reorder failed': props<{ errors: Error[] }>(),
        'delete stage': props<{ data: ColumnDeleteModel }>(),
        'delete stage success': props<{ data: ColumnDeleteResponseModel }>(),
        'delete stage failed': props<{ errors: Error[] }>(),
    },
});
