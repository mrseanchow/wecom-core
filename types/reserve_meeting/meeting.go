package reserve_meeting

import "github.com/mrseanchow/wecom-core/types/common"

// CreateMeetingRequest 创建预约会议请求
type CreateMeetingRequest struct {
	AdminUserID     string     `json:"admin_userid"`
	Title           string     `json:"title"`
	MeetingStart    int64      `json:"meeting_start,omitempty"`
	MeetingDuration int        `json:"meeting_duration,omitempty"`
	Description     string     `json:"description,omitempty"`
	Location        string     `json:"location,omitempty"`
	AgentID         uint32     `json:"agentid,omitempty"`
	Invitees        *Invitees  `json:"invitees,omitempty"`
	Guests          []Guest    `json:"guests,omitempty"`
	Settings        *Settings  `json:"settings,omitempty"`
	CalID           string     `json:"cal_id,omitempty"`
	Reminders       *Reminders `json:"reminders,omitempty"`
}

// CreateMeetingResponse 创建预约会议响应
type CreateMeetingResponse struct {
	common.Response
	MeetingID   string   `json:"meetingid"`
	ExcessUsers []string `json:"excess_users,omitempty"`
	MeetingCode string   `json:"meeting_code,omitempty"`
	MeetingLink string   `json:"meeting_link,omitempty"`
}

// UpdateMeetingRequest 修改预约会议请求
type UpdateMeetingRequest struct {
	MeetingID       string     `json:"meetingid"`
	Title           string     `json:"title,omitempty"`
	MeetingStart    int64      `json:"meeting_start,omitempty"`
	MeetingDuration int        `json:"meeting_duration,omitempty"`
	Description     string     `json:"description,omitempty"`
	Location        string     `json:"location,omitempty"`
	Invitees        *Invitees  `json:"invitees,omitempty"`
	Settings        *Settings  `json:"settings,omitempty"`
	CalID           string     `json:"cal_id,omitempty"`
	Reminders       *Reminders `json:"reminders,omitempty"`
}

// UpdateMeetingResponse 修改预约会议响应
type UpdateMeetingResponse struct {
	common.Response
	ExcessUsers []string `json:"excess_users,omitempty"`
}

// CancelMeetingRequest 取消预约会议请求
type CancelMeetingRequest struct {
	MeetingID    string `json:"meetingid"`
	SubMeetingID string `json:"sub_meetingid,omitempty"`
}

// GetMeetingInfoRequest 获取会议详情请求
type GetMeetingInfoRequest struct {
	MeetingID    string `json:"meetingid,omitempty"`
	MeetingCode  string `json:"meeting_code,omitempty"`
	SubMeetingID string `json:"sub_meetingid,omitempty"`
}

// GetMeetingInfoResponse 获取会议详情响应
type GetMeetingInfoResponse struct {
	common.Response
	AdminUserID         string           `json:"admin_userid"`
	Title               string           `json:"title"`
	MeetingStart        int64            `json:"meeting_start"`
	MeetingDuration     int              `json:"meeting_duration"`
	Description         string           `json:"description"`
	Location            string           `json:"location"`
	MainDepartment      uint32           `json:"main_department"`
	Status              uint32           `json:"status"`
	MeetingType         uint32           `json:"meeting_type"`
	Attendees           *Attendees       `json:"attendees"`
	Settings            *MeetingSettings `json:"settings"`
	CalID               string           `json:"cal_id"`
	Reminders           *Reminders       `json:"reminders"`
	MeetingCode         string           `json:"meeting_code"`
	MeetingLink         string           `json:"meeting_link"`
	HasVote             bool             `json:"has_vote"`
	HasMoreSubMeeting   uint32           `json:"has_more_sub_meeting"`
	RemainSubMeetings   uint32           `json:"remain_sub_meetings"`
	CurrentSubMeetingID string           `json:"current_sub_meetingid"`
	SubMeetings         []SubMeeting     `json:"sub_meetings"`
	Guests              []Guest          `json:"guests"`
	SubRepeatList       []SubRepeatInfo  `json:"sub_repeat_list"`
}

