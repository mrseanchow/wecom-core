package qrcode

import "github.com/mrseanchow/wecom-core/types/common"

// GetJoinQRCodeResponse 获取加入企业二维码响�?
type GetJoinQRCodeResponse struct {
	common.Response
	// JoinQRCode 二维码链接，有效�?�?
	JoinQRCode string `json:"join_qrcode"`
}

// BatchInviteRequest 邀请成员请�?
type BatchInviteRequest struct {
	// User 成员ID列表，最多支�?000�?
	User []string `json:"user,omitempty"`
	// Party 部门ID列表，最多支�?00�?
	Party []int `json:"party,omitempty"`
	// Tag 标签ID列表，最多支�?00�?
	Tag []int `json:"tag,omitempty"`
}

// BatchInviteResponse 邀请成员响�?
type BatchInviteResponse struct {
	common.Response
	// InvalidUser 非法成员列表
	InvalidUser []string `json:"invaliduser,omitempty"`
	// InvalidParty 非法部门列表
	InvalidParty []int `json:"invalidparty,omitempty"`
	// InvalidTag 非法标签列表
	InvalidTag []int `json:"invalidtag,omitempty"`
}

