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

