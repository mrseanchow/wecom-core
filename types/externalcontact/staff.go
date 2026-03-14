package externalcontact

import "github.com/mrseanchow/wecom-core/types/common"

// GetFollowUserListResponse 获取配置了客户联系功能的成员列表响应
type GetFollowUserListResponse struct {
	common.Response
	FollowUser []string `json:"follow_user"`
}

