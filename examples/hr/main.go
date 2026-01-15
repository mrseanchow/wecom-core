package main

import (
	"context"
	"fmt"
	"log"

	"github.com/shuaidd/wecom-core"
	"github.com/shuaidd/wecom-core/config"
	"github.com/shuaidd/wecom-core/pkg/logger"
	"github.com/shuaidd/wecom-core/types/hr"
)

func main() {
	client, err := wecom.New(
		config.WithCorpID("your_corp_id"),
		config.WithAgent("hr", 123456, "your_agent_secret", "HR应用"),
		config.WithLogger(logger.NewStdLogger()),
		config.WithRetry(3),
	)
	if err != nil {
		log.Fatalf("Failed to create wecom client: %v", err)
	}

	ctx := context.Background()

	fmt.Println("=== 获取员工字段配置 ===")
	fieldsResp, err := client.HR.GetFields(ctx)
	if err != nil {
		log.Printf("Failed to get fields: %v", err)
	} else {
		fmt.Printf("字段组数量: %d\n", len(fieldsResp.GroupList))
		for _, group := range fieldsResp.GroupList {
			fmt.Printf("  字段组: ID=%d, Name=%s\n", group.GroupID, group.GroupName)
			for _, field := range group.FieldList {
				fmt.Printf("    字段: ID=%d, Name=%s, Type=%d, ValueType=%d, Must=%v\n",
					field.FieldID, field.FieldName, field.FieldType, field.ValueType, field.IsMust)
			}
		}
	}

	fmt.Println("\n=== 获取员工花名册信息 ===")
	getStaffResp, err := client.HR.GetStaffInfo(ctx, &hr.GetStaffInfoRequest{
		UserID: "zhangsan",
		GetAll: false,
		FieldIDs: []hr.FieldIDWithIdx{
			{FieldID: 11004, SubIdx: 0},
		},
	})
	if err != nil {
		log.Printf("Failed to get staff info: %v", err)
	} else {
		fmt.Printf("字段信息数量: %d\n", len(getStaffResp.FieldInfo))
		for _, field := range getStaffResp.FieldInfo {
			fmt.Printf("  字段: ID=%d, SubIdx=%d, Result=%d, ValueType=%d\n",
				field.FieldID, field.SubIdx, field.Result, field.ValueType)
		}
	}

	fmt.Println("\n=== 更新员工花名册信息 ===")
	name := "张三"
	updateResp, err := client.HR.UpdateStaffInfo(ctx, &hr.UpdateStaffInfoRequest{
		UserID: "zhangsan",
		UpdateItems: []hr.UpdateItem{
			{
				FieldID:     11001,
				SubIdx:      0,
				ValueString: &name,
			},
		},
	})
	if err != nil {
		log.Printf("Failed to update staff info: %v", err)
	} else {
		fmt.Printf("更新结果数量: %d\n", len(updateResp.UpdateResults))
		for _, result := range updateResp.UpdateResults {
			fmt.Printf("  字段: ID=%d, SubIdx=%d, Result=%d\n",
				result.FieldID, result.SubIdx, result.Result)
		}
	}

	fmt.Println("\n=== 删除员工字段组 ===")
	removeResp, err := client.HR.UpdateStaffInfo(ctx, &hr.UpdateStaffInfoRequest{
		UserID: "zhangsan",
		RemoveItems: []hr.RemoveItem{
			{
				GroupType: hr.GroupTypeEducation,
				SubIdx:    1,
			},
		},
	})
	if err != nil {
		log.Printf("Failed to remove staff field group: %v", err)
	} else {
		fmt.Printf("删除结果数量: %d\n", len(removeResp.RemoveResults))
		for _, result := range removeResp.RemoveResults {
			fmt.Printf("  字段组: Type=%d, SubIdx=%d, Result=%d\n",
				result.GroupType, result.SubIdx, result.Result)
		}
	}

	fmt.Println("\n=== 插入员工字段组 ===")
	gender := uint32(1)
	insertResp, err := client.HR.UpdateStaffInfo(ctx, &hr.UpdateStaffInfoRequest{
		UserID: "zhangsan",
		InsertItems: []hr.InsertItem{
			{
				GroupType: hr.GroupTypeEducation,
				Item: []hr.UpdateItem{
					{
						FieldID:     14001,
						SubIdx:      0,
						ValueString: &name,
					},
					{
						FieldID:     14002,
						SubIdx:      0,
						ValueUint32: &gender,
					},
				},
			},
		},
	})
	if err != nil {
		log.Printf("Failed to insert staff field group: %v", err)
	} else {
		fmt.Printf("插入结果数量: %d\n", len(insertResp.InsertResults))
		for _, result := range insertResp.InsertResults {
			fmt.Printf("  字段组: Type=%d, Idx=%d, Result=%d\n",
				result.GroupType, result.Idx, result.Result)
		}
	}

	fmt.Println("\n=== 示例完成 ===")
}
