package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	Model

	//TODO allow only alpha numeric for nickname
	Nickname string `validate:"required,max=64"`
	Password []byte `validate:"required"`
}

func (u *User) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(struct {
		ID        uint      `json:"id"`
		CreatedOn time.Time `json:"created_on"`
		UpdatedOn time.Time `json:"updated_on"`
		Nickname  string    `json:"nickname"`
	}{
		u.ID, u.CreatedOn, u.UpdatedOn, u.Nickname,
	})
	if err != nil {
		err = fmt.Errorf("failed to marshal User. %w", err)
	}

	return b, err
}
