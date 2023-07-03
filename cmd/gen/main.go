package main

import (
	_ "github.com/SpeedCrash100/go-login-service/internal/config"
	"github.com/SpeedCrash100/go-login-service/internal/database"
	"github.com/SpeedCrash100/go-login-service/pkg/models"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "internal/database/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	db := database.NewUserDB()
	g.UseDB(db)

	g.ApplyInterface(func(database.UserQuerier) {}, models.User{})

	g.Execute()
}
