package hr

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/hr"
)

func (s *Service) GetFields(ctx context.Context) (*hr.GetFieldsResponse, error) {
	return client.GetAndUnmarshal[hr.GetFieldsResponse](s.client, ctx, "/cgi-bin/hr/get_fields", nil)
}
