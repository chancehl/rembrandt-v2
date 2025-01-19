package db

import (
	"context"
	"errors"
	"time"

	"github.com/chancehl/rembrandt-v2/internal/cache"
	"github.com/jackc/pgx/v5"
)

type Client struct {
	url   string
	cache *cache.InMemoryCache
}

func NewClient(url string, c *cache.InMemoryCache) (*Client, error) {
	return &Client{cache: c, url: url}, nil
}

func (c *Client) GetSubscription(guildID string) (*Subscription, error) {
	conn, err := pgx.Connect(context.Background(), c.url)
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	var subscription Subscription
	if err := conn.QueryRow(context.Background(), "SELECT guild_id FROM subscriptions WHERE guild_id = $1", guildID).Scan(&subscription.GuildID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &subscription, nil
}

func (c *Client) CreateSubscription(guildID, channelID, memberID string) (*string, error) {
	conn, err := pgx.Connect(context.Background(), c.url)
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	var id string
	statement := "INSERT INTO subscriptions (guild_id, channel_id, created_by, created_on, last_modified_by, last_modified, active) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	if err := conn.QueryRow(context.Background(), statement, guildID, channelID, memberID, time.Now(), memberID, time.Now(), true).Scan(&id); err != nil {
		return nil, err
	}

	return &id, nil
}

func (c *Client) DeactivateSusbcription(guildID, memberID string) error {
	conn, err := pgx.Connect(context.Background(), c.url)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	return conn.QueryRow(context.Background(), "UPDATE subscriptions SET active = false, last_modified_by = $2, last_modified = $3 WHERE guild_id = $1", guildID, memberID, time.Now()).Scan()
}

func (c *Client) ActivateSubscription(guildID, memberID string) error {
	conn, err := pgx.Connect(context.Background(), c.url)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	return conn.QueryRow(context.Background(), "UPDATE subscriptions SET active = true, last_modified_by = $2, last_modified = $3 WHERE guild_id = $1", guildID, memberID, time.Now()).Scan()
}
