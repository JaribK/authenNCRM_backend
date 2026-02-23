package infrastructures

import (
	"authenncrm/config"
	"authenncrm/internal/entities"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase(cfg *config.Config) (*gorm.DB, error) {

	useAutoMigrate := true

	dsn := fmt.Sprintf(`
			host=%s
			port=%s 
			user=%s 
			password=%s 
			dbname=%s 
			sslmode=disable`,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
		DryRun: false,
	})
	if err != nil {
		return nil, err
	}

	if err := migrateDatabase(db, useAutoMigrate); err != nil {
		return nil, err
	}

	log.Println("Connected to the database successfully.")

	return db, nil
}

func migrateDatabase(db *gorm.DB, isEnabled bool) error {
	if !isEnabled {
		return nil
	}

	db.AutoMigrate(
		&entities.User{},
		&entities.Role{},
		&entities.PointLogs{},
	)

	log.Println("Database migration completed successfully.")

	return nil
}
