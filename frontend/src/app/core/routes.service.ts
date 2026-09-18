import { inject, Injectable } from '@angular/core';
import { selectProjects } from '@entities/project';
import { selectWorkspaces } from '@entities/workspace';
import { Store } from '@ngrx/store';


@Injectable()
export class AppRoutes {
    private readonly store = inject(Store);

    private readonly workspaceId = this.store.selectSignal(selectWorkspaces.currentWorkspaceId);

    private readonly projectId = this.store.selectSignal(selectProjects.currentProjectId);

    private projectRoute(): unknown[] {
        return ['/app', 'w', this.workspaceId(), 'p', this.projectId()];
    }

    projectsCreate(): unknown[] {
        return ['/app', 'w', this.workspaceId(), 'projects', 'create']
    }

    backlog(): unknown[] {
        return [...this.projectRoute(), 'backlog'];
    }

    boards(): unknown[] {
        return [...this.projectRoute(), 'boards'];
    }

    board(id: number): unknown[] {
        return [...this.projectRoute(), 'board', id];
    }

    boardsCreate(): unknown[] {
        return [...this.projectRoute(), 'board', 'create'];
    }

    createTask(): unknown[] {
        return [...this.projectRoute(), 'task', 'create'];
    }

    task(boardId: number, taskId: number): unknown[] {
        return [...this.board(boardId), 'task', taskId];
    }
}
