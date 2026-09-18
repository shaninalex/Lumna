import type { EntityState } from "@ngrx/entity";
import { createEntityAdapter } from "@ngrx/entity";
import { createReducer, on } from "@ngrx/store";
import type { TaskModel } from "./task.model";
import { actionTask } from "./task.actions";

export type TaskState = EntityState<TaskModel>
export const taskAdapter = createEntityAdapter<TaskModel>({
    // NOTE: each task can be in multiple boards, so global sorter is unusable
    // sortComparer: (a, b) => a.order - b.order,
});
const initialState = taskAdapter.getInitialState();

export const taskReducer = createReducer(
    initialState,
    on(actionTask.getListSuccess, (state, {tasks}) =>
        taskAdapter.upsertMany(tasks, state)
    ),
    on(actionTask.createSuccess, (state, {task}) =>
        taskAdapter.addOne(task, state)
    ),
    on(actionTask.setTask, (state, {task}) =>
        taskAdapter.upsertOne(task, state)
    ),
    on(actionTask.assignTaskSuccess, (state, {action}) => {
        const _t = state.entities[action.task_id]
        if (!_t) {
            return state
        }
        const assignees = [..._t.assignees]
        const idx = assignees.indexOf(action.identity_id)
        if (idx > -1) {
            assignees.splice(idx, 1);
        } else {
            assignees.push(action.identity_id);
        }
        return taskAdapter.updateOne({id: action.task_id, changes: {assignees}}, state)
    })
);