// GetUserMeetingIDsRequest 获取成员会议ID列表请求
type GetUserMeetingIDsRequest struct {
	UserID    string `json:"userid"`
	Cursor    string `json:"cursor,omitempty"`
	BeginTime int64  `json:"begin_time,omitempty"`
	EndTime   int64  `json:"end_time,omitempty"`
	Limit     uint32 `json:"limit,omitempty"`
}

// GetUserMeetingIDsResponse 获取成员会议ID列表响应
type GetUserMeetingIDsResponse struct {
	common.Response
	NextCursor    string   `json:"next_cursor,omitempty"`
	MeetingIDList []string `json:"meetingid_list,omitempty"`
}

// CheckDeviceInMeetingRequest 获取成员设备是否入会请求
type CheckDeviceInMeetingRequest struct {
	UserID         string   `json:"userid"`
	InstanceIDList []int32  `json:"instance_id_list,omitempty"`
	MeetingIDList  []string `json:"meetingid_list"`
}

// CheckDeviceInMeetingResponse 获取成员设备是否入会响应
type CheckDeviceInMeetingResponse struct {
	common.Response
	ResultList []CheckDeviceResult `json:"result_list,omitempty"`
}

// CheckDeviceResult 设备检查结�?
type CheckDeviceResult struct {
	MeetingID  string `json:"meetingid"`
	InstanceID int32  `json:"instance_id"`
}

// Invitees 受邀成员
type Invitees struct {
	UserID []string `json:"userid,omitempty"`
}

// Guest 会议嘉宾
type Guest struct {
	Area        string `json:"area"`
	PhoneNumber string `json:"phone_number"`
	GuestName   string `json:"guest_name,omitempty"`
}

// Settings 会议配置
type Settings struct {
	Password                  string     `json:"password,omitempty"`
	EnableWaitingRoom         bool       `json:"enable_waiting_room,omitempty"`
	AllowEnterBeforeHost      bool       `json:"allow_enter_before_host,omitempty"`
	EnableEnterMute           uint32     `json:"enable_enter_mute,omitempty"`
	AllowUnmuteSelf           bool       `json:"allow_unmute_self,omitempty"`
	MuteAll                   bool       `json:"mute_all,omitempty"`
	AllowExternalUser         bool       `json:"allow_external_user,omitempty"`
	EnableScreenWatermark     bool       `json:"enable_screen_watermark,omitempty"`
	WatermarkType             uint32     `json:"watermark_type,omitempty"`
	AutoRecordType            string     `json:"auto_record_type,omitempty"`
	AttendeeJoinAutoRecord    bool       `json:"attendee_join_auto_record,omitempty"`
	EnableHostPauseAutoRecord bool       `json:"enable_host_pause_auto_record,omitempty"`
	EnableInterpreter         bool       `json:"enable_interpreter,omitempty"`
	EnableEnroll              bool       `json:"enable_enroll,omitempty"`
	EnableHostKey             bool       `json:"enable_host_key,omitempty"`
	HostKey                   string     `json:"host_key,omitempty"`
	Hosts                     *Hosts     `json:"hosts,omitempty"`
	RemindScope               uint32     `json:"remind_scope,omitempty"`
	RingUsers                 *RingUsers `json:"ring_users,omitempty"`
}

// Hosts 会议主持人列�?
type Hosts struct {
	UserID []string `json:"userid,omitempty"`
}

// RingUsers 指定响铃的成员列�?
type RingUsers struct {
	UserID []string `json:"userid,omitempty"`
}

// Reminders 重复会议相关配置
type Reminders struct {
	IsRepeat         uint32   `json:"is_repeat,omitempty"`
	RepeatType       uint32   `json:"repeat_type,omitempty"`
	IsCustomRepeat   uint32   `json:"is_custom_repeat,omitempty"`
	RepeatUntilType  uint32   `json:"repeat_until_type,omitempty"`
	RepeatUntilCount uint32   `json:"repeat_until_count,omitempty"`
	RepeatUntil      int64    `json:"repeat_until,omitempty"`
	RepeatInterval   uint32   `json:"repeat_interval,omitempty"`
	RepeatDayOfWeek  []uint32 `json:"repeat_day_of_week,omitempty"`
	RepeatDayOfMonth []uint32 `json:"repeat_day_of_month,omitempty"`
	RemindBefore     []uint32 `json:"remind_before,omitempty"`
}

