import { Component } from '@angular/core';
import { MainLayout } from '@core/layout';
import { MatDivider } from '@angular/material/divider';
import { ProjectEditFeature, ProjectRemoveFeature } from '@features/project';

@Component({
    imports: [MainLayout, ProjectEditFeature, ProjectRemoveFeature, MatDivider],
    selector: 'lu-project-settigns',
    template: `
        <lu-main-layout>
            <div class="container-fluid py-4">
                <lu-project-edit-feature />
                <mat-divider class="my-4" />
                <lu-project-remove-feature />
            </div>
        </lu-main-layout>`,
})
export class ProjectSettingsPage {
}
