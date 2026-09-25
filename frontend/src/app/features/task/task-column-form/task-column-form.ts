import {
    Component,
    DestroyRef,
    ElementRef,
    HostListener,
    inject,
    input,
    OnInit,
    signal,
    WritableSignal
} from '@angular/core';
import { Actions, ofType } from '@ngrx/effects';
import { actionUI } from '@core/store/ui';
import { tap } from 'rxjs';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { actionTask, TaskCreateModel, taskCreateModelDefault } from '@entities/task';
import { form, FormField } from '@angular/forms/signals';
import { Store } from '@ngrx/store';

@Component({
    imports: [FormField],
    selector: 'lu-task-column-form',
    styleUrl: './task-column-form.css',
    templateUrl: './task-column-form.html',
})
export class TaskColumnForm implements OnInit {
    position = input.required<string>();
    columnId = input.required<number>();

    boardId = input.required<number>();
    taskCount = input.required<number>();
    projectId = input<number | null>();
    nextPosition = input.required<number>();

    visible = signal(false);
    taskEditFormModel: WritableSignal<TaskCreateModel> = signal<TaskCreateModel>(taskCreateModelDefault);
    taskEditForm = form(this.taskEditFormModel);
    private actions$ = inject(Actions);
    private destroyRef = inject(DestroyRef);
    private elementRef = inject(ElementRef);
    private store = inject(Store);

    submit(event: Event) {
        event.preventDefault();
        const pr = this.projectId()
        if (!pr) { return }
        const data = this.taskEditFormModel()
        data.position = this.nextPosition()
        data.projectId = pr;
        data.columnId = this.columnId()
        data.boardId = this.boardId()
        this.store.dispatch(actionTask.create({data}));
        this.visible.set(false);
    }

    ngOnInit() {
        this.actions$.pipe(
            ofType(actionUI.taskColumnForm),
            takeUntilDestroyed(this.destroyRef),
            tap(action => {
                this.visible.set(
                    this.position() === action.position &&
                    this.columnId() === action.columnId
                );
                this.taskEditForm().reset();
                this.taskEditFormModel.set(taskCreateModelDefault)
            })
        ).subscribe();
    }

    @HostListener('document:click', ['$event'])
    onDocumentClick(event: MouseEvent) {
        if (this.visible() && !this.elementRef.nativeElement.contains(event.target)) {
            this.visible.set(false);
        }
    }
}
