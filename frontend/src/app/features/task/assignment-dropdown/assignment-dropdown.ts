import { Component, inject, Input } from '@angular/core';
import { CdkMenu, CdkMenuItem, CdkMenuTrigger } from '@angular/cdk/menu';
import { Store } from '@ngrx/store';
import { actionTask } from '@entities/task/model/task.actions';
import { TaskModel } from '@entities/task/model';
import { NgIcon, provideIcons } from '@ng-icons/core';
import { reiconUserCircle } from '@ng-icons/reicon';
import { Member, selectMember } from '@entities/member';
import { FormsModule } from '@angular/forms';

@Component({
    selector: 'lu-assignment-dropdown',
    imports: [CdkMenu, CdkMenuItem, CdkMenuTrigger, NgIcon, FormsModule],
    providers: [provideIcons({reiconUserCircle})],
    template: `
        <div
            class="text-muted d-inline-flex align-items-center lh-1"
            [cdkMenuTriggerFor]="assignmentMenu"
        >
            @if (task.assignees !== null && task.assignees.length > 0) {
                <img src="images/7.png" alt="" class="rounded-circle" style="width: 16px"/>
            } @else {
                <ng-icon name="reiconUserCircle"/>
            }
        </div>
        <ng-template #assignmentMenu>
            <div class="card" cdkMenu>
                <input type="text" class="form-control border-0" placeholder="Search..." [(ngModel)]="search"/>
                <div class="list-group list-group-flush">
                    @for (member of printMembers(); track member.id) {
                        <button
                            type="button"
                            class="list-group-item list-group-item-action"
                            cdkMenuItem
                            (click)="assign(task.id, member.id)"
                        >
                            <img src="images/7.png" alt="" class="rounded-circle" style="width: 16px"/>
                            {{ member.fullName }}
                        </button>
                    }
                </div>
            </div>
        </ng-template>
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
