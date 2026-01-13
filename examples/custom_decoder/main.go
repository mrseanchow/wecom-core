package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/shuaidd/wecom-core"
	"github.com/shuaidd/wecom-core/config"
	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/pkg/logger"
	"github.com/shuaidd/wecom-core/types/common"
	"github.com/shuaidd/wecom-core/types/media"
)

type ExampleData struct {
	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}

type NoCache struct {
}

func NewNoCache() *NoCache {
	return &NoCache{}
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

func main1() {

	// 自定义解码器示例：用于解析形如 {"errcode":0,"errmsg":"","data":{...}} 的响应，
	// 将内部的 data 字段解码到目标对象。
	decoder := func(data []byte, v any) error {
		var env struct {
			ErrCode int             `json:"errcode"`
			ErrMsg  string          `json:"errmsg"`
			Data    json.RawMessage `json:"data"`
		}
		if err := json.Unmarshal(data, &env); err != nil {
			return fmt.Errorf("unmarshal envelope: %w", err)
		}
		if env.ErrCode != 0 {
			return fmt.Errorf("api error: %d %s", env.ErrCode, env.ErrMsg)
		}
		// 如果服务器直接返回业务对象（没有 data 包装），回退到直接解码整个 body
		if len(env.Data) == 0 {
			return json.Unmarshal(data, v)
		}
		return json.Unmarshal(env.Data, v)
	}

	// 1) 注入到 Client（注意：示例中未提供真实的 tokenManager/retryExecutor，因此下面的 Do 调用为注释）
	c := client.New("https://api.example.com", 10*time.Second, nil, nil, nil)
	c.SetDecoder(decoder)

	// 如果你已经为 Client 提供了有效的 tokenManager 和 retryExecutor，
	// 则可以直接使用 DoAndUnmarshal（下方为示例代码，运行前请确保 tokenManager/retryExecutor 不为 nil）
	/*
		req := client.NewRequest(client.MethodGet, "/v1/example")
		res, err := client.DoAndUnmarshal[ExampleData](c, context.Background(), req)
		if err != nil {
			fmt.Println("DoAndUnmarshal error:", err)
			return
		}
		fmt.Printf("result from server: %+v\n", res)
	*/

	// 2) 直接使用自定义解码器解析已知响应体（便于单元测试 / 本地调试）
	sampleBody := []byte(`{"errcode":0,"errmsg":"","data":{"foo":"hello","bar":123}}`)
	var out ExampleData
	if err := decoder(sampleBody, &out); err != nil {
		fmt.Println("decoder error:", err)
		return
	}
	fmt.Printf("decoded result: %+v\n", out)
}

func buildCli() (*wecom.Client, error) {
	cli, err := wecom.New(
		config.WithBaseURL("https://fzapi.shinyway.com"),
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
			req.Header.Add("access-token", "eyJhbGciOiJIUzI1NiJ9.eyJvbGRVbml0SWQiOiIyMDIyMDMyMjA2NjY2MiIsInVuaXROYW1lIjoi56CU5Y-R6YOoIiwibG9naW5fdGltZSI6IjIwMjYtMDEtMTIgMjE6MzE6MzIiLCJjb21wYW55TmFtZSI6IuaWsOmAmuWbvemZhSIsImlkIjoiMTIwcm4iLCJvbGRJZCI6IjIwMTcwNDEwMDIyNzE3IiwiYWNjb3VudCI6IkREU0hVQUkiLCJsb2dpbl90aW1lc3RhbXAiOjE3NjgyMjQ2OTIyNTEsIm9sZENvbXBhbnlJZCI6Inh0Z2oiLCJ1c2VybmFtZSI6IuW4heWGrOWGrCIsInVhTUQ1IjoiNjg0ZmFjM2Q4ZTU5NTg0NTY0MGU1MDdhOTEyMmJkNTUifQ.MqMa1pA76_nG5axYfrdJfw6XJkuznAEXKuwoWpdobiI")
			req.Header.Add("agent-name", "boss-customer")
			return nil
		}),
	)
	if err != nil {
		fmt.Println("create wecom client error:", err)
		return nil, err
	}

	return cli, nil
}

func main() {
	getUser()
	upload()
}

func getUser() {
	cli, _ := buildCli()
	ctx := wecom.WithAgentName(context.Background(), "boss-customer")
	user, err := cli.Contact.GetUser(ctx, "20170410022717")
	if err != nil {
		fmt.Println("get user error:", err)
		return
	}
	fmt.Printf("user: %+v\n", user)
}

func upload() {
	cli, _ := buildCli()
	// 读取文件到内存
	data, err := os.ReadFile("1.png")
	if err != nil {
		fmt.Println("read file error:", err)
		return
	}

	// 将 []byte 包装为 io.Reader
	reader := bytes.NewReader(data)

	ctx := wecom.WithAgentName(context.Background(), "boss-customer")
	// 调用上传方法并处理返回值/错误
	if data, err := cli.Media.UploadMediaFromReader(ctx, media.MediaTypeImage, reader, "1.png"); err != nil {
		fmt.Println("upload error:", err)
		return
	} else {
		fmt.Printf("upload success: mediaId %s %+v \n", data.MediaID, data)
	}
}
