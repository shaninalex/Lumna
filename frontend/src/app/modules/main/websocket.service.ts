import type { OnDestroy } from '@angular/core';
import { inject, Injectable } from '@angular/core';
import type { HttpErrorResponse } from '@angular/common/http';
import type { Observable, Subscription } from 'rxjs';
import { Subject } from 'rxjs';
import type { WebSocketSubject } from 'rxjs/webSocket';
import { webSocket } from 'rxjs/webSocket';
import { SessionApi } from '@core/store';

const MAX_RECONNECT_DELAY = 30_000;

function wsUrl(): string {
    const protocol = location.protocol === 'https:' ? 'wss:' : 'ws:';
    return `${protocol}//${location.host}/ws`;
}

@Injectable()
export class WebSocketService implements OnDestroy {
    private sessionApi = inject(SessionApi);
    private messages$ = new Subject<unknown>();
    private socket$?: WebSocketSubject<unknown>;
    private socketSub?: Subscription;
    private refreshSub?: Subscription;
    private reconnectTimer?: ReturnType<typeof setTimeout>;
    private attempt = 0;
    private stopped = false;

    constructor() {
        this.connect();
    }

    getMessages(): Observable<unknown> {
        return this.messages$.asObservable();
    }

    closeConnection(): void {
        this.stopped = true;
        clearTimeout(this.reconnectTimer);
        this.refreshSub?.unsubscribe();
        this.socketSub?.unsubscribe();
        this.socket$?.complete();
    }

    ngOnDestroy(): void {
        this.closeConnection();
        this.messages$.complete();
    }

    private connect(): void {
        this.socket$ = webSocket({
            url: wsUrl(),
            openObserver: { next: () => (this.attempt = 0) },
        });
        this.socketSub = this.socket$.subscribe({
            next: (message) => this.messages$.next(message),
            error: () => this.scheduleReconnect(),
            complete: () => this.scheduleReconnect(),
        });
    }

    private scheduleReconnect(): void {
        if (this.stopped) {
            return;
        }
        const delay = Math.min(MAX_RECONNECT_DELAY, 1000 * 2 ** this.attempt);
        this.attempt++;
        this.reconnectTimer = setTimeout(() => {
            this.refreshSub = this.sessionApi.refresh().subscribe({
                next: () => this.connect(),
                error: (err: HttpErrorResponse) => {
                    if (err.status === 401) {
                        this.stopped = true;
                        return;
                    }
                    this.scheduleReconnect();
                },
            });
        }, delay);
    }
}
