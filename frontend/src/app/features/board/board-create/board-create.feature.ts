import { Component, DestroyRef, inject, signal, ChangeDetectionStrategy } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { form, FormField, required } from '@angular/forms/signals';
import { Router } from '@angular/router';
import { AppRoutes } from '@core';
import type { BoardPayloadModel } from '@entities/board';
import { actionBoard } from '@entities/board';
import { selectProjects } from '@entities/project';
import { Actions, ofType } from '@ngrx/effects';
import { Store } from '@ngrx/store';
import { tap } from 'rxjs';

@Component({
    selector: 'lu-board-create-feature',
    imports: [FormField],
    changeDetection: ChangeDetectionStrategy.Eager,
    template: `
        <form (submit)="onSubmit($event)">
            <div class="mb-4">
                <label for="board_title" class="form-label">Board title</label>
                <input
                    type="text"
                    class="form-control"
                    id="board_title"
                    [formField]="pForm.title"
                />
            </div>

            <div>
                <button class="btn btn-primary" type="submit">Create</button>
            </div>
        </form>
    `,
})
export class BoardCreateFeature {
    private store = inject(Store);
    private actions$ = inject(Actions);
    private destroyRef = inject(DestroyRef);
    private router = inject(Router);
    private appRoutes = inject(AppRoutes);
    private currentProjectId = this.store.selectSignal(selectProjects.currentProjectId);

    pFormModel = signal<BoardPayloadModel>({ title: '', projectId: 0 });
    pForm = form(this.pFormModel, (schemaPath) => required(schemaPath.title));

    constructor() {
        const _currentProjectId = this.currentProjectId();
        if (_currentProjectId) this.pFormModel().projectId = _currentProjectId;

        this.actions$
            .pipe(ofType(actionBoard.createFailed), takeUntilDestroyed(this.destroyRef))
            .subscribe((action) => console.log(action));
        this.actions$
            .pipe(
                ofType(actionBoard.createSuccess),
                takeUntilDestroyed(this.destroyRef),
                tap(() => {
                    this.router.navigate(this.appRoutes.boards());
                }),
            )
            .subscribe();
    }

    onSubmit(event: Event): void {
        event.preventDefault();
        this.store.dispatch(actionBoard.create({ data: this.pFormModel() }));
    }
}
