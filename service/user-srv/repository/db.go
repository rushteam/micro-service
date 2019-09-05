package repository

import (
	"context"
	"database/sql"

	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

//Database ..
type Database interface {
	// All db.Database methods are available on this session.
	db.Database
	// All SQLBuilder methods are available on this session.
	sqlbuilder.SQLBuilder
	// Context returns the context used as default for queries on this session
	// and for new transactions.  If no context has been set, a default
	// context.Background() is returned.
	Context() context.Context
	SetTxOptions(sql.TxOptions)
	// TxOptions returns the defaultx TxOptions.
	TxOptions() *sql.TxOptions
}
