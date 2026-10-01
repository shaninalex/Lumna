import { Component, input } from '@angular/core';
import { ActivityModel } from '../../model/activity.model';
import { TimeAgoPipe } from '@shared/utils';

@Component({
    selector: 'lu-activity-list-group-item',
    imports: [TimeAgoPipe],
    template: `
        <small class="text-muted">
            {{ activity().createdAt | timeAgo }}
        </small>
        <div>
            {{ activity().message }}
        </div>
    `,
    host: {
        class: 'list-group-item',
    }
})
export class ActivityListGroupItemComponent {
    activity = input.required<ActivityModel>()
}
