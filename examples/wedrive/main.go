package main

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/mrseanchow/wecom-core"
	"github.com/mrseanchow/wecom-core/types/wedrive"
)

// examples/wedrive/main.go: 分块上传示例
//
// 说明�?
// - 请在运行前配置好 wecom 客户端（通过 config.Option 或环境变量，按项�?README 的方式）�?
// - 本示例把文件�?2MB 分块并依次上传，最后调用上传完成接口�?
// - 为简单起见，分块累积 sha 使用对前 N 个字节做 sha1(sum(data[:end])) 的方式来生成（与官网描述一致的“累�?sha”在行为上等价，用于示例）�?
//
// 使用�?
//
//	go run ./examples/wedrive/main.go /path/to/file
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run ./examples/wedrive/main.go /path/to/file")
		return
	}
	path := os.Args[1]

	// 创建客户端（根据需要传�?config.Option�?
	client, err := wecom.New()
	if err != nil {
		fmt.Printf("wecom.New error: %v\n", err)
		return
	}

	ctx := context.Background()

	// 读取文件（示例用途，请在生产中使用流式读取以节省内存�?
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("read file error: %v\n", err)
		return
	}

	const blockSize = 2 * 1024 * 1024 // 2MB
	var blockSHAs []string
	total := len(data)
	for offset := 0; offset < total; offset += blockSize {
		end := offset + blockSize
		if end > total {
			end = total
		}
		// 计算累积 sha：sha1(data[:end])
		h := sha1.Sum(data[:end])
		blockSHAs = append(blockSHAs, fmt.Sprintf("%x", h))
	}

	// 1) 初始化分块上�?
	initReq := &wedrive.FileUploadInitRequest{
		FileName:       "example-" + fmt.Sprintf("%d", os.Getpid()),
		Size:           uint64(len(data)),
		BlockSHA:       blockSHAs,
		SkipPushCard:   true,
		SelectedTicket: "",
	}
	initResp, err := client.Wedrive.UploadInit(ctx, initReq)
	if err != nil {
		fmt.Printf("UploadInit error: %v\n", err)
		return
	}
	if initResp.HitExist {
		fmt.Printf("hit exist, fileid=%s\n", initResp.FileID)
		return
	}
	uploadKey := initResp.UploadKey
	fmt.Printf("upload_key=%s\n", uploadKey)

	// 2) 依次上传每个分块
	partIndex := int32(1)
	for offset := 0; offset < total; offset += blockSize {
		end := offset + blockSize
		if end > total {
			end = total
		}
		part := data[offset:end]
		partB64 := base64.StdEncoding.EncodeToString(part)

		partReq := &wedrive.FileUploadPartRequest{
			UploadKey:         uploadKey,
			Index:             partIndex,
			FileBase64Content: partB64,
		}
		_, err := client.Wedrive.UploadPart(ctx, partReq)
		if err != nil {
			fmt.Printf("UploadPart index=%d error: %v\n", partIndex, err)
			return
		}
		fmt.Printf("uploaded part %d (%d bytes)\n", partIndex, len(part))
		partIndex++
	}

	// 3) 上传完成通知
	finishReq := &wedrive.FileUploadFinishRequest{
		UploadKey: uploadKey,
	}
	finishResp, err := client.Wedrive.UploadFinish(ctx, finishReq)
	if err != nil {
		fmt.Printf("UploadFinish error: %v\n", err)
		return
	}
	fmt.Printf("upload finished, fileid=%s\n", finishResp.FileID)
}

