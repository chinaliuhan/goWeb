package db

import "database/sql"

func init()  {
	sql.Register("",SQLiteDriver{})
}

