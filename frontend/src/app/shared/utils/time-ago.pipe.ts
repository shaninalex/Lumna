import type { PipeTransform } from '@angular/core';
import { Pipe } from '@angular/core';
import { formatDistance } from 'date-fns';

@Pipe({
    name: 'timeAgo',
})
export class TimeAgoPipe implements PipeTransform {
    transform(value: Date | undefined | null): unknown {
        if (!value) {
            return '';
        }
        return formatDistance(value, new Date(), { addSuffix: true });
    }
}
