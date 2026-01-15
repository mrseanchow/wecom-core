package meetingroom

import "github.com/shuaidd/wecom-core/types/common"

// GetBookingInfoRequest 查询会议室预定信息请求
type GetBookingInfoRequest struct {
	MeetingRoomID int64  `json:"meetingroom_id,omitempty"`
	StartTime     int64  `json:"start_time,omitempty"`
	EndTime       int64  `json:"end_time,omitempty"`
	City          string `json:"city,omitempty"`
	Building      string `json:"building,omitempty"`
	Floor         string `json:"floor,omitempty"`
}

// Schedule 预定情况
type Schedule struct {
	BookingID  string `json:"booking_id,omitempty"`
	ScheduleID string `json:"schedule_id,omitempty"`
	StartTime  int64  `json:"start_time,omitempty"`
	EndTime    int64  `json:"end_time,omitempty"`
	Booker     string `json:"booker,omitempty"`
	Status     int    `json:"status,omitempty"`
}

// BookingInfo 会议室预订信息
type BookingInfo struct {
	MeetingRoomID int64      `json:"meetingroom_id,omitempty"`
	Schedule      []Schedule `json:"schedule,omitempty"`
}

// GetBookingInfoResponse 查询会议室预定信息响应
type GetBookingInfoResponse struct {
	common.Response
	BookingList []BookingInfo `json:"booking_list,omitempty"`
}

// BookMeetingRoomRequest 预定会议室请求
type BookMeetingRoomRequest struct {
	MeetingRoomID int64    `json:"meetingroom_id"`
	Subject       string   `json:"subject,omitempty"`
	StartTime     int64    `json:"start_time"`
	EndTime       int64    `json:"end_time"`
	Booker        string   `json:"booker"`
	Attendees     []string `json:"attendees,omitempty"`
}

// BookMeetingRoomResponse 预定会议室响应
type BookMeetingRoomResponse struct {
	common.Response
	BookingID  string `json:"booking_id,omitempty"`
	ScheduleID string `json:"schedule_id,omitempty"`
}

// BookByScheduleRequest 通过日程预定会议室请求
type BookByScheduleRequest struct {
	MeetingRoomID int64  `json:"meetingroom_id"`
	ScheduleID    string `json:"schedule_id"`
	Booker        string `json:"booker"`
}

// BookByScheduleResponse 通过日程预定会议室响应
type BookByScheduleResponse struct {
	common.Response
	BookingID    string  `json:"booking_id,omitempty"`
	ConflictDate []int64 `json:"conflict_date,omitempty"`
}

// BookByMeetingRequest 通过会议预定会议室请求
type BookByMeetingRequest struct {
	MeetingRoomID int64  `json:"meetingroom_id"`
	MeetingID     string `json:"meetingid"`
	Booker        string `json:"booker"`
}

// BookByMeetingResponse 通过会议预定会议室响应
type BookByMeetingResponse struct {
	common.Response
	BookingID    string  `json:"booking_id,omitempty"`
	ConflictDate []int64 `json:"conflict_date,omitempty"`
}

// CancelBookRequest 取消预定会议室请求
type CancelBookRequest struct {
	BookingID    string `json:"booking_id"`
	KeepSchedule int    `json:"keep_schedule,omitempty"`
	CancelDate   int64  `json:"cancel_date,omitempty"`
}

// CancelBookResponse 取消预定会议室响应
type CancelBookResponse struct {
	common.Response
}

// GetBookInfoRequest 根据预定ID查询预定详情请求
type GetBookInfoRequest struct {
	MeetingRoomID int64  `json:"meetingroom_id"`
	BookingID     string `json:"booking_id"`
}

// BookSchedule 预定详情
type BookSchedule struct {
	BookingID       string `json:"booking_id,omitempty"`
	MasterBookingID string `json:"master_booking_id,omitempty"`
	ScheduleID      string `json:"schedule_id,omitempty"`
	StartTime       int64  `json:"start_time,omitempty"`
	EndTime         int64  `json:"end_time,omitempty"`
	Booker          string `json:"booker,omitempty"`
	Status          int    `json:"status,omitempty"`
}

// GetBookInfoResponse 根据预定ID查询预定详情响应
type GetBookInfoResponse struct {
	common.Response
	MeetingRoomID int64         `json:"meetingroom_id,omitempty"`
	Schedule      *BookSchedule `json:"schedule,omitempty"`
}
