package living

import "github.com/mrseanchow/wecom-core/types/common"

// ActivityDetail 活动直播详情
type ActivityDetail struct {
	Description string   `json:"description,omitempty"`
	ImageList   []string `json:"image_list,omitempty"`
}

// CreateRequest 创建预约直播请求
type CreateRequest struct {
	AnchorUserID         string          `json:"anchor_userid"`
	Theme                string          `json:"theme"`
	LivingStart          int64           `json:"living_start"`
	LivingDuration       int             `json:"living_duration"`
	Description          string          `json:"description,omitempty"`
	Type                 int             `json:"type,omitempty"`
	AgentID              int64           `json:"agentid,omitempty"`
	RemindTime           int             `json:"remind_time,omitempty"`
	ActivityCoverMediaID string          `json:"activity_cover_mediaid,omitempty"`
	ActivityShareMediaID string          `json:"activity_share_mediaid,omitempty"`
	ActivityDetail       *ActivityDetail `json:"activity_detail,omitempty"`
}

// CreateResponse 创建预约直播返回
type CreateResponse struct {
	common.Response
	LivingID string `json:"livingid,omitempty"`
}

// ModifyRequest 修改预约直播请求
type ModifyRequest struct {
	LivingID       string `json:"livingid"`
	Theme          string `json:"theme,omitempty"`
	LivingStart    int64  `json:"living_start,omitempty"`
	LivingDuration int    `json:"living_duration,omitempty"`
	Description    string `json:"description,omitempty"`
	Type           int    `json:"type,omitempty"`
	RemindTime     int    `json:"remind_time,omitempty"`
}

// CommonEmptyResponse �?errcode/errmsg 的返�?
type CommonEmptyResponse = common.Response

// DeleteReplayRequest 删除直播回放请求
type DeleteReplayRequest struct {
	LivingID string `json:"livingid"`
}

// CancelRequest 取消预约直播请求
type CancelRequest struct {
	LivingID string `json:"livingid"`
}

// GetLivingCodeRequest 获取微信观看直播凭证请求
type GetLivingCodeRequest struct {
	LivingID string `json:"livingid"`
	OpenID   string `json:"openid"`
}

// GetLivingCodeResponse 获取微信观看直播凭证返回
type GetLivingCodeResponse struct {
	common.Response
	LivingCode string `json:"living_code,omitempty"`
}

// GetUserAllLivingIDRequest 获取成员直播ID列表请求
type GetUserAllLivingIDRequest struct {
	UserID string `json:"userid"`
	Cursor string `json:"cursor,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

// GetUserAllLivingIDResponse 获取成员直播ID列表返回
type GetUserAllLivingIDResponse struct {
	common.Response
	NextCursor   string   `json:"next_cursor,omitempty"`
	LivingIDList []string `json:"livingid_list,omitempty"`
}

// GetWatchStatRequest 获取直播观看明细请求
type GetWatchStatRequest struct {
	LivingID string `json:"livingid"`
	NextKey  string `json:"next_key,omitempty"`
}

// UserWatchStat 企业成员观看统计
type UserWatchStat struct {
	UserID                string `json:"userid,omitempty"`
	WatchTime             int    `json:"watch_time,omitempty"`
	IsComment             int    `json:"is_comment,omitempty"`
	IsMic                 int    `json:"is_mic,omitempty"`
	InvitorUserID         string `json:"invitor_userid,omitempty"`
	InvitorExternalUserID string `json:"invitor_external_userid,omitempty"`
}

// ExternalWatchStat 外部成员观看统计
type ExternalWatchStat struct {
	ExternalUserID        string `json:"external_userid,omitempty"`
	Type                  int    `json:"type,omitempty"`
	Name                  string `json:"name,omitempty"`
	WatchTime             int    `json:"watch_time,omitempty"`
	IsComment             int    `json:"is_comment,omitempty"`
	IsMic                 int    `json:"is_mic,omitempty"`
	InvitorUserID         string `json:"invitor_userid,omitempty"`
	InvitorExternalUserID string `json:"invitor_external_userid,omitempty"`
}

// StatInfo 观看统计信息
type StatInfo struct {
	Users         []UserWatchStat     `json:"users,omitempty"`
	ExternalUsers []ExternalWatchStat `json:"external_users,omitempty"`
}

// GetWatchStatResponse 获取直播观看明细返回
type GetWatchStatResponse struct {
	common.Response
	Ending   int       `json:"ending,omitempty"`
	NextKey  string    `json:"next_key,omitempty"`
	StatInfo *StatInfo `json:"stat_info,omitempty"`
}

// GetLivingInfoResponse 获取直播详情返回
type GetLivingInfoResponse struct {
	common.Response
	LivingInfo *struct {
		Theme                 string `json:"theme,omitempty"`
		LivingStart           int64  `json:"living_start,omitempty"`
		LivingDuration        int    `json:"living_duration,omitempty"`
		Status                int    `json:"status,omitempty"`
		ReserveStart          int64  `json:"reserve_start,omitempty"`
		ReserveLivingDuration int    `json:"reserve_living_duration,omitempty"`
		Description           string `json:"description,omitempty"`
		AnchorUserID          string `json:"anchor_userid,omitempty"`
		MainDepartment        int64  `json:"main_department,omitempty"`
		ViewerNum             int    `json:"viewer_num,omitempty"`
		CommentNum            int    `json:"comment_num,omitempty"`
		MicNum                int    `json:"mic_num,omitempty"`
		OpenReplay            int    `json:"open_replay,omitempty"`
		ReplayStatus          int    `json:"replay_status,omitempty"`
		Type                  int    `json:"type,omitempty"`
		PushStreamURL         string `json:"push_stream_url,omitempty"`
		OnlineCount           int    `json:"online_count,omitempty"`
		SubscribeCount        int    `json:"subscribe_count,omitempty"`
	} `json:"living_info,omitempty"`
}

// GetLivingShareInfoRequest 获取跳转小程序商城的直播观众信息请求
type GetLivingShareInfoRequest struct {
	WWShareCode string `json:"ww_share_code"`
}

// GetLivingShareInfoResponse 获取跳转小程序商城的直播观众信息返回
type GetLivingShareInfoResponse struct {
	common.Response
	LivingID              string `json:"livingid,omitempty"`
	ViewerUserID          string `json:"viewer_userid,omitempty"`
	ViewerExternalUserID  string `json:"viewer_external_userid,omitempty"`
	InvitorUserID         string `json:"invitor_userid,omitempty"`
	InvitorExternalUserID string `json:"invitor_external_userid,omitempty"`
}

