package model

import "time"

// Brand database model
type Brand struct {
	ID        int64     `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	CreatedBy string    `gorm:"column:created_by"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	UpdatedBy string    `gorm:"column:updated_by"`
}
