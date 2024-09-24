package models

import "gorm.io/gorm"

type Node struct {
	gorm.Model
	Title    string  `json:"title"`
	Content  string  `json:"content"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	ParentID uint    `json:"parent_id"`
}