import { createActionGroup, props } from '@ngrx/store';
import { TaskAssignModel, TaskCreateModel, TaskEditModel, TaskListQueryModel, TaskModel } from './task.model';
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
        'update task': props<{ data: TaskEditModel }>(),
        'update failed': props<{ errors: Error[] }>(),
        'assign task': props<{ action: TaskAssignModel }>(),
        'assign task success': props<{ action: TaskAssignModel }>(),
        'assign task failed': props<{ errors: Error[] }>(),
        'delete task': props<{ taskId: number }>(),
        'delete task success': props<{ taskId: number }>(),
    },
});
