package meeting

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/common"
	"github.com/shuaidd/wecom-core/types/meeting"
)
const (
	EetingRealcontrolCloseScreenShareURL = "/cgi-bin/meeting/realcontrol/close_screen_share"
	EetingRealcontrolDismissURL = "/cgi-bin/meeting/realcontrol/dismiss"
	EetingRealcontrolKickoutUsersURL = "/cgi-bin/meeting/realcontrol/kickout_users"
	EetingRealcontrolManageWaitingRoomUsersURL = "/cgi-bin/meeting/realcontrol/manage_waiting_room_users"
	EetingRealcontrolMuteUserURL = "/cgi-bin/meeting/realcontrol/mute_user"
	EetingRealcontrolSetCohostURL = "/cgi-bin/meeting/realcontrol/set_cohost"
	EetingRealcontrolSetNicknamesURL = "/cgi-bin/meeting/realcontrol/set_nicknames"
	EetingRealcontrolSetURL = "/cgi-bin/meeting/realcontrol/set"
	EetingRealcontrolSwitchUserVideoURL = "/cgi-bin/meeting/realcontrol/switch_user_video"
)

// KickoutUsers 移出成员
// 文档: https://developer.work.weixin.qq.com/document/path/...
func (s *Service) KickoutUsers(ctx context.Context, req *meeting.KickoutUsersRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, EetingRealcontrolKickoutUsersURL, req)
	return err
}

// SetNicknames 修改成员在会中显示的昵称
// 文档: https://developer.work.weixin.qq.com/document/path/...
func (s *Service) SetNicknames(ctx context.Context, req *meeting.SetNicknamesRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, EetingRealcontrolSetNicknamesURL, req)
	return err
}

// MuteUser 静音成员
// 文档: https://developer.work.weixin.qq.com/document/path/...
func (s *Service) MuteUser(ctx context.Context, req *meeting.MuteUserRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, EetingRealcontrolMuteUserURL, req)
	return err
}

// Dismiss 结束会议
// 文档: https://developer.work.weixin.qq.com/document/path/...
func (s *Service) Dismiss(ctx context.Context, req *meeting.DismissMeetingRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, EetingRealcontrolDismissURL, req)
	return err
}

// SetCohost 管理联席主持人
// 文档: https://developer.work.weixin.qq.com/document/path/...
func (s *Service) SetCohost(ctx context.Context, req *meeting.SetCohostRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, EetingRealcontrolSetCohostURL, req)
	return err
}

// SetMeetingSettings 管理会中设置
// 文档: https://developer.work.weixin.qq.com/document/path/...
func (s *Service) SetMeetingSettings(ctx context.Context, req *meeting.MeetingSettingsRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, EetingRealcontrolSetURL, req)
	return err
}

// ManageWaitingRoomUsers 管理等候室成员
// 文档: https://developer.work.weixin.qq.com/document/path/...
func (s *Service) ManageWaitingRoomUsers(ctx context.Context, req *meeting.ManageWaitingRoomUsersRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, EetingRealcontrolManageWaitingRoomUsersURL, req)
	return err
}

// SwitchUserVideo 关闭或开启成员视频
// 文档: https://developer.work.weixin.qq.com/document/path/...
func (s *Service) SwitchUserVideo(ctx context.Context, req *meeting.SwitchUserVideoRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, EetingRealcontrolSwitchUserVideoURL, req)
	return err
}

// CloseScreenShare 关闭成员屏幕共享
// 文档: https://developer.work.weixin.qq.com/document/path/...
func (s *Service) CloseScreenShare(ctx context.Context, req *meeting.CloseScreenShareRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, EetingRealcontrolCloseScreenShareURL, req)
	return err
}
