package tests

import (
	"context"
	"testing"

	"github.com/shuaidd/wecom-core"
	"github.com/shuaidd/wecom-core/types/agent"
)

func TestAgentCreateMenu(t *testing.T) {

	t.Run("default config", func(t *testing.T) {
		cli, err := wecom.Default()
		if err != nil {
			t.Errorf("获取cli失败 %+v", err)
			return
		}
		agentID := 1000050
		ctx := wecom.WithAgentName(context.Background(), "bosszs")
		err = cli.Agent.CreateMenu(ctx, &agent.CreateMenuRequest{
			AgentID: agentID,
		})
		if err != nil {
			t.Errorf("创建菜单失败 %+v", err)
		}
	})
}
