package internal

import (
	"time"

	"gorm.io/gorm"
)

type Secret struct {
	gorm.Model
	Hash      string    `gorm:"hash"`
	FileName  string    `gorm:"filename"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
