import { Component, inject } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { WebSocketService } from '@root/src/app/modules/main/websocket.service';

@Component({
    selector: 'lu-main',
    imports: [RouterOutlet],

    template: `<router-outlet />`,
})
export class MainComponent {
    constructor() {
        // Instantiating the service opens the websocket connection.
        inject(WebSocketService);
    }
}
