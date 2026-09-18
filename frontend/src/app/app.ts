import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import packageJson from "@root/package.json";

@Component({
    selector: 'lu-root',
    imports: [RouterOutlet],
    template: `
        <router-outlet />
        <div class="app-version">
            v{{ version }}
        </div>
    `,
})
export class App {
    version: string = packageJson.version;
}
