package wedoc

import (
	"context"

	"github.com/mrseanchow/wecom-core/internal/client"
	"github.com/mrseanchow/wecom-core/types/wedoc"
)

const (
	batchAddVipURL = "/cgi-bin/wedoc/vip/batch_add"
	listVipURL     = "/cgi-bin/wedoc/vip/list"
	batchDelVipURL = "/cgi-bin/wedoc/vip/batch_del"
)

// BatchAddVip 分配高级功能账号
// 该接口用于分配应用可见范围内企业成员的高级功�?
func (s *Service) BatchAddVip(ctx context.Context, req *wedoc.BatchAddVipRequest) (*wedoc.BatchAddVipResponse, error) {
	return client.PostAndUnmarshal[wedoc.BatchAddVipResponse](s.client, ctx, batchAddVipURL, req)
}

// ListVip 获取高级功能账号列表
// 该接口用于查询企业已分配高级功能且在应用可见范围的账号列�?
func (s *Service) ListVip(ctx context.Context, req *wedoc.ListVipRequest) (*wedoc.ListVipResponse, error) {
	return client.PostAndUnmarshal[wedoc.ListVipResponse](s.client, ctx, listVipURL, req)
}

// BatchDelVip 取消高级功能账号
// 该接口用于撤销分配应用可见范围企业成员的高级功�?
func (s *Service) BatchDelVip(ctx context.Context, req *wedoc.BatchDelVipRequest) (*wedoc.BatchDelVipResponse, error) {
	return client.PostAndUnmarshal[wedoc.BatchDelVipResponse](s.client, ctx, batchDelVipURL, req)
}

