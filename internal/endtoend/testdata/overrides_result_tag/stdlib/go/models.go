// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package querytest

import (
	"database/sql"

	"github.com/google/uuid"
)

type Account struct {
	ID    uuid.UUID `sometagtype:"some_value"`
	State sql.NullString
}

type UsersAccount struct {
	Id2  uuid.UUID
	Name sql.NullString
}
