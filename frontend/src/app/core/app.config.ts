import type { ApplicationConfig } from '@angular/core';
import { isDevMode, provideBrowserGlobalErrorListeners } from '@angular/core';
import { provideHttpClient, withInterceptors, withXhr } from '@angular/common/http';
import { provideRouterStore } from '@ngrx/router-store';
import { provideRouter, withComponentInputBinding } from '@angular/router';
import { provideStore } from '@ngrx/store';
import { provideEffects } from '@ngrx/effects';

import { routes } from './app.routes';
import { apiInterceptor } from './api.interceptor';
import { rootEffects, rootReducers } from './store';
import { provideStoreDevtools } from '@ngrx/store-devtools';

export const appConfig: ApplicationConfig = {
    providers: [
        provideBrowserGlobalErrorListeners(),
        provideHttpClient(withXhr(), withInterceptors([apiInterceptor])),
        provideRouter(routes, withComponentInputBinding()),
        provideEffects(rootEffects),
        provideStore(rootReducers),
        provideRouterStore(),

        provideStoreDevtools({
            maxAge: 25,
            logOnly: !isDevMode(),
            autoPause: true,
            trace: false,
            traceLimit: 75,
            connectInZone: true,
        }),
    ],
};
