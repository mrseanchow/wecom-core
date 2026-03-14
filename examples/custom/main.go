package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mrseanchow/wecom-core"
	"github.com/mrseanchow/wecom-core/config"
	"github.com/mrseanchow/wecom-core/pkg/logger"
)

// 示例：如何调用未封装的企业微�?API
// 完全复用 SDK �?token 管理和重试逻辑

func main() {
	// 创建企业微信客户�?
	client, err := wecom.New(
		config.WithCorpID("your_corp_id"),
		config.WithDebug(true),
		config.WithLogger(logger.NewStdLogger()),
	)
	if err != nil {
		log.Fatalf("Failed to create wecom client: %v", err)
	}

	ctx := context.Background()

	// 示例1: 使用 CustomGet - 手动解析响应
	fmt.Println("=== 示例1: CustomGet - 手动解析响应 ===")
	exampleCustomGet(client, ctx)

	// 示例2: 使用 CustomPost - 手动解析响应
	fmt.Println("\n=== 示例2: CustomPost - 手动解析响应 ===")
	exampleCustomPost(client, ctx)

	// 示例3: 使用 CustomGetAndUnmarshal - 自动解析响应（推荐）
	fmt.Println("\n=== 示例3: CustomGetAndUnmarshal - 自动解析响应 ===")
	exampleCustomGetAndUnmarshal(client, ctx)

	// 示例4: 使用 CustomPostAndUnmarshal - 自动解析响应（推荐）
	fmt.Println("\n=== 示例4: CustomPostAndUnmarshal - 自动解析响应 ===")
	exampleCustomPostAndUnmarshal(client, ctx)
}

// exampleCustomGet 示例：使�?CustomGet 手动解析响应
func exampleCustomGet(client *wecom.Client, ctx context.Context) {
	// 调用未封装的 GET 接口
	resp, err := client.CustomGet(ctx, "/cgi-bin/getcallbackip", nil)
	if err != nil {
		log.Printf("Failed to get callback IP: %v", err)
		return
	}

	// 手动解析响应
	type GetCallbackIPResponse struct {
		wecom.CommonResponse
		IPList []string `json:"ip_list"`
	}

	var result GetCallbackIPResponse
	if err := resp.Unmarshal(&result); err != nil {
		log.Printf("Failed to unmarshal response: %v", err)
		return
	}

	fmt.Printf("Callback IP List: %v\n", result.IPList)
}

// exampleCustomPost 示例：使�?CustomPost 手动解析响应
func exampleCustomPost(client *wecom.Client, ctx context.Context) {
	// 定义请求�?
	type SendMessageRequest struct {
		ToUser  string `json:"touser"`
		MsgType string `json:"msgtype"`
		AgentID int    `json:"agentid"`
		Text    struct {
			Content string `json:"content"`
		} `json:"text"`
	}

	req := SendMessageRequest{
		ToUser:  "UserID1",
		MsgType: "text",
		AgentID: 1000001,
	}
	req.Text.Content = "这是一条测试消�?

	// 调用未封装的 POST 接口
	resp, err := client.CustomPost(ctx, "/cgi-bin/message/send", req)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
		return
	}

	// 手动解析响应
	type SendMessageResponse struct {
		wecom.CommonResponse
		InvalidUser  string `json:"invaliduser"`
		InvalidParty string `json:"invalidparty"`
		InvalidTag   string `json:"invalidtag"`
		MsgID        string `json:"msgid"`
	}

	var result SendMessageResponse
	if err := resp.Unmarshal(&result); err != nil {
		log.Printf("Failed to unmarshal response: %v", err)
		return
	}

	fmt.Printf("Message sent successfully, MsgID: %s\n", result.MsgID)
}

// exampleCustomGetAndUnmarshal 示例：使�?CustomGetAndUnmarshal 自动解析响应（推荐）
func exampleCustomGetAndUnmarshal(client *wecom.Client, ctx context.Context) {
	// 定义响应类型
	type GetCallbackIPResponse struct {
		wecom.CommonResponse
		IPList []string `json:"ip_list"`
	}

	// 调用接口并自动解析响应（推荐使用这种方式�?
	result, err := wecom.CustomGetAndUnmarshal[GetCallbackIPResponse](
		client,
		ctx,
		"/cgi-bin/getcallbackip",
		nil, // 无查询参�?
	)
	if err != nil {
		log.Printf("Failed to get callback IP: %v", err)
		return
	}

	fmt.Printf("Callback IP List: %v\n", result.IPList)
}

// exampleCustomPostAndUnmarshal 示例：使�?CustomPostAndUnmarshal 自动解析响应（推荐）
func exampleCustomPostAndUnmarshal(client *wecom.Client, ctx context.Context) {
	// 定义请求类型
	type SendMessageRequest struct {
		ToUser  string `json:"touser"`
		MsgType string `json:"msgtype"`
		AgentID int    `json:"agentid"`
		Text    struct {
			Content string `json:"content"`
		} `json:"text"`
	}

	// 定义响应类型
	type SendMessageResponse struct {
		wecom.CommonResponse
		InvalidUser  string `json:"invaliduser"`
		InvalidParty string `json:"invalidparty"`
		InvalidTag   string `json:"invalidtag"`
		MsgID        string `json:"msgid"`
	}

	// 准备请求
	req := SendMessageRequest{
		ToUser:  "UserID1",
		MsgType: "text",
		AgentID: 1000001,
	}
	req.Text.Content = "这是一条测试消�?

	// 调用接口并自动解析响应（推荐使用这种方式�?
	result, err := wecom.CustomPostAndUnmarshal[SendMessageResponse](
		client,
		ctx,
		"/cgi-bin/message/send",
		req,
		nil,
	)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
		return
	}

	fmt.Printf("Message sent successfully, MsgID: %s\n", result.MsgID)
}

