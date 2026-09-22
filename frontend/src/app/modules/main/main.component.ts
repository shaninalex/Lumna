import { Component, inject, OnInit } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { WebSocketService } from '@root/src/app/modules/main/websocket.service';

@Component({
    selector: 'lu-main',
    imports: [RouterOutlet],
    template: `<router-outlet />`,
})
export class MainComponent implements OnInit {
    private ws = inject(WebSocketService);

    ngOnInit() {
        this.ws.getMessages().subscribe(messages => {
            console.log(messages);
        })
    }
}
