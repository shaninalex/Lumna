import { Injectable } from '@angular/core';
import { webSocket, WebSocketSubject } from 'rxjs/webSocket';
import { Observable } from 'rxjs';

@Injectable()
export class WebSocketService {
    private socket$: WebSocketSubject<unknown>;

    constructor() {
        console.log('WebSocketService');
        this.socket$ = webSocket(`ws://localhost:8000/ws`);
    }

    sendMessage(message: string): void {
        this.socket$.next(message);
    }

    getMessages(): Observable<unknown> {
        return this.socket$.asObservable();
    }

    closeConnection() {
        this.socket$.complete();
    }
}
