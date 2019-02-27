package expsql

import (
	"database/sql"
)

// Option represents an option for a db creation.
type Option func(*DB)

// WithLogger sets a custom logger。
func WithLogger(logger Logger) Option {
	return func(db *DB) {
		db.logger = logger
	}
}

// DB extends a sql.DB for a safe transaction operation.
type DB struct {
	*sql.DB
	logger Logger
}

// NewDB creates a new DB.
func NewDB(
	db *sql.DB,
	opts ...Option,
) *DB {
	instance := &DB{
		DB:     db,
		logger: &defaultLogger{},
	}
	for _, opt := range opts {
		opt(instance)
	}
	return instance
}

// WithTx creates a new transaction and handles rollback/commit based on the
// error object returned by the `TxFn`
func (db *DB) WithTx(fn TxFn) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and re panic
			err = tx.Rollback()
			if err != nil {
				db.logger.SetRollbackError(err)
			}
			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			err2 := tx.Rollback()
			if err2 != nil {
				db.logger.SetRollbackError(err)
			}
		} else {
			// all good, commit
			err = tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}
