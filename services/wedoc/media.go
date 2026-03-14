package wedoc

import (
	"context"

	"github.com/mrseanchow/wecom-core/internal/client"
	"github.com/mrseanchow/wecom-core/types/wedoc"
)

const (
	imageUploadURL = "/cgi-bin/wedoc/image_upload"
)

// ImageUpload 上传文档图片
// 该接口用于上传图�?
func (s *Service) ImageUpload(ctx context.Context, req *wedoc.ImageUploadRequest) (*wedoc.ImageUploadResponse, error) {
	return client.PostAndUnmarshal[wedoc.ImageUploadResponse](s.client, ctx, imageUploadURL, req)
}

