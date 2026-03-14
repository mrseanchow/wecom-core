package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mrseanchow/wecom-core"
	"github.com/mrseanchow/wecom-core/config"
	"github.com/mrseanchow/wecom-core/pkg/logger"
)

func main() {
	// 创建企业微信客户�?
	client, err := wecom.New(
		config.WithCorpID("wwd8b4a0d94c30bcb7"),
		// config.WithCorpSecret("c024lT7BImnNiBawyKheY5NhbtuzH0vxS2y3e_fAR-s"),
		config.WithDebug(true),
		config.WithLogger(logger.NewStdLogger()),
	)
	if err != nil {
		log.Fatalf("Failed to create wecom client: %v", err)
	}

	ctx := wecom.WithTraceID(context.Background(), "sss")
	// 读取成员信息
	user, err := client.Contact.GetUser(ctx, "ddshuai")
	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}

	fmt.Printf("成员信息: UserID=%s, Name=%s, Mobile=%s\n",
		user.UserID, user.Name, user.Mobile)
}

