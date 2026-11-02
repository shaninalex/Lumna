import { Component } from '@angular/core';
import { WorkspaceRemoveFeature } from '@features';

@Component({
    imports: [WorkspaceRemoveFeature],
    selector: 'lu-workspace-remove',
    template: `
        <lu-workspace-remove-feature/>`,
})
export class WorkspaceRemovePage {
}
