import { Component, inject, input } from '@angular/core';
import { Router } from '@angular/router';
import { TaskDetailViewView } from '@features';
import { AppRoutes } from '@core';
import { ModalLayout } from '@core/layout';

@Component({
    selector: 'lu-task-detail-page',
    imports: [TaskDetailViewView, ModalLayout],
    template: `
        <lu-modal-layout (closed)="close()">
            <div class="card">
                <div class="card-body">
                    <lu-task-detail-view [taskId]="taskId()"/>
                </div>
            </div>
        </lu-modal-layout>
    `,
})
export class TaskDetailPage {
    taskId = input.required({ transform: (id: string) => Number(id) });

    private readonly router = inject(Router);

    private readonly appRoutes = inject(AppRoutes);

    close(): void {
        this.router.navigateByUrl(this.appRoutes.closeModal());
    }
}
