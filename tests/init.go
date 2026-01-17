package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/shuaidd/wecom-core"
	"github.com/shuaidd/wecom-core/config"
	"github.com/shuaidd/wecom-core/pkg/logger"
	"github.com/shuaidd/wecom-core/types/common"
)

type NoCache struct {
}

func (n *NoCache) Get(ctx context.Context, key string) (string, time.Time, error) {
	expireAt, _ := time.Parse("2006-01-02 15:04:05", "2027-01-01 00:00:00")
	return "ww", expireAt, nil
}

func (n *NoCache) Set(ctx context.Context, key string, token string, expireAt time.Time) error {
	return nil
}

func (n *NoCache) Delete(ctx context.Context, key string) error {
	return nil
}

func init() {
	godotenv.Load()

	baseURL := os.Getenv("WECOM_BASE_URL")
	accessToken := os.Getenv("WECOM_ACCESS_TOKEN")

	if baseURL == "" || accessToken == "" {
		panic("baseURL or accessToken is empty")
	}

	wecom.MustInit(
		config.WithBaseURL(baseURL),
		config.WithCorpID("sss"),
		config.WithAgent("boss-customer", 1, "xx", "boss客户"),
		config.WithCache(&NoCache{}),
		config.WithToken(false),
		config.WithTimeout(120*time.Second),
		config.WithDecoder(func(data []byte, v any) error {
			var env struct {
				Status string          `json:"status"`
				Msg    string          `json:"msg"`
				Data   json.RawMessage `json:"data"`
			}
			if err := json.Unmarshal(data, &env); err != nil {
				return fmt.Errorf("unmarshal envelope: %w", err)
			}
			if env.Status != "1" {
				fmt.Printf("接口调用失败 errcode: %s errmsg: %s", env.Status, env.Msg)
				return fmt.Errorf("api error: status: %s msg: %s", env.Status, env.Msg)
			}
			if len(env.Data) == 0 {
				return json.Unmarshal(data, v)
			}

			result := json.Unmarshal(env.Data, v)
			if resp, ok := v.(common.Response); ok {
				errCode, _ := strconv.Atoi(env.Status)
				resp.ErrCode = errCode
				resp.ErrMsg = env.Msg
				logger.F("errcode", resp.ErrCode)
				logger.F("errmsg", resp.ErrMsg)
			}
			return result
		}),
		config.WithRequestInterceptor(func(ctx context.Context, req *http.Request, body any) error {
			agentName := wecom.GetAgentName(ctx)
			if agentName == "" {
				agentName = "boss-customer"
			}
			req.Header.Add("Auth-Type", "access-token")
			req.Header.Add("access-token", accessToken)
			req.Header.Add("agent-name", agentName)
			return nil
		}),
	)
	fmt.Println("初始化完成")
}
