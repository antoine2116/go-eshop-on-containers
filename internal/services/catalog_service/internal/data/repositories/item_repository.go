package repositories

import (
	"context"
	"errors"
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/utils"
	"github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ItemRepository interface {
	GetAllItems(ctx context.Context, query *utils.PaginationQuery, ids []int) (*utils.PaginationResult[*models.Item], error)
	GetItemById(ctx context.Context, id int) (*models.Item, error)
	GetItemsWithName(ctx context.Context, query *utils.PaginationQuery, name string) (*utils.PaginationResult[*models.Item], error)
}

type itemRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewItemRepository(logger *zap.Logger, db *gorm.DB) ItemRepository {
	return &itemRepository{
		logger: logger,
		db:     db,
	}
}

func (r *itemRepository) GetAllItems(
	ctx context.Context,
	query *utils.PaginationQuery,
	ids []int,
) (*utils.PaginationResult[*models.Item], error) {
	var items []*models.Item
	var count int64

	err := r.db.WithContext(ctx).Model(items).Count(&count).Error
	if err != nil {
		return nil, err
	}

	dbQuery := r.db.WithContext(ctx).
		Offset(query.GetOffset()).
		Limit(query.GetLimit())

	if len(ids) > 0 {
		dbQuery.Where("id IN ?", ids)
	}

	err = dbQuery.Find(&items).Error
	if err != nil {
		return nil, err
	}

	return utils.NewPaginationResult[*models.Item](query.PageIndex, query.PageSize, count, items), nil
}

func (r *itemRepository) GetItemById(
	ctx context.Context,
	id int,
) (*models.Item, error) {
	var item *models.Item

	err := r.db.WithContext(ctx).First(&item, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (r *itemRepository) GetItemsWithName(
	ctx context.Context,
	query *utils.PaginationQuery,
	name string,
) (*utils.PaginationResult[*models.Item], error) {
	var items []*models.Item
	var count int64

	err := r.db.WithContext(ctx).
		Model(items).
		Where("name LIKE ?", name+"%").
		Count(&count).Error
	if err != nil {
		return nil, err
	}

	dbQuery := r.db.WithContext(ctx).
		Where("name LIKE ?", name+"%").
		Offset(query.GetOffset()).
		Limit(query.GetLimit())

	err = dbQuery.Find(&items).Error
	if err != nil {
		return nil, err
	}

	return utils.NewPaginationResult[*models.Item](query.PageIndex, query.PageSize, count, items), nil
}
