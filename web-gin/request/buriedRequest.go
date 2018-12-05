package buried_request


type AddBuriedJson struct {
	SourceId   string `form:"source_id"`
	UserId     string `form:"user_id" binding:"required"`
	CrmUserId  string `form:"crm_user_id"`
	GroupId    string `form:"group_id"`
	CrmGroupId string `form:"crm_group_id"`
	GroupLink  string `form:"group_link"`
	SourceType string `form:"source_type"`
}