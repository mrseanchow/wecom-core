package email

import (
	"context"

	"github.com/mrseanchow/wecom-core/internal/client"
	"github.com/mrseanchow/wecom-core/types/email"
)

const (
	getAuthCodeListURL = "/cgi-bin/exmail/publicmail/get_auth_code_list"
	deleteAuthCodeURL  = "/cgi-bin/exmail/publicmail/delete_auth_code"
)

// GetAuthCodeList 获取客户端专用密码列�?
func (s *Service) GetAuthCodeList(ctx context.Context, publicMailID uint32) (*email.GetAuthCodeListResponse, error) {
	req := &email.GetAuthCodeListRequest{
		ID: publicMailID,
	}
	return client.PostAndUnmarshal[email.GetAuthCodeListResponse](s.client, ctx, getAuthCodeListURL, req)
}

// DeleteAuthCode 删除客户端专用密�?
func (s *Service) DeleteAuthCode(ctx context.Context, publicMailID, authCodeID uint32) (*email.DeleteAuthCodeResponse, error) {
	req := &email.DeleteAuthCodeRequest{
		ID:         publicMailID,
		AuthCodeID: authCodeID,
	}
	return client.PostAndUnmarshal[email.DeleteAuthCodeResponse](s.client, ctx, deleteAuthCodeURL, req)
}

