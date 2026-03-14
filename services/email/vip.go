package email

import (
	"context"

	"github.com/mrseanchow/wecom-core/internal/client"
	"github.com/mrseanchow/wecom-core/types/email"
)

// BatchAddVIP 分配高级功能账号
// 该接口可以为在应用可见范围的企业成员分配高级功能�?
//
// 文档: https://developer.work.weixin.qq.com/document/path/95848
func (s *Service) BatchAddVIP(ctx context.Context, req *email.BatchAddVIPRequest) (*email.BatchAddVIPResponse, error) {
	return client.PostAndUnmarshal[email.BatchAddVIPResponse](s.client, ctx, "/cgi-bin/exmail/vip/batch_add", req)
}

// BatchDelVIP 取消高级功能账号
// 该接口用于撤销分配应用可见范围的企业成员的高级功能�?
//
// 文档: https://developer.work.weixin.qq.com/document/path/95849
func (s *Service) BatchDelVIP(ctx context.Context, req *email.BatchDelVIPRequest) (*email.BatchDelVIPResponse, error) {
	return client.PostAndUnmarshal[email.BatchDelVIPResponse](s.client, ctx, "/cgi-bin/exmail/vip/batch_del", req)
}

// ListVIP 获取高级功能账号列表
// 该接口可以查询企业已分配高级功能且在应用可见范围的账号列表�?
//
// 文档: https://developer.work.weixin.qq.com/document/path/95850
func (s *Service) ListVIP(ctx context.Context, req *email.ListVIPRequest) (*email.ListVIPResponse, error) {
	return client.PostAndUnmarshal[email.ListVIPResponse](s.client, ctx, "/cgi-bin/exmail/vip/list", req)
}

