import type { PipeTransform } from '@angular/core';
import { Pipe } from '@angular/core';

@Pipe({
    name: 'trim',
})
export class TrimPipe implements PipeTransform {
    transform(value: string, len: number): unknown {
        if (value.length <= len || len > value.length) {
            return value;
        }

        return value.substring(0, len) + '...';
    }
}
