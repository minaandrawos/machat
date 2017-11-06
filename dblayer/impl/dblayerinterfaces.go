package impl

type SQLQuery interface {
	getAddClientQuery() (string, bool)
	getClientIdQuery() string
	getAddRoomQuery() (string, bool)
	getRoomIdQuery() string
	getAddLogQuery() string
}
