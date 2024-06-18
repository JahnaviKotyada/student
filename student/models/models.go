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

type Class struct {
	gorm.Model
	ClassID   uint      `json:"class_id"`
	ClassName string    `json:"class_name"`
	StudentID uint      `json:"student_id"`
	Students  []Student `gorm:"-" json:"-"`
}

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

// MarshalJSON custom marshal function for School struct
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

// UnmarshalJSON custom unmarshal function for School struct
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

// MarshalJSON custom marshal function for Class struct
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

// UnmarshalJSON custom unmarshal function for Class struct
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

// MarshalJSON custom marshal function for Student struct
func (s *Student) MarshalJSON() ([]byte, error) {
	type Alias Student
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(*s),
	})
}

// UnmarshalJSON custom unmarshal function for Student struct
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
