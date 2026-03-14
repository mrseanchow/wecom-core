package externalcontact

import "github.com/mrseanchow/wecom-core/types/common"

// AddJoinWayRequest 配置客户群进群方式请�?
type AddJoinWayRequest struct {
	Scene          int      `json:"scene"`
	Remark         string   `json:"remark,omitempty"`
	AutoCreateRoom int      `json:"auto_create_room,omitempty"`
	RoomBaseName   string   `json:"room_base_name,omitempty"`
	RoomBaseID     int      `json:"room_base_id,omitempty"`
	ChatIDList     []string `json:"chat_id_list"`
	State          string   `json:"state,omitempty"`
	MarkSource     bool     `json:"mark_source,omitempty"`
}

// AddJoinWayResponse 配置客户群进群方式响�?
type AddJoinWayResponse struct {
	common.Response
	ConfigID string `json:"config_id"`
}

// GetJoinWayRequest 获取客户群进群方式配置请�?
type GetJoinWayRequest struct {
	ConfigID string `json:"config_id"`
}

// JoinWay 客户群进群方式配置详�?
type JoinWay struct {
	ConfigID       string   `json:"config_id"`
	Scene          int      `json:"scene"`
	Remark         string   `json:"remark,omitempty"`
	AutoCreateRoom int      `json:"auto_create_room,omitempty"`
	RoomBaseName   string   `json:"room_base_name,omitempty"`
	RoomBaseID     int      `json:"room_base_id,omitempty"`
	ChatIDList     []string `json:"chat_id_list"`
	QRCode         string   `json:"qr_code,omitempty"`
	State          string   `json:"state,omitempty"`
	MarkSource     bool     `json:"mark_source,omitempty"`
}

// GetJoinWayResponse 获取客户群进群方式配置响�?
type GetJoinWayResponse struct {
	common.Response
	JoinWay JoinWay `json:"join_way"`
}

// UpdateJoinWayRequest 更新客户群进群方式配置请�?
type UpdateJoinWayRequest struct {
	ConfigID       string   `json:"config_id"`
	Scene          int      `json:"scene"`
	Remark         string   `json:"remark,omitempty"`
	AutoCreateRoom int      `json:"auto_create_room,omitempty"`
	RoomBaseName   string   `json:"room_base_name,omitempty"`
	RoomBaseID     int      `json:"room_base_id,omitempty"`
	ChatIDList     []string `json:"chat_id_list"`
	State          string   `json:"state,omitempty"`
	MarkSource     bool     `json:"mark_source,omitempty"`
}

// DeleteJoinWayRequest 删除客户群进群方式配置请�?
type DeleteJoinWayRequest struct {
	ConfigID string `json:"config_id"`
}

