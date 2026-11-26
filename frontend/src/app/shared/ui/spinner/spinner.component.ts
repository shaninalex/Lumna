import { Component, input } from '@angular/core';
import { MatProgressSpinner } from '@angular/material/progress-spinner';

@Component({
    imports: [MatProgressSpinner],
    selector: 'lu-spinner',
    template: `
        @if (loading()) {
            <mat-progress-spinner mode="indeterminate" [diameter]="diameter()" />
        }`
})
export class SpinnerComponent {
    loading = input.required();
    diameter = input(18);
}
