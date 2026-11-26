import { Component, DestroyRef, inject, linkedSignal } from '@angular/core';
import { MatButton } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInput } from '@angular/material/input';
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
    imports: [FormField, MatButton, MatFormFieldModule, MatInput],
    template: `
        <form (submit)="onSubmit($event)">
            <mat-form-field class="w-100 mb-3">
                <mat-label>Board title</mat-label>
                <input matInput type="text" id="board_title" [formField]="pForm.title" />
            </mat-form-field>

            <div>
                <button matButton="filled" type="submit">Create</button>
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

    pFormModel = linkedSignal<number | null | undefined, BoardPayloadModel>({
        source: this.currentProjectId,
        computation: (projectId, previous) => ({
            title: previous?.value.title ?? '',
            projectId: projectId ?? 0,
        }),
    });
    pForm = form(this.pFormModel, (schemaPath) => required(schemaPath.title));

    constructor() {
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
