import { Component, inject, Input, ChangeDetectionStrategy } from '@angular/core';
import { RouterLink } from '@angular/router';
import type { ProjectModel } from '@entities/project/model';
import { selectWorkspaces } from '@entities/workspace/model';
import { Store } from '@ngrx/store';

@Component({
    selector: 'lu-project-card',
    imports: [RouterLink],
    changeDetection: ChangeDetectionStrategy.Eager,
    template: `
        <div class="card h-100">
            <div class="card-body">
                <div class="d-flex justify-content-between align-items-start mb-3">
                    <div>
                        <h5 class="card-title mb-1">
                            <a
                                [routerLink]="[
                                    '/app/w',
                                    currentWorkspaceId() || '',
                                    'p',
                                    project.id,
                                ]"
                            >
                                {{ project.title }}
                            </a>
                        </h5>

                        <div class="text-muted small">WEB</div>
                    </div>

                    <span class="badge text-bg-success"> Active </span>
                </div>

                <p class="card-text text-muted">
                    Placeholder description for the project. Brief summary of what this project is
                    about.
                </p>

                <div class="d-flex justify-content-between text-muted small">
                    <span>12 members</span>
                    <span>245 tasks</span>
                </div>
            </div>
        </div>
    `,
})
export class ProjectCardComponent {
    @Input() project: ProjectModel;

    private store = inject(Store);
    currentWorkspaceId = this.store.selectSignal(selectWorkspaces.currentWorkspaceId);
}
