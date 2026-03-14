package updown

import (
	"github.com/mrseanchow/wecom-core/internal/client"
)

// Service 上下游服�?
type Service struct {
	client *client.Client
}

// NewService 创建上下游服�?
func NewService(c *client.Client) *Service {
	return &Service{
		client: c,
	}
}

