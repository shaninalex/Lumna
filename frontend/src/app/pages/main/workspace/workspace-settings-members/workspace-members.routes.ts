import { Routes } from '@angular/router';
import { MembersListPage } from './members-list';
import { MembersInvitationsPage } from './members-invitations';

export const routes: Routes = [
    {
        path: "",
        component: MembersListPage,
    },
    {
        path: "invitations",
        component: MembersInvitationsPage,
    }
]
