package db

import "time"

type Subscription struct {
	GuildID        string    `db:"guild_id"`
	ChannelID      string    `db:"channel_id"`
	CreatedBy      string    `db:"created_by"`
	CreatedOn      time.Time `db:"created_on"`
	LastModifiedOn time.Time `db:"last_modified"`
	LastModifiedBy string    `db:"last_modified_by"`
	Active         bool      `db:"active"`
}
