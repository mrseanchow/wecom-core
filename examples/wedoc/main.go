package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mrseanchow/wecom-core"
	"github.com/mrseanchow/wecom-core/config"
	"github.com/mrseanchow/wecom-core/pkg/logger"
	"github.com/mrseanchow/wecom-core/types/wedoc"
)

func main() {
	// 1. 创建企业微信客户�?
	client, err := wecom.New(
		config.WithCorpID("your_corp_id"),
		config.WithLogger(logger.NewStdLogger()),
		config.WithRetry(3),
	)
	if err != nil {
		log.Fatalf("Failed to create wecom client: %v", err)
	}

	ctx := context.Background()

	// ========== 文档管理示例 ==========

	// 2. 新建文档
	fmt.Println("=== 新建文档 ===")
	createDocResp, err := client.Wedoc.CreateDoc(ctx, &wedoc.CreateDocRequest{
		DocType:    3, // 3:文档 4:表格 10:智能表格
		DocName:    "测试文档",
		AdminUsers: []string{"zhangsan"},
	})
	if err != nil {
		log.Printf("Failed to create doc: %v", err)
	} else {
		fmt.Printf("文档创建成功: DocID=%s, URL=%s\n", createDocResp.DocID, createDocResp.URL)
	}

	// 保存docid供后续使�?
	docID := ""
	if createDocResp != nil {
		docID = createDocResp.DocID
	}

	// 3. 获取文档基础信息
	if docID != "" {
		fmt.Println("\n=== 获取文档基础信息 ===")
		docInfo, err := client.Wedoc.GetDocBaseInfo(ctx, &wedoc.GetDocBaseInfoRequest{
			DocID: docID,
		})
		if err != nil {
			log.Printf("Failed to get doc info: %v", err)
		} else {
			fmt.Printf("文档信息: DocID=%s, Name=%s, Type=%d\n",
				docInfo.DocBaseInfo.DocID,
				docInfo.DocBaseInfo.DocName,
				docInfo.DocBaseInfo.DocType)
		}
	}

	// 4. 重命名文�?
	if docID != "" {
		fmt.Println("\n=== 重命名文�?===")
		err = client.Wedoc.RenameDoc(ctx, &wedoc.RenameDocRequest{
			DocID:   docID,
			NewName: "新测试文档名�?,
		})
		if err != nil {
			log.Printf("Failed to rename doc: %v", err)
		} else {
			fmt.Println("文档重命名成�?)
		}
	}

	// 5. 分享文档
	if docID != "" {
		fmt.Println("\n=== 分享文档 ===")
		shareResp, err := client.Wedoc.ShareDoc(ctx, &wedoc.ShareDocRequest{
			DocID: docID,
		})
		if err != nil {
			log.Printf("Failed to share doc: %v", err)
		} else {
			fmt.Printf("文档分享链接: %s\n", shareResp.ShareURL)
		}
	}

	// ========== 收集表管理示�?==========

	// 6. 创建收集�?
	fmt.Println("\n=== 创建收集�?===")
	createFormResp, err := client.Wedoc.CreateForm(ctx, &wedoc.CreateFormRequest{
		FormInfo: wedoc.FormInfo{
			FormTitle: "员工信息收集�?,
			FormDesc:  "请填写您的个人信�?,
			FormQuestion: wedoc.FormQuestion{
				Items: []wedoc.QuestionItem{
					{
						QuestionID: 1,
						Title:      "您的姓名",
						Pos:        1,
						Status:     1,
						ReplyType:  1, // 1:文本
						MustReply:  true,
					},
					{
						QuestionID: 2,
						Title:      "您的部门",
						Pos:        2,
						Status:     1,
						ReplyType:  2, // 2:单�?
						MustReply:  true,
						OptionItem: []wedoc.OptionItem{
							{Key: 1, Value: "技术部", Status: 1},
							{Key: 2, Value: "产品�?, Status: 1},
							{Key: 3, Value: "运营�?, Status: 1},
						},
					},
					{
						QuestionID: 3,
						Title:      "您的爱好（多选）",
						Pos:        3,
						Status:     1,
						ReplyType:  3, // 3:多�?
						MustReply:  false,
						OptionItem: []wedoc.OptionItem{
							{Key: 1, Value: "运动", Status: 1},
							{Key: 2, Value: "阅读", Status: 1},
							{Key: 3, Value: "旅游", Status: 1},
							{Key: 4, Value: "音乐", Status: 1},
						},
					},
				},
			},
			FormSetting: wedoc.FormSetting{
				FillOutAuth:     0,     // 0:所有人可填�?
				AllowMultiFill:  false, // 不允许多次填�?
				CanAnonymous:    false, // 不允许匿�?
				CanNotifySubmit: true,  // 有回复时提醒
			},
		},
	})
	if err != nil {
		log.Printf("Failed to create form: %v", err)
	} else {
		fmt.Printf("收集表创建成�? FormID=%s\n", createFormResp.FormID)
	}

	// 保存formid供后续使�?
	formID := ""
	if createFormResp != nil {
		formID = createFormResp.FormID
	}

	// 7. 获取收集表信�?
	if formID != "" {
		fmt.Println("\n=== 获取收集表信�?===")
		formInfo, err := client.Wedoc.GetFormInfo(ctx, &wedoc.GetFormInfoRequest{
			FormID: formID,
		})
		if err != nil {
			log.Printf("Failed to get form info: %v", err)
		} else {
			fmt.Printf("收集表信�? FormID=%s, Title=%s, 问题�?%d\n",
				formInfo.FormInfo.FormID,
				formInfo.FormInfo.FormTitle,
				len(formInfo.FormInfo.FormQuestion.Items))
		}
	}

	// 8. 编辑收集表（修改设置�?
	if formID != "" {
		fmt.Println("\n=== 编辑收集�?===")
		err = client.Wedoc.ModifyForm(ctx, &wedoc.ModifyFormRequest{
			Oper:   2, // 2:全量修改设置
			FormID: formID,
			FormInfo: wedoc.FormInfo{
				FormSetting: wedoc.FormSetting{
					FillOutAuth:     0,
					AllowMultiFill:  true, // 修改为允许多次填�?
					CanAnonymous:    false,
					CanNotifySubmit: true,
				},
			},
		})
		if err != nil {
			log.Printf("Failed to modify form: %v", err)
		} else {
			fmt.Println("收集表编辑成�?)
		}
	}

	// 9. 获取收集表统计信�?
	if formID != "" {
		// 注意：需要先获取收集表信息以获得repeated_id
		formInfo, err := client.Wedoc.GetFormInfo(ctx, &wedoc.GetFormInfoRequest{
			FormID: formID,
		})
		if err == nil && len(formInfo.FormInfo.RepeatedID) > 0 {
			fmt.Println("\n=== 获取收集表统计信�?===")
			repeatedID := formInfo.FormInfo.RepeatedID[0]
			statistic, err := client.Wedoc.GetFormStatistic(ctx, &wedoc.GetFormStatisticRequest{
				RepeatedID: repeatedID,
				ReqType:    1, // 1:只获取统计结�?
			})
			if err != nil {
				log.Printf("Failed to get form statistic: %v", err)
			} else {
				fmt.Printf("统计信息: 已填写次�?%d, 已填写人�?%d, 未填写人�?%d\n",
					statistic.FillCnt,
					statistic.FillUserCnt,
					statistic.UnfillUserCnt)
			}

			// 10. 获取已提交列表（示例�?
			fmt.Println("\n=== 获取已提交列�?===")
			submitList, err := client.Wedoc.GetFormStatistic(ctx, &wedoc.GetFormStatisticRequest{
				RepeatedID: repeatedID,
				ReqType:    2,          // 2:获取已提交列�?
				StartTime:  1700000000, // 示例时间�?
				EndTime:    1800000000, // 示例时间�?
				Limit:      20,
			})
			if err != nil {
				log.Printf("Failed to get submit list: %v", err)
			} else {
				fmt.Printf("已提交人�? %d\n", len(submitList.SubmitUsers))
			}

			// 11. 读取收集表答案（示例�?
			if len(submitList.SubmitUsers) > 0 {
				fmt.Println("\n=== 读取收集表答�?===")
				var answerIDs []uint64
				for _, user := range submitList.SubmitUsers {
					answerIDs = append(answerIDs, user.AnswerID)
					if len(answerIDs) >= 10 { // 最多读�?0�?
						break
					}
				}

				answers, err := client.Wedoc.GetFormAnswer(ctx, &wedoc.GetFormAnswerRequest{
					RepeatedID: repeatedID,
					AnswerIDs:  answerIDs,
				})
				if err != nil {
					log.Printf("Failed to get form answers: %v", err)
				} else {
					fmt.Printf("答案数量: %d\n", len(answers.Answer.AnswerList))
					for _, answer := range answers.Answer.AnswerList {
						fmt.Printf("  - 答案ID=%d, 用户=%s, 问题回答�?%d\n",
							answer.AnswerID,
							answer.UserName,
							len(answer.Reply.Items))
					}
				}
			}
		}
	}

	// 12. 删除文档（清理）
	if docID != "" {
		fmt.Println("\n=== 删除文档 ===")
		err = client.Wedoc.DeleteDoc(ctx, &wedoc.DeleteDocRequest{
			DocID: docID,
		})
		if err != nil {
			log.Printf("Failed to delete doc: %v", err)
		} else {
			fmt.Println("文档删除成功")
		}
	}

	// 13. 删除收集表（使用DeleteDoc接口，传入FormID�?
	if formID != "" {
		fmt.Println("\n=== 删除收集�?===")
		err = client.Wedoc.DeleteDoc(ctx, &wedoc.DeleteDocRequest{
			FormID: formID,
		})
		if err != nil {
			log.Printf("Failed to delete form: %v", err)
		} else {
			fmt.Println("收集表删除成�?)
		}
	}

	fmt.Println("\n=== 示例完成 ===")
}