// SubMeeting 周期性子会议
type SubMeeting struct {
	SubMeetingID string `json:"sub_meetingid"`
	Status       uint32 `json:"status"`
	StartTime    int64  `json:"start_time"`
	EndTime      int64  `json:"end_time"`
	Title        string `json:"title"`
	RepeatID     string `json:"repeat_id"`
}

// SubRepeatInfo 周期性会议分段信�?
type SubRepeatInfo struct {
	RepeatID         uint32   `json:"repeat_id"`
	RepeatType       uint32   `json:"repeat_type"`
	RepeatUntilType  uint32   `json:"repeat_until_type"`
	RepeatUntilCount uint32   `json:"repeat_until_count,omitempty"`
	RepeatUntil      int64    `json:"repeat_until,omitempty"`
	RepeatInterval   uint32   `json:"repeat_interval"`
	IsCustomRepeat   uint32   `json:"is_custom_repeat"`
	RepeatDayOfWeek  []uint32 `json:"repeat_day_of_week,omitempty"`
	RepeatDayOfMonth []uint32 `json:"repeat_day_of_month,omitempty"`
}

// Attendees 会议成员
type Attendees struct {
	Member          []AttendeeMember          `json:"member,omitempty"`
	TmpExternalUser []AttendeeTmpExternalUser `json:"tmp_external_user,omitempty"`
}

// AttendeeMember 企业内部成员
type AttendeeMember struct {
	UserID         string `json:"userid"`
	Status         uint32 `json:"status"`
	FirstJoinTime  int64  `json:"first_join_time"`
	LastQuitTime   int64  `json:"last_quit_time"`
	TotalJoinCount uint32 `json:"total_join_count"`
	CumulativeTime uint32 `json:"cumulative_time"`
}

// AttendeeTmpExternalUser 会中参会的外部联系人
type AttendeeTmpExternalUser struct {
	TmpExternalUserID string `json:"tmp_external_userid"`
	Status            uint32 `json:"status"`
	FirstJoinTime     int64  `json:"first_join_time"`
	LastQuitTime      int64  `json:"last_quit_time"`
	TotalJoinCount    uint32 `json:"total_join_count"`
	CumulativeTime    uint32 `json:"cumulative_time"`
}

// MeetingSettings 会议配置（获取详情返回）
type MeetingSettings struct {
	RemindScope               uint32     `json:"remind_scope"`
	NeedPassword              bool       `json:"need_password"`
	Password                  string     `json:"password"`
	EnableWaitingRoom         bool       `json:"enable_waiting_room"`
	AllowEnterBeforeHost      bool       `json:"allow_enter_before_host"`
	EnableEnterMute           uint32     `json:"enable_enter_mute"`
	EnableEnterMuteType       uint32     `json:"enable_enter_mute_type"`
	AllowUnmuteSelf           bool       `json:"allow_unmute_self"`
	AllowExternalUser         bool       `json:"allow_external_user"`
	EnableScreenWatermark     bool       `json:"enable_screen_watermark"`
	WatermarkType             uint32     `json:"watermark_type"`
	AutoRecordType            string     `json:"auto_record_type"`
	AttendeeJoinAutoRecord    bool       `json:"attendee_join_auto_record"`
	EnableHostPauseAutoRecord bool       `json:"enable_host_pause_auto_record"`
	EnableDocUploadPermission bool       `json:"enable_doc_upload_permission"`
	EnableEnroll              bool       `json:"enable_enroll"`
	EnableHostKey             bool       `json:"enable_host_key"`
	HostKey                   string     `json:"host_key"`
	Hosts                     *Hosts     `json:"hosts"`
	CurrentHosts              *Hosts     `json:"current_hosts"`
	CoHosts                   *Hosts     `json:"co_hosts"`
	RingUsers                 *RingUsers `json:"ring_users"`
}

