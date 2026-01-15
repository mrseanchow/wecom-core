package meetingroom

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/meetingroom"
)

// Service 会议室服务
type Service struct {
	client *client.Client
}

// NewService 创建会议室服务实例
func NewService(c *client.Client) *Service {
	return &Service{client: c}
}

// Add 添加会议室
// 文档: docs/会议室管理.md
func (s *Service) Add(ctx context.Context, req *meetingroom.AddMeetingRoomRequest) (*meetingroom.AddMeetingRoomResponse, error) {
	return client.PostAndUnmarshal[meetingroom.AddMeetingRoomResponse](s.client, ctx, "/cgi-bin/oa/meetingroom/add", req)
}

// List 查询会议室
// 文档: docs/会议室管理.md
func (s *Service) List(ctx context.Context, req *meetingroom.ListMeetingRoomsRequest) (*meetingroom.ListMeetingRoomsResponse, error) {
	return client.PostAndUnmarshal[meetingroom.ListMeetingRoomsResponse](s.client, ctx, "/cgi-bin/oa/meetingroom/list", req)
}

// Edit 编辑会议室
// 文档: docs/会议室管理.md
func (s *Service) Edit(ctx context.Context, req *meetingroom.EditMeetingRoomRequest) error {
	_, err := client.PostAndUnmarshal[meetingroom.EditMeetingRoomResponse](s.client, ctx, "/cgi-bin/oa/meetingroom/edit", req)
	return err
}

// Delete 删除会议室
// 文档: docs/会议室管理.md
func (s *Service) Delete(ctx context.Context, meetingRoomID int64) error {
	req := &meetingroom.DeleteMeetingRoomRequest{
		MeetingRoomID: meetingRoomID,
	}
	_, err := client.PostAndUnmarshal[meetingroom.DeleteMeetingRoomResponse](s.client, ctx, "/cgi-bin/oa/meetingroom/del", req)
	return err
}
