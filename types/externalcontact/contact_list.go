package externalcontact

import "github.com/mrseanchow/wecom-core/types/common"

// GetContactListRequest 获取已服务的外部联系人请�?
type GetContactListRequest struct {
	Cursor string `json:"cursor,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

// ContactInfo 外部联系人信�?
type ContactInfo struct {
	IsCustomer     bool   `json:"is_customer"`
	TmpOpenID      string `json:"tmp_openid"`
	ExternalUserID string `json:"external_userid,omitempty"`
	Name           string `json:"name,omitempty"`
	FollowUserID   string `json:"follow_userid,omitempty"`
	ChatID         string `json:"chat_id,omitempty"`
	ChatName       string `json:"chat_name,omitempty"`
	AddTime        int64  `json:"add_time"`
}

// GetContactListResponse 获取已服务的外部联系人响�?
type GetContactListResponse struct {
	common.Response
	InfoList   []ContactInfo `json:"info_list"`
	NextCursor string        `json:"next_cursor,omitempty"`
}

