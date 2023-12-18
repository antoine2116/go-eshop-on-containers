package dtos

type GetItemByIdRequestDto struct {
	ItemId int `uri:"id" biding:"required"`
}
