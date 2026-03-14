package externalcontact

import "github.com/mrseanchow/wecom-core/types/common"

// OwnerFilter 群主过滤
type OwnerFilter struct {
	UserIDList []string `json:"userid_list,omitempty"`
}

// ListGroupChatRequest 获取客户群列表请�?
type ListGroupChatRequest struct {
	StatusFilter int          `json:"status_filter,omitempty"`
	OwnerFilter  *OwnerFilter `json:"owner_filter,omitempty"`
	Cursor       string       `json:"cursor,omitempty"`
	Limit        int          `json:"limit"`
}

// GroupChatItem 客户群列表项
type GroupChatItem struct {
	ChatID string `json:"chat_id"`
	Status int    `json:"status"`
}

// ListGroupChatResponse 获取客户群列表响�?
type ListGroupChatResponse struct {
	common.Response
	GroupChatList []GroupChatItem `json:"group_chat_list"`
	NextCursor    string          `json:"next_cursor,omitempty"`
}

// GetGroupChatRequest 获取客户群详情请�?
type GetGroupChatRequest struct {
	ChatID   string `json:"chat_id"`
	NeedName int    `json:"need_name,omitempty"`
}

// GroupChatInvitor 邀请�?
type GroupChatInvitor struct {
	UserID string `json:"userid"`
}

// GroupChatMember 客户群成�?
type GroupChatMember struct {
	UserID        string            `json:"userid"`
	Type          int               `json:"type"`
	UnionID       string            `json:"unionid,omitempty"`
	JoinTime      int64             `json:"join_time"`
	JoinScene     int               `json:"join_scene"`
	Invitor       *GroupChatInvitor `json:"invitor,omitempty"`
	GroupNickname string            `json:"group_nickname,omitempty"`
	Name          string            `json:"name,omitempty"`
}

// GroupChatAdmin 客户群管理员
type GroupChatAdmin struct {
	UserID string `json:"userid"`
}

// GroupChat 客户群详�?
type GroupChat struct {
	ChatID        string            `json:"chat_id"`
	Name          string            `json:"name"`
	Owner         string            `json:"owner"`
	CreateTime    int64             `json:"create_time"`
	Notice        string            `json:"notice,omitempty"`
	MemberList    []GroupChatMember `json:"member_list"`
	AdminList     []GroupChatAdmin  `json:"admin_list,omitempty"`
	MemberVersion string            `json:"member_version,omitempty"`
}

// GetGroupChatResponse 获取客户群详情响�?
type GetGroupChatResponse struct {
	common.Response
	GroupChat GroupChat `json:"group_chat"`
}

// OpenGIDToChatIDRequest 客户群opengid转换请求
type OpenGIDToChatIDRequest struct {
	OpenGID string `json:"opengid"`
}

// OpenGIDToChatIDResponse 客户群opengid转换响应
type OpenGIDToChatIDResponse struct {
	common.Response
	ChatID string `json:"chat_id"`
}

// TransferGroupChatRequest 分配离职成员的客户群请求
type TransferGroupChatRequest struct {
	ChatIDList []string `json:"chat_id_list"`
	NewOwner   string   `json:"new_owner"`
}

// FailedChat 转群失败条目
type FailedChat struct {
	ChatID  string `json:"chat_id"`
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// TransferGroupChatResponse 分配离职成员的客户群响应
type TransferGroupChatResponse struct {
	common.Response
	FailedChatList []FailedChat `json:"failed_chat_list,omitempty"`
}

