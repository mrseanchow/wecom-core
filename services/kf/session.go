package kf

import (
	"context"

	"github.com/mrseanchow/wecom-core/internal/client"
	"github.com/mrseanchow/wecom-core/types/kf"
)

// GetServiceState 获取会话状�?
// 文档: https://developer.work.weixin.qq.com/document/path/94669
func (s *Service) GetServiceState(ctx context.Context, req *kf.GetServiceStateRequest) (*kf.GetServiceStateResponse, error) {
	return client.PostAndUnmarshal[kf.GetServiceStateResponse](s.client, ctx, "/cgi-bin/kf/service_state/get", req)
}

// TransServiceState 变更会话状�?
// 文档: https://developer.work.weixin.qq.com/document/path/94669
func (s *Service) TransServiceState(ctx context.Context, req *kf.TransServiceStateRequest) (*kf.TransServiceStateResponse, error) {
	return client.PostAndUnmarshal[kf.TransServiceStateResponse](s.client, ctx, "/cgi-bin/kf/service_state/trans", req)
}

