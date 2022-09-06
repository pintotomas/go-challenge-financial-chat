package migrations

import (
	"fmt"
	"go-challenge-financial-chat/app/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	migrations = append(migrations, &gormigrate.Migration{
		ID: "2022090601",
		Migrate: func(tx *gorm.DB) error {
			fmt.Println("Running migration create_user")
			type User struct {
				models.Model
				Nickname string
				Password string
			}

			return tx.AutoMigrate(&User{})
		},
		Rollback: func(tx *gorm.DB) error {
			fmt.Println("Rollback migration create_user")
			type User struct{}

			return tx.Migrator().DropTable(&User{})
		},
	})
}
