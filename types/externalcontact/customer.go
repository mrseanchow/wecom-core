package externalcontact

import "github.com/mrseanchow/wecom-core/types/common"

// ExternalProfile 外部联系人的自定义展示信�?
type ExternalProfile struct {
	ExternalAttr []ExternalAttr `json:"external_attr,omitempty"`
}

// ExternalAttr 外部联系人的自定义展示信息属�?
type ExternalAttr struct {
	Type        int        `json:"type"`
	Name        string     `json:"name"`
	Text        *TextAttr  `json:"text,omitempty"`
	Web         *WebAttr   `json:"web,omitempty"`
	Miniprogram *MinipAttr `json:"miniprogram,omitempty"`
}

// TextAttr 文本类型的属�?
type TextAttr struct {
	Value string `json:"value"`
}

// WebAttr 网页类型的属�?
type WebAttr struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

// MinipAttr 小程序类型的属�?
type MinipAttr struct {
	AppID    string `json:"appid"`
	PagePath string `json:"pagepath"`
	Title    string `json:"title"`
}

// ExternalContact 外部联系人（客户）信�?
type ExternalContact struct {
	ExternalUserID  string          `json:"external_userid"`
	Name            string          `json:"name"`
	Position        string          `json:"position,omitempty"`
	Avatar          string          `json:"avatar,omitempty"`
	CorpName        string          `json:"corp_name,omitempty"`
	CorpFullName    string          `json:"corp_full_name,omitempty"`
	Type            int             `json:"type"`
	Gender          int             `json:"gender"`
	UnionID         string          `json:"unionid,omitempty"`
	ExternalProfile ExternalProfile `json:"external_profile,omitempty"`
}

// FollowUserTag 跟进人为客户打的标签
type FollowUserTag struct {
	GroupName string `json:"group_name"`
	TagName   string `json:"tag_name"`
	TagID     string `json:"tag_id,omitempty"`
	Type      int    `json:"type"`
}

// WechatChannels 视频号信�?
type WechatChannels struct {
	Nickname string `json:"nickname"`
	Source   int    `json:"source"`
}

// FollowUser 跟进人信�?
type FollowUser struct {
	UserID         string          `json:"userid"`
	Remark         string          `json:"remark,omitempty"`
	Description    string          `json:"description,omitempty"`
	CreateTime     int64           `json:"createtime"`
	Tags           []FollowUserTag `json:"tags,omitempty"`
	RemarkCorpName string          `json:"remark_corp_name,omitempty"`
	RemarkMobiles  []string        `json:"remark_mobiles,omitempty"`
	OperUserID     string          `json:"oper_userid,omitempty"`
	AddWay         int             `json:"add_way"`
	WechatChannels *WechatChannels `json:"wechat_channels,omitempty"`
	State          string          `json:"state,omitempty"`
}

// FollowInfo 企业成员客户跟进信息（用于批量接口）
type FollowInfo struct {
	UserID         string          `json:"userid"`
	Remark         string          `json:"remark,omitempty"`
	Description    string          `json:"description,omitempty"`
	CreateTime     int64           `json:"createtime"`
	TagID          []string        `json:"tag_id,omitempty"`
	RemarkCorpName string          `json:"remark_corp_name,omitempty"`
	RemarkMobiles  []string        `json:"remark_mobiles,omitempty"`
	OperUserID     string          `json:"oper_userid,omitempty"`
	AddWay         int             `json:"add_way"`
	WechatChannels *WechatChannels `json:"wechat_channels,omitempty"`
	State          string          `json:"state,omitempty"`
}

// ListExternalContactResponse 获取客户列表响应
type ListExternalContactResponse struct {
	common.Response
	ExternalUserID []string `json:"external_userid"`
}

// GetExternalContactResponse 获取客户详情响应
type GetExternalContactResponse struct {
	common.Response
	ExternalContact ExternalContact `json:"external_contact"`
	FollowUser      []FollowUser    `json:"follow_user"`
	NextCursor      string          `json:"next_cursor,omitempty"`
}

// UpdateRemarkRequest 修改客户备注信息请求
type UpdateRemarkRequest struct {
	UserID           string   `json:"userid"`
	ExternalUserID   string   `json:"external_userid"`
	Remark           string   `json:"remark,omitempty"`
	Description      string   `json:"description,omitempty"`
	RemarkCompany    string   `json:"remark_company,omitempty"`
	RemarkMobiles    []string `json:"remark_mobiles,omitempty"`
	RemarkPicMediaID string   `json:"remark_pic_mediaid,omitempty"`
}

// BatchGetByUserRequest 批量获取客户详情请求
type BatchGetByUserRequest struct {
	UserIDList []string `json:"userid_list"`
	Cursor     string   `json:"cursor,omitempty"`
	Limit      int      `json:"limit,omitempty"`
}

// ExternalContactItem 批量获取客户详情�?
type ExternalContactItem struct {
	ExternalContact ExternalContact `json:"external_contact"`
	FollowInfo      FollowInfo      `json:"follow_info"`
}

// FailInfo 失败信息
type FailInfo struct {
	UnlicensedUserIDList []string `json:"unlicensed_userid_list,omitempty"`
}

// BatchGetByUserResponse 批量获取客户详情响应
type BatchGetByUserResponse struct {
	common.Response
	ExternalContactList []ExternalContactItem `json:"external_contact_list"`
	NextCursor          string                `json:"next_cursor,omitempty"`
	FailInfo            *FailInfo             `json:"fail_info,omitempty"`
}

