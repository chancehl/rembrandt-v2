CREATE TABLE IF NOT EXISTS subscriptions (
	guild_id VARCHAR(255) PRIMARY KEY,
	channel_id VARCHAR(255) NOT NULL,
	created_by VARCHAR(255) NOT NULL,
	created_on DATE NOT NULL,
	created_by VARCHAR(255) NOT NULL,
	active BOOL NOT NULL,
	last_modified DATE,
	last_modified_by VARCHAR(255)
);