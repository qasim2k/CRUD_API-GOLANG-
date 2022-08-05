package main

import "gorm.io/gorm"

type Cafe struct {
	gorm.Model
	Cafename  string  `json:"cafenmae"`
	Location  string  `json:"location"`
	Cafeowner string  `json:"cafeowner"`
	Contact   float32 `json:"contact"`
	Email     string  `json:"email"`
}
