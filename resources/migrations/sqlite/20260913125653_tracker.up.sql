CREATE TABLE scopes
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    project_id  INTEGER NOT NULL,
    name        TEXT    NOT NULL,
    description TEXT,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME,

    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE
);

CREATE TABLE stages
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    scope_id    INTEGER NOT NULL,
    name        TEXT    NOT NULL,
    description TEXT,
    category    TEXT,
    position    NUMERIC,
    wip_limit   INTEGER,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME,

    FOREIGN KEY (scope_id) REFERENCES scopes (id) ON DELETE CASCADE
);

CREATE TABLE sprints
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    project_id  INTEGER NOT NULL,
    name        TEXT    NOT NULL,
    description TEXT,
    start_date  DATETIME,
    end_date    DATETIME,
    is_active   BOOLEAN NOT NULL DEFAULT FALSE,

    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE
);

CREATE TABLE work_items
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    project_id  INTEGER NOT NULL,
    type        TEXT    NOT NULL DEFAULT 'task',
    parent_id   INTEGER,
    title       TEXT    NOT NULL,
    description TEXT,
    scope_id    INTEGER,
    stage_id    INTEGER,
    rank        NUMERIC,
    priority    TEXT,
    sprint_id   INTEGER,
    estimate    TEXT,
    created_at  DATETIME         DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME,

    FOREIGN KEY (parent_id) REFERENCES work_items(id) ON DELETE SET NULL,
    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE,
    FOREIGN KEY (sprint_id) REFERENCES sprints (id) ON DELETE SET NULL,
    FOREIGN KEY (scope_id) REFERENCES scopes (id) ON DELETE SET NULL,
    FOREIGN KEY (stage_id) REFERENCES stages (id) ON DELETE SET NULL
);

CREATE TABLE work_items_assignees
(
    work_item_id INTEGER NOT NULL,
    identity_id INTEGER NOT NULL,

    PRIMARY KEY (work_item_id, identity_id),
    FOREIGN KEY (work_item_id) REFERENCES work_items (id) ON DELETE CASCADE,
    FOREIGN KEY (identity_id) REFERENCES identities (id) ON DELETE CASCADE
);

CREATE TABLE work_item_relations
(
    source_id  INTEGER NOT NULL,
    target_id  INTEGER NOT NULL,
    type       TEXT    NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (source_id, target_id, type),
    FOREIGN KEY (source_id) REFERENCES work_items (id) ON DELETE CASCADE,
    FOREIGN KEY (target_id) REFERENCES work_items (id) ON DELETE CASCADE,
    CHECK (source_id != target_id)
);


