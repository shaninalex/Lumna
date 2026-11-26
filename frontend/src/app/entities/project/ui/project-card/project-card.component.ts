import { Component, inject, Input } from '@angular/core';
import { RouterLink } from '@angular/router';
import { MatCardModule } from '@angular/material/card';
import type { ProjectModel } from '@entities/project/model';
import { selectWorkspaces } from '@entities/workspace/model';
import { Store } from '@ngrx/store';

@Component({
    selector: 'lu-project-card',
    imports: [RouterLink, MatCardModule],

    template: `
        <mat-card class="h-100">
            <mat-card-content>
                <div class="d-flex justify-content-between align-items-start mb-3">
                    <div>
                        <h5 class="mb-1">
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

                    <span class="lu-tag lu-tag--success">Active</span>
                </div>

                <p class="text-muted">
                    Placeholder description for the project. Brief summary of what this project is
                    about.
                </p>

                <div class="d-flex justify-content-between text-muted small">
                    <span>12 members</span>
                    <span>245 tasks</span>
                </div>
            </mat-card-content>
        </mat-card>
    `,
})
export class ProjectCardComponent {
    @Input() project: ProjectModel;

    private store = inject(Store);
    currentWorkspaceId = this.store.selectSignal(selectWorkspaces.currentWorkspaceId);
}
