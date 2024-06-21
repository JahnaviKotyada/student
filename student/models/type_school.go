package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type School struct {
	gorm.Model
	Name    string `json:"name"`
	ClassID uint   `json:"class_id"`
	Class   Class  `gorm:"-" json:"-"`
}

func (s *School) MarshalJSON() ([]byte, error) {
	type Alias School
	return json.Marshal(&struct {
		Alias
		Class Class `json:"class"`
	}{
		Alias: (Alias)(*s),
		Class: s.Class,
	})
}

func (s *School) UnmarshalJSON(data []byte) error {
	type Alias School
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
