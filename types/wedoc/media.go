package wedoc

import "github.com/mrseanchow/wecom-core/types/common"

// ==================== 上传文档图片 ====================

// ImageUploadRequest 上传文档图片请求
type ImageUploadRequest struct {
	DocID         string `json:"docid"`          // 文档ID，通过新建文档接口创建后获�?
	Base64Content string `json:"base64_content"` // base64之后的图片内�?
}

// ImageUploadResponse 上传文档图片响应
type ImageUploadResponse struct {
	common.Response
	URL    string `json:"url"`    // 图片的url
	Height int64  `json:"height"` // 图片的高
	Width  int64  `json:"width"`  // 图片的宽
	Size   int64  `json:"size"`   // 图片的大�?
}

