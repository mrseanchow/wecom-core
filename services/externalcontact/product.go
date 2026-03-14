package externalcontact

import (
	"context"

	"github.com/mrseanchow/wecom-core/internal/client"
	"github.com/mrseanchow/wecom-core/types/common"
	"github.com/mrseanchow/wecom-core/types/externalcontact"
)

// AddProductAlbum 创建商品图册
// 企业和第三方应用可以通过此接口增加商�?
// 文档: https://developer.work.weixin.qq.com/document/path/95096
func (s *Service) AddProductAlbum(ctx context.Context, req *externalcontact.AddProductAlbumRequest) (*externalcontact.AddProductAlbumResponse, error) {
	return client.PostAndUnmarshal[externalcontact.AddProductAlbumResponse](s.client, ctx, "/cgi-bin/externalcontact/add_product_album", req)
}

// GetProductAlbum 获取商品图册
// 企业和第三方应用可以通过此接口获取商品信�?
// 文档: https://developer.work.weixin.qq.com/document/path/95096
func (s *Service) GetProductAlbum(ctx context.Context, req *externalcontact.GetProductAlbumRequest) (*externalcontact.GetProductAlbumResponse, error) {
	return client.PostAndUnmarshal[externalcontact.GetProductAlbumResponse](s.client, ctx, "/cgi-bin/externalcontact/get_product_album", req)
}

// GetProductAlbumList 获取商品图册列表
// 企业和第三方应用可以通过此接口导出商�?
// 文档: https://developer.work.weixin.qq.com/document/path/95096
func (s *Service) GetProductAlbumList(ctx context.Context, req *externalcontact.GetProductAlbumListRequest) (*externalcontact.GetProductAlbumListResponse, error) {
	return client.PostAndUnmarshal[externalcontact.GetProductAlbumListResponse](s.client, ctx, "/cgi-bin/externalcontact/get_product_album_list", req)
}

// UpdateProductAlbum 编辑商品图册
// 企业和第三方应用可以通过此接口修改商品信�?
// 文档: https://developer.work.weixin.qq.com/document/path/95096
func (s *Service) UpdateProductAlbum(ctx context.Context, req *externalcontact.UpdateProductAlbumRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/externalcontact/update_product_album", req)
	return err
}

// DeleteProductAlbum 删除商品图册
// 企业和第三方应用可以通过此接口删除商品信�?
// 文档: https://developer.work.weixin.qq.com/document/path/95096
func (s *Service) DeleteProductAlbum(ctx context.Context, req *externalcontact.DeleteProductAlbumRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/externalcontact/delete_product_album", req)
	return err
}

