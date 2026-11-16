import { ProjectEffects } from "@entities/project";
import { UserEffects } from "@entities/user";
import { WorkspaceEffects } from "@entities/workspace";
import { MainEffects } from "./main.effects";
import { TaskEffects } from "@entities/task";
import { BoardEffects } from "@entities/board";
import { ColumnEffects } from "@entities/column";
import { KanbanEffects } from "@widgets/kanban-widget"
import { MemberEffects } from '@entities/member/model';
import { InvitationEffects } from '@entities/invitation/model';
import { RefreshTokenEffects } from '@entities/refresh_token';

export const mainEffects = [
    // effects
    TaskEffects,
    MainEffects,
    UserEffects,
    WorkspaceEffects,
    ProjectEffects,
    BoardEffects,
    ColumnEffects,
    MemberEffects,
    InvitationEffects,
    RefreshTokenEffects,

    // widgets
    KanbanEffects,
];
