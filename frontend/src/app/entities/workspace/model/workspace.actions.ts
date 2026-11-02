import { createActionGroup, emptyProps, props } from '@ngrx/store';
import type { Error } from '@shared/models';
import type { WorkspaceCreateModel, WorkspaceModel } from './workspace.model';

export const actionWorkspace = createActionGroup({
    source: 'Workspace',
    events: {
        'get list': emptyProps(),
        'set list': props<{ list: WorkspaceModel[] }>(),
        create: props<{ data: WorkspaceCreateModel }>(),
        created: props<{ data: WorkspaceModel }>(),
        'create failed': props<{ errors: Error[] }>(),
        'data requested': props<{ initiator?: unknown }>(),
        'set current': props<{ id: number | null }>(),

        remove: props<{ id: number }>(),
        removeSuccess: props<{ id: number }>(),
        removeFailed: props<{ errors: Error[] }>(),
    },
});
