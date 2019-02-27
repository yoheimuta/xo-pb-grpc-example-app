package expmysql

import (
	"database/sql"

	"github.com/yoheimuta/xo-example-app/infra/expsql"
)

// Client represents a MySQL client.
type Client struct {
	db *expsql.DB
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
		db: expsql.NewDB(db),
	}, nil
}

// Close closes the database.
func (c *Client) Close() error {
	return c.db.Close()
}
