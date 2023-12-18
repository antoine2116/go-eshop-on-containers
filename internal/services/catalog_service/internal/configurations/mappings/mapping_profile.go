package mappings

import (
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/mapper"
	dtosV1 "github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/dtos/v1"
	"github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/models"
)

func ConfigureMappings() error {
	err := mapper.CreateMap[*models.Item, *dtosV1.ItemDto]()
	if err != nil {
		return err
	}

	err = mapper.CreateMap[*dtosV1.ItemDto, *models.Item]()
	if err != nil {
		return err
	}

	err = mapper.CreateMap[*models.Type, *dtosV1.TypeDto]()
	if err != nil {
		return err
	}

	err = mapper.CreateMap[*dtosV1.TypeDto, *models.Type]()
	if err != nil {
		return err
	}

	err = mapper.CreateMap[*models.Brand, *dtosV1.BrandDto]()
	if err != nil {
		return err
	}

	err = mapper.CreateMap[*dtosV1.BrandDto, *models.Brand]()
	if err != nil {
		return err
	}

	return nil
}
