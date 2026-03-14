package security

import "github.com/mrseanchow/wecom-core/types/common"

// GetFileOperRecordRequest 获取文件操作记录请求
type GetFileOperRecordRequest struct {
	StartTime  int64      `json:"start_time"`            // 开始时�?
	EndTime    int64      `json:"end_time"`              // 结束时间，范围不能超�?4�?
	UserIDList []string   `json:"userid_list,omitempty"` // 需要查询的文件操作者的userid，单次最�?00�?
	Operation  *Operation `json:"operation,omitempty"`   // 操作类型过滤
	Cursor     string     `json:"cursor,omitempty"`      // 分页游标
	Limit      int        `json:"limit,omitempty"`       // 限制返回的条数，最�?000
}

// Operation 操作类型
type Operation struct {
	Type   int `json:"type,omitempty"`   // 操作类型
	Source int `json:"source,omitempty"` // 操作来源
}

// GetFileOperRecordResponse 获取文件操作记录响应
type GetFileOperRecordResponse struct {
	common.Response
	HasMore    bool             `json:"has_more"`    // 是否还有更多数据
	NextCursor string           `json:"next_cursor"` // 下一次调用的cursor�?
	RecordList []FileOperRecord `json:"record_list"` // 记录列表
}

// FileOperRecord 文件操作记录
type FileOperRecord struct {
	Time          int64         `json:"time"`                     // 操作时间
	UserID        string        `json:"userid,omitempty"`         // 企业用户账号id
	ExternalUser  *ExternalUser `json:"external_user,omitempty"`  // 企业外部人员账号信息
	Operation     *Operation    `json:"operation"`                // 操作类型
	FileInfo      string        `json:"file_info"`                // 文件操作说明
	FileMD5       string        `json:"file_md5,omitempty"`       // 文件的MD5
	FileSize      int64         `json:"file_size,omitempty"`      // 文件大小（字节）
	ApplicantName string        `json:"applicant_name,omitempty"` // 申请人的名字
	DeviceType    int           `json:"device_type,omitempty"`    // 设备类型
	DeviceCode    string        `json:"device_code,omitempty"`    // 设备编码
}

// ExternalUser 企业外部人员信息
type ExternalUser struct {
	Type     int    `json:"type"`                // 用户类型�?：微信用户；2：企业微信用�?
	Name     string `json:"name"`                // 用户�?
	CorpName string `json:"corp_name,omitempty"` // 企业名称
}

