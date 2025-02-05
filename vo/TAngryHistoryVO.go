package vo

import (
	"time"
)

type AngryHistoryRecordVo struct {
	ID      uint      `gorm:"primaryKey;autoIncrement;comment:'唯一标识'" json:"id"`
	Time    time.Time `gorm:"not null;comment:'记录时间'" json:"time"`
	Records string    `gorm:"type:varchar(255);not null;comment:'详细记录'" json:"records"`
}
