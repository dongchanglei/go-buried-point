package service

import (
	"gitlab.kanche.com/kkche/go-buried-point/web-gin/models"
)

type BuriedPoint struct {
	SourceId   string
	UserId     string
	CrmUserId  string
	GroupId    string
	CrmGroupId string
	GroupLink  string
	SourceType string
	PageNum    int
	PageSize   int
}

func (a *BuriedPoint) Add() error {
	buried := map[string]interface{}{
		"source_id":    a.SourceId,
		"user_id":      a.UserId,
		"crm_user_id":  a.CrmUserId,
		"group_id":     a.GroupId,
		"crm_group_id": a.CrmGroupId,
		"group_link":   a.GroupLink,
		"source_type":  a.SourceType,
	}

	if err := models.AddBuriedPoint(buried); err != nil {
		return err
	}

	return nil
}

func (a *BuriedPoint) GetAll() ([]*models.BuriedPointModel, error) {
	var burieds []*models.BuriedPointModel
	burieds, err := models.GetBuriedPoints(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}

	return burieds, nil
}

func (a *BuriedPoint) Count() (int, error) {
	return models.GetBuriedPointTotal(a.getMaps())
}

func (a *BuriedPoint) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	return maps
}
