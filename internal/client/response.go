package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mrseanchow/wecom-core/internal/errors"
	"github.com/mrseanchow/wecom-core/types/common"
)

// Response 表示一个HTTP响应
type Response struct {
	// HTTPResponse 原始HTTP响应
	HTTPResponse *http.Response
	// CommonResponse 通用响应字段
	common.Response
	// Body 响应体（原始字节�?
	Body []byte
}

// ParseResponse 解析HTTP响应
func ParseResponse(httpResp *http.Response) (*Response, error) {
	resp := &Response{
		HTTPResponse: httpResp,
	}

	// 读取响应�?
	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	resp.Body = body

	// 解析通用响应字段
	if err := json.Unmarshal(body, &resp.Response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// 检查错误码
	if resp.ErrCode != 0 {
		return resp, errors.New(resp.ErrCode, resp.ErrMsg)
	}

	return resp, nil
}

// Unmarshal 将响应体解析到目标对�?
func (r *Response) Unmarshal(v any) error {
	if err := json.Unmarshal(r.Body, v); err != nil {
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}
	return nil
}

// IsSuccess 判断响应是否成功
func (r *Response) IsSuccess() bool {
	return r.ErrCode == 0
}

