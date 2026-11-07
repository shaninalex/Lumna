import { Component, input } from '@angular/core';
import { reiconRefresh } from '@ng-icons/reicon';
import { NgIcon, provideIcons } from '@ng-icons/core';

@Component({
    imports: [NgIcon],
    providers: [provideIcons({reiconRefresh})],
    selector: 'lu-spinner',
    styleUrl: "./spinner.component.scss",
    template: `
        @if (loading()) {
            <div class="spinner">
                <ng-icon name="reiconRefresh"/>
            </div>
        }`
})
export class SpinnerComponent {
    loading = input.required();
}
