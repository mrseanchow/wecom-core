package email

import (
	"context"
	"fmt"
	"net/url"

	"github.com/mrseanchow/wecom-core/internal/client"
	"github.com/mrseanchow/wecom-core/types/email"
)

const (
	createPublicMailURL = "/cgi-bin/exmail/publicmail/create"
	updatePublicMailURL = "/cgi-bin/exmail/publicmail/update"
	deletePublicMailURL = "/cgi-bin/exmail/publicmail/delete"
	getPublicMailURL    = "/cgi-bin/exmail/publicmail/get"
	searchPublicMailURL = "/cgi-bin/exmail/publicmail/search"
)

// CreatePublicMail 创建公共邮箱
func (s *Service) CreatePublicMail(ctx context.Context, req *email.CreatePublicMailRequest) (*email.CreatePublicMailResponse, error) {
	return client.PostAndUnmarshal[email.CreatePublicMailResponse](s.client, ctx, createPublicMailURL, req)
}

// UpdatePublicMail 更新公共邮箱
func (s *Service) UpdatePublicMail(ctx context.Context, req *email.UpdatePublicMailRequest) (*email.UpdatePublicMailResponse, error) {
	return client.PostAndUnmarshal[email.UpdatePublicMailResponse](s.client, ctx, updatePublicMailURL, req)
}

// DeletePublicMail 删除公共邮箱
func (s *Service) DeletePublicMail(ctx context.Context, id uint32) (*email.DeletePublicMailResponse, error) {
	req := &email.DeletePublicMailRequest{
		ID: id,
	}
	return client.PostAndUnmarshal[email.DeletePublicMailResponse](s.client, ctx, deletePublicMailURL, req)
}

// GetPublicMail 获取公共邮箱详情
func (s *Service) GetPublicMail(ctx context.Context, idList []uint32) (*email.GetPublicMailResponse, error) {
	req := &email.GetPublicMailRequest{
		IDList: idList,
	}
	return client.PostAndUnmarshal[email.GetPublicMailResponse](s.client, ctx, getPublicMailURL, req)
}

// SearchPublicMail 模糊搜索公共邮箱
// fuzzy: 1-开启模糊搜�?0-获取全部公共邮箱
// emailKeyword: 公共邮箱名称或邮箱地址(fuzzy=1时有�?
func (s *Service) SearchPublicMail(ctx context.Context, fuzzy uint32, emailKeyword string) (*email.SearchPublicMailResponse, error) {
	query := url.Values{}
	query.Set("fuzzy", fmt.Sprintf("%d", fuzzy))
	if emailKeyword != "" {
		query.Set("email", emailKeyword)
	}

	return client.GetAndUnmarshal[email.SearchPublicMailResponse](s.client, ctx, searchPublicMailURL, query)
}

// ListAllPublicMail 获取全部公共邮箱
func (s *Service) ListAllPublicMail(ctx context.Context) (*email.SearchPublicMailResponse, error) {
	return s.SearchPublicMail(ctx, 0, "")
}

