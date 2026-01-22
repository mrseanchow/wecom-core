package tests

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
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
	data, err := os.ReadFile("/var/folders/t3/y6vxtfbn76b5n3t0hzdc6myw0000gn/T/material_3332207419.png")
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

func TestUploadMedia(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}

	// 调用上传方法并处理返回值/错误
	if data, err := cli.Media.UploadMedia(context.Background(), media.MediaTypeImage, "/var/folders/t3/y6vxtfbn76b5n3t0hzdc6myw0000gn/T/material_3332207419.png"); err != nil {
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

// 新增：从远程下载文件并上传到素材库
func TestUploadMediaFromRemote(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}

	// 远程下载地址
	url := "https://fzapi.shinyway.com/boss-base-svc/storage/download/20260117100001"
	// 使用带鉴权的 HTTP 请求，添加 access-token header
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("request error:", err)
		return
	}
	// 从环境变量读取 access-token（请根据实际情况替换）
	req.Header.Set("access-token", os.Getenv("WECOM_ACCESS_TOKEN"))
	req.Header.Add("Auth-Type", "access-token")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("download error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("download status:", resp.Status)
		return
	}

	// 尝试从 Content-Disposition 或 URL 中推断文件名
	filename := "5.png"

	ctx := wecom.WithAgentName(context.Background(), "boss-customer")
	// 直接以 resp.Body 作为 io.Reader 上传
	if data, err := cli.Media.UploadMediaFromReader(ctx, media.MediaTypeImage, resp.Body, filename); err != nil {
		fmt.Println("upload error:", err)
		return
	} else {
		fmt.Printf("upload success: mediaId %s %+v \n", data.MediaID, data)
	}
}
