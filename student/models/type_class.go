package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	ClassID   uint      `json:"class_id"`
	ClassName string    `json:"class_name"`
	StudentID uint      `json:"student_id"`
	Students  []Student `gorm:"-" json:"-"`
}

func (c *Class) MarshalJSON() ([]byte, error) {
	type Alias Class
	return json.Marshal(&struct {
		Alias
		Students []Student `json:"students"`
	}{
		Alias:    (Alias)(*c),
		Students: c.Students,
	})
}

func (c *Class) UnmarshalJSON(data []byte) error {
	type Alias Class
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
