package database

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

/**
 * gormLogger returns a configured GORM logger instance.
 * The prefix is redundant, but it improves code readability without conflict with other loggers.
 */
func gormLogger() logger.Interface {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  true,          // Disable color
		},
	)

	return newLogger
}
