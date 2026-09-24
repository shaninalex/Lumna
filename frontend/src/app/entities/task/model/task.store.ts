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
        const task = state.entities[action.taskId];
        if (!task) {
            return state;
        }

        const assignees = new Set(task.assignees ?? []);
        if (assignees.has(action.identityId)) {
            assignees.delete(action.identityId);
        } else {
            assignees.add(action.identityId);
        }

        return taskAdapter.updateOne({id: action.taskId, changes: {assignees: [...assignees]}}, state);
    }),
    on(actionTask.deleteTaskSuccess, (state, {taskId}) => taskAdapter.removeOne(taskId, state))
);
