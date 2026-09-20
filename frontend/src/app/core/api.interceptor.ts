import type {
    HttpRequest,
    HttpHandlerFn,
    HttpEvent,
    HttpErrorResponse} from '@angular/common/http';
import { inject } from '@angular/core';
import type { Observable} from 'rxjs';
import { BehaviorSubject, catchError, filter, switchMap, take, throwError } from 'rxjs';
import { SessionApi } from '@core/store';

let isRefreshing = false;
const refreshSubject = new BehaviorSubject<boolean | null>(null);

export function apiInterceptor(
    req: HttpRequest<unknown>,
    next: HttpHandlerFn,
): Observable<HttpEvent<unknown>> {
    const sessionAPI = inject(SessionApi);
    const authReq = req.clone({
        withCredentials: true,
    });

    return next(authReq).pipe(
        catchError((error: HttpErrorResponse) => {
            if (error.status === 401 && !authReq.url.includes('/http/v1/auth/refresh')) {
                return handle401(authReq, next, sessionAPI);
            }
            return throwError(() => error);
        }),
    );
}

function handle401(
    req: HttpRequest<unknown>,
    next: HttpHandlerFn,
    sessionApi: SessionApi
): Observable<HttpEvent<unknown>> {
    if (!isRefreshing) {
        isRefreshing = true;
        refreshSubject.next(null);

        return sessionApi.refresh().pipe(
            switchMap(() => {
                isRefreshing = false;
                refreshSubject.next(true);
                return next(
                    req.clone({
                        withCredentials: true,
                    }),
                );
            }),
            catchError((err) => {
                isRefreshing = false;
                refreshSubject.next(false);
                return throwError(() => err);
            }),
        );
    }

    return refreshSubject.pipe(
        filter((result) => result === true),
        take(1),
        switchMap(() =>
            next(
                req.clone({
                    withCredentials: true,
                }),
            ),
        ),
    );
}
