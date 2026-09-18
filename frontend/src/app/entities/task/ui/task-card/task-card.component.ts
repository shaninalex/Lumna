import { Component, inject, Input } from '@angular/core';
import { RouterLink } from '@angular/router';
import type { TaskModel } from '@entities/task/model';
import { TrimPipe } from '@shared/utils';
import { DatePipe } from '@angular/common';
import { AppRoutes } from '@core';
import { AssignmentDropdown } from '../assignment-dropdown';

@Component({
    selector: 'lu-task-card',
    imports: [RouterLink, TrimPipe, DatePipe, AssignmentDropdown],
    template: `
        <div class="card text-decoration-none text-body">
            <div class="card-body">
                <!--
                    <div class="d-flex justify-content-between mb-2">
                        <span class="badge text-bg-primary"> Feature </span>

                        <small class="text-muted"> #FEAT-212 </small>
                    </div>
                -->
                <h6 class="mb-2">
                    <a [routerLink]="routeService.task(task.board_id, task.id)">
                        {{ task.title }}
                    </a>
                </h6>

                @if (task.body !== '') {
                    <p class="small text-muted mb-3">{{ task.body | trim: 65 }}</p>
                }
                <div class="d-flex justify-content-between align-items-center">
                    <div class="d-flex gap-2 align-items-start">
                        <lu-assignment-dropdown [task]="task" />
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
    `,
})
export class TaskCardComponent {
    @Input({ required: true }) task: TaskModel

    readonly routeService = inject(AppRoutes);
}
