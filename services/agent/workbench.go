package agent

import (
	"context"

	"github.com/mrseanchow/wecom-core/internal/client"
	"github.com/mrseanchow/wecom-core/types/agent"
)

// SetWorkbenchTemplate 设置应用在工作台展示的模�?
func (s *Service) SetWorkbenchTemplate(ctx context.Context, req *agent.SetWorkbenchTemplateRequest) error {
	_, err := client.PostAndUnmarshal[agent.SetWorkbenchTemplateResponse](s.client, ctx, "/cgi-bin/agent/set_workbench_template", req)
	return err
}

// GetWorkbenchTemplate 获取应用在工作台展示的模�?
func (s *Service) GetWorkbenchTemplate(ctx context.Context, agentID int) (*agent.GetWorkbenchTemplateResponse, error) {
	return client.PostAndUnmarshal[agent.GetWorkbenchTemplateResponse](s.client, ctx, "/cgi-bin/agent/get_workbench_template", agent.GetWorkbenchTemplateRequest{
		AgentID: agentID,
	})
}

// SetWorkbenchData 设置应用在用户工作台展示的数�?
func (s *Service) SetWorkbenchData(ctx context.Context, req *agent.SetWorkbenchDataRequest) error {
	_, err := client.PostAndUnmarshal[agent.SetWorkbenchDataResponse](s.client, ctx, "/cgi-bin/agent/set_workbench_data", req)
	return err
}

// BatchSetWorkbenchData 批量设置应用在用户工作台展示的数�?
func (s *Service) BatchSetWorkbenchData(ctx context.Context, req *agent.BatchSetWorkbenchDataRequest) error {
	_, err := client.PostAndUnmarshal[agent.BatchSetWorkbenchDataResponse](s.client, ctx, "/cgi-bin/agent/batch_set_workbench_data", req)
	return err
}

// GetWorkbenchData 获取应用在用户工作台展示的数�?
func (s *Service) GetWorkbenchData(ctx context.Context, agentID int, userID string) (*agent.GetWorkbenchDataResponse, error) {
	return client.PostAndUnmarshal[agent.GetWorkbenchDataResponse](s.client, ctx, "/cgi-bin/agent/get_workbench_data", agent.GetWorkbenchDataRequest{
		AgentID: agentID,
		UserID:  userID,
	})
}

