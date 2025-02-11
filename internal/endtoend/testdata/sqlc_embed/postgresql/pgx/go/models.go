// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package querytest

import (
	"database/sql"
)

type BazUser struct {
	ID   int32  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type Post struct {
	ID     int32 `db:"id" json:"id"`
	UserID int32 `db:"user_id" json:"user_id"`
}

type User struct {
	ID   int32         `db:"id" json:"id"`
	Name string        `db:"name" json:"name"`
	Age  sql.NullInt32 `db:"age" json:"age"`
}
