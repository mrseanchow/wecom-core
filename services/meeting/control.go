package meeting

import (
	"context"

	"github.com/mrseanchow/wecom-core/internal/client"
	"github.com/mrseanchow/wecom-core/types/common"
	"github.com/mrseanchow/wecom-core/types/meeting"
)

// KickoutUsers 移出成员
// 文档: https://developer.work.weixin.qq.com/document/path/...
func (s *Service) KickoutUsers(ctx context.Context, req *meeting.KickoutUsersRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/meeting/realcontrol/kickout_users", req)
	return err
}

// SetNicknames 修改成员在会中显示的昵称
// 文档: https://developer.work.weixin.qq.com/document/path/...
func (s *Service) SetNicknames(ctx context.Context, req *meeting.SetNicknamesRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/meeting/realcontrol/set_nicknames", req)
	return err
}

// MuteUser 静音成员
// 文档: https://developer.work.weixin.qq.com/document/path/...
func (s *Service) MuteUser(ctx context.Context, req *meeting.MuteUserRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/meeting/realcontrol/mute_user", req)
	return err
}

// Dismiss 结束会议
// 文档: https://developer.work.weixin.qq.com/document/path/...
func (s *Service) Dismiss(ctx context.Context, req *meeting.DismissMeetingRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/meeting/realcontrol/dismiss", req)
	return err
}

// SetCohost 管理联席主持�?
// 文档: https://developer.work.weixin.qq.com/document/path/...
func (s *Service) SetCohost(ctx context.Context, req *meeting.SetCohostRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/meeting/realcontrol/set_cohost", req)
	return err
}

// SetMeetingSettings 管理会中设置
// 文档: https://developer.work.weixin.qq.com/document/path/...
func (s *Service) SetMeetingSettings(ctx context.Context, req *meeting.MeetingSettingsRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/meeting/realcontrol/set", req)
	return err
}

// ManageWaitingRoomUsers 管理等候室成员
// 文档: https://developer.work.weixin.qq.com/document/path/...
func (s *Service) ManageWaitingRoomUsers(ctx context.Context, req *meeting.ManageWaitingRoomUsersRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/meeting/realcontrol/manage_waiting_room_users", req)
	return err
}

// SwitchUserVideo 关闭或开启成员视�?
// 文档: https://developer.work.weixin.qq.com/document/path/...
func (s *Service) SwitchUserVideo(ctx context.Context, req *meeting.SwitchUserVideoRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/meeting/realcontrol/switch_user_video", req)
	return err
}

// CloseScreenShare 关闭成员屏幕共享
// 文档: https://developer.work.weixin.qq.com/document/path/...
func (s *Service) CloseScreenShare(ctx context.Context, req *meeting.CloseScreenShareRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/meeting/realcontrol/close_screen_share", req)
	return err
}

