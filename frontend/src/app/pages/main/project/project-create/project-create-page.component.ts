import { Component, inject, ChangeDetectionStrategy } from '@angular/core';
import { ProjectCreateFeature } from '@features/project';
import { GlobalLayout } from '@core/layout';
import { UiService } from '@shared/ui';
import { AppRoutes } from '@core';

@Component({
    selector: 'lu-project-create-page',
    imports: [GlobalLayout, ProjectCreateFeature],
    changeDetection: ChangeDetectionStrategy.Eager,
    template: `
        <lu-global-layout>
            <div class="container py-4">
                <div class="d-flex justify-content-between align-items-center mb-4">
                    <h2 class="mb-1">Create Project</h2>
                </div>

                <lu-project-create-feature />
            </div>
        </lu-global-layout>
    `,
})
export class ProjectCreatePage {
    private ui = inject(UiService);
    readonly appRoutes = inject(AppRoutes);

    constructor() {
        this.ui.setPageTitle('Create Project');
    }
}
