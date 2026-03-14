package message

import (
	"context"

	"github.com/mrseanchow/wecom-core/internal/client"
	"github.com/mrseanchow/wecom-core/types/message"
)

// Send 发送应用消�?
// 文档: https://developer.work.weixin.qq.com/document/path/90236
func (s *Service) Send(ctx context.Context, req *message.SendMessageRequest) (*message.SendMessageResponse, error) {
	return client.PostAndUnmarshal[message.SendMessageResponse](s.client, ctx, "/cgi-bin/message/send", req)
}

