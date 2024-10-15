package repository

import (
	"context"
	"database/sql"
	"go_restful_api/model/domain"
)

type CategoryRepository interface {
	Create(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, categoryId int)
}
