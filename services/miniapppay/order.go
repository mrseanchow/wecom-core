package miniapppay

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/common"
	"github.com/shuaidd/wecom-core/types/miniapppay"
)

func (s *Service) CreateOrder(ctx context.Context, req *miniapppay.CreateOrderRequest) (*miniapppay.CreateOrderResponse, error) {
	return client.PostAndUnmarshal[miniapppay.CreateOrderResponse](s.client, ctx, "/cgi-bin/miniapppay/create_order", req)
}

func (s *Service) GetOrder(ctx context.Context, req *miniapppay.GetOrderRequest) (*miniapppay.GetOrderResponse, error) {
	return client.PostAndUnmarshal[miniapppay.GetOrderResponse](s.client, ctx, "/cgi-bin/miniapppay/get_order", req)
}

func (s *Service) CloseOrder(ctx context.Context, req *miniapppay.CloseOrderRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/miniapppay/close_order", req)
	return err
}
