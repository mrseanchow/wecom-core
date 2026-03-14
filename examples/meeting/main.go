package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mrseanchow/wecom-core"
	"github.com/mrseanchow/wecom-core/config"
	"github.com/mrseanchow/wecom-core/pkg/logger"
	"github.com/mrseanchow/wecom-core/types/meeting"
)

func main() {
	client, err := wecom.New(
		config.WithCorpID("your_corp_id"),
		config.WithLogger(logger.NewStdLogger()),
	)
	if err != nil {
		log.Fatalf("Failed to create wecom client: %v", err)
	}

	ctx := context.Background()

	// 创建预约会议
	createResp, err := client.Meeting.Create(ctx, &meeting.CreateMeetingRequest{
		AdminUserID:     "zhangsan",
		Title:           "产品评审会议",
		MeetingStart:    time.Now().Unix() + 3600,
		MeetingDuration: 3600,
		Description:     "2.0版本产品评审",
		Location:        "10�?005会议�?,
		Invitees: &meeting.Invitees{
			UserID: []string{"lisi", "wangwu"},
		},
		Settings: &meeting.Settings{
			RemindScope:           3,
			Password:              "123456",
			EnableWaitingRoom:     true,
			AllowEnterBeforeHost:  true,
			EnableEnterMute:       1,
			EnableScreenWatermark: false,
			Hosts: &meeting.Hosts{
				UserID: []string{"zhangsan"},
			},
		},
	})
	if err != nil {
		log.Fatalf("创建会议失败: %v", err)
	}
	fmt.Printf("会议创建成功: MeetingID=%s\n", createResp.MeetingID)

	// 获取会议详情
	info, err := client.Meeting.GetInfo(ctx, createResp.MeetingID)
	if err != nil {
		log.Fatalf("获取会议详情失败: %v", err)
	}
	fmt.Printf("会议标题: %s, 状�? %d\n", info.Title, info.Status)

	// 修改会议
	err = client.Meeting.Update(ctx, &meeting.UpdateMeetingRequest{
		MeetingID: createResp.MeetingID,
		Title:     "更新后的会议标题",
		Location:  "11楼会议室",
		Invitees: &meeting.Invitees{
			UserID: []string{"lisi", "wangwu", "zhaoliu"},
		},
	})
	if err != nil {
		log.Fatalf("修改会议失败: %v", err)
	}
	fmt.Println("会议已修�?)

	// 获取成员会议列表
	listResp, err := client.Meeting.GetUserMeetingIDs(ctx, &meeting.GetUserMeetingIDsRequest{
		UserID:    "zhangsan",
		Cursor:    "0",
		Limit:     100,
		BeginTime: time.Now().Add(-7 * 24 * time.Hour).Unix(),
		EndTime:   time.Now().Add(7 * 24 * time.Hour).Unix(),
	})
	if err != nil {
		log.Fatalf("获取会议列表失败: %v", err)
	}
	fmt.Printf("成员会议ID列表: %v\n", listResp.MeetingIDList)

	// 取消会议
	err = client.Meeting.Cancel(ctx, createResp.MeetingID)
	if err != nil {
		log.Fatalf("取消会议失败: %v", err)
	}
	fmt.Println("会议已取�?)
}

