package reserve_meeting

import "github.com/mrseanchow/wecom-core/types/common"

// GetInviteesRequest 获取会议受邀成员列表请求
type GetInviteesRequest struct {
	MeetingID string `json:"meetingid"`
	Cursor    string `json:"cursor,omitempty"`
}

// GetInviteesResponse 获取会议受邀成员列表响应
type GetInviteesResponse struct {
	common.Response
	HasMore    bool      `json:"has_more"`
	NextCursor string    `json:"next_cursor,omitempty"`
	Invitees   []Invitee `json:"invitees,omitempty"`
}

// Invitee 受邀成员
type Invitee struct {
	UserID string `json:"userid"`
}

// SetInviteesRequest 更新会议受邀成员列表请求
type SetInviteesRequest struct {
	MeetingID string    `json:"meetingid"`
	Invitees  []Invitee `json:"invitees,omitempty"`
}

// GetGuestsRequest 获取会议嘉宾列表请求
type GetGuestsRequest struct {
	MeetingID string `json:"meetingid"`
}

// GetGuestsResponse 获取会议嘉宾列表响应
type GetGuestsResponse struct {
	common.Response
	MeetingID   string  `json:"meetingid"`
	MeetingCode string  `json:"meeting_code"`
	Title       string  `json:"title"`
	Guests      []Guest `json:"guests"`
}

// SetGuestsRequest 更新会议嘉宾列表请求
type SetGuestsRequest struct {
	MeetingID string  `json:"meetingid"`
	Guests    []Guest `json:"guests"`
}

// GetQualityRequest 获取会议健康度请�?
type GetQualityRequest struct {
	MeetingID    string `json:"meetingid"`
	SubMeetingID string `json:"sub_meetingid,omitempty"`
	StartTime    int64  `json:"start_time"`
	Cursor       string `json:"cursor,omitempty"`
	Limit        int32  `json:"limit,omitempty"`
}

// GetQualityResponse 获取会议健康度响�?
type GetQualityResponse struct {
	common.Response
	Quality            int32             `json:"quality"`
	AudioQuality       int32             `json:"audio_quality"`
	VideoQuality       int32             `json:"video_quality"`
	ScreenShareQuality int32             `json:"screen_share_quality"`
	NetworkQuality     int32             `json:"network_quality"`
	Problems           []string          `json:"problems,omitempty"`
	Attendees          []QualityAttendee `json:"attendees,omitempty"`
	NextCursor         string            `json:"next_cursor,omitempty"`
	HasMore            bool              `json:"has_more"`
}

// QualityAttendee 参会人员健康�?
type QualityAttendee struct {
	UserID             string   `json:"userid"`
	TmpOpenID          string   `json:"tmp_openid"`
	InstanceID         int32    `json:"instance_id"`
	Quality            int32    `json:"quality"`
	AudioQuality       int32    `json:"audio_quality"`
	VideoQuality       int32    `json:"video_quality"`
	ScreenShareQuality int32    `json:"screen_share_quality"`
	NetworkQuality     int32    `json:"network_quality"`
	Problems           []string `json:"problems,omitempty"`
}

// GetRealtimeAttendeeListRequest 获取实时会中成员列表请求
type GetRealtimeAttendeeListRequest struct {
	MeetingID    string `json:"meetingid"`
	SubMeetingID string `json:"sub_meetingid,omitempty"`
	Cursor       string `json:"cursor,omitempty"`
	Limit        uint32 `json:"limit,omitempty"`
}

// GetRealtimeAttendeeListResponse 获取实时会中成员列表响应
type GetRealtimeAttendeeListResponse struct {
	common.Response
	Status     string            `json:"status"`
	HasMore    bool              `json:"has_more"`
	NextCursor string            `json:"next_cursor,omitempty"`
	Attendees  []MeetingAttendee `json:"attendees,omitempty"`
}

// MeetingAttendee 参会�?
type MeetingAttendee struct {
	UserID            string `json:"userid"`
	TmpOpenID         string `json:"tmp_openid"`
	JoinTime          string `json:"join_time"`
	InstanceID        uint32 `json:"instance_id"`
	Role              uint32 `json:"role"`
	JoinType          uint32 `json:"join_type"`
	AudioState        bool   `json:"audio_state"`
	VideoState        bool   `json:"video_state,omitempty"`
	ScreenSharedState bool   `json:"screen_shared_state,omitempty"`
}

