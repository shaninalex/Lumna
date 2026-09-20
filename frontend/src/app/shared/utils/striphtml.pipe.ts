import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
    name: 'strip_html'
})
export class StripHtmlPipe implements PipeTransform {
    transform(value: string): string {
        return value.replace(/<.*?>/g, '');
    }
}
