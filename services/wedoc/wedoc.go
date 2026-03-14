package wedoc

import (
	"github.com/mrseanchow/wecom-core/internal/client"
)

// Service 微文档服�?
type Service struct {
	client *client.Client
}

// New 创建微文档服务实�?
func New(c *client.Client) *Service {
	return &Service{client: c}
}