// GetAttendeeListRequest 获取已参会成员列表请�?
type GetAttendeeListRequest struct {
	MeetingID    string `json:"meetingid"`
	SubMeetingID string `json:"sub_meetingid,omitempty"`
	StartTime    uint32 `json:"start_time,omitempty"`
	EndTime      uint32 `json:"end_time,omitempty"`
	Cursor       string `json:"cursor,omitempty"`
	Limit        uint32 `json:"limit,omitempty"`
}

// GetAttendeeListResponse 获取已参会成员列表响�?
type GetAttendeeListResponse struct {
	common.Response
	HasMore    bool              `json:"has_more"`
	NextCursor string            `json:"next_cursor,omitempty"`
	Attendees  []HistoryAttendee `json:"attendees,omitempty"`
}

// HistoryAttendee 历史参会�?
type HistoryAttendee struct {
	TmpOpenID         string `json:"tmp_openid"`
	UserID            string `json:"userid"`
	JoinTime          string `json:"join_time"`
	QuitTime          string `json:"quit_time"`
	InstanceID        uint32 `json:"instance_id"`
	Role              uint32 `json:"role"`
	WebinarRole       uint32 `json:"webinar_role,omitempty"`
	JoinType          uint32 `json:"join_type"`
	Net               string `json:"net,omitempty"`
	AudioState        bool   `json:"audio_state"`
	VideoState        bool   `json:"video_state,omitempty"`
	ScreenSharedState bool   `json:"screen_shared_state,omitempty"`
	CustomerData      string `json:"customer_data,omitempty"`
}

// WaitingRoomGetCurrentUserListRequest 获取实时等候室成员列表请求
type WaitingRoomGetCurrentUserListRequest struct {
	MeetingID string `json:"meetingid"`
	Limit     int32  `json:"limit,omitempty"`
	Cursor    string `json:"cursor,omitempty"`
}

// WaitingRoomGetCurrentUserListResponse 获取实时等候室成员列表响应
type WaitingRoomGetCurrentUserListResponse struct {
	common.Response
	HasMore    bool              `json:"has_more"`
	NextCursor string            `json:"next_cursor,omitempty"`
	UserList   []WaitingRoomUser `json:"user_list,omitempty"`
}

// WaitingRoomGetUserListRequest 获取等候室成员记录请求
type WaitingRoomGetUserListRequest struct {
	MeetingID string `json:"meetingid"`
	Limit     int32  `json:"limit,omitempty"`
	Cursor    string `json:"cursor,omitempty"`
}

// WaitingRoomGetUserListResponse 获取等候室成员记录响应
type WaitingRoomGetUserListResponse struct {
	common.Response
	HasMore    bool                     `json:"has_more"`
	NextCursor string                   `json:"next_cursor,omitempty"`
	UserList   []WaitingRoomHistoryUser `json:"user_list,omitempty"`
}

// WaitingRoomUser 等候室用户
type WaitingRoomUser struct {
	UserID       string `json:"userid"`
	InstanceID   int32  `json:"instance_id"`
	CustomerData string `json:"customer_data,omitempty"`
	TmpOpenID    string `json:"tmp_openid"`
}

// WaitingRoomHistoryUser 等候室历史用户
type WaitingRoomHistoryUser struct {
	UserID     string `json:"userid,omitempty"`
	TmpOpenID  string `json:"tmp_openid"`
	InstanceID int32  `json:"instance_id"`
	JoinTime   int64  `json:"join_time"`
	QuitTime   int64  `json:"quit_time"`
}

// CreateCustomerShortURLRequest 创建用户专属参会链接请求
type CreateCustomerShortURLRequest struct {
	MeetingID    string `json:"meetingid"`
	CustomerData string `json:"customer_data"`
}

// CreateCustomerShortURLResponse 创建用户专属参会链接响应
type CreateCustomerShortURLResponse struct {
	common.Response
	MeetingShortURLCustomerData []CustomerData `json:"meeting_short_url_customer_data"`
}

// GetCustomerShortURLRequest 获取用户专属参会链接请求
type GetCustomerShortURLRequest struct {
	MeetingID string `json:"meetingid"`
}

// GetCustomerShortURLResponse 获取用户专属参会链接响应
type GetCustomerShortURLResponse struct {
	common.Response
	MeetingShortURLCustomerDataList []CustomerData `json:"meeting_short_url_customer_data_list"`
}

// CustomerData 用户专属数据
type CustomerData struct {
	CustomerData    string `json:"customer_data"`
	MeetingShortURL string `json:"meeting_short_url"`
}

