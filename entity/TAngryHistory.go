package entity

import (
	"gorm.io/gorm"
	"time"
)

type AngryHistoryRecord struct {
	ID        uint           `gorm:"primaryKey;autoIncrement;comment:'唯一标识'" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null;comment:'记录名称'" json:"name"`
	Time      time.Time      `gorm:"not null;comment:'记录时间'" json:"time"`
	Score     float64        `gorm:"type:decimal(10,2);not null;comment:'得分'" json:"score"`
	Records   string         `gorm:"type:varchar(255);not null;comment:'详细记录'" json:"records"`
	CreatedAt time.Time      // GORM 默认会自动管理 `CreatedAt`
	UpdatedAt time.Time      // GORM 默认会自动管理 `UpdatedAt`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
