import { Component, input } from '@angular/core';
import { ActivityModel } from '../../model/activity.model';
import { ActivityListGroupItemComponent } from '../activity-list-group-item';

@Component({
    selector: 'lu-activity-list',
    imports: [
        ActivityListGroupItemComponent,
    ],
    template: `
        <div class="card">
            <div class="card-header">
                Activity
            </div>
            <div class="list-group list-group-flush">
                @if (activities().length > 0) {
                    @for (activity of activities(); track activity.id) {
                        <lu-activity-list-group-item [activity]="activity"/>
                    }
                }   @else {
                    <div class="list-group-item">
                        No activities yet.
                    </div>
                }
            </div>
        </div>
    `
})
export class ActivityListComponent {
    activities = input.required<ActivityModel[]>()
}
