CREATE TABLE activities
(
	id          INTEGER PRIMARY KEY AUTOINCREMENT,
	identity_id INTEGER,
	entity_id   INTEGER NOT NULL,
	entity_type TEXT    NOT NULL,
	event_type  TEXT    NOT NULL,
	content     TEXT    NOT NULL,
	created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,

	FOREIGN KEY (identity_id)
		REFERENCES identities (id)
		ON DELETE CASCADE
);
