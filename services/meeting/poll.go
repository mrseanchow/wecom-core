package meeting

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/common"
	"github.com/shuaidd/wecom-core/types/meeting"
)
const (
	EetingPollCreateThemeURL = "/cgi-bin/meeting/poll/create_theme"
	EetingPollDeleteURL = "/cgi-bin/meeting/poll/delete"
	EetingPollFinishURL = "/cgi-bin/meeting/poll/finish"
	EetingPollGetPollDetailURL = "/cgi-bin/meeting/poll/get_poll_detail"
	EetingPollGetPollListURL = "/cgi-bin/meeting/poll/get_poll_list"
	EetingPollGetThemeInfoURL = "/cgi-bin/meeting/poll/get_theme_info"
	EetingPollStartURL = "/cgi-bin/meeting/poll/start"
	EetingPollUpdateThemeURL = "/cgi-bin/meeting/poll/update_theme"
)

// CreatePollTheme 创建会议投票主题
// 文档: https://developer.work.weixin.qq.com/document/path/...
func (s *Service) CreatePollTheme(ctx context.Context, req *meeting.CreatePollThemeRequest) (*meeting.CreatePollThemeResponse, error) {
	return client.PostAndUnmarshal[meeting.CreatePollThemeResponse](s.client, ctx, EetingPollCreateThemeURL, req)
}

// UpdatePollTheme 修改会议投票主题
// 文档: https://developer.work.weixin.qq.com/document/path/...
func (s *Service) UpdatePollTheme(ctx context.Context, req *meeting.UpdatePollThemeRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, EetingPollUpdateThemeURL, req)
	return err
}

// DeletePoll 删除会议投票
// 文档: https://developer.work.weixin.qq.com/document/path/...
func (s *Service) DeletePoll(ctx context.Context, req *meeting.DeletePollRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, EetingPollDeleteURL, req)
	return err
}

// FinishPoll 结束会议投票
// 文档: https://developer.work.weixin.qq.com/document/path/...
func (s *Service) FinishPoll(ctx context.Context, req *meeting.FinishPollRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, EetingPollFinishURL, req)
	return err
}

// GetPollThemeInfo 获取会议投票主题信息
// 文档: https://developer.work.weixin.qq.com/document/path/...
func (s *Service) GetPollThemeInfo(ctx context.Context, req *meeting.GetPollThemeInfoRequest) (*meeting.GetPollThemeInfoResponse, error) {
	return client.PostAndUnmarshal[meeting.GetPollThemeInfoResponse](s.client, ctx, EetingPollGetThemeInfoURL, req)
}

// GetPollDetail 获取会议投票详情
// 文档: https://developer.work.weixin.qq.com/document/path/...
func (s *Service) GetPollDetail(ctx context.Context, req *meeting.GetPollDetailRequest) (*meeting.GetPollDetailResponse, error) {
	return client.PostAndUnmarshal[meeting.GetPollDetailResponse](s.client, ctx, EetingPollGetPollDetailURL, req)
}

// GetPollList 获取会议投票列表
// 文档: https://developer.work.weixin.qq.com/document/path/...
func (s *Service) GetPollList(ctx context.Context, req *meeting.GetPollListRequest) (*meeting.GetPollListResponse, error) {
	return client.PostAndUnmarshal[meeting.GetPollListResponse](s.client, ctx, EetingPollGetPollListURL, req)
}

// StartPoll 发起会议投票
// 文档: https://developer.work.weixin.qq.com/document/path/...
func (s *Service) StartPoll(ctx context.Context, req *meeting.StartPollRequest) (*meeting.StartPollResponse, error) {
	return client.PostAndUnmarshal[meeting.StartPollResponse](s.client, ctx, EetingPollStartURL, req)
}
