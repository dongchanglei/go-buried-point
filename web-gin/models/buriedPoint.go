package models

import (
	"gitlab.kanche.com/kkche/go-buried-point/web-gin/db/mysql"
	"github.com/jinzhu/gorm"
)

//埋点实体
type BuriedPointModel struct {
	mysql.BaseModel

	SourceId 	  string `json:"source_id" gorm:"index"`
	UserId        string `json:"user_id"`
	CrmUserId     string `json:"crm_user_id"`
	GroupId       string `json:"group_id"`
	CrmGroupId    string `json:"crm_group_id"`
	GroupLink     string `json:"group_link"`
	SourceType    string `json:"source_type"`
}

func AddBuriedPoint(data map[string]interface{}) error {
	buriedPoint := BuriedPointModel{
		SourceId:       data["source_id"].(string),
		UserId:         data["user_id"].(string),
		CrmUserId:      data["crm_user_id"].(string),
		GroupId:        data["group_id"].(string),
		CrmGroupId:     data["crm_group_id"].(string),
		GroupLink:      data["group_link"].(string),
		SourceType:     data["source_type"].(string),
	}

	if err := mysql.MysqlDb.Create(&buriedPoint).Error; err != nil {
		return err
	}

	return nil
}

func GetBuriedPointTotal(maps interface{}) (int, error) {
	var count int
	if err := mysql.MysqlDb.Model(&BuriedPointModel{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func GetBuriedPoints(pageNum int, pageSize int, maps interface{}) ([]*BuriedPointModel, error) {
	var buriedPoint []*BuriedPointModel
	err := mysql.MysqlDb.Model(&BuriedPointModel{}).Where(maps).Offset(pageNum).Limit(pageSize).Find(&buriedPoint).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return buriedPoint, nil
}