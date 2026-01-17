package agent

import (
	"context"
	"fmt"
	"net/url"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/agent"
)

// CreateMenu 创建菜单
func (s *Service) CreateMenu(ctx context.Context, req *agent.CreateMenuRequest) error {
	// 构建请求，需要在query参数中传递agentid，在body中传递菜单内容
	_, err := client.PostAndUnmarshal[agent.CreateMenuResponse](s.client, ctx, fmt.Sprintf("/cgi-bin/menu/create?agentid=%d", req.AgentID), req)
	return err
}

// GetMenu 获取菜单
func (s *Service) GetMenu(ctx context.Context, agentID int) (*agent.GetMenuResponse, error) {
	query := url.Values{}
	query.Set("agentid", fmt.Sprintf("%d", agentID))

	return client.GetAndUnmarshal[agent.GetMenuResponse](s.client, ctx, "/cgi-bin/menu/get", query)
}

// DeleteMenu 删除菜单
func (s *Service) DeleteMenu(ctx context.Context, agentID int) error {
	query := url.Values{}
	query.Set("agentid", fmt.Sprintf("%d", agentID))

	_, err := client.GetAndUnmarshal[agent.DeleteMenuResponse](s.client, ctx, "/cgi-bin/menu/delete", query)
	return err
}
