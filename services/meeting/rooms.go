package meeting

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/common"
	"github.com/shuaidd/wecom-core/types/meeting"
)
const (
	EetingRoomsBookURL = "/cgi-bin/meeting/rooms/book"
	EetingRoomsCallURL = "/cgi-bin/meeting/rooms/call"
	EetingRoomsCancelCallURL = "/cgi-bin/meeting/rooms/cancel_call"
	EetingRoomsGetConfigURL = "/cgi-bin/meeting/rooms/get_config"
	EetingRoomsGetInfoURL = "/cgi-bin/meeting/rooms/get_info"
	EetingRoomsGetInventoryURL = "/cgi-bin/meeting/rooms/get_inventory"
	EetingRoomsGetResponseStatusURL = "/cgi-bin/meeting/rooms/get_response_status"
	EetingRoomsListControllersURL = "/cgi-bin/meeting/rooms/list_controllers"
	EetingRoomsListDevicesURL = "/cgi-bin/meeting/rooms/list_devices"
	EetingRoomsListMeetingsURL = "/cgi-bin/meeting/rooms/list_meetings"
	EetingRoomsListURL = "/cgi-bin/meeting/rooms/list"
	EetingRoomsReleaseURL = "/cgi-bin/meeting/rooms/release"
)

// ListMeetingRooms 获取企业下的Rooms会议室列表
// 文档: docs/获取Rooms会议室列表.md
func (s *Service) ListMeetingRooms(ctx context.Context, req *meeting.ListMeetingRoomsRequest) (*meeting.ListMeetingRoomsResponse, error) {
	return client.PostAndUnmarshal[meeting.ListMeetingRoomsResponse](s.client, ctx, EetingRoomsListURL, req)
}

// ListMeetings 获取指定Rooms会议室下的会议列表
// 文档: docs/获取Rooms会议室下的会议列表.md
func (s *Service) ListMeetings(ctx context.Context, req *meeting.ListMeetingsRequest) (*meeting.ListMeetingsResponse, error) {
	return client.PostAndUnmarshal[meeting.ListMeetingsResponse](s.client, ctx, EetingRoomsListMeetingsURL, req)
}

// GetRoomInfo 根据 Rooms 会议室 ID 获取会议室详情
// 文档: docs/获取Rooms会议室详情.md
func (s *Service) GetRoomInfo(ctx context.Context, req *meeting.GetRoomInfoRequest) (*meeting.GetRoomInfoResponse, error) {
	return client.PostAndUnmarshal[meeting.GetRoomInfoResponse](s.client, ctx, EetingRoomsGetInfoURL, req)
}

// GetRoomConfig 获取 Rooms 会议室配置项
// 文档: docs/获取Rooms会议室配置项.md
func (s *Service) GetRoomConfig(ctx context.Context, req *meeting.GetRoomConfigRequest) (*meeting.GetRoomConfigResponse, error) {
	return client.PostAndUnmarshal[meeting.GetRoomConfigResponse](s.client, ctx, EetingRoomsGetConfigURL, req)
}

// GetInventory 获取企业购买的 Rooms 会议室资源
// 文档: docs/获取Rooms会议室资源.md
func (s *Service) GetInventory(ctx context.Context) (*meeting.GetInventoryResponse, error) {
	return client.PostAndUnmarshal[meeting.GetInventoryResponse](s.client, ctx, EetingRoomsGetInventoryURL, nil)
}

// ListControllers 获取控制器列表
// 文档: docs/获取控制器列表.md
func (s *Service) ListControllers(ctx context.Context, req *meeting.ListControllersRequest) (*meeting.ListControllersResponse, error) {
	return client.PostAndUnmarshal[meeting.ListControllersResponse](s.client, ctx, EetingRoomsListControllersURL, req)
}

// ListDevices 获取设备列表
// 文档: docs/获取设备列表.md
func (s *Service) ListDevices(ctx context.Context, req *meeting.ListDevicesRequest) (*meeting.ListDevicesResponse, error) {
	return client.PostAndUnmarshal[meeting.ListDevicesResponse](s.client, ctx, EetingRoomsListDevicesURL, req)
}

// BookRooms 预定 Rooms 会议室
// 文档: docs/预定Rooms会议室.md
func (s *Service) BookRooms(ctx context.Context, req *meeting.BookRoomsRequest) (*meeting.BookRoomsResponse, error) {
	return client.PostAndUnmarshal[meeting.BookRoomsResponse](s.client, ctx, EetingRoomsBookURL, req)
}

// ReleaseRooms 通过会议 ID 释放 Rooms 会议室
// 文档: docs/释放Rooms会议室.md
func (s *Service) ReleaseRooms(ctx context.Context, req *meeting.ReleaseRoomsRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, EetingRoomsReleaseURL, req)
	return err
}

// CallRoom 呼叫 Rooms 会议室
// 文档: docs/呼叫Rooms会议室.md
func (s *Service) CallRoom(ctx context.Context, req *meeting.CallRoomRequest) (*meeting.CallRoomResponse, error) {
	return client.PostAndUnmarshal[meeting.CallRoomResponse](s.client, ctx, EetingRoomsCallURL, req)
}

// CancelCall 取消呼叫 Rooms 会议室
// 文档: docs/取消呼叫Rooms会议室.md
func (s *Service) CancelCall(ctx context.Context, req *meeting.CancelCallRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, EetingRoomsCancelCallURL, req)
	return err
}

// GetResponseStatus 获取 Rooms 会议室应答状态
// 文档: docs/获取Rooms会议室应答状态.md
func (s *Service) GetResponseStatus(ctx context.Context, req *meeting.GetResponseStatusRequest) (*meeting.GetResponseStatusResponse, error) {
	return client.PostAndUnmarshal[meeting.GetResponseStatusResponse](s.client, ctx, EetingRoomsGetResponseStatusURL, req)
}
