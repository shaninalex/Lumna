import {Component, DestroyRef, inject} from '@angular/core';
import { SidebarComponent, HeaderComponent } from '@widgets';
import {Actions, ofType} from '@ngrx/effects';
import {actionToggleSidebar} from '@core';
import {takeUntilDestroyed} from '@angular/core/rxjs-interop';
import {NgClass} from '@angular/common';

@Component({
    selector: 'lu-main-layout',
    imports: [SidebarComponent, HeaderComponent, NgClass],
    styleUrl: './main.layout.css',
    standalone: true,
    template: `
        <div class="dashboard" [ngClass]="{ 'sidebar-closed': hideSidebar }">
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
    hideSidebar = false;

    private actions$ = inject(Actions);
    private ref = inject(DestroyRef);

    constructor() {
        this.actions$
            .pipe(
                ofType(actionToggleSidebar),
                takeUntilDestroyed(this.ref),
            ).subscribe(() => this.hideSidebar = !this.hideSidebar);
    }
}
