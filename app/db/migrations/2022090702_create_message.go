package migrations

import (
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	"go-challenge-financial-chat/app/models"
	"gorm.io/gorm"
)

func init() {
	migrations = append(migrations, &gormigrate.Migration{
		ID: "2022090702",
		Migrate: func(tx *gorm.DB) error {
			fmt.Println("Running migration create_message")
			type Message struct {
				models.Model
				Text string
			}

			return tx.AutoMigrate(&Message{})
		},
		Rollback: func(tx *gorm.DB) error {
			fmt.Println("Rollback migration create_message")
			type Message struct{}

			return tx.Migrator().DropTable(&Message{})
		},
	})
}
