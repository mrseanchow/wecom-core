package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

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
	wecom.MustInit(
		config.WithBaseURL("https://fzapi.shinyway.com"),
		config.WithCorpID("sss"),
		config.WithAgent("boss-customer", 1, "ss", "boss客户"),
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
				return fmt.Errorf("api error: status: %s msg: %s", env.Status, env.Msg)
			}
			// 如果服务器直接返回业务对象（没有 data 包装），回退到直接解码整个 body
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
			req.Header.Add("Auth-Type", "access-token")
			req.Header.Add("access-token", "eyJhbGciOiJIUzI1NiJ9.eyJvbGRVbml0SWQiOiIyMDIyMDMyMjA2NjY2MiIsInVuaXROYW1lIjoi56CU5Y-R6YOoIiwibG9naW5fdGltZSI6IjIwMjYtMDEtMTYgMjE6MzU6NTYiLCJjb21wYW55TmFtZSI6IuaWsOmAmuWbvemZhSIsImlkIjoiMTIwcm4iLCJ1c2VyTG9naW5LZXkiOiJnYXRld2F5OnVzZXI6bG9naW46djM6MjAxNzA0MTAwMjI3MTc6MTc2ODU3MDU1NjI5OCIsIm9sZElkIjoiMjAxNzA0MTAwMjI3MTciLCJhY2NvdW50IjoiRERTSFVBSSIsImxvZ2luX3RpbWVzdGFtcCI6MTc2ODU3MDU1NjI5OCwib2xkQ29tcGFueUlkIjoieHRnaiIsInVzZXJuYW1lIjoi5biF5Yas5YasIiwidWFNRDUiOiI4OWRiNzI5Y2ZjZGMxMjkxMTFmMDE3YjBlN2FjMzI0YSJ9.83MRSSyVIhXR-l1PbSZ6alleHEMWmtf6W30_kFtwp3w")
			req.Header.Add("agent-name", "boss-customer")
			return nil
		}),
	)
	fmt.Println("初始化完成")
}
