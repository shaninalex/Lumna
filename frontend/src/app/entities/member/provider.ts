import { createFeature } from "@ngrx/store";
import { memberReducer } from './model'

export const memberFeature = createFeature({
    name: 'member',
    reducer: memberReducer,
});
