import { Component, inject, Input } from '@angular/core';
import { Store } from '@ngrx/store';
import { actionTask } from '@entities/task/model/task.actions';
import { TaskModel } from '@entities/task/model';
import { Member, selectMember } from '@entities/member';
import { FormsModule } from '@angular/forms';
import { MatIconButton } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIcon } from '@angular/material/icon';
import { MatInput } from '@angular/material/input';
import { MatMenuModule } from '@angular/material/menu';

@Component({
    selector: 'lu-assignment-dropdown',
    imports: [FormsModule, MatIconButton, MatFormFieldModule, MatIcon, MatInput, MatMenuModule],
    template: `
        <button
            matIconButton
            type="button"
            class="assignment-trigger text-muted"
            [matMenuTriggerFor]="assignmentMenu"
            aria-label="Assign task"
        >
            @if (task.assignees !== null && task.assignees.length > 0) {
                <img src="images/7.png" alt="" class="rounded-circle avatar"/>
            } @else {
                <mat-icon>account_circle</mat-icon>
            }
        </button>
        <mat-menu #assignmentMenu="matMenu">
            <div class="px-2 pt-1" (click)="$event.stopPropagation()" (keydown)="$event.stopPropagation()">
                <mat-form-field class="w-100">
                    <mat-icon matPrefix>search</mat-icon>
                    <input matInput type="text" placeholder="Search..." [(ngModel)]="search"/>
                </mat-form-field>
            </div>
            @for (member of printMembers(); track member.id) {
                <button mat-menu-item type="button" (click)="assign(task.id, member.id)">
                    <img src="images/7.png" alt="" class="rounded-circle avatar me-2"/>
                    <span>{{ member.fullName }}</span>
                </button>
            }
        </mat-menu>
    `,
    styles: `
        .assignment-trigger {
            --mat-icon-button-state-layer-size: 28px;
            padding: 2px;
        }

        .avatar {
            width: 20px;
            height: 20px;
            vertical-align: middle;
        }
    `,
})
export class AssignmentDropdown {
    @Input() task: TaskModel;
    search: string = ''

    private store = inject(Store);
    readonly members = this.store.selectSignal(selectMember.all)

    printMembers(): Member[] {
        if (this.search === "") return this.members()
        return this.members().filter(m => m.fullName.toLowerCase().includes(this.search.toLowerCase()))
    }

    assign(taskId: number, identityId: number) {
        this.store.dispatch(
            actionTask.assignTask({action: {taskId: taskId, identityId: identityId}}),
        );
    }
}
