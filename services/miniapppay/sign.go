package miniapppay

import (
	"context"

	"github.com/mrseanchow/wecom-core/internal/client"
	"github.com/mrseanchow/wecom-core/types/miniapppay"
)

func (s *Service) GetSign(ctx context.Context, req *miniapppay.GetSignRequest) (*miniapppay.GetSignResponse, error) {
	return client.PostAndUnmarshal[miniapppay.GetSignResponse](s.client, ctx, "/cgi-bin/miniapppay/get_sign", req)
}
