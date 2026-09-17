import { inject, Injectable } from '@angular/core';
import { selectProjects } from '@entities/project';
import { selectWorkspaces } from '@entities/workspace';
import { Store } from '@ngrx/store';
import { ActivatedRoute, PRIMARY_OUTLET, Router, UrlTree } from '@angular/router';

/**
 * Path of the route the `modal` outlet routes are declared under. Auxiliary routes are addressed
 * relative to their parent, so links that open a modal have to be built relative to this route.
 */
export const PROJECT_ROUTE_PATH = 'p/:projectId';

/** Name of the outlet declared in the root component for routes rendered on top of a page. */
export const MODAL_OUTLET = 'modal';

@Injectable()
export class AppRoutes {
    private readonly store = inject(Store);
    private readonly router = inject(Router);
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

    /**
     * Task detail lives in the `modal` outlet, so it cannot be addressed by a plain path — it has
     * to be added as an auxiliary segment next to whatever page is currently open, e.g.
     * `/app/w/1/p/2/(board/3//modal:task/7)`.
     */
    task(id: number): UrlTree {
        return this.modalRoute(['task', id]);
    }

    closeModal(): UrlTree {
        return this.modalRoute(null);
    }

    editTask(id: string): unknown[] {
        return [...this.projectRoute(), 'task', id, 'edit'];
    }

    private modalRoute(commands: unknown[] | null): UrlTree {
        const relativeTo = this.activeProjectRoute();

        if (!relativeTo) {
            return this.router.createUrlTree([
                ...this.projectRoute(),
                { outlets: { [MODAL_OUTLET]: commands } },
            ]);
        }

        return this.router.createUrlTree(
            [{ outlets: { [MODAL_OUTLET]: commands } }],
            { relativeTo },
        );
    }

    private activeProjectRoute(): ActivatedRoute | null {
        let route: ActivatedRoute | undefined = this.router.routerState.root;

        while (route) {
            if (route.routeConfig?.path === PROJECT_ROUTE_PATH) {
                return route;
            }

            route = route.children.find((child) => child.outlet === PRIMARY_OUTLET);
        }

        return null;
    }
}
