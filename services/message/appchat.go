package message

import (
	"context"
	"net/url"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/common"
	"github.com/shuaidd/wecom-core/types/message"
)
const (
	EdocSmartsheetGroupchatListURL = "/cgi-bin/wedoc/smartsheet/groupchat/list"
	PpchatGetURL = "/cgi-bin/appchat/get"
	PpchatSendURL = "/cgi-bin/appchat/send"
	PpchatUpdateURL = "/cgi-bin/appchat/update"
)

// SendAppChat 应用推送消息
// 文档: https://developer.work.weixin.qq.com/document/path/90248
func (s *Service) SendAppChat(ctx context.Context, req *message.AppChatSendRequest) (*common.Response, error) {
	return client.PostAndUnmarshal[common.Response](s.client, ctx, PpchatSendURL, req)
}

// GetAppChat 获取群聊会话
// 文档: https://developer.work.weixin.qq.com/document/path/90247
func (s *Service) GetAppChat(ctx context.Context, chatID string) (*message.GetAppChatResponse, error) {
	query := url.Values{}
	query.Set("chatid", chatID)
	return client.GetAndUnmarshal[message.GetAppChatResponse](s.client, ctx, PpchatGetURL, query)
}

// UpdateAppChat 修改群聊会话
// 文档: https://developer.work.weixin.qq.com/document/path/90246
func (s *Service) UpdateAppChat(ctx context.Context, req *message.UpdateAppChatRequest) (*common.Response, error) {
	return client.PostAndUnmarshal[common.Response](s.client, ctx, PpchatUpdateURL, req)
}

// ListSmartsheetGroupChat 获取群聊列表
// 该接口可以查询通过智能表格自动化创建的群聊列表
// 文档: https://developer.work.weixin.qq.com/document/path/98149
func (s *Service) ListSmartsheetGroupChat(ctx context.Context, req *message.ListSmartsheetGroupChatRequest) (*message.ListSmartsheetGroupChatResponse, error) {
	return client.PostAndUnmarshal[message.ListSmartsheetGroupChatResponse](s.client, ctx, EdocSmartsheetGroupchatListURL, req)
}
