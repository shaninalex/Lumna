import { Component, inject } from '@angular/core';
import { SidebarComponent, HeaderComponent } from '@widgets';
import { Store } from '@ngrx/store';
import { selectUI } from '@core/store/ui';

@Component({
    selector: 'lu-main-layout',
    imports: [SidebarComponent, HeaderComponent],
    styleUrl: './main.layout.scss',
    template: `
        <div class="dashboard" [class.sidebar-closed]="!sidebarOpen()">
            <div class="dashboard-header">
                <lu-header/>
            </div>
            <div class="dashboard-sidebar">
                <lu-sidebar/>
            </div>
            <div class="dashboard-content">
                <ng-content/>
            </div>
        </div>
    `,
})
export class MainLayout {
    private store = inject(Store);
    sidebarOpen = this.store.selectSignal(selectUI.sidebarOpen);
}
