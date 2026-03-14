package security

import (
	"context"

	"github.com/mrseanchow/wecom-core/internal/client"
	"github.com/mrseanchow/wecom-core/types/security"
)

// ImportTrustDevice 导入可信企业设备
// 文档: https://developer.work.weixin.qq.com/document/path/94706
func (s *Service) ImportTrustDevice(ctx context.Context, req *security.ImportTrustDeviceRequest) (*security.ImportTrustDeviceResponse, error) {
	return client.PostAndUnmarshal[security.ImportTrustDeviceResponse](s.client, ctx, "/cgi-bin/security/trustdevice/import", req)
}

// ListTrustDevice 获取设备信息
// 文档: https://developer.work.weixin.qq.com/document/path/94706
func (s *Service) ListTrustDevice(ctx context.Context, req *security.ListTrustDeviceRequest) (*security.ListTrustDeviceResponse, error) {
	return client.PostAndUnmarshal[security.ListTrustDeviceResponse](s.client, ctx, "/cgi-bin/security/trustdevice/list", req)
}

// GetDeviceByUser 获取成员使用设备
// 文档: https://developer.work.weixin.qq.com/document/path/94706
func (s *Service) GetDeviceByUser(ctx context.Context, req *security.GetDeviceByUserRequest) (*security.GetDeviceByUserResponse, error) {
	return client.PostAndUnmarshal[security.GetDeviceByUserResponse](s.client, ctx, "/cgi-bin/security/trustdevice/get_by_user", req)
}

// DeleteTrustDevice 删除设备信息
// 文档: https://developer.work.weixin.qq.com/document/path/94706
func (s *Service) DeleteTrustDevice(ctx context.Context, req *security.DeleteTrustDeviceRequest) (*security.DeleteTrustDeviceResponse, error) {
	return client.PostAndUnmarshal[security.DeleteTrustDeviceResponse](s.client, ctx, "/cgi-bin/security/trustdevice/delete", req)
}

// ApproveTrustDevice 确认为可信设�?
// 文档: https://developer.work.weixin.qq.com/document/path/94706
func (s *Service) ApproveTrustDevice(ctx context.Context, req *security.ApproveTrustDeviceRequest) (*security.ApproveTrustDeviceResponse, error) {
	return client.PostAndUnmarshal[security.ApproveTrustDeviceResponse](s.client, ctx, "/cgi-bin/security/trustdevice/approve", req)
}

// RejectTrustDevice 驳回可信设备申请
// 文档: https://developer.work.weixin.qq.com/document/path/94706
func (s *Service) RejectTrustDevice(ctx context.Context, req *security.RejectTrustDeviceRequest) (*security.RejectTrustDeviceResponse, error) {
	return client.PostAndUnmarshal[security.RejectTrustDeviceResponse](s.client, ctx, "/cgi-bin/security/trustdevice/reject", req)
}

