import { Component } from '@angular/core';
import { TaskCreateFullFeature } from '@features/task';
import { MainLayout } from '@core/layout';

@Component({
    selector: 'lu-task-create-page',
    imports: [MainLayout, TaskCreateFullFeature],
    templateUrl: './task-create.page.html',
})
export class TaskCreatePage {}
