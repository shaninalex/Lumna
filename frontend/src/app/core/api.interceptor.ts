import type {
    HttpRequest,
    HttpHandlerFn,
    HttpEvent,
    HttpErrorResponse} from '@angular/common/http';
import {
    HttpClient
} from '@angular/common/http';
import { inject } from '@angular/core';
import type { Observable} from 'rxjs';
import { BehaviorSubject, catchError, filter, switchMap, take, throwError } from 'rxjs';

let isRefreshing = false;
const refreshSubject = new BehaviorSubject<boolean | null>(null);

export function apiInterceptor(
    req: HttpRequest<unknown>,
    next: HttpHandlerFn,
): Observable<HttpEvent<unknown>> {
    const http = inject(HttpClient);
    const authReq = req.clone({
        withCredentials: true,
    });

    return next(authReq).pipe(
        catchError((error: HttpErrorResponse) => {
            if (error.status === 401 && !authReq.url.includes('/http/v1/auth/refresh')) {
                return handle401(authReq, next, http);
            }

            // TODO: handle other errors:
            // - 403
            // - 500
            // - 400

            return throwError(() => error);
        }),
    );
}

function handle401(
    req: HttpRequest<unknown>,
    next: HttpHandlerFn,
    http: HttpClient,
): Observable<HttpEvent<unknown>> {
    if (!isRefreshing) {
        isRefreshing = true;
        refreshSubject.next(null);

        return http.post('/api/v1/auth/refresh', null, { withCredentials: true }).pipe(
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
