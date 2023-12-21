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
	GetAllItems(ctx context.Context, pagination *utils.PaginationQuery, ids []int) (*utils.PaginationResult[*models.Item], error)
	GetItemById(ctx context.Context, id int) (*models.Item, error)
	GetItemsWithName(ctx context.Context, pagination *utils.PaginationQuery, name string) (*utils.PaginationResult[*models.Item], error)
	GetItemsByTypeIdAndBrandId(ctx context.Context, pagination *utils.PaginationQuery, typeId, brandId int) (*utils.PaginationResult[*models.Item], error)
	GetItemsByBrandId(ctx context.Context, pagination *utils.PaginationQuery, brandId int) (*utils.PaginationResult[*models.Item], error)
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
	pagination *utils.PaginationQuery,
	ids []int,
) (*utils.PaginationResult[*models.Item], error) {
	var items []*models.Item
	var count int64

	err := r.db.WithContext(ctx).Model(items).Count(&count).Error
	if err != nil {
		return nil, err
	}

	query := r.db.WithContext(ctx).
		Preload("Brand").Preload("Type").
		Offset(pagination.GetOffset()).
		Limit(pagination.GetLimit())

	if len(ids) > 0 {
		query.Where("id IN ?", ids)
	}

	err = query.Find(&items).Error
	if err != nil {
		return nil, err
	}

	return utils.NewPaginationResult[*models.Item](pagination.PageIndex, pagination.PageSize, count, items), nil
}

func (r *itemRepository) GetItemById(
	ctx context.Context,
	id int,
) (*models.Item, error) {
	var item *models.Item

	query := r.db.WithContext(ctx).
		Preload("Brand").Preload("Type")

	err := query.First(&item, id).Error
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
	pagination *utils.PaginationQuery,
	name string,
) (*utils.PaginationResult[*models.Item], error) {
	var items []*models.Item
	var count int64

	root := r.db.Where("name LIKE ?", name+"%")

	err := root.WithContext(ctx).Model(items).Count(&count).Error
	if err != nil {
		return nil, err
	}

	query := root.WithContext(ctx).
		Preload("Brand").Preload("Type").
		Offset(pagination.GetOffset()).
		Limit(pagination.GetLimit())

	err = query.Find(&items).Error
	if err != nil {
		return nil, err
	}

	return utils.NewPaginationResult[*models.Item](pagination.PageIndex, pagination.PageSize, count, items), nil
}

func (r *itemRepository) GetItemsByTypeIdAndBrandId(
	ctx context.Context,
	pagination *utils.PaginationQuery,
	typeId, brandId int,
) (*utils.PaginationResult[*models.Item], error) {
	var items []*models.Item
	var count int64

	root := r.db.Where("type_id = ?", typeId)
	if brandId != 0 {
		root.Where("brand_id = ?", brandId)
	}

	err := root.WithContext(ctx).Model(items).Count(&count).Error
	if err != nil {
		return nil, err
	}

	query := root.WithContext(ctx).
		Preload("Brand").Preload("Type").
		Offset(pagination.GetOffset()).
		Limit(pagination.GetLimit())

	err = query.Find(&items).Error
	if err != nil {
		return nil, err
	}

	return utils.NewPaginationResult[*models.Item](pagination.PageIndex, pagination.PageSize, count, items), nil
}

func (r *itemRepository) GetItemsByBrandId(
	ctx context.Context,
	pagination *utils.PaginationQuery,
	brandId int,
) (*utils.PaginationResult[*models.Item], error) {
	var items []*models.Item
	var count int64

	root := r.db.Where("brand_id = ?", brandId)

	err := root.WithContext(ctx).Model(items).Count(&count).Error
	if err != nil {
		return nil, err
	}

	query := root.WithContext(ctx).
		Preload("Brand").Preload("Type").
		Offset(pagination.GetOffset()).
		Limit(pagination.GetLimit())

	err = query.Find(&items).Error
	if err != nil {
		return nil, err
	}

	return utils.NewPaginationResult[*models.Item](pagination.PageIndex, pagination.PageSize, count, items), nil
}
