package gormPostgres

import (
	"fmt"
	customLogger "github.com/antoine2116/go-eshop-on-containers/internal/pkg/gorm/logger"
	"go.uber.org/zap"

	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/gorm/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGorm(config *config.GormOptions, logger *zap.Logger) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		config.Host,
		config.User,
		config.Password,
		config.Db,
		config.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: customLogger.NewGormLogger(logger),
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
