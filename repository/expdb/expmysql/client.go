package expmysql

import (
	"database/sql"
)

// Client represents a MySQL client.
type Client struct {
	db *sql.DB
}

// NewClient creates a new Client.
func NewClient(
	dataSourceName string,
) (*Client, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &Client{
		db: db,
	}, nil
}

// Close closes the database.
func (c *Client) Close() error {
	return c.db.Close()
}
