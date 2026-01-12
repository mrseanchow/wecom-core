package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/shuaidd/wecom-core/internal/client"
)

type ExampleData struct {
	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}

func main() {
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
