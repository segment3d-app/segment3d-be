// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: assets.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createAsset = `-- name: CreateAsset :one
INSERT INTO "assets" (
        uid,
        title,
        slug,
        status,
        "assetUrl",
        "assetType",
        "thumbnailUrl",
        "isPrivate",
        likes
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, uid, title, slug, "assetUrl", "assetType", "thumbnailUrl", "gaussianUrl", "pointCloudUrl", "isPrivate", status, likes, "createdAt", "updatedAt"
`

type CreateAssetParams struct {
	Uid          uuid.NullUUID `json:"uid"`
	Title        string        `json:"title"`
	Slug         string        `json:"slug"`
	Status       string        `json:"status"`
	AssetUrl     string        `json:"assetUrl"`
	AssetType    string        `json:"assetType"`
	ThumbnailUrl string        `json:"thumbnailUrl"`
	IsPrivate    bool          `json:"isPrivate"`
	Likes        int32         `json:"likes"`
}

func (q *Queries) CreateAsset(ctx context.Context, arg CreateAssetParams) (Assets, error) {
	row := q.db.QueryRowContext(ctx, createAsset,
		arg.Uid,
		arg.Title,
		arg.Slug,
		arg.Status,
		arg.AssetUrl,
		arg.AssetType,
		arg.ThumbnailUrl,
		arg.IsPrivate,
		arg.Likes,
	)
	var i Assets
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.Title,
		&i.Slug,
		&i.AssetUrl,
		&i.AssetType,
		&i.ThumbnailUrl,
		&i.GaussianUrl,
		&i.PointCloudUrl,
		&i.IsPrivate,
		&i.Status,
		&i.Likes,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllAssets = `-- name: GetAllAssets :many
SELECT a.id, a.uid, a.title, a.slug, a."assetUrl", a."assetType", a."thumbnailUrl", a."gaussianUrl", a."pointCloudUrl", a."isPrivate", a.status, a.likes, a."createdAt", a."updatedAt",
    u.name,
    u.avatar,
    u.email
FROM "assets" AS a
    LEFT JOIN "users" AS u ON u.uid = a.uid
ORDER BY a."createdAt" DESC
`

type GetAllAssetsRow struct {
	ID            uuid.UUID      `json:"id"`
	Uid           uuid.NullUUID  `json:"uid"`
	Title         string         `json:"title"`
	Slug          string         `json:"slug"`
	AssetUrl      string         `json:"assetUrl"`
	AssetType     string         `json:"assetType"`
	ThumbnailUrl  string         `json:"thumbnailUrl"`
	GaussianUrl   sql.NullString `json:"gaussianUrl"`
	PointCloudUrl sql.NullString `json:"pointCloudUrl"`
	IsPrivate     bool           `json:"isPrivate"`
	Status        string         `json:"status"`
	Likes         int32          `json:"likes"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	Name          sql.NullString `json:"name"`
	Avatar        sql.NullString `json:"avatar"`
	Email         sql.NullString `json:"email"`
}

func (q *Queries) GetAllAssets(ctx context.Context) ([]GetAllAssetsRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllAssets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllAssetsRow{}
	for rows.Next() {
		var i GetAllAssetsRow
		if err := rows.Scan(
			&i.ID,
			&i.Uid,
			&i.Title,
			&i.Slug,
			&i.AssetUrl,
			&i.AssetType,
			&i.ThumbnailUrl,
			&i.GaussianUrl,
			&i.PointCloudUrl,
			&i.IsPrivate,
			&i.Status,
			&i.Likes,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Avatar,
			&i.Email,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAssetsById = `-- name: GetAssetsById :one
SELECT id, uid, title, slug, "assetUrl", "assetType", "thumbnailUrl", "gaussianUrl", "pointCloudUrl", "isPrivate", status, likes, "createdAt", "updatedAt"
FROM "assets"
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetAssetsById(ctx context.Context, id uuid.UUID) (Assets, error) {
	row := q.db.QueryRowContext(ctx, getAssetsById, id)
	var i Assets
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.Title,
		&i.Slug,
		&i.AssetUrl,
		&i.AssetType,
		&i.ThumbnailUrl,
		&i.GaussianUrl,
		&i.PointCloudUrl,
		&i.IsPrivate,
		&i.Status,
		&i.Likes,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAssetsBySlug = `-- name: GetAssetsBySlug :one
SELECT id, uid, title, slug, "assetUrl", "assetType", "thumbnailUrl", "gaussianUrl", "pointCloudUrl", "isPrivate", status, likes, "createdAt", "updatedAt"
FROM "assets"
WHERE slug = $1
LIMIT 1
`

func (q *Queries) GetAssetsBySlug(ctx context.Context, slug string) (Assets, error) {
	row := q.db.QueryRowContext(ctx, getAssetsBySlug, slug)
	var i Assets
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.Title,
		&i.Slug,
		&i.AssetUrl,
		&i.AssetType,
		&i.ThumbnailUrl,
		&i.GaussianUrl,
		&i.PointCloudUrl,
		&i.IsPrivate,
		&i.Status,
		&i.Likes,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAssetsByUid = `-- name: GetAssetsByUid :many
SELECT id, uid, title, slug, "assetUrl", "assetType", "thumbnailUrl", "gaussianUrl", "pointCloudUrl", "isPrivate", status, likes, "createdAt", "updatedAt"
FROM "assets"
WHERE uid = $1
ORDER BY "createdAt" DESC
`

func (q *Queries) GetAssetsByUid(ctx context.Context, uid uuid.NullUUID) ([]Assets, error) {
	rows, err := q.db.QueryContext(ctx, getAssetsByUid, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Assets{}
	for rows.Next() {
		var i Assets
		if err := rows.Scan(
			&i.ID,
			&i.Uid,
			&i.Title,
			&i.Slug,
			&i.AssetUrl,
			&i.AssetType,
			&i.ThumbnailUrl,
			&i.GaussianUrl,
			&i.PointCloudUrl,
			&i.IsPrivate,
			&i.Status,
			&i.Likes,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMyAssets = `-- name: GetMyAssets :many
SELECT id, uid, title, slug, "assetUrl", "assetType", "thumbnailUrl", "gaussianUrl", "pointCloudUrl", "isPrivate", status, likes, "createdAt", "updatedAt"
FROM "assets"
WHERE uid = $1
ORDER BY "createdAt" DESC
`

func (q *Queries) GetMyAssets(ctx context.Context, uid uuid.NullUUID) ([]Assets, error) {
	rows, err := q.db.QueryContext(ctx, getMyAssets, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Assets{}
	for rows.Next() {
		var i Assets
		if err := rows.Scan(
			&i.ID,
			&i.Uid,
			&i.Title,
			&i.Slug,
			&i.AssetUrl,
			&i.AssetType,
			&i.ThumbnailUrl,
			&i.GaussianUrl,
			&i.PointCloudUrl,
			&i.IsPrivate,
			&i.Status,
			&i.Likes,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSlug = `-- name: GetSlug :many
SELECT slug
FROM "assets"
WHERE slug LIKE $1
ORDER BY "createdAt" ASC
`

func (q *Queries) GetSlug(ctx context.Context, slug string) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getSlug, slug)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var slug string
		if err := rows.Scan(&slug); err != nil {
			return nil, err
		}
		items = append(items, slug)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
