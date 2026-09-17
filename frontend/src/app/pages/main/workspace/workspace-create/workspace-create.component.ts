import { Component } from '@angular/core';
import { WorkspaceCreateFeature } from '@features/workspace';
import { GlobalLayout } from '@core/layout';

@Component({
    selector: 'lu-workspace-create-page',
    imports: [GlobalLayout, WorkspaceCreateFeature],
    template: `
        <lu-global-layout>
            <div class="container py-4">
                <div class="d-flex justify-content-between align-items-center mb-4">
                    <h2 class="mb-1">Create Workspace</h2>
                </div>
                <lu-workspace-create-feature />
            </div>
        </lu-global-layout>
    `,
})
export class WorkspaceCreateComponent {}
