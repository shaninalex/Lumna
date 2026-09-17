import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { provideState } from '@ngrx/store';
import { provideEffects } from '@ngrx/effects';

import { routes } from './main.routes';
import { AppRoutes } from '@core';
import { mainEffects } from './store';
import { WorkspaceApi, workspaceFeature } from '@entities/workspace';
import { ProjectApi, projectFeature } from '@entities/project';
import { UserApi, userFeature } from '@entities/user';
import { TaskApi, taskFeature } from '@entities/task';
import { ColumnApi, columnFeature } from '@entities/column';
import { BoardApi, boardFeature } from '@entities/board';
import { KanbanApi } from '@features/board/kanban-board/api';

@NgModule({
    declarations: [],
    imports: [
        CommonModule,
        RouterModule.forChild(routes),
    ],
    providers: [
        AppRoutes,
        WorkspaceApi,
        ProjectApi,
        UserApi,
        TaskApi,
        BoardApi,
        ColumnApi,
        KanbanApi,

        provideEffects(mainEffects),

        provideState(workspaceFeature),
        provideState(projectFeature),
        provideState(userFeature),
        provideState(taskFeature),
        provideState(boardFeature),
        provideState(columnFeature),
    ],
})
export class MainModule {
}
