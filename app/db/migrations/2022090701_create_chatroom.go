package migrations

import (
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	"go-challenge-financial-chat/app/models"
	"gorm.io/gorm"
)

func init() {
	migrations = append(migrations, &gormigrate.Migration{
		ID: "2022090701",
		Migrate: func(tx *gorm.DB) error {
			fmt.Println("Running migration create_chatroom")
			type ChatRoom struct {
				models.Model
				Name string
			}

			return tx.AutoMigrate(&ChatRoom{})
		},
		Rollback: func(tx *gorm.DB) error {
			fmt.Println("Rollback migration create_chatroom")
			type ChatRoom struct{}

			return tx.Migrator().DropTable(&ChatRoom{})
		},
	})
}
