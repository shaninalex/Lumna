import { Component } from '@angular/core';
import { MainLayout } from '@core/layout';
import { ProjectEditFeature, ProjectRemoveFeature } from '@features/project';

@Component({
    imports: [MainLayout, ProjectEditFeature, ProjectRemoveFeature],
    selector: 'lu-project-settigns',
    template: `
        <lu-main-layout>
            <div class="container-fluid py-4">
                <lu-project-edit />
                <hr>
                <lu-project-remove />
            </div>
        </lu-main-layout>`,
})
export class ProjectSettingsPage {
}
