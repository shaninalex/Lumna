import { isPast } from 'date-fns';

export function isInPast(d: Date): boolean {
    return isPast(d)
}
