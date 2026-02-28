package miniapppay

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/miniapppay"
)
const (
	IniapppayGetBillURL = "/cgi-bin/miniapppay/get_bill"
)

func (s *Service) GetBill(ctx context.Context, req *miniapppay.GetBillRequest) (*miniapppay.GetBillResponse, error) {
	return client.PostAndUnmarshal[miniapppay.GetBillResponse](s.client, ctx, IniapppayGetBillURL, req)
}
