-----------------------------
-- IDENTITIES AND CREDENTIALS
-----------------------------

CREATE TABLE identities
(
    id         integer PRIMARY KEY AUTOINCREMENT,
    full_name  text     NOT NULL,
    email      text     NOT NULL,
    active     numeric,
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime NULL,

    CONSTRAINT uni_identities_email UNIQUE (email)
);

CREATE TABLE credentials
(
    id               integer PRIMARY KEY AUTOINCREMENT,
    identity_id      integer NOT NULL,
    provider         text    NOT NULL,
    provider_user_id text,
    email            text,
    password_hash    text,
    created_at       datetime DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (identity_id) REFERENCES identities (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE refresh_tokens
(
    id          integer PRIMARY KEY AUTOINCREMENT,
    identity_id integer  NOT NULL,
    hash        text     NOT NULL,
    client_id   text,
    scopes      text,
    expires_at  datetime NOT NULL,
    revoked     numeric  NOT NULL DEFAULT false,
    created_at  datetime          DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (identity_id) REFERENCES identities (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE user_auth_history
(
    id          integer PRIMARY KEY AUTOINCREMENT,
    event_type  text NOT NULL,
    created_at  datetime DEFAULT CURRENT_TIMESTAMP
);

-----------------------------
-- SYSTEM
-----------------------------

CREATE TABLE emails
(
    id          integer PRIMARY KEY AUTOINCREMENT,
    from_email  VARCHAR,
    to_email    VARCHAR,
    subject     VARCHAR,
    body        VARCHAR,
    format      VARCHAR,
    sender_name VARCHAR,
    headers     VARCHAR,
    status      VARCHAR
        CHECK (status IN ('pending', 'running', 'success', 'repeat', 'error', 'skipped'))
                         DEFAULT 'pending',
    cc          VARCHAR,
    bcc         VARCHAR,
    reply_to    VARCHAR,
    sent_at     DATETIME,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    meta        text NULL
);

-----------------------------
-- WORKSPACES AND USER MANAGEMENT
-----------------------------

CREATE TABLE invitations
(
    id          integer PRIMARY KEY AUTOINCREMENT,
    email       text,
    token_hash  text,
    state       text,
    role        text,
    created_at  datetime DEFAULT CURRENT_TIMESTAMP,
    valid_until datetime NOT NULL,
    meta        text     NULL,

    CONSTRAINT uni_invitations_email UNIQUE (email)
);

CREATE TABLE workspaces
(
    id          integer PRIMARY KEY AUTOINCREMENT,
    title       text     NOT NULL,
    active      numeric,
    owner_email text     NULL,
    created_at  datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at  datetime NULL,

    CONSTRAINT workspace_title_unique UNIQUE (owner_email, title)
);

CREATE TABLE identity_workspaces
(
    identity_id     INTEGER NOT NULL,
    workspace_id    INTEGER NOT NULL,
    created_at      DATETIME DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (identity_id, workspace_id),
    FOREIGN KEY (identity_id) REFERENCES identities(id) ON DELETE CASCADE,
    FOREIGN KEY (workspace_id) REFERENCES workspaces(id) ON DELETE CASCADE
);


-----------------------------
-- WORK
-----------------------------

/*

Core entities:
    Workspace
    Project
    Board
    Column
    Task

Relationships:
    Project -> Workspace
    Board -> Project
    Column -> Board
    Taks -> Column

Events/History:
    TaskHistory

Value objects/attributes:
    Color
    Tag/Label
    Badge
    (badge is not a tag/label. Badge can have logic in it, while label is just a lable)
*/

CREATE TABLE projects
(
    id           INTEGER PRIMARY KEY AUTOINCREMENT,
    title        TEXT    NOT NULL,
    workspace_id INTEGER NOT NULL,
    owner_id     INTEGER NULL,
    meta         TEXT    NULL,
    created_at   DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at   DATETIME,

    FOREIGN KEY (workspace_id) REFERENCES workspaces (id) ON DELETE CASCADE,
    FOREIGN KEY (owner_id) REFERENCES identities (id) ON DELETE SET NULL
);

CREATE TABLE boards
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    title       TEXT    NOT NULL,
    project_id  INTEGER NOT NULL,
    meta        TEXT    NULL,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME,

    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE
);

CREATE TABLE columns
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    title       TEXT    NOT NULL,
    board_id    INTEGER NOT NULL,
    meta        TEXT    NULL,
    position    INTEGER NOT NULL DEFAULT 0,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME,

    -- Enables composite FK enforcement from board_tasks
    UNIQUE (board_id, id),
    FOREIGN KEY (board_id) REFERENCES boards (id) ON DELETE CASCADE
);

CREATE TABLE tasks
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    title       TEXT    NOT NULL,
    body        TEXT    NULL,
    completed   BOOLEAN DEFAULT 0,
    project_id  INTEGER NOT NULL, -- task always belongs to some project!
    meta        TEXT    NULL,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME,
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE
);

CREATE TABLE task_owners
(
    task_id     INTEGER NOT NULL,
    user_id     INTEGER NOT NULL,

    UNIQUE (task_id),
    PRIMARY KEY (task_id, user_id),
    FOREIGN KEY (task_id) REFERENCES tasks (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES identities (id) ON DELETE CASCADE
);

CREATE TABLE task_assignees
(
    task_id     INTEGER NOT NULL,
    user_id     INTEGER NOT NULL,

    PRIMARY KEY (task_id, user_id),
    FOREIGN KEY (task_id) REFERENCES tasks (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES identities (id) ON DELETE CASCADE
);

CREATE TABLE board_tasks
(
    board_id    INTEGER NOT NULL,
    task_id     INTEGER NOT NULL,
    column_id   INTEGER NULL,
    position    INTEGER NOT NULL,

    PRIMARY KEY (board_id, task_id),
    FOREIGN KEY (board_id) REFERENCES boards (id) ON DELETE CASCADE,
    FOREIGN KEY (task_id) REFERENCES tasks (id) ON DELETE CASCADE,
    FOREIGN KEY (board_id, column_id) REFERENCES columns (board_id, id) ON DELETE SET NULL
);

CREATE TABLE entity_events
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    identity_id INTEGER NULL,
    entity_id   INTEGER NULL,
    entity_type TEXT NULL,
    event_type  TEXT NOT NULL,
    data        TEXT NULL,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (identity_id) REFERENCES identities (id) ON DELETE SET NULL
);

CREATE INDEX idx_board_tasks_layout ON board_tasks (board_id, column_id, position);
