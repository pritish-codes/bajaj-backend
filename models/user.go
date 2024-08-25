package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserId                   string   `json:"user_id" gorm:"primaryKey"`
	Email                    string   `json:"email"`
	RollNumber               string   `json:"roll_number"`
	Numbers                  []string `json:"numbers"`
	Alphabets                []string `json:"alphabets"`
	HighestLowercaseAlphabet []string `json:"highest_lowercase_alphabet"`
}
