package kf

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/kf"
)
const (
	FServiceStateGetURL = "/cgi-bin/kf/service_state/get"
	FServiceStateTransURL = "/cgi-bin/kf/service_state/trans"
)

// GetServiceState 获取会话状态
// 文档: https://developer.work.weixin.qq.com/document/path/94669
func (s *Service) GetServiceState(ctx context.Context, req *kf.GetServiceStateRequest) (*kf.GetServiceStateResponse, error) {
	return client.PostAndUnmarshal[kf.GetServiceStateResponse](s.client, ctx, FServiceStateGetURL, req)
}

// TransServiceState 变更会话状态
// 文档: https://developer.work.weixin.qq.com/document/path/94669
func (s *Service) TransServiceState(ctx context.Context, req *kf.TransServiceStateRequest) (*kf.TransServiceStateResponse, error) {
	return client.PostAndUnmarshal[kf.TransServiceStateResponse](s.client, ctx, FServiceStateTransURL, req)
}
