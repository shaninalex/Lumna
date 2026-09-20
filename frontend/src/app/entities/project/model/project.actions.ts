import { createActionGroup, props } from '@ngrx/store';
import type { ProjectModel, ProjectCreateModel } from './project.model';
import type { Error } from '@shared/models';

export const actionProject = createActionGroup({
    source: 'Project',
    events: {
        create: props<{ payload: ProjectCreateModel }>(),
        'create succefull': props<{ project: ProjectModel }>(),
        'create failed': props<{ errors: Error[] }>(),

        delete: props<{ project_id: number }>(),
        'delete succefull': props<{ projectId: number }>(),
        'delete failed': props<{ errors: Error[] }>(),

        'get list': props<{ workspaceId: number }>(),
        'set list': props<{ projects: ProjectModel[] }>(),

        patch: props<{ id: number; data: ProjectCreateModel }>(),
        'patch successfull': props<{ data: ProjectCreateModel }>(),
        'patch failed': props<{ errors: Error[] }>(),

        'set current': props<{ id: number | null }>(),
    },
});
