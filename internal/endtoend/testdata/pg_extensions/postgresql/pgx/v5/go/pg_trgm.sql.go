// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: pg_trgm.sql

package querytest

import (
	"context"
)

const wordSimilarity = `-- name: WordSimilarity :one
SELECT word_similarity('word', 'two words')
`

func (q *Queries) WordSimilarity(ctx context.Context) (float32, error) {
	row := q.db.QueryRow(ctx, wordSimilarity)
	var word_similarity float32
	err := row.Scan(&word_similarity)
	return word_similarity, err
}
