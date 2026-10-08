CREATE TABLE email_queue
(
	id         INTEGER PRIMARY KEY AUTOINCREMENT,
	status     TEXT     NOT NULL default 'pending',
	subject    TEXT     NOT NULL,
	content    TEXT     NOT NULL,
	type       TEXT     NULL,
	attempts   INTEGER  NULL,
	receivers  TEXT     NOT NULL,
	message    TEXT     NULL,
	send_at    DATETIME NULL,
	created_at DATETIME          DEFAULT CURRENT_TIMESTAMP
);
