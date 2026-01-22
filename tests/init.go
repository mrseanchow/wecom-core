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
	expireAt, _ := time.Parse("2006-01-02 15:04:05", "2026-01-19 14:42:52")
	return "RUFbN2gMw878NbFT2utpcrSvqI0QyJBqR5umxXNGnsGYd9RVviBvYoZf95KTqFngWBfYsl6ig-6iKDdyE1-07bOhYfBg2CuomAyzU71AlUoumloNQU4D6AEYEjODQxJI6S7rLzTfyZxhgSq705_ydrrHzJzzgHZqhSP0cS6El9nZ1qKyjNw3dsNVNoGvj_iUj5eCIhhm60udqkXludgniw", expireAt, nil
}

func (n *NoCache) Set(ctx context.Context, key string, token string, expireAt time.Time) error {
	fmt.Println(key, token, expireAt)
	return nil
}

func (n *NoCache) Delete(ctx context.Context, key string) error {
	return nil
}

func initByProxy() {

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
	fmt.Println("使用代理服务访问企微-初始化完成")
}

func initByDirect() {
	corpID := os.Getenv("CORP_ID")
	addressBookSecret := os.Getenv("ADDRESS_BOOK_SECRET")
	bossZSSecret := os.Getenv("BOSS_ZS_SECRET")

	wecom.MustInit(
		config.WithCorpID(corpID),
		config.WithAgent("bosszs", 1000050, bossZSSecret, "boss助手"),
		config.WithAgent("addressBook", 1, addressBookSecret, "通讯录助手"),
		config.WithCache(&NoCache{}),
		config.WithLogger(logger.NewStdLogger()),
		config.WithTimeout(20*time.Second),
	)
	fmt.Println("使用直联访问企微-初始化完成")
}

func init() {
	godotenv.Load()
	proxy := os.Getenv("proxy") == "true"
	if proxy {
		initByProxy()
	} else {
		initByDirect()
	}
}
