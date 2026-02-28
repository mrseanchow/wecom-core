package miniapppay

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/miniapppay"
)
const (
	IniapppayGetRefundDetailURL = "/cgi-bin/miniapppay/get_refund_detail"
	IniapppayRefundURL = "/cgi-bin/miniapppay/refund"
)

func (s *Service) Refund(ctx context.Context, req *miniapppay.RefundRequest) (*miniapppay.RefundResponse, error) {
	return client.PostAndUnmarshal[miniapppay.RefundResponse](s.client, ctx, IniapppayRefundURL, req)
}

func (s *Service) GetRefund(ctx context.Context, req *miniapppay.GetRefundRequest) (*miniapppay.GetRefundResponse, error) {
	return client.PostAndUnmarshal[miniapppay.GetRefundResponse](s.client, ctx, IniapppayGetRefundDetailURL, req)
}
