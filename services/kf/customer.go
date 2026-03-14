package kf

import (
	"context"
	"net/url"

	"github.com/mrseanchow/wecom-core/internal/client"
	"github.com/mrseanchow/wecom-core/types/common"
	"github.com/mrseanchow/wecom-core/types/kf"
)

// BatchGetCustomer 批量获取客户基础信息
// 文档: https://developer.work.weixin.qq.com/document/path/95159
func (s *Service) BatchGetCustomer(ctx context.Context, req *kf.BatchGetCustomerRequest) (*kf.BatchGetCustomerResponse, error) {
	return client.PostAndUnmarshal[kf.BatchGetCustomerResponse](s.client, ctx, "/cgi-bin/kf/customer/batchget", req)
}

// GetUpgradeServiceConfig 获取配置的专员与客户�?
// 文档: https://developer.work.weixin.qq.com/document/path/94674
func (s *Service) GetUpgradeServiceConfig(ctx context.Context) (*kf.GetUpgradeServiceConfigResponse, error) {
	return client.GetAndUnmarshal[kf.GetUpgradeServiceConfigResponse](s.client, ctx, "/cgi-bin/kf/customer/get_upgrade_service_config", url.Values{})
}

// UpgradeService 为客户升级为专员或客户群服务
// 文档: https://developer.work.weixin.qq.com/document/path/94675
func (s *Service) UpgradeService(ctx context.Context, req *kf.UpgradeServiceRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/kf/customer/upgrade_service", req)
	return err
}

// CancelUpgradeService 为客户取消推�?
// 文档: https://developer.work.weixin.qq.com/document/path/94676
func (s *Service) CancelUpgradeService(ctx context.Context, req *kf.CancelUpgradeServiceRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/kf/customer/cancel_upgrade_service", req)
	return err
}

