package externalpay

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/externalpay"
)

type Service struct {
	client *client.Client
}

func NewService(c *client.Client) *Service {
	return &Service{client: c}
}

func (s *Service) UploadImage(ctx context.Context, imageData []byte, contentType string) (*externalpay.UploadImageResponse, error) {
	return client.PostMultipartAndUnmarshal[externalpay.UploadImageResponse](s.client, ctx, "/cgi-bin/miniapppay/upload_image", imageData, contentType)
}

func (s *Service) GetMerchant(ctx context.Context, req *externalpay.GetMerchantRequest) (*externalpay.GetMerchantResponse, error) {
	return client.PostAndUnmarshal[externalpay.GetMerchantResponse](s.client, ctx, "/cgi-bin/externalpay/getmerchant", req)
}

func (s *Service) SetMchUseScope(ctx context.Context, req *externalpay.SetMchUseScopeRequest) (*externalpay.SetMchUseScopeResponse, error) {
	return client.PostAndUnmarshal[externalpay.SetMchUseScopeResponse](s.client, ctx, "/cgi-bin/externalpay/set_mch_use_scope", req)
}

func (s *Service) GetFundFlow(ctx context.Context, req *externalpay.GetFundFlowRequest) (*externalpay.GetFundFlowResponse, error) {
	return client.PostAndUnmarshal[externalpay.GetFundFlowResponse](s.client, ctx, "/cgi-bin/externalpay/get_fund_flow", req)
}

func (s *Service) GetBillList(ctx context.Context, req *externalpay.GetBillListRequest) (*externalpay.GetBillListResponse, error) {
	return client.PostAndUnmarshal[externalpay.GetBillListResponse](s.client, ctx, "/cgi-bin/externalpay/get_bill_list", req)
}

func (s *Service) GetPaymentInfo(ctx context.Context, req *externalpay.GetPaymentInfoRequest) (*externalpay.GetPaymentInfoResponse, error) {
	return client.PostAndUnmarshal[externalpay.GetPaymentInfoResponse](s.client, ctx, "/cgi-bin/externalpay/get_payment_info", req)
}

func (s *Service) ApplyMch(ctx context.Context, req *externalpay.ApplyMchRequest) (*externalpay.ApplyMchResponse, error) {
	return client.PostAndUnmarshal[externalpay.ApplyMchResponse](s.client, ctx, "/cgi-bin/miniapppay/apply_mch", req)
}

func (s *Service) GetApplymentStatus(ctx context.Context, req *externalpay.GetApplymentStatusRequest) (*externalpay.GetApplymentStatusResponse, error) {
	return client.PostAndUnmarshal[externalpay.GetApplymentStatusResponse](s.client, ctx, "/cgi-bin/miniapppay/get_applyment_status", req)
}
