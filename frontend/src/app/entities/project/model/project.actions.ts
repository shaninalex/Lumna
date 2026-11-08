import { createActionGroup, props } from '@ngrx/store';
import type { ProjectModel, ProjectCreateModel } from './project.model';
import type { Error } from '@shared/models';

export const actionProject = createActionGroup({
    source: 'Project',
    events: {
        create: props<{ payload: ProjectCreateModel }>(),
        createSuccessful: props<{ project: ProjectModel }>(),
        createFailed: props<{ errors: Error[] }>(),

        delete: props<{ projectId: number }>(),
        deleteSuccessful: props<{ projectId: number }>(),
        deleteFailed: props<{ errors: Error[] }>(),

        getList: props<{ workspaceId: number }>(),
        setList: props<{ projects: ProjectModel[] }>(),

        patch: props<{ id: number; data: ProjectCreateModel }>(),
        patchSuccessful: props<{ data: ProjectCreateModel }>(),
        patchFailed: props<{ errors: Error[] }>(),

        setCurrent: props<{ id: number | null }>(),
    },
});
