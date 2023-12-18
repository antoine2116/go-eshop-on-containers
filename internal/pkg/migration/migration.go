package migration

import (
	"fmt"

	"emperror.dev/errors"
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/migration/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/zap"
)

type Runner interface {
	Up() error
	Down() error
}

type goMigrateMigrator struct {
	config  *config.MigrationOptions
	logger  *zap.Logger
	migrate *migrate.Migrate
}

func NewMigrator(config *config.MigrationOptions, logger *zap.Logger) (Runner, error) {
	db := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Db,
	)

	src := fmt.Sprintf("file://%s", config.MigrationsDir)

	m, err := migrate.New(src, db)
	if err != nil {
		return nil, errors.WrapIf(err, "Failed to initialize migrator")
	}

	return &goMigrateMigrator{
		config:  config,
		logger:  logger,
		migrate: m,
	}, nil
}

func (m *goMigrateMigrator) Up() error {
	if m.config.SkipMigrations {
		m.logger.Info("Database migration skipped")
		return nil
	}

	err := m.migrate.Up()

	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	}

	if err != nil {
		return errors.WrapIf(err, "Failed to migrate database")
	}

	m.logger.Info("Migration finished")

	return nil
}

func (m *goMigrateMigrator) Down() error {
	if m.config.SkipMigrations {
		m.logger.Info("Database migration skipped")
		return nil
	}

	err := m.migrate.Up()

	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	}

	if err != nil {
		return errors.WrapIf(err, "Failed to migrate database")
	}

	m.logger.Info("Migration finished")

	return nil
}
