package tests

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/shuaidd/wecom-core"
	"github.com/shuaidd/wecom-core/types/media"
)

func TestUploadMediaFromReader(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}

	// 读取文件到内存
	data, err := os.ReadFile("二维码.png")
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

func TestUploadMediaFromFile(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}

	// 读取文件到内存
	data, err := os.Open("圆桌会PPT内容.md")
	if err != nil {
		fmt.Println("read file error:", err)
		return
	}

	defer data.Close()

	ctx := wecom.WithAgentName(context.Background(), "boss-customer")
	// 调用上传方法并处理返回值/错误
	if data, err := cli.Media.UploadMediaFromReader(ctx, media.MediaTypeFile, data, "圆桌会PPT内容.md"); err != nil {
		fmt.Println("upload error:", err)
		return
	} else {
		fmt.Printf("upload success: mediaId %s %+v \n", data.MediaID, data)
	}
}
