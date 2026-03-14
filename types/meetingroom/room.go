package meetingroom

import "github.com/mrseanchow/wecom-core/types/common"

// Coordinate 坐标信息
type Coordinate struct {
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
}

// Range 使用范围
type Range struct {
	UserList       []string `json:"user_list,omitempty"`
	DepartmentList []int64  `json:"department_list,omitempty"`
}

// AddMeetingRoomRequest 添加会议室请求
type AddMeetingRoomRequest struct {
	Name       string      `json:"name"`
	Capacity   int         `json:"capacity"`
	City       string      `json:"city,omitempty"`
	Building   string      `json:"building,omitempty"`
	Floor      string      `json:"floor,omitempty"`
	Equipment  []int       `json:"equipment,omitempty"`
	Coordinate *Coordinate `json:"coordinate,omitempty"`
	Range      *Range      `json:"range,omitempty"`
}

// AddMeetingRoomResponse 添加会议室响应
type AddMeetingRoomResponse struct {
	common.Response
	MeetingRoomID int64 `json:"meetingroom_id,omitempty"`
}

// ListMeetingRoomsRequest 查询会议室请求
type ListMeetingRoomsRequest struct {
	City      string `json:"city,omitempty"`
	Building  string `json:"building,omitempty"`
	Floor     string `json:"floor,omitempty"`
	Equipment []int  `json:"equipment,omitempty"`
}

// MeetingRoom 会议室信息
type MeetingRoom struct {
	MeetingRoomID int64       `json:"meetingroom_id,omitempty"`
	Name          string      `json:"name,omitempty"`
	Capacity      int         `json:"capacity,omitempty"`
	City          string      `json:"city,omitempty"`
	Building      string      `json:"building,omitempty"`
	Floor         string      `json:"floor,omitempty"`
	Equipment     []int       `json:"equipment,omitempty"`
	Coordinate    *Coordinate `json:"coordinate,omitempty"`
	NeedApproval  int         `json:"need_approval,omitempty"`
	Range         *Range      `json:"range,omitempty"`
}

// ListMeetingRoomsResponse 查询会议室响应
type ListMeetingRoomsResponse struct {
	common.Response
	MeetingRoomList []MeetingRoom `json:"meetingroom_list,omitempty"`
}

// EditMeetingRoomRequest 编辑会议室请求
type EditMeetingRoomRequest struct {
	MeetingRoomID int64       `json:"meetingroom_id"`
	Name          string      `json:"name,omitempty"`
	Capacity      int         `json:"capacity,omitempty"`
	City          string      `json:"city,omitempty"`
	Building      string      `json:"building,omitempty"`
	Floor         string      `json:"floor,omitempty"`
	Equipment     []int       `json:"equipment,omitempty"`
	Coordinate    *Coordinate `json:"coordinate,omitempty"`
	Range         *Range      `json:"range,omitempty"`
}

// EditMeetingRoomResponse 编辑会议室响应
type EditMeetingRoomResponse struct {
	common.Response
}

// DeleteMeetingRoomRequest 删除会议室请求
type DeleteMeetingRoomRequest struct {
	MeetingRoomID int64 `json:"meetingroom_id"`
}

// DeleteMeetingRoomResponse 删除会议室响应
type DeleteMeetingRoomResponse struct {
	common.Response
}
