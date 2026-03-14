package agent

import (
	"context"
	"fmt"
	"net/url"

	"github.com/mrseanchow/wecom-core/internal/client"
	"github.com/mrseanchow/wecom-core/types/agent"
)

// Get 获取指定的应用详�?
func (s *Service) Get(ctx context.Context, agentID int) (*agent.GetAgentResponse, error) {
	query := url.Values{}
	query.Set("agentid", fmt.Sprintf("%d", agentID))

	return client.GetAndUnmarshal[agent.GetAgentResponse](s.client, ctx, "/cgi-bin/agent/get", query)
}

// List 获取access_token对应的应用列�?
func (s *Service) List(ctx context.Context) (*agent.ListAgentResponse, error) {
	return client.GetAndUnmarshal[agent.ListAgentResponse](s.client, ctx, "/cgi-bin/agent/list", nil)
}

// Set 设置应用
func (s *Service) Set(ctx context.Context, req *agent.SetAgentRequest) error {
	_, err := client.PostAndUnmarshal[agent.SetAgentResponse](s.client, ctx, "/cgi-bin/agent/set", req)
	return err
}

