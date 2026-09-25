import { Component, effect, inject, input } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AppRoutes } from '@core';
import { ModalLayout } from '@core/layout';
import { TaskDetailView } from '@widgets';
import { UiService } from '@shared/ui';
import { Store } from '@ngrx/store';
import { selectTasks } from '@entities/task';

@Component({
    selector: 'lu-task-detail-page',
    imports: [TaskDetailView, ModalLayout],

    template: `
        <lu-modal-layout (closed)="close()">
            <lu-task-detail-view [taskId]="taskId()" />
        </lu-modal-layout>
    `,
})
export class TaskDetailPage {
    taskId = input.required({
        transform: (id: string) => Number(id),
    });

    private store = inject(Store);
    private router = inject(Router);
    private activatedRoute = inject(ActivatedRoute);
    private appRoutes = inject(AppRoutes);
    private ui = inject(UiService);

    constructor() {
        effect(() => {
            const t = this.store.selectSignal(selectTasks.byId(this.taskId()));
            this.ui.setPageTitle(`Task: ${t()?.title}`);
        });
    }

    close(): void {
        const params = this.activatedRoute.parent?.snapshot.params;
        if (params && params['boardId'] !== undefined) {
            this.router.navigate(this.appRoutes.board(Number(params['boardId'])));
        }
    }
}
