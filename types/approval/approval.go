package approval

import (
	"encoding/json"

	"github.com/mrseanchow/wecom-core/types/common"
)

// SetUserVacationQuotaRequest 修改成员假期余额请求
type SetUserVacationQuotaRequest struct {
	UserID       string `json:"userid"`
	VacationID   int    `json:"vacation_id"`
	LeftDuration int64  `json:"leftduration"`
	TimeAttr     int    `json:"time_attr"`
	Remarks      string `json:"remarks,omitempty"`
}

// CreateTemplateRequest 创建审批模板请求
type CreateTemplateRequest struct {
	TemplateName    []LangText      `json:"template_name"`
	TemplateContent json.RawMessage `json:"template_content"`
}

// CreateTemplateResponse 创建审批模板响应
type CreateTemplateResponse struct {
	common.Response
	TemplateID string `json:"template_id,omitempty"`
}

// UpdateTemplateRequest 更新审批模板请求
type UpdateTemplateRequest struct {
	TemplateID      string          `json:"template_id"`
	TemplateName    []LangText      `json:"template_name"`
	TemplateContent json.RawMessage `json:"template_content"`
}

// GetTemplateDetailRequest 获取审批模板详情请求
type GetTemplateDetailRequest struct {
	TemplateID string `json:"template_id"`
}

// GetTemplateDetailResponse 获取审批模板详情响应（template_content 以原�?JSON 返回�?
type GetTemplateDetailResponse struct {
	common.Response
	TemplateNames   []LangText      `json:"template_names,omitempty"`
	TemplateContent json.RawMessage `json:"template_content,omitempty"`
}

// ApplyEventRequest 提交审批申请请求（applyevent�?
type ApplyEventRequest struct {
	CreatorUserID       string          `json:"creator_userid,omitempty"`
	TemplateID          string          `json:"template_id,omitempty"`
	UseTemplateApprover int             `json:"use_template_approver,omitempty"`
	ChooseDepartment    int             `json:"choose_department,omitempty"`
	Process             json.RawMessage `json:"process,omitempty"`
	ApplyData           json.RawMessage `json:"apply_data,omitempty"`
	SummaryList         json.RawMessage `json:"summary_list,omitempty"`
}

// ApplyEventResponse 提交审批申请响应
type ApplyEventResponse struct {
	common.Response
	SpNo string `json:"sp_no,omitempty"`
}

// GetApprovalInfoRequest 批量获取审批单号请求
type GetApprovalInfoRequest struct {
	StartTime string           `json:"starttime"`
	EndTime   string           `json:"endtime"`
	NewCursor string           `json:"new_cursor,omitempty"`
	Size      int              `json:"size,omitempty"`
	Filters   []ApprovalFilter `json:"filters,omitempty"`
}

// ApprovalFilter 筛选条�?
type ApprovalFilter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// GetApprovalInfoResponse 批量获取审批单号响应
type GetApprovalInfoResponse struct {
	common.Response
	SpNoList      []string `json:"sp_no_list,omitempty"`
	NewNextCursor string   `json:"new_next_cursor,omitempty"`
}

// GetCorpVacConfigResponse 获取企业假期管理配置响应
type GetCorpVacConfigResponse struct {
	common.Response
	Lists []VacationConf `json:"lists,omitempty"`
}

// VacationConf 假期配置项（保留常用字段，其它使�?RawMessage�?
type VacationConf struct {
	ID             int             `json:"id,omitempty"`
	Name           string          `json:"name,omitempty"`
	TimeAttr       int             `json:"time_attr,omitempty"`
	DurationType   int             `json:"duration_type,omitempty"`
	QuotaAttr      json.RawMessage `json:"quota_attr,omitempty"`
	PerdayDuration int64           `json:"perday_duration,omitempty"`
	ExpireRule     json.RawMessage `json:"expire_rule,omitempty"`
}

// GetUserVacationQuotaRequest 获取成员假期余额请求
type GetUserVacationQuotaRequest struct {
	UserID string `json:"userid"`
}

// GetUserVacationQuotaResponse 获取成员假期余额响应
type GetUserVacationQuotaResponse struct {
	common.Response
	Lists []UserVacationQuota `json:"lists,omitempty"`
}

// UserVacationQuota 成员假期余额�?
type UserVacationQuota struct {
	ID                 int    `json:"id,omitempty"`
	AssignDuration     int64  `json:"assignduration,omitempty"`
	UsedDuration       int64  `json:"usedduration,omitempty"`
	LeftDuration       int64  `json:"leftduration,omitempty"`
	VacationName       string `json:"vacationname,omitempty"`
	RealAssignDuration int64  `json:"real_assignduration,omitempty"`
}

// GetApprovalDetailRequest 获取审批申请详情请求
type GetApprovalDetailRequest struct {
	SpNo string `json:"sp_no"`
}

// GetApprovalDetailResponse 获取审批申请详情响应
// 由于审批详情结构较复杂，这里�?info 字段保留为原�?JSON，调用方可以自行解析
type GetApprovalDetailResponse struct {
	common.Response
	Info json.RawMessage `json:"info,omitempty"`
}

// GetApprovalDataOldRequest 旧接口：获取审批数据请求
type GetApprovalDataOldRequest struct {
	StartTime int64  `json:"starttime"`
	EndTime   int64  `json:"endtime"`
	NextSpNum string `json:"next_spnum,omitempty"`
}

// GetApprovalDataOldResponse 旧接口返回（data 字段原样返回�?
type GetApprovalDataOldResponse struct {
	common.Response
	Count     int             `json:"count,omitempty"`
	Total     int             `json:"total,omitempty"`
	NextSpNum json.RawMessage `json:"next_spnum,omitempty"`
	Data      json.RawMessage `json:"data,omitempty"`
}

// GetOpenApprovalDataRequest 自建应用查询审批单状态（getopenapprovaldata�?
type GetOpenApprovalDataRequest struct {
	ThirdNo string `json:"thirdNo"`
}

// GetOpenApprovalDataResponse 自建应用查询响应（data 原始�?
type GetOpenApprovalDataResponse struct {
	common.Response
	Data json.RawMessage `json:"data,omitempty"`
}

// LangText 简单多语言文本
type LangText struct {
	Text string `json:"text,omitempty"`
	Lang string `json:"lang,omitempty"`
}

