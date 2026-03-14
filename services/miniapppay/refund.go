package miniapppay

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/miniapppay"
)

func (s *Service) Refund(ctx context.Context, req *miniapppay.RefundRequest) (*miniapppay.RefundResponse, error) {
	return client.PostAndUnmarshal[miniapppay.RefundResponse](s.client, ctx, "/cgi-bin/miniapppay/refund", req)
}

func (s *Service) GetRefund(ctx context.Context, req *miniapppay.GetRefundRequest) (*miniapppay.GetRefundResponse, error) {
	return client.PostAndUnmarshal[miniapppay.GetRefundResponse](s.client, ctx, "/cgi-bin/miniapppay/get_refund_detail", req)
}
