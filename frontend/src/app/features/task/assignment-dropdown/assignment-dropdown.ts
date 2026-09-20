import { Component, inject, Input } from '@angular/core';
import { CdkMenu, CdkMenuItem, CdkMenuTrigger } from '@angular/cdk/menu';
import { Store } from '@ngrx/store';
import { actionTask } from '@entities/task/model/task.actions';
import { TaskModel } from '@entities/task/model';
import { selectUser } from '@entities/user';
import { filter } from 'rxjs';
import { AsyncPipe } from '@angular/common';

@Component({
    selector: 'lu-assignment-dropdown',
    imports: [CdkMenu, CdkMenuItem, CdkMenuTrigger, AsyncPipe],
    template: `
        <div class="text-muted d-inline-flex align-items-center lh-1" [cdkMenuTriggerFor]="assignmentMenu">
            @if (task.assignees !== null && task.assignees.length > 0) {
                <img src="images/7.png" alt="" class="rounded-circle" style="width: 16px">
            } @else {
                <i class="fa-regular fa-circle-user"></i>
            }
        </div>
        <ng-template #assignmentMenu>
            <div class="card" cdkMenu>
                @if (user$ | async; as user) {
                    <button
                        class="btn btn-outline-secondary border-0 border-bottom text-start"
                        cdkMenuItem
                        (click)="assign(task.id, user.id)"
                    >
                        <img src="images/7.png" alt="" class="rounded-circle" style="width: 16px">
                        Assign yourself
                    </button>
                }
                <!-- TODO: users list -->
                <input type="text" class="form-control border-0" placeholder="Search...">
                <div class="list-group list-group-flush">
                    <button type="button" class="list-group-item list-group-item-action" cdkMenuItem>
                        <img src="images/7.png" alt="" class="rounded-circle" style="width: 16px">
                        Garry Hudini
                    </button>
                    <button type="button" class="list-group-item list-group-item-action" cdkMenuItem>
                        <img src="images/7.png" alt="" class="rounded-circle" style="width: 16px">
                        Garry Hudini
                    </button>
                    <button type="button" class="list-group-item list-group-item-action" cdkMenuItem>
                        <img src="images/7.png" alt="" class="rounded-circle" style="width: 16px">
                        Garry Hudini
                    </button>
                    <button type="button" class="list-group-item list-group-item-action" cdkMenuItem>
                        <img src="images/7.png" alt="" class="rounded-circle" style="width: 16px">
                        Garry Hudini
                    </button>
                    <button type="button" class="list-group-item list-group-item-action" cdkMenuItem>
                        <img src="images/7.png" alt="" class="rounded-circle" style="width: 16px">
                        Garry Hudini
                    </button>
                </div>
            </div>
        </ng-template>
    `,
})
export class AssignmentDropdown {
    @Input() task: TaskModel;
    private store = inject(Store);
    readonly user$ = this.store.select(selectUser.user).pipe(filter(user => !!user));

    assign(taskId: number, identityId: number) {
        this.store.dispatch(actionTask.assignTask({action: {task_id: taskId, identity_id: identityId}}))
    }
}
