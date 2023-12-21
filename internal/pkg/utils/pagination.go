package utils

import (
	"emperror.dev/errors"
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/mapper"
	"github.com/gin-gonic/gin"
	"strconv"
)

type PaginationResult[T any] struct {
	PageIndex int   `json:"pageIndex"`
	PageSize  int   `json:"pageSize"`
	Count     int64 `json:"count"`
	Data      []T   `json:"data"`
}

func NewPaginationResult[T any](pageIndex, pageSize int, count int64, data []T) *PaginationResult[T] {
	return &PaginationResult[T]{
		PageIndex: pageIndex,
		PageSize:  pageSize,
		Count:     count,
		Data:      data,
	}
}

const (
	defaultPageIndex = 0
	defaultPageSize  = 10
)

type PaginationQuery struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

func (q *PaginationQuery) SetPageIndex(index string) error {
	if index == "" {
		q.PageIndex = defaultPageIndex
		return nil
	}

	n, err := strconv.Atoi(index)
	if err != nil {
		return err
	}

	q.PageIndex = n

	return nil
}

func (q *PaginationQuery) SetPageSize(size string) error {
	if size == "" {
		q.PageSize = defaultPageSize
		return nil
	}

	n, err := strconv.Atoi(size)
	if err != nil {
		return err
	}

	q.PageSize = n

	return nil
}

func (q *PaginationQuery) GetOffset() int {
	return (q.PageIndex) * q.PageSize
}

func (q *PaginationQuery) GetLimit() int {
	return q.PageSize
}

func GetPaginationQueryFromCtx(c *gin.Context) (*PaginationQuery, error) {
	q := &PaginationQuery{}

	index := c.Query("pageIndex")
	if err := q.SetPageIndex(index); err != nil {
		return nil, err
	}

	size := c.Query("pageSize")
	if err := q.SetPageSize(size); err != nil {
		return nil, err
	}

	return q, nil
}

func PaginationResultToPaginationResultDto[TDto any, TModel any](
	paginationResult *PaginationResult[TModel],
) (*PaginationResult[TDto], error) {
	if paginationResult == nil {
		return nil, errors.New("listResult is nil")
	}

	items, err := mapper.Map[[]TDto](paginationResult.Data)
	if err != nil {
		return nil, err
	}

	return &PaginationResult[TDto]{
		PageIndex: paginationResult.PageIndex,
		PageSize:  paginationResult.PageSize,
		Count:     paginationResult.Count,
		Data:      items,
	}, nil
}
