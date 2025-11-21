package internal

import (
	"time"

	"gorm.io/gorm"
)

type Secret struct {
	gorm.Model
	Hash      string
	FileName  string
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
