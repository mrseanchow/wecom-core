package miniapppay

import (
	"github.com/shuaidd/wecom-core/internal/client"
)

type Service struct {
	client *client.Client
}

func NewService(c *client.Client) *Service {
	return &Service{
		client: c,
	}
}
