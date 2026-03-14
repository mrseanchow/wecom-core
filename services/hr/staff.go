package hr

import (
	"context"

	"github.com/mrseanchow/wecom-core/internal/client"
	"github.com/mrseanchow/wecom-core/types/hr"
)

func (s *Service) GetStaffInfo(ctx context.Context, req *hr.GetStaffInfoRequest) (*hr.GetStaffInfoResponse, error) {
	return client.PostAndUnmarshal[hr.GetStaffInfoResponse](s.client, ctx, "/cgi-bin/hr/get_staff_info", req)
}

func (s *Service) UpdateStaffInfo(ctx context.Context, req *hr.UpdateStaffInfoRequest) (*hr.UpdateStaffInfoResponse, error) {
	return client.PostAndUnmarshal[hr.UpdateStaffInfoResponse](s.client, ctx, "/cgi-bin/hr/update_staff_info", req)
}
