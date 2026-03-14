package meetingroom

import (
	"context"

	"github.com/mrseanchow/wecom-core/internal/client"
	"github.com/mrseanchow/wecom-core/types/meetingroom"
)

// GetBookingInfo 查询会议室的预定信息
// 文档: docs/会议室预定管�?md
func (s *Service) GetBookingInfo(ctx context.Context, req *meetingroom.GetBookingInfoRequest) (*meetingroom.GetBookingInfoResponse, error) {
	return client.PostAndUnmarshal[meetingroom.GetBookingInfoResponse](s.client, ctx, "/cgi-bin/oa/meetingroom/get_booking_info", req)
}

// Book 预定会议�?
// 文档: docs/会议室预定管�?md
func (s *Service) Book(ctx context.Context, req *meetingroom.BookMeetingRoomRequest) (*meetingroom.BookMeetingRoomResponse, error) {
	return client.PostAndUnmarshal[meetingroom.BookMeetingRoomResponse](s.client, ctx, "/cgi-bin/oa/meetingroom/book", req)
}

// BookBySchedule 通过日程预定会议�?
// 文档: docs/会议室预定管�?md
func (s *Service) BookBySchedule(ctx context.Context, req *meetingroom.BookByScheduleRequest) (*meetingroom.BookByScheduleResponse, error) {
	return client.PostAndUnmarshal[meetingroom.BookByScheduleResponse](s.client, ctx, "/cgi-bin/oa/meetingroom/book_by_schedule", req)
}

// BookByMeeting 通过会议预定会议�?
// 文档: docs/会议室预定管�?md
func (s *Service) BookByMeeting(ctx context.Context, req *meetingroom.BookByMeetingRequest) (*meetingroom.BookByMeetingResponse, error) {
	return client.PostAndUnmarshal[meetingroom.BookByMeetingResponse](s.client, ctx, "/cgi-bin/oa/meetingroom/book_by_meeting", req)
}

// CancelBook 取消预定会议�?
// 文档: docs/会议室预定管�?md
func (s *Service) CancelBook(ctx context.Context, req *meetingroom.CancelBookRequest) error {
	_, err := client.PostAndUnmarshal[meetingroom.CancelBookResponse](s.client, ctx, "/cgi-bin/oa/meetingroom/cancel_book", req)
	return err
}

// GetBookInfo 根据会议室预定ID查询预定详情
// 文档: docs/会议室预定管�?md
func (s *Service) GetBookInfo(ctx context.Context, req *meetingroom.GetBookInfoRequest) (*meetingroom.GetBookInfoResponse, error) {
	return client.PostAndUnmarshal[meetingroom.GetBookInfoResponse](s.client, ctx, "/cgi-bin/oa/meetingroom/bookinfo/get", req)
}

