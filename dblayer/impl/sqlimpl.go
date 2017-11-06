package impl

import (
	"database/sql"
)

type SQLImpl struct {
	*sql.DB
	SQLQuery
}

func (sqlimpl *SQLImpl) AddClient(client string, room int64) (int64, error) {
	query, fetchid := sqlimpl.getAddClientQuery()
	rows, err := sqlimpl.Exec(query, client, room)
	if err != nil {
		return 0, err
	}
	if fetchid {
		row := sqlimpl.QueryRow(sqlimpl.getClientIdQuery(), client)
		var id int64
		err = row.Scan(&id)
		return id, err
	}
	return rows.LastInsertId()
}

func (sqlimpl *SQLImpl) AddRoom(room string) (int64, error) {
	query, fetchid := sqlimpl.getAddRoomQuery()
	rows, err := sqlimpl.Exec(query, room)
	if err != nil {
		return 0, err
	}
	if fetchid {
		row := sqlimpl.QueryRow(sqlimpl.getRoomIdQuery(), room)
		var id int64
		err = row.Scan(&id)
		return id, err
	}
	return rows.LastInsertId()
}

func (sqlimpl *SQLImpl) AddLog(log string, client int64, room int64) error {
	_, err := sqlimpl.Exec(sqlimpl.getAddLogQuery(), client, room)
	return err
}
