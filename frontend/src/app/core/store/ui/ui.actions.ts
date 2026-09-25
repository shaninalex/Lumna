import { createActionGroup, props } from '@ngrx/store';

export const actionUI = createActionGroup({
    source: 'UI',
    events: {
        'sidebar state': props<{ state: boolean }>(),
        'task column form': props<{ columnId: number; position: 'top' | 'bottom' }>(),
    },
});
