package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type Message struct {
	Model

	Text string `validate:"required,max=64"`
}

func (c *Message) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(struct {
		ID        uint      `json:"id"`
		CreatedOn time.Time `json:"created_on"`
		UpdatedOn time.Time `json:"updated_on"`
		Text      string    `json:"text"`
	}{
		c.ID, c.CreatedOn, c.UpdatedOn, c.Text,
	})
	if err != nil {
		err = fmt.Errorf("failed to marshal Message. %w", err)
	}

	return b, err
}
