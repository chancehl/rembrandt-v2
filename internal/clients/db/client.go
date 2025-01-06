package db

import (
	"context"
	"fmt"
	"os"

	"github.com/chancehl/rembrandt-v2/internal/cache"
	"github.com/jackc/pgx/v5"
)

type Client struct {
	conn  *pgx.Conn
	cache *cache.InMemoryCache
}

func NewClient(c *cache.InMemoryCache) (*Client, error) {
	conn, err := pgx.Connect(context.TODO(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("could not create connection to db: %+v", err)
	}
	return &Client{conn: conn, cache: c}, nil
}
