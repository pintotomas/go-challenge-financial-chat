package models

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Model

	//TODO allow only alpha numeric for nickname
	Nickname string `validate:"required,max=64"`
	Password string `validate:"required"`
}

func (u *User) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(struct {
		ID       uint   `json:"id"`
		Nickname string `json:"nickname"`
	}{
		u.ID, u.Nickname,
	})
	if err != nil {
		err = fmt.Errorf("failed to marshal User. %w", err)
	}

	return b, err
}
