CREATE TABLE notifications
(
	id          INTEGER PRIMARY KEY AUTOINCREMENT,
	identity_id INTEGER,
	type        TEXT NOT NULL,
	content     TEXT NOT NULL,
	ref_id      INTEGER,
	priority    TEXT,
	read_at     DATETIME,
	created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,

	FOREIGN KEY (identity_id)
		REFERENCES identities (id)
		ON DELETE CASCADE
);
