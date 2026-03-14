package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mrseanchow/wecom-core"
	"github.com/mrseanchow/wecom-core/config"
	ext "github.com/mrseanchow/wecom-core/types/externalcontact"
)

func main() {
	ctx := context.Background()

	// 替换为真实的 CorpID �?CorpSecret
	client, err := wecom.New(config.WithCorpID("YOUR_CORP_ID"), config.WithDebug(true))
	if err != nil {
		log.Fatalf("create client error: %v", err)
	}

	svc := client.ExternalContact

	// 1. 获取待分配的离职成员列表
	guReq := &ext.GetUnassignedListRequest{PageSize: 100}
	guResp, err := svc.GetUnassignedList(ctx, guReq)
	if err != nil {
		log.Fatalf("GetUnassignedList error: %v", err)
	}
	fmt.Printf("GetUnassignedList response: %+v\n", guResp)

	// 2. 分配离职成员的客�?
	tcReq := &ext.TransferCustomerRequest{
		HandoverUserID:  "old_userid",
		TakeoverUserID:  "new_userid",
		ExternalUserIDs: []string{"woAJ2GCAAAXtWyujaWJHDDGi0mACBBBB"},
	}
	tcResp, err := svc.TransferCustomer(ctx, tcReq)
	if err != nil {
		log.Fatalf("TransferCustomer error: %v", err)
	}
	fmt.Printf("TransferCustomer response: %+v\n", tcResp)

	// 3. 分配离职成员的客户群
	tgReq := &ext.TransferGroupChatRequest{
		ChatIDList: []string{"wrOgQhDgAAcwMTB7YmDkbeBsgT_AAAA"},
		NewOwner:   "new_userid",
	}
	tgResp, err := svc.TransferGroupChat(ctx, tgReq)
	if err != nil {
		log.Fatalf("TransferGroupChat error: %v", err)
	}
	fmt.Printf("TransferGroupChat response: %+v\n", tgResp)
}

