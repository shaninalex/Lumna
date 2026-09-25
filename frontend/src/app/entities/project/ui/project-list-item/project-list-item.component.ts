import type { OnInit } from '@angular/core';
import { Component, computed, inject, Input, ChangeDetectionStrategy } from '@angular/core';
import type { ProjectModel } from '@entities/project/model';
import { selectWorkspaces } from '@entities/workspace/model';
import { Store } from '@ngrx/store';
import { RouterLink } from '@angular/router';
import { TrimPipe } from '@shared/utils';
import { selectTasks } from '@entities/task/model/task.selectors';
import type { Observable } from 'rxjs';
import { AsyncPipe } from '@angular/common';

@Component({
    selector: 'lu-project-list-item',
    imports: [RouterLink, TrimPipe, AsyncPipe],
    changeDetection: ChangeDetectionStrategy.Eager,
    template: `
        <a
            cdkMenuItem
            [routerLink]="['/app/w', currentWorkspaceId() || '', 'p', project.id]"
            class="d-flex justify-content-between align-items-center gap-2"
        >
            <div class="project-icon">
                {{ project.title | trim: 1 }}
            </div>
            <div class="me-auto">
                {{ project.title }}
            </div>

            @if (tasksCount$ | async; as count) {
                <span class="badge text-bg-primary rounded-pill">{{ count }}</span>
            }
        </a>
    `,
})
export class ProjectListItemComponent implements OnInit {
    @Input() project: ProjectModel;
    private store = inject(Store);

    currentWorkspaceId = this.store.selectSignal(selectWorkspaces.currentWorkspaceId);
    tasksCount$: Observable<number>;

    ngOnInit(): void {
        computed(() => {
            const i = this.currentWorkspaceId();
            if (i && i > 0) {
                this.tasksCount$ = this.store.select(selectTasks.countByProjectId(i));
            }
        });
    }
}
