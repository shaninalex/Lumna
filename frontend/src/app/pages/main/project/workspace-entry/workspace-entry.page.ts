import { Component } from '@angular/core';
import { MainLayout } from '@core/layout';

@Component({
    selector: 'lu-workspace-entry-page',
    imports: [MainLayout],
    template: `
        <lu-main-layout>
            <h1>Workspace Entry</h1>
        </lu-main-layout>
    `,
})
export class WorkspaceEntryPage {}
