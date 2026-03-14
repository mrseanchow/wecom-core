package reserve_meeting

import (
	"context"

	"github.com/mrseanchow/wecom-core/internal/client"
	"github.com/mrseanchow/wecom-core/types/common"
	"github.com/mrseanchow/wecom-core/types/reserve_meeting"
)

// Service 预约会议高级管理服务
type Service struct {
	client *client.Client
}

// NewService 创建预约会议高级管理服务实例
func NewService(c *client.Client) *Service {
	return &Service{client: c}
}

// Create 创建预约会议
// 文档: https://developer.work.weixin.qq.com/document/path/93624
func (s *Service) Create(ctx context.Context, req *reserve_meeting.CreateMeetingRequest) (*reserve_meeting.CreateMeetingResponse, error) {
	return client.PostAndUnmarshal[reserve_meeting.CreateMeetingResponse](s.client, ctx, "/cgi-bin/meeting/create", req)
}

// Update 修改预约会议
// 文档: https://developer.work.weixin.qq.com/document/path/93625
func (s *Service) Update(ctx context.Context, req *reserve_meeting.UpdateMeetingRequest) error {
	_, err := client.PostAndUnmarshal[reserve_meeting.UpdateMeetingResponse](s.client, ctx, "/cgi-bin/meeting/update", req)
	return err
}

// Cancel 取消预约会议
// 文档: https://developer.work.weixin.qq.com/document/path/93626
func (s *Service) Cancel(ctx context.Context, meetingID string) error {
	req := &reserve_meeting.CancelMeetingRequest{
		MeetingID: meetingID,
	}
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/meeting/cancel", req)
	return err
}

// CancelWithSubMeeting 取消周期性子会议
// 文档: https://developer.work.weixin.qq.com/document/path/93626
func (s *Service) CancelWithSubMeeting(ctx context.Context, meetingID, subMeetingID string) error {
	req := &reserve_meeting.CancelMeetingRequest{
		MeetingID:    meetingID,
		SubMeetingID: subMeetingID,
	}
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/meeting/cancel", req)
	return err
}

// GetInfo 获取会议详情
// 文档: https://developer.work.weixin.qq.com/document/path/93628
func (s *Service) GetInfo(ctx context.Context, meetingID string) (*reserve_meeting.GetMeetingInfoResponse, error) {
	req := &reserve_meeting.GetMeetingInfoRequest{
		MeetingID: meetingID,
	}
	return client.PostAndUnmarshal[reserve_meeting.GetMeetingInfoResponse](s.client, ctx, "/cgi-bin/meeting/get_info", req)
}

// GetInfoByCode 通过会议号获取会议详�?
// 文档: https://developer.work.weixin.qq.com/document/path/93628
func (s *Service) GetInfoByCode(ctx context.Context, meetingCode string) (*reserve_meeting.GetMeetingInfoResponse, error) {
	req := &reserve_meeting.GetMeetingInfoRequest{
		MeetingCode: meetingCode,
	}
	return client.PostAndUnmarshal[reserve_meeting.GetMeetingInfoResponse](s.client, ctx, "/cgi-bin/meeting/get_info", req)
}

// GetSubMeetingInfo 获取周期性子会议详情
// 文档: https://developer.work.weixin.qq.com/document/path/93628
func (s *Service) GetSubMeetingInfo(ctx context.Context, meetingID, subMeetingID string) (*reserve_meeting.GetMeetingInfoResponse, error) {
	req := &reserve_meeting.GetMeetingInfoRequest{
		MeetingID:    meetingID,
		SubMeetingID: subMeetingID,
	}
	return client.PostAndUnmarshal[reserve_meeting.GetMeetingInfoResponse](s.client, ctx, "/cgi-bin/meeting/get_info", req)
}

// GetUserMeetingIDs 获取成员会议ID列表
// 文档: https://developer.work.weixin.qq.com/document/path/93854
func (s *Service) GetUserMeetingIDs(ctx context.Context, req *reserve_meeting.GetUserMeetingIDsRequest) (*reserve_meeting.GetUserMeetingIDsResponse, error) {
	return client.PostAndUnmarshal[reserve_meeting.GetUserMeetingIDsResponse](s.client, ctx, "/cgi-bin/meeting/get_user_meetingid", req)
}

// CheckDeviceInMeeting 获取成员设备是否入会
// 文档: https://developer.work.weixin.qq.com/document/path/95026
func (s *Service) CheckDeviceInMeeting(ctx context.Context, req *reserve_meeting.CheckDeviceInMeetingRequest) (*reserve_meeting.CheckDeviceInMeetingResponse, error) {
	return client.PostAndUnmarshal[reserve_meeting.CheckDeviceInMeetingResponse](s.client, ctx, "/cgi-bin/meeting/check_device_in_meeting", req)
}

// GetQuality 获取会议健康�?
// 文档: https://developer.work.weixin.qq.com/document/path/94060
func (s *Service) GetQuality(ctx context.Context, req *reserve_meeting.GetQualityRequest) (*reserve_meeting.GetQualityResponse, error) {
	return client.PostAndUnmarshal[reserve_meeting.GetQualityResponse](s.client, ctx, "/cgi-bin/meeting/get_quality", req)
}

