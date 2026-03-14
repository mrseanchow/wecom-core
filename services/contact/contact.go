package contact

import (
	"github.com/mrseanchow/wecom-core/internal/client"
)

// Service 通讯录服�?
type Service struct {
	client *client.Client
}

// NewService 创建通讯录服�?
func NewService(c *client.Client) *Service {
	return &Service{
		client: c,
	}
}

