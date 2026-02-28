package hr

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/hr"
)
const (
	RGetStaffInfoURL = "/cgi-bin/hr/get_staff_info"
	RUpdateStaffInfoURL = "/cgi-bin/hr/update_staff_info"
)

func (s *Service) GetStaffInfo(ctx context.Context, req *hr.GetStaffInfoRequest) (*hr.GetStaffInfoResponse, error) {
	return client.PostAndUnmarshal[hr.GetStaffInfoResponse](s.client, ctx, RGetStaffInfoURL, req)
}

func (s *Service) UpdateStaffInfo(ctx context.Context, req *hr.UpdateStaffInfoRequest) (*hr.UpdateStaffInfoResponse, error) {
	return client.PostAndUnmarshal[hr.UpdateStaffInfoResponse](s.client, ctx, RUpdateStaffInfoURL, req)
}
