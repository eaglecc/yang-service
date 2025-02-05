package xutils

import (
	"strconv"
	"time"
	"yang-service/commons/constparam"
	"yang-service/entity"
	"yang-service/vo"
)

type XUtilServ struct {
}

// GetAngryHistory 读取数据库angry_history_records表，获取历史记录
func (s *XUtilServ) GetAngryHistory() (recordsVo []vo.AngryHistoryRecordVo, err error) {
	var records []entity.AngryHistoryRecord
	result := constparam.YANG_DB.Order("time desc").Find(&records) //TODO： 修改为在Dao层查数据
	if result.Error != nil {
		return nil, result.Error
	}

	for _, record := range records {
		recordsVo = append(recordsVo, vo.AngryHistoryRecordVo{
			ID:      record.ID,
			Time:    record.Time,
			Records: record.Records,
		})
	}

	return recordsVo, nil
}

func (s *XUtilServ) AddAngryHistory(record vo.AngryHistoryRecordVo) (int, error) {
	scoreVal, _ := strconv.ParseFloat(record.Score, 64)
	// Map vo.AngryHistoryRecordVo to entity.AngryHistoryRecord
	newRecord := entity.AngryHistoryRecord{
		Time:    time.Now().UTC(),
		Score:   scoreVal,
		Records: record.Records,
	}

	// Save the new record to the database
	result := constparam.YANG_DB.Create(&newRecord) //TODO： 修改为在Dao层存数据
	if result.Error != nil {
		return 0, result.Error
	}
	return 1, nil
}
