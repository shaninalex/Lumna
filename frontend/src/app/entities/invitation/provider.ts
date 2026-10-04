import { createFeature } from "@ngrx/store";
import { invitationReducer } from "./model";

export const invitationFeature = createFeature({
    name: 'invitation',
    reducer: invitationReducer,
});
