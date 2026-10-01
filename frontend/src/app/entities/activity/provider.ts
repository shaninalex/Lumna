import { createFeature } from "@ngrx/store";
import { activityReducer } from "./model";

export const activityFeature = createFeature({
    name: 'activity',
    reducer: activityReducer,
});
