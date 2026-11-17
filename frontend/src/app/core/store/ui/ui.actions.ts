import { createActionGroup, props } from '@ngrx/store';

export const actionUI = createActionGroup({
    source: 'UI',
    events: {
        sidebarState: props<{ state: boolean }>(),
        taskColumnForm: props<{ columnId: number; position: 'top' | 'bottom' }>(),
    },
});
