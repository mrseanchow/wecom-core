package miniapppay

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/common"
	"github.com/shuaidd/wecom-core/types/miniapppay"
)
const (
	IniapppayCloseOrderURL = "/cgi-bin/miniapppay/close_order"
	IniapppayCreateOrderURL = "/cgi-bin/miniapppay/create_order"
	IniapppayGetOrderURL = "/cgi-bin/miniapppay/get_order"
)

func (s *Service) CreateOrder(ctx context.Context, req *miniapppay.CreateOrderRequest) (*miniapppay.CreateOrderResponse, error) {
	return client.PostAndUnmarshal[miniapppay.CreateOrderResponse](s.client, ctx, IniapppayCreateOrderURL, req)
}

func (s *Service) GetOrder(ctx context.Context, req *miniapppay.GetOrderRequest) (*miniapppay.GetOrderResponse, error) {
	return client.PostAndUnmarshal[miniapppay.GetOrderResponse](s.client, ctx, IniapppayGetOrderURL, req)
}

func (s *Service) CloseOrder(ctx context.Context, req *miniapppay.CloseOrderRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, IniapppayCloseOrderURL, req)
	return err
}
