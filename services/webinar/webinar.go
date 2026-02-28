package webinar

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/common"
	"github.com/shuaidd/wecom-core/types/webinar"
)
const (
	EetingWebinarCancelURL = "/cgi-bin/meeting/webinar/cancel"
	EetingWebinarEnrollApproveURL = "/cgi-bin/meeting/webinar/enroll/approve"
	EetingWebinarEnrollDeleteURL = "/cgi-bin/meeting/webinar/enroll/delete"
	EetingWebinarEnrollGetConfigURL = "/cgi-bin/meeting/webinar/enroll/get_config"
	EetingWebinarEnrollListURL = "/cgi-bin/meeting/webinar/enroll/list"
	EetingWebinarEnrollQueryByTmpOpenidURL = "/cgi-bin/meeting/webinar/enroll/query_by_tmp_openid"
	EetingWebinarEnrollSetConfigURL = "/cgi-bin/meeting/webinar/enroll/set_config"
	EetingWebinarGetURL = "/cgi-bin/meeting/webinar/get"
	EetingWebinarListGuestURL = "/cgi-bin/meeting/webinar/list_guest"
	EetingWebinarUpdateURL = "/cgi-bin/meeting/webinar/update"
	EetingWebinarUpdateWarmUpURL = "/cgi-bin/meeting/webinar/update_warm_up"
)

type Service struct {
	client *client.Client
}

func NewService(c *client.Client) *Service {
	return &Service{client: c}
}

func (s *Service) Update(ctx context.Context, req *webinar.UpdateWebinarRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, EetingWebinarUpdateURL, req)
	return err
}

func (s *Service) SetConfig(ctx context.Context, req *webinar.SetWebinarConfigRequest) (*webinar.SetWebinarConfigResponse, error) {
	return client.PostAndUnmarshal[webinar.SetWebinarConfigResponse](s.client, ctx, EetingWebinarEnrollSetConfigURL, req)
}

func (s *Service) DeleteEnroll(ctx context.Context, req *webinar.DeleteWebinarEnrollRequest) (*webinar.DeleteWebinarEnrollResponse, error) {
	return client.PostAndUnmarshal[webinar.DeleteWebinarEnrollResponse](s.client, ctx, EetingWebinarEnrollDeleteURL, req)
}

func (s *Service) Cancel(ctx context.Context, meetingID string) error {
	req := &webinar.CancelWebinarRequest{
		MeetingID: meetingID,
	}
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, EetingWebinarCancelURL, req)
	return err
}

func (s *Service) ApproveEnroll(ctx context.Context, req *webinar.ApproveWebinarEnrollRequest) (*webinar.ApproveWebinarEnrollResponse, error) {
	return client.PostAndUnmarshal[webinar.ApproveWebinarEnrollResponse](s.client, ctx, EetingWebinarEnrollApproveURL, req)
}

func (s *Service) UpdateWarmUp(ctx context.Context, req *webinar.UpdateWebinarWarmUpRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, EetingWebinarUpdateWarmUpURL, req)
	return err
}

func (s *Service) ListGuest(ctx context.Context, req *webinar.ListWebinarGuestRequest) (*webinar.ListWebinarGuestResponse, error) {
	return client.PostAndUnmarshal[webinar.ListWebinarGuestResponse](s.client, ctx, EetingWebinarListGuestURL, req)
}

func (s *Service) QueryEnrollByTmpOpenID(ctx context.Context, req *webinar.QueryWebinarEnrollByTmpOpenIDRequest) (*webinar.QueryWebinarEnrollByTmpOpenIDResponse, error) {
	return client.PostAndUnmarshal[webinar.QueryWebinarEnrollByTmpOpenIDResponse](s.client, ctx, EetingWebinarEnrollQueryByTmpOpenidURL, req)
}

func (s *Service) ListEnroll(ctx context.Context, req *webinar.ListWebinarEnrollRequest) (*webinar.ListWebinarEnrollResponse, error) {
	return client.PostAndUnmarshal[webinar.ListWebinarEnrollResponse](s.client, ctx, EetingWebinarEnrollListURL, req)
}

func (s *Service) GetConfig(ctx context.Context, meetingID string) (*webinar.GetWebinarConfigResponse, error) {
	req := &webinar.GetWebinarConfigRequest{
		MeetingID: meetingID,
	}
	return client.PostAndUnmarshal[webinar.GetWebinarConfigResponse](s.client, ctx, EetingWebinarEnrollGetConfigURL, req)
}

func (s *Service) GetInfo(ctx context.Context, req *webinar.GetWebinarRequest) (*webinar.GetWebinarResponse, error) {
	return client.PostAndUnmarshal[webinar.GetWebinarResponse](s.client, ctx, EetingWebinarGetURL, req)
}
