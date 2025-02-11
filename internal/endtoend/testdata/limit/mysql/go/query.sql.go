// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package querytest

import (
	"context"
)

const limitMe = `-- name: LimitMe :exec
UPDATE foo SET bar='baz' LIMIT ?
`

func (q *Queries) LimitMe(ctx context.Context, limit int32) error {
	_, err := q.db.ExecContext(ctx, limitMe, limit)
	return err
}

const limitMeToo = `-- name: LimitMeToo :exec
DELETE FROM foo LIMIT ?
`

func (q *Queries) LimitMeToo(ctx context.Context, limit int32) error {
	_, err := q.db.ExecContext(ctx, limitMeToo, limit)
	return err
}
