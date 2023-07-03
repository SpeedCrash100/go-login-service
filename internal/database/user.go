package database

import (
	"github.com/SpeedCrash100/go-login-service/pkg/consts"
	"github.com/SpeedCrash100/go-login-service/pkg/models"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func NewUserDB() *gorm.DB {
	db_string := viper.GetString(consts.CONFIG_USER_DB_CONN_STRING)

	gorm, err := gorm.Open(sqlite.Open(db_string), &gorm.Config{
		Logger: newLogger(),
	})

	if err != nil {
		log.Fatal().Err(err).Msg("database connection fail")
	}

	return gorm
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(models.User{})
}

type UserQuerier interface {
	// SELECT * FROM @@table WHERE id=@id
	GetByID(id int) (gen.T, error)

	// SELECT * FROM @@table WHERE username=@username
	GetByUsername(username string) (gen.T, error)
}
