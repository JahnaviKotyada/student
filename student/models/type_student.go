package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	StudentID int     `json:"student_id"`
	Name      string  `json:"name"`
	Marks     int     `json:"marks"`
	Address   Address `gorm:"embedded;type:jsonb" json:"address"`
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
}

func (s *Student) MarshalJSON() ([]byte, error) {
	type Alias Student
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(*s),
	})
}

func (s *Student) UnmarshalJSON(data []byte) error {
	type Alias Student
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
