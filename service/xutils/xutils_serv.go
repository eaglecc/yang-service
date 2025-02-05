package xutils

import (
	"yang-service/commons/constparam"
	"yang-service/entity"
	"yang-service/vo"
)

type XUtilServ struct {
}

// GetAngryHistory 读取数据库angry_history_records表，获取历史记录
func (s *XUtilServ) GetAngryHistory() (recordsVo []vo.AngryHistoryRecordVo, err error) {
	var records []entity.AngryHistoryRecord
	result := constparam.YANG_DB.Find(&records) //TODO： 修改为在Dao层查数据
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
