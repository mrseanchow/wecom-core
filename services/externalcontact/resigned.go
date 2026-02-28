package externalcontact

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/externalcontact"
)
const (
	XternalcontactGetUnassignedListURL = "/cgi-bin/externalcontact/get_unassigned_list"
	XternalcontactResignedTransferCustomerURL = "/cgi-bin/externalcontact/resigned/transfer_customer"
)

// GetUnassignedList 获取待分配的离职成员列表
// 文档: https://developer.work.weixin.qq.com/document/path/93963
func (s *Service) GetUnassignedList(ctx context.Context, req *externalcontact.GetUnassignedListRequest) (*externalcontact.GetUnassignedListResponse, error) {
	return client.PostAndUnmarshal[externalcontact.GetUnassignedListResponse](s.client, ctx, XternalcontactGetUnassignedListURL, req)
}

// TransferCustomer 分配离职成员的客户
// 文档: https://developer.work.weixin.qq.com/document/path/93964
func (s *Service) TransferCustomer(ctx context.Context, req *externalcontact.TransferCustomerRequest) (*externalcontact.TransferCustomerResponse, error) {
	return client.PostAndUnmarshal[externalcontact.TransferCustomerResponse](s.client, ctx, XternalcontactResignedTransferCustomerURL, req)
}
