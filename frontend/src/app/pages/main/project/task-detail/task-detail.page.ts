import { Component, inject, input } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AppRoutes } from '@core';
import { ModalLayout } from '@core/layout';
import { TaskDetailViewView } from '@entities/task';

@Component({
    selector: 'lu-task-detail-page',
    imports: [TaskDetailViewView, ModalLayout],
    template: `
        <lu-modal-layout (closed)="close()">
            <lu-task-detail-view [taskId]="taskId()"/>
        </lu-modal-layout>
    `,
})
export class TaskDetailPage {
    private readonly router = inject(Router);
    private readonly activatedRoute = inject(ActivatedRoute);
    private readonly appRoutes = inject(AppRoutes);

    taskId = input.required({transform: (id: string) => Number(id)});

    close(): void {
        const params = this.activatedRoute.parent?.snapshot.params;
        if (params && params['boardId'] !== undefined) {
            this.router.navigate(this.appRoutes.board(Number(params['boardId'])))
        }
    }
}
