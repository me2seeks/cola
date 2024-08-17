package models

type Example struct {
	ID     uint   `json:"id" gorm:"primary_key"`        // ID
	Name   string `json:"name"`                         // Name
	Status string `json:"status" gorm:"default:active"` // Status, active or inactive
}
