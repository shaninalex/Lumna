import { createActionGroup, props } from '@ngrx/store';
import type { KanbanMoveColumn, KanbanMoveTask, KanbanTransferTask } from './kanban.models';

export const actionKanban = createActionGroup({
    source: 'Kanban',
    events: {
        dropColumn: props<{ event: KanbanMoveColumn }>(),
        moveTask: props<{ event: KanbanMoveTask }>(),
        transferTask: props<{ event: KanbanTransferTask }>(),
    },
});
