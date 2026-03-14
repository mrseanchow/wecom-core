package calendar

import (
	"context"

	"github.com/mrseanchow/wecom-core/internal/client"
	"github.com/mrseanchow/wecom-core/types/calendar"
	"github.com/mrseanchow/wecom-core/types/common"
)

const (
	addCalendarURL    = "/cgi-bin/oa/calendar/add"
	getCalendarURL    = "/cgi-bin/oa/calendar/get"
	updateCalendarURL = "/cgi-bin/oa/calendar/update"
	deleteCalendarURL = "/cgi-bin/oa/calendar/del"
)

// CreateCalendar 创建日历
// 该接口用于通过应用在企业内创建一个日�?
func (s *Service) CreateCalendar(ctx context.Context, req *calendar.CreateCalendarRequest) (*calendar.CreateCalendarResponse, error) {
	return client.PostAndUnmarshal[calendar.CreateCalendarResponse](s.client, ctx, addCalendarURL, req)
}

// GetCalendar 获取日历详情
// 该接口用于获取应用在企业内创建的日历信息
func (s *Service) GetCalendar(ctx context.Context, req *calendar.GetCalendarRequest) (*calendar.GetCalendarResponse, error) {
	return client.PostAndUnmarshal[calendar.GetCalendarResponse](s.client, ctx, getCalendarURL, req)
}

// UpdateCalendar 更新日历
// 该接口用于修改指定日历的信息
// 注意，更新操作是覆盖式，而不是增量式
func (s *Service) UpdateCalendar(ctx context.Context, req *calendar.UpdateCalendarRequest) (*calendar.UpdateCalendarResponse, error) {
	return client.PostAndUnmarshal[calendar.UpdateCalendarResponse](s.client, ctx, updateCalendarURL, req)
}

// DeleteCalendar 删除日历
// 该接口用于删除指定日�?
func (s *Service) DeleteCalendar(ctx context.Context, req *calendar.DeleteCalendarRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, deleteCalendarURL, req)
	return err
}

