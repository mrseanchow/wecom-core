package wedoc

import (
	"context"

	"github.com/mrseanchow/wecom-core/internal/client"
	"github.com/mrseanchow/wecom-core/types/common"
	"github.com/mrseanchow/wecom-core/types/wedoc"
)

const (
	createDocURL      = "/cgi-bin/wedoc/create_doc"
	getDocBaseInfoURL = "/cgi-bin/wedoc/get_doc_base_info"
	deleteDocURL      = "/cgi-bin/wedoc/del_doc"
	renameDocURL      = "/cgi-bin/wedoc/rename_doc"
	shareDocURL       = "/cgi-bin/wedoc/doc_share"
)

// CreateDoc 新建文档
// 该接口用于新建文档、表格及智能表格
func (s *Service) CreateDoc(ctx context.Context, req *wedoc.CreateDocRequest) (*wedoc.CreateDocResponse, error) {
	return client.PostAndUnmarshal[wedoc.CreateDocResponse](s.client, ctx, createDocURL, req)
}

// GetDocBaseInfo 获取文档基础信息
// 该接口用于获取指定文档、表格、智能表格及收集表的基础信息
func (s *Service) GetDocBaseInfo(ctx context.Context, req *wedoc.GetDocBaseInfoRequest) (*wedoc.GetDocBaseInfoResponse, error) {
	return client.PostAndUnmarshal[wedoc.GetDocBaseInfoResponse](s.client, ctx, getDocBaseInfoURL, req)
}

// DeleteDoc 删除文档
// 该接口用于删除指定文档、表格、智能表格及收集�?
func (s *Service) DeleteDoc(ctx context.Context, req *wedoc.DeleteDocRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, deleteDocURL, req)
	return err
}

// RenameDoc 重命名文�?
// 该接口用于对指定文档、表格、智能表格及收集表进行重命名
func (s *Service) RenameDoc(ctx context.Context, req *wedoc.RenameDocRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, renameDocURL, req)
	return err
}

// ShareDoc 分享文档
// 该接口用于获取文档、表格、智能表格及收集表的分享链接
func (s *Service) ShareDoc(ctx context.Context, req *wedoc.ShareDocRequest) (*wedoc.ShareDocResponse, error) {
	return client.PostAndUnmarshal[wedoc.ShareDocResponse](s.client, ctx, shareDocURL, req)
}

