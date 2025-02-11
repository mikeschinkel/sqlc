// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package querytest

import (
	"context"
)

const getLatestVersion = `-- name: GetLatestVersion :one
WITH RECURSIVE search_tree(id, chain_id, chain_counter) AS (
	SELECT base.id, base.id AS chain_id, 0 as chain_counter
	FROM versions AS base
	WHERE base.previous_version_id IS NULL
  	UNION ALL
	SELECT v.id, search_tree.chain_id, search_tree.chain_counter + 1
	FROM versions AS v
	INNER JOIN search_tree ON search_tree.id = v.previous_version_id
)
SELECT DISTINCT ON (search_tree.chain_id) 
	search_tree.id
FROM search_tree   
ORDER BY search_tree.chain_id, chain_counter DESC
`

func (q *Queries) GetLatestVersion(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, getLatestVersion)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const getLatestVersionWithSubquery = `-- name: GetLatestVersionWithSubquery :one
SELECT id
FROM versions
WHERE versions.id IN (
  WITH RECURSIVE search_tree(id, chain_id, chain_counter) AS (
	SELECT base.id, base.id AS chain_id, 0 as chain_counter
	FROM versions AS base
	WHERE versions.previous_version_id IS NULL
	UNION ALL
	SELECT v.id, search_tree.chain_id, search_tree.chain_counter + 1
	FROM versions AS v
	INNER JOIN search_tree ON search_tree.id = v.previous_version_id 
  )
  SELECT DISTINCT ON (search_tree.chain_id) 
	search_tree.id
  FROM search_tree   
  ORDER BY search_tree.chain_id, chain_counter DESC
)
`

func (q *Queries) GetLatestVersionWithSubquery(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, getLatestVersionWithSubquery)
	var id int64
	err := row.Scan(&id)
	return id, err
}
