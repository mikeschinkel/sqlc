// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package querytest

import (
	"database/sql"
)

type User struct {
	ID        uint64
	FirstName string
	LastName  sql.NullString
}
