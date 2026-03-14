package meetingroom

import (
	"context"

	"github.com/mrseanchow/wecom-core/internal/client"
	"github.com/mrseanchow/wecom-core/types/meetingroom"
)

// Service 会议室服�?
type Service struct {
	client *client.Client
}

// NewService 创建会议室服务实�?
func NewService(c *client.Client) *Service {
	return &Service{client: c}
}

// Add 添加会议�?
// 文档: docs/会议室管�?md
func (s *Service) Add(ctx context.Context, req *meetingroom.AddMeetingRoomRequest) (*meetingroom.AddMeetingRoomResponse, error) {
	return client.PostAndUnmarshal[meetingroom.AddMeetingRoomResponse](s.client, ctx, "/cgi-bin/oa/meetingroom/add", req)
}

// List 查询会议�?
// 文档: docs/会议室管�?md
func (s *Service) List(ctx context.Context, req *meetingroom.ListMeetingRoomsRequest) (*meetingroom.ListMeetingRoomsResponse, error) {
	return client.PostAndUnmarshal[meetingroom.ListMeetingRoomsResponse](s.client, ctx, "/cgi-bin/oa/meetingroom/list", req)
}

// Edit 编辑会议�?
// 文档: docs/会议室管�?md
func (s *Service) Edit(ctx context.Context, req *meetingroom.EditMeetingRoomRequest) error {
	_, err := client.PostAndUnmarshal[meetingroom.EditMeetingRoomResponse](s.client, ctx, "/cgi-bin/oa/meetingroom/edit", req)
	return err
}

// Delete 删除会议�?
// 文档: docs/会议室管�?md
func (s *Service) Delete(ctx context.Context, meetingRoomID int64) error {
	req := &meetingroom.DeleteMeetingRoomRequest{
		MeetingRoomID: meetingRoomID,
	}
	_, err := client.PostAndUnmarshal[meetingroom.DeleteMeetingRoomResponse](s.client, ctx, "/cgi-bin/oa/meetingroom/del", req)
	return err
}

