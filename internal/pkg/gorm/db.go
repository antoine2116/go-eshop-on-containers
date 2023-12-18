package gormPostgres

import (
	"fmt"

	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/gorm/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGorm(config *config.GormOptions) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		config.Host,
		config.User,
		config.Password,
		config.Db,
		config.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate()

	return db, nil
}
