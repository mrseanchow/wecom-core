package miniapppay

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/miniapppay"
)
const (
	IniapppayGetSignURL = "/cgi-bin/miniapppay/get_sign"
)

func (s *Service) GetSign(ctx context.Context, req *miniapppay.GetSignRequest) (*miniapppay.GetSignResponse, error) {
	return client.PostAndUnmarshal[miniapppay.GetSignResponse](s.client, ctx, IniapppayGetSignURL, req)
}
