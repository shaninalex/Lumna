import { Component, inject, Input } from '@angular/core';
import { RouterLink } from '@angular/router';
import { StripHtmlPipe, TrimPipe } from '@shared/utils';
import { DatePipe } from '@angular/common';
import { AppRoutes } from '@core';
import { CdkMenu, CdkMenuItem, CdkMenuTrigger } from '@angular/cdk/menu';

import { actionTask, TaskModel } from '../../model';
import { Store } from '@ngrx/store';


@Component({
    selector: 'lu-task-card',
    imports: [RouterLink, TrimPipe, DatePipe, CdkMenu, CdkMenuItem, CdkMenuTrigger, StripHtmlPipe],
    styleUrl: './task-card.component.css',
    template: `
        <div class="card text-decoration-none text-body">
            <div class="card-body p-2">
                <!--
                    <div class="d-flex justify-content-between mb-2">
                        <span class="badge text-bg-primary"> Feature </span>

                        <small class="text-muted"> #FEAT-212 </small>
                    </div>
                -->
                <div class="d-flex justify-content-between mb-2 align-items-start">
                    <h6 class="flex-grow-1 mb-0">
                        <a class="task-card-link" [routerLink]="routeService.task(task.board_id, task.id)">
                            {{ task.title }}
                        </a>
                    </h6>
                    <button [cdkMenuTriggerFor]="taskMenu" class="btn btn-sm p-1 lh-1">
                        <i class="fa-solid fa-ellipsis"></i>
                    </button>
                </div>

                @if (task.body) {
                    <p class="small text-muted mb-2">{{ task.body | strip_html | trim: 65 }}</p>
                }

                <div class="d-flex justify-content-between align-items-center">
                    <div class="d-flex gap-2 align-items-start">
                        <!-- Slot for feature actions (assignment, status change, etc.) -->
                        <ng-content select="[assignmentSlot]"/>

                        @if (task.due_to) {
                            <span class="badge rounded-pill text-bg-secondary">
                                <i class="fa-regular fa-calendar"></i>
                                {{ task.due_to | date: 'd MMM' }}
                            </span>
                        }
                    </div>
                    @if (task.body !== '') {
                        <i class="fa-solid fa-align-left"></i>
                    }
                </div>
            </div>
        </div>

        <ng-template #taskMenu>
            <div class="list-group" cdkMenu>
                <button cdkMenuItem type="button" class="list-group-item list-group-item-action" (click)="delete()">
                    Delete
                </button>
            </div>
        </ng-template>
    `,
})
export class TaskCardComponent {
    @Input({required: true}) task: TaskModel
    readonly routeService = inject(AppRoutes);
    private store = inject(Store);

    delete(): void {
        this.store.dispatch(actionTask.deleteTask({taskId: this.task.id}))
    }
}
