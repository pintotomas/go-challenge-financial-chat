package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type ChatRoom struct {
	Model

	//TODO allow only letters
	Name string `validate:"required,max=64"`
}

func (c *ChatRoom) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(struct {
		ID        uint      `json:"id"`
		CreatedOn time.Time `json:"created_on"`
		UpdatedOn time.Time `json:"updated_on"`
		Name      string    `json:"name"`
	}{
		c.ID, c.CreatedOn, c.UpdatedOn, c.Name,
	})
	if err != nil {
		err = fmt.Errorf("failed to marshal ChatRoom. %w", err)
	}

	return b, err
}
