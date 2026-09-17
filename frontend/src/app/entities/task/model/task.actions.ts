import { createActionGroup, props } from '@ngrx/store';
import type { TaskCreateModel, TaskListQueryModel, TaskModel } from './task.model';
import type { Error } from '@shared/models';

export const actionTask = createActionGroup({
    source: 'Task',
    events: {
        'get list': props<{ query: TaskListQueryModel }>(),
        'get list success': props<{ tasks: TaskModel[] }>(),
        'get list failed': props<{ errors: Error[] }>(),
        create: props<{ data: TaskCreateModel }>(),
        'create failed': props<{ errors: Error[] }>(),
        'create success': props<{ task: TaskModel }>(),
        'set task': props<{ task: TaskModel }>(),
    },
});
