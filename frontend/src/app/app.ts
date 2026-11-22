import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

@Component({
    selector: 'lu-root',
    imports: [RouterOutlet],

    template: `
        <router-outlet />
        <div class="app-version">v{{ version }}</div>
    `,
})
export class App {
    version: string = "todo";
}
