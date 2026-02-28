package hr

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/hr"
)
const (
	RGetFieldsURL = "/cgi-bin/hr/get_fields"
)

func (s *Service) GetFields(ctx context.Context) (*hr.GetFieldsResponse, error) {
	return client.GetAndUnmarshal[hr.GetFieldsResponse](s.client, ctx, RGetFieldsURL, nil)
}
