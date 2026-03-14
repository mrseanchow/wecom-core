package security

import "github.com/mrseanchow/wecom-core/types/common"

// ListAdminOperLogRequest 获取管理端操作日志请�?
type ListAdminOperLogRequest struct {
	StartTime int64  `json:"start_time"`          // 开始时�?
	EndTime   int64  `json:"end_time"`            // 结束时间，范围不能超�?�?
	OperType  int    `json:"oper_type,omitempty"` // 操作类型，不填表示全�?
	UserID    string `json:"userid,omitempty"`    // 操作者userid
	Cursor    string `json:"cursor,omitempty"`    // 分页游标
	Limit     int    `json:"limit,omitempty"`     // 最大记录数，默�?00，最�?00
}

// ListAdminOperLogResponse 获取管理端操作日志响�?
type ListAdminOperLogResponse struct {
	common.Response
	HasMore    bool              `json:"has_more"`    // 是否还有下一�?
	NextCursor string            `json:"next_cursor"` // 下一页的分页游标
	RecordList []AdminOperRecord `json:"record_list"` // 记录列表
}

// AdminOperRecord 管理端操作记�?
type AdminOperRecord struct {
	Time       int64  `json:"time"`        // 操作时间
	UserID     string `json:"userid"`      // 操作者userid
	OperType   int    `json:"oper_type"`   // 操作类型
	DetailType int    `json:"detail_type"` // 操作行为
	DetailInfo string `json:"detail_info"` // 相关数据
	IP         string `json:"ip"`          // 操作者ip
}

