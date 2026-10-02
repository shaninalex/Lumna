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

ALTER TABLE identity_workspaces
	ADD COLUMN role TEXT NOT NULL
		DEFAULT 'member';
