package miniapppay

import (
	"context"

	"github.com/mrseanchow/wecom-core/internal/client"
	"github.com/mrseanchow/wecom-core/types/miniapppay"
)

func (s *Service) GetBill(ctx context.Context, req *miniapppay.GetBillRequest) (*miniapppay.GetBillResponse, error) {
	return client.PostAndUnmarshal[miniapppay.GetBillResponse](s.client, ctx, "/cgi-bin/miniapppay/get_bill", req)
}

