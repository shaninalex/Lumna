import { Component, inject, Input } from '@angular/core';
import { RouterLink } from '@angular/router';
import { StripHtmlPipe, TrimPipe } from '@shared/utils';
import { DatePipe } from '@angular/common';
import { AppRoutes } from '@core';
import { MatIconButton } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatIcon } from '@angular/material/icon';
import { MatMenuModule } from '@angular/material/menu';

import { actionTask, TaskModel } from '../../model';
import { Store } from '@ngrx/store';

@Component({
    selector: 'lu-task-card',
    imports: [RouterLink, TrimPipe, DatePipe, StripHtmlPipe, MatCardModule, MatIconButton, MatIcon, MatMenuModule],
    styleUrl: './task-card.component.scss',
    template: `
        <mat-card>
            <mat-card-content class="p-2">
                <!--
                    <div class="d-flex justify-content-between mb-2">
                        <span class="lu-tag lu-tag--primary"> Feature </span>

                        <small class="text-muted"> #FEAT-212 </small>
                    </div>
                -->
                <div class="d-flex justify-content-between mb-2 align-items-start gap-1">
                    <h6 class="flex-grow-1 mb-0">
                        <a
                            class="task-card-link"
                            [routerLink]="routeService.task(task.boardId, task.id)"
                        >
                            {{ task.title }}
                        </a>
                    </h6>
                    <button
                        matIconButton
                        type="button"
                        class="task-menu-trigger"
                        [matMenuTriggerFor]="taskMenu"
                        aria-label="Task menu"
                    >
                        <mat-icon>menu</mat-icon>
                    </button>
                </div>

                @if (task.body) {
                    <p class="small text-muted mb-2">{{ task.body | strip_html | trim: 65 }}</p>
                }

                <div class="d-flex justify-content-between align-items-center">
                    <div class="d-flex gap-2 align-items-center">
                        <!-- Slot for feature actions (assignment, status change, etc.) -->
                        <ng-content select="[assignmentSlot]" />

                        @if (task.dueTo) {
                            <span class="lu-tag lu-tag--neutral">
                                <mat-icon>calendar_today</mat-icon>
                                {{ task.dueTo | date: 'd MMM' }}
                            </span>
                        }
                    </div>
                </div>
            </mat-card-content>
        </mat-card>

        <mat-menu #taskMenu="matMenu">
            <button mat-menu-item type="button" (click)="delete()">
                <mat-icon>delete</mat-icon>
                <span>Delete</span>
            </button>
        </mat-menu>
    `,

})
export class TaskCardComponent {
    @Input({ required: true }) task: TaskModel;
    readonly routeService = inject(AppRoutes);
    private store = inject(Store);

    delete(): void {
        this.store.dispatch(actionTask.deleteTask({ taskId: this.task.id }));
    }
}
