package impl

import (
	//for the sqlite driver
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteimpl struct {
	*SQLImpl
}

func NewSQLiteImpl(dsname string) (*SQLImpl, error) {
	db, err := sql.Open("sqlite3", dsname)
	return &SQLImpl{
		DB:       db,
		SQLQuery: new(sqliteimpl),
	}, err
}

func (sqlite *sqliteimpl) getAddClientQuery() (string, bool) {
	return "", false
}

func (sqlite *sqliteimpl) getClientIdQuery() string {
	return ""

}

func (sqlite *sqliteimpl) getAddRoomQuery() (string, bool) {
	return "", false
}

func (sqlite *sqliteimpl) getRoomIdQuery() string {
	return ""
}

func (sqlite *sqliteimpl) getAddLogQuery() string {
	return ""
}
