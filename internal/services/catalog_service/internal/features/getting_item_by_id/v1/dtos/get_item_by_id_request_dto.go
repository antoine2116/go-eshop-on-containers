package dtos

type GetItemByIdRequestDto struct {
	Id int `uri:"id" biding:"required"`
}
