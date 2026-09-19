import { OnInit, Component, DestroyRef, inject, Input, signal } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { FormField, form, required } from '@angular/forms/signals';
import { actionsColumns } from '@entities/column/model';
import { Actions, ofType } from '@ngrx/effects';
import { Store } from '@ngrx/store';
import { Error } from '@shared/models';
import { tap } from 'rxjs';

@Component({
    selector: 'lu-new-column-form',
    imports: [FormField],
    template: `
        @if (openedForm()) {
            <form (submit)="submit($event)" class="w-[280px]">
                <div class="mb-2">
                    <input
                        class="form-control w-full"
                        placeholder="Column name"
                        [formField]="columnForm.title"
                    />
                    @if (columnForm.title().touched()) {
                        <ul class="error-list">
                            @for (error of columnForm.title().errors(); track error) {
                                <li class="text-danger">
                                    <small>{{ error.message }}</small>
                                </li>
                            }
                        </ul>
                    }
                </div>

                <div class="d-flex gap-2">
                    <button
                        class="btn btn-sm btn-primary"
                        (click)="submit($event)"
                        [disabled]="columnForm().invalid()"
                    >
                        @if (loading()) {
                            Processing...
                        } @else {
                            Create
                        }
                    </button>
                    <button class="btn btn-sm btn-secondary" (click)="this.openedForm.set(false)">
                        cancel
                    </button>
                </div>
            </form>
        } @else {
            <button class="btn btn-outline-secondary btn-sm" (click)="openForm()">
                Create new column
            </button>
        }
    `,
    host: { class: 'flex-shrink-0' },
})
export class NewColumnFormComponent implements OnInit {
    private actions$ = inject(Actions);
    private store = inject(Store);
    private destroyRef = inject(DestroyRef);

    @Input() board_id: number;
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
            board_id: this.board_id,
            order: this.position,
        };

        this.store.dispatch(actionsColumns.create({ payload }));
        this.columnForm().reset();
    }

    openForm(): void {
        this.openedForm.set(true);
    }
}
