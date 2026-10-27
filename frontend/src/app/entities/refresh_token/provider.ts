import { createFeature } from "@ngrx/store";
import { refreshTokenReducer } from './model'

export const refreshTokenFeature = createFeature({
    name: 'refresh_token',
    reducer: refreshTokenReducer,
});
