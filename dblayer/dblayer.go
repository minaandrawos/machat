package dblayer

import (
	"github.com/minaandrawos/machat/dblayer/impl"
)

func NewDBLayer(db string, con string) (*impl.SQLImpl, error) {
	switch db {
	case "sqlite":
		return impl.NewSQLiteImpl(con)
	}
	return impl.NewSQLiteImpl(con)
}
