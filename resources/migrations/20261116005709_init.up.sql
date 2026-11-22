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
	identity_id  INTEGER NOT NULL,
	workspace_id INTEGER NOT NULL,
	created_at   DATETIME DEFAULT CURRENT_TIMESTAMP,
	role         TEXT    NOT NULL
						  DEFAULT 'member',

	PRIMARY KEY (identity_id, workspace_id),
	FOREIGN KEY (identity_id) REFERENCES identities (id) ON DELETE CASCADE,
	FOREIGN KEY (workspace_id) REFERENCES workspaces (id) ON DELETE CASCADE
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

CREATE TABLE IF NOT EXISTS "journal"
(
	id          INTEGER PRIMARY KEY AUTOINCREMENT,
	actor_id    INTEGER NULL,
	entity_type TEXT    NOT NULL,
	entity_id   INTEGER NOT NULL,
	event_type  TEXT    NOT NULL,
	data        TEXT    NULL,
	created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,

	FOREIGN KEY (actor_id) REFERENCES identities (id) ON DELETE SET NULL
);

CREATE INDEX journal_entity_idx ON journal (entity_type, entity_id, created_at);

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

CREATE TABLE work_items_assignees
(
	work_item_id INTEGER NOT NULL,
	identity_id  INTEGER NOT NULL,

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

CREATE TABLE IF NOT EXISTS "work_items"
(
	id          INTEGER PRIMARY KEY AUTOINCREMENT,
	project_id  INTEGER  NOT NULL,
	type        TEXT     NOT NULL DEFAULT 'task',
	parent_id   INTEGER,
	title       TEXT     NOT NULL,
	description TEXT,
	scope_id    INTEGER,
	stage_id    INTEGER,
	rank        NUMERIC,
	priority    TEXT,
	sprint_id   INTEGER,
	estimate    TEXT,
	created_at  DATETIME          DEFAULT CURRENT_TIMESTAMP,
	updated_at  DATETIME,
	due_to      DATETIME NULL,

	FOREIGN KEY (parent_id) REFERENCES "work_items" (id) ON DELETE SET NULL,
	FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE,
	FOREIGN KEY (sprint_id) REFERENCES sprints (id) ON DELETE SET NULL,
	FOREIGN KEY (scope_id) REFERENCES scopes (id) ON DELETE SET NULL,
	FOREIGN KEY (stage_id) REFERENCES stages (id) ON DELETE SET NULL
);

CREATE TABLE notifications
(
	id          INTEGER PRIMARY KEY AUTOINCREMENT,
	type        TEXT NOT NULL,
	content     TEXT NOT NULL,
	ref_id      INTEGER,
	ref_type    TEXT,
	priority    TEXT,
	created_at  DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE workspace_invitations
(
	id           INTEGER PRIMARY KEY AUTOINCREMENT,
	workspace_id INTEGER  NOT NULL,
	email        TEXT     NOT NULL,
	role         TEXT     NOT NULL DEFAULT 'member',
	token_hash   TEXT     NOT NULL,
	invited_by   INTEGER  NULL,
	expires_at   DATETIME NOT NULL,
	accepted_at  DATETIME NULL,
	revoked_at   DATETIME NULL,
	created_at   DATETIME          DEFAULT CURRENT_TIMESTAMP,

	FOREIGN KEY (workspace_id) REFERENCES workspaces (id) ON DELETE CASCADE,
	FOREIGN KEY (invited_by) REFERENCES identities (id) ON DELETE SET NULL
);

CREATE UNIQUE INDEX uni_ws_invitation_pending
	ON workspace_invitations (workspace_id, email)
	WHERE accepted_at IS NULL AND revoked_at IS NULL;

CREATE TABLE email_queue
(
	id         INTEGER PRIMARY KEY AUTOINCREMENT,
	status     TEXT     NOT NULL default 'pending',
	subject    TEXT     NOT NULL,
	html       TEXT     NOT NULL,
	text       TEXT     NOT NULL,
	type       TEXT     NULL,
	attempts   INTEGER  NULL,
	receivers  TEXT     NOT NULL,
	message    TEXT     NULL,
	send_at    DATETIME NULL,
	created_at DATETIME          DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE notification_recipients
(
	id              INTEGER PRIMARY KEY AUTOINCREMENT,
	identity_id     INTEGER NOT NULL,
	notification_id INTEGER NOT NULL,
	read_at         DATETIME DEFAULT NULL,

	CONSTRAINT notification_recipients_unique UNIQUE(notification_id, identity_id),

	FOREIGN KEY (identity_id)
		REFERENCES identities (id)
		ON DELETE CASCADE,

	FOREIGN KEY (notification_id)
		REFERENCES notifications (id)
		ON DELETE CASCADE
);

CREATE INDEX notification_recipients_inbox_idx ON notification_recipients (identity_id, read_at);
