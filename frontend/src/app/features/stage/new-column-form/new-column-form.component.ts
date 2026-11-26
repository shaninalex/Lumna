import {
    OnInit,
    Component,
    DestroyRef,
    inject,
    Input,
    signal,
} from '@angular/core';
import { MatButton } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInput } from '@angular/material/input';
import { MatIcon } from '@angular/material/icon';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { FormField, form, required } from '@angular/forms/signals';
import { actionsColumns } from '@entities/column/model';
import { Actions, ofType } from '@ngrx/effects';
import { Store } from '@ngrx/store';
import { Error } from '@shared/models';
import { tap } from 'rxjs';

@Component({
    selector: 'lu-new-column-form',
    imports: [FormField, MatButton, MatFormFieldModule, MatIcon, MatInput],
    template: `
        @if (openedForm()) {
            <form (submit)="submit($event)" class="new-column-form">
                <mat-form-field class="w-100 mb-2">
                    <mat-label>Column name</mat-label>
                    <input matInput [formField]="columnForm.title" />
                </mat-form-field>

                <div class="d-flex gap-2">
                    <button
                        matButton
                        (click)="submit($event)"
                        [disabled]="columnForm().invalid()"
                    >
                        @if (loading()) {
                            Processing...
                        } @else {
                            Create
                        }
                    </button>
                    <button matButton type="button" (click)="this.openedForm.set(false)">
                        Cancel
                    </button>
                </div>
            </form>
        } @else {
            <button matButton="tonal" type="button" (click)="openForm()">
                <mat-icon>add</mat-icon> Create new column
            </button>
        }
    `,

    styles: `
        .new-column-form {
            width: 280px;
        }
    `,
    host: { class: 'flex-shrink-0' },
})
export class NewColumnFormComponent implements OnInit {
    private actions$ = inject(Actions);
    private store = inject(Store);
    private destroyRef = inject(DestroyRef);

    @Input() boardId: number;
    @Input() position: number;

    openedForm = signal<boolean>(false);
    loading = signal(false);
    errors = signal<Error[]>([]);
    columnFormModel = signal<{ title: string }>({ title: '' });
    columnForm = form(this.columnFormModel, (schemaPath) => {
        required(schemaPath.title, { message: 'Name is required' });
    });

    ngOnInit() {
        this.actions$
            .pipe(
                takeUntilDestroyed(this.destroyRef),
                ofType(actionsColumns.createSuccess),
                tap(() => {
                    this.loading.set(false);
                    this.openedForm.set(false);
                    this.columnForm().value.set({ title: '' });
                    this.columnForm().reset();
                }),
            )
            .subscribe();

        this.actions$
            .pipe(
                takeUntilDestroyed(this.destroyRef),
                ofType(actionsColumns.createFailed),
                tap((data) => {
                    this.errors.set(data.errors);
                    this.loading.set(false);
                }),
            )
            .subscribe();
    }

    submit(event: Event) {
        event.preventDefault();
        const formData = this.columnFormModel();

        if (!formData.title) return;

        const payload = {
            title: formData.title,
            boardId: this.boardId,
            order: this.position,
        };

        this.store.dispatch(actionsColumns.create({ payload }));
        this.columnForm().reset();
    }

    openForm(): void {
        this.openedForm.set(true);
    }
}
