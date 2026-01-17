package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/shuaidd/wecom-core"
	"github.com/shuaidd/wecom-core/types/wedoc"
)

func TestCreateDoc(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	resp, err := cli.Wedoc.CreateDoc(context.Background(), &wedoc.CreateDocRequest{
		DocType:    10,
		DocName:    "测试智能文档",
		AdminUsers: []string{"20170410022717"},
	})
	//响应：&{URL:https://doc.weixin.qq.com/doc/w3_AFwAABSbALsCNaOnqx7cIR569FBfI_a?scode=AAkAFgfQAAcgICoIbJAFwAABSbALs
	// DocID:dcHdADPmk_Tm5skaL18An-Qx5-NcQEnuhTrmrAak1nruxaJHFHpN19ke9SxqTB1vaiIkCwuz3Fqyz-cevQMtQjmQ}
	//响应：&{URL:https://doc.weixin.qq.com/smartsheet/s3_AFwAABSbALsCN1DZ3y9LEStSWm2SO_a?scode=AAkAFgfQAAc4ljH8zrAFwAABSbALs
	// DocID:dcVtxay44NP1VpJjHVPWUX1KTo-AANc5bUfNRRucCLcjgcJY4y89dvZRRKXKER6y1l5rlQeFNcn-F3vHpWrNXAxg} 错误：<nil>
	fmt.Printf("响应：%+v 错误：%+v\n", resp, err)
}

func TestGetDocBaseInfo(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	resp, err := cli.Wedoc.GetDocBaseInfo(context.Background(), &wedoc.GetDocBaseInfoRequest{
		DocID: "dcHdADPmk_Tm5skaL18An-Qx5-NcQEnuhTrmrAak1nruxaJHFHpN19ke9SxqTB1vaiIkCwuz3Fqyz-cevQMtQjmQ",
	})
	fmt.Printf("响应：%+v 错误：%+v\n", resp, err)
}

func TestRenameDoc(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	err = cli.Wedoc.RenameDoc(context.Background(), &wedoc.RenameDocRequest{
		DocID:   "dcHdADPmk_Tm5skaL18An-Qx5-NcQEnuhTrmrAak1nruxaJHFHpN19ke9SxqTB1vaiIkCwuz3Fqyz-cevQMtQjmQ",
		NewName: "重命名后的测试文档",
	})
	fmt.Printf("错误：%+v\n", err)
}

func TestShareDoc(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	resp, err := cli.Wedoc.ShareDoc(context.Background(), &wedoc.ShareDocRequest{
		DocID: "dcHdADPmk_Tm5skaL18An-Qx5-NcQEnuhTrmrAak1nruxaJHFHpN19ke9SxqTB1vaiIkCwuz3Fqyz-cevQMtQjmQ",
	})
	fmt.Printf("响应：%+v 错误：%+v\n", resp, err)
}

func TestDeleteDoc(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	err = cli.Wedoc.DeleteDoc(context.Background(), &wedoc.DeleteDocRequest{
		DocID: "dcHdADPmk_Tm5skaL18An-Qx5-NcQEnuhTrmrAak1nruxaJHFHpN19ke9SxqTB1vaiIkCwuz3Fqyz-cevQMtQjmQ",
	})
	fmt.Printf("错误：%+v\n", err)
}

func TestAddRecords(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	resp, err := cli.Wedoc.AddRecords(context.Background(), &wedoc.AddRecordsRequest{
		DocID:   "test_doc_id",
		SheetID: "test_sheet_id",
		KeyType: "field_title",
		Records: []wedoc.AddRecord{
			{
				Values: map[string]interface{}{
					"姓名": "张三",
					"年龄": 25,
					"部门": "技术部",
				},
			},
		},
	})
	fmt.Printf("响应：%+v 错误：%+v\n", resp, err)
}

func TestGetRecords(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	resp, err := cli.Wedoc.GetRecords(context.Background(), &wedoc.GetRecordsRequest{
		DocID:   "dcVtxay44NP1VpJjHVPWUX1KTo-AANc5bUfNRRucCLcjgcJY4y89dvZRRKXKER6y1l5rlQeFNcn-F3vHpWrNXAxg",
		SheetID: "q979lj",
		KeyType: "field_title",
		Limit:   100,
	})
	fmt.Printf("响应：%+v 错误：%+v\n", resp, err)
}

func TestUpdateRecords(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	resp, err := cli.Wedoc.UpdateRecords(context.Background(), &wedoc.UpdateRecordsRequest{
		DocID:   "dcVtxay44NP1VpJjHVPWUX1KTo-AANc5bUfNRRucCLcjgcJY4y89dvZRRKXKER6y1l5rlQeFNcn-F3vHpWrNXAxg",
		SheetID: "q979lj",
		KeyType: "field_title",
		Records: []wedoc.UpdateRecord{
			{
				RecordID: "test_record_id",
				Values: map[string]interface{}{
					"姓名": "李四",
					"年龄": 30,
				},
			},
		},
	})
	fmt.Printf("响应：%+v 错误：%+v\n", resp, err)
}

func TestDeleteRecords(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	err = cli.Wedoc.DeleteRecords(context.Background(), &wedoc.DeleteRecordsRequest{
		DocID:     "dcVtxay44NP1VpJjHVPWUX1KTo-AANc5bUfNRRucCLcjgcJY4y89dvZRRKXKER6y1l5rlQeFNcn-F3vHpWrNXAxg",
		SheetID:   "q979lj",
		RecordIDs: []string{"test_record_id"},
	})
	fmt.Printf("错误：%+v\n", err)
}

func TestAddFields(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	resp, err := cli.Wedoc.AddFields(context.Background(), &wedoc.AddFieldsRequest{
		DocID:   "test_doc_id",
		SheetID: "test_sheet_id",
		Fields: []wedoc.AddField{
			{
				FieldTitle: "姓名",
				FieldType:  "text",
			},
			{
				FieldTitle: "年龄",
				FieldType:  "number",
				PropertyNumber: &wedoc.NumberFieldProperty{
					DecimalPlaces: 0,
				},
			},
		},
	})
	fmt.Printf("响应：%+v 错误：%+v\n", resp, err)
}

func TestGetFields(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	resp, err := cli.Wedoc.GetFields(context.Background(), &wedoc.GetFieldsRequest{
		DocID:   "dcVtxay44NP1VpJjHVPWUX1KTo-AANc5bUfNRRucCLcjgcJY4y89dvZRRKXKER6y1l5rlQeFNcn-F3vHpWrNXAxg",
		SheetID: "q979lj",
	})
	fmt.Printf("响应：%+v 错误：%+v\n", resp, err)
}

func TestUpdateFields(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	resp, err := cli.Wedoc.UpdateFields(context.Background(), &wedoc.UpdateFieldsRequest{
		DocID:   "test_doc_id",
		SheetID: "test_sheet_id",
		Fields: []wedoc.UpdateField{
			{
				FieldID:    "test_field_id",
				FieldTitle: "新字段名",
			},
		},
	})
	fmt.Printf("响应：%+v 错误：%+v\n", resp, err)
}

func TestDeleteFields(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	err = cli.Wedoc.DeleteFields(context.Background(), &wedoc.DeleteFieldsRequest{
		DocID:    "test_doc_id",
		SheetID:  "test_sheet_id",
		FieldIDs: []string{"test_field_id"},
	})
	fmt.Printf("错误：%+v\n", err)
}

func TestAddView(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	resp, err := cli.Wedoc.AddView(context.Background(), &wedoc.AddViewRequest{
		DocID:     "test_doc_id",
		SheetID:   "test_sheet_id",
		ViewTitle: "测试视图",
		ViewType:  "grid",
	})
	fmt.Printf("响应：%+v 错误：%+v\n", resp, err)
}

func TestGetViews(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	resp, err := cli.Wedoc.GetViews(context.Background(), &wedoc.GetViewsRequest{
		DocID:   "dcVtxay44NP1VpJjHVPWUX1KTo-AANc5bUfNRRucCLcjgcJY4y89dvZRRKXKER6y1l5rlQeFNcn-F3vHpWrNXAxg",
		SheetID: "q979lj",
	})
	fmt.Printf("响应：%+v 错误：%+v\n", resp, err)
}

func TestUpdateView(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	resp, err := cli.Wedoc.UpdateView(context.Background(), &wedoc.UpdateViewRequest{
		DocID:     "test_doc_id",
		SheetID:   "test_sheet_id",
		ViewID:    "test_view_id",
		ViewTitle: "更新的视图名称",
	})
	fmt.Printf("响应：%+v 错误：%+v\n", resp, err)
}

func TestDeleteView(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	err = cli.Wedoc.DeleteView(context.Background(), &wedoc.DeleteViewRequest{
		DocID:   "test_doc_id",
		SheetID: "test_sheet_id",
		ViewID:  "test_view_id",
	})
	fmt.Printf("错误：%+v\n", err)
}

func TestAddSheet(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	resp, err := cli.Wedoc.AddSheet(context.Background(), &wedoc.AddSheetRequest{
		DocID: "test_doc_id",
		Properties: &wedoc.SheetProperty{
			Title: "测试子表",
			Index: 0,
		},
	})
	fmt.Printf("响应：%+v 错误：%+v\n", resp, err)
}

func TestGetSheet(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	resp, err := cli.Wedoc.GetSheet(context.Background(), &wedoc.GetSheetRequest{
		DocID: "dcVtxay44NP1VpJjHVPWUX1KTo-AANc5bUfNRRucCLcjgcJY4y89dvZRRKXKER6y1l5rlQeFNcn-F3vHpWrNXAxg",
	})
	fmt.Printf("响应：%+v 错误：%+v\n", resp, err)
}

func TestUpdateSheet(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	resp, err := cli.Wedoc.UpdateSheet(context.Background(), &wedoc.UpdateSheetRequest{
		DocID:   "test_doc_id",
		SheetID: "test_sheet_id",
		Properties: &wedoc.SheetProperty{
			Title: "更新的子表名称",
			Index: 0,
		},
	})
	fmt.Printf("响应：%+v 错误：%+v\n", resp, err)
}

func TestDeleteSheet(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	err = cli.Wedoc.DeleteSheet(context.Background(), &wedoc.DeleteSheetRequest{
		DocID:   "test_doc_id",
		SheetID: "test_sheet_id",
	})
	fmt.Printf("错误：%+v\n", err)
}

func TestAddFieldGroup(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	resp, err := cli.Wedoc.AddFieldGroup(context.Background(), &wedoc.AddFieldGroupRequest{
		DocID:   "test_doc_id",
		SheetID: "test_sheet_id",
		Name:    "基本信息",
		Children: []wedoc.FieldGroupChildren{
			{FieldID: "field1"},
			{FieldID: "field2"},
		},
	})
	fmt.Printf("响应：%+v 错误：%+v\n", resp, err)
}

func TestGetFieldGroup(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	resp, err := cli.Wedoc.GetFieldGroup(context.Background(), &wedoc.GetFieldGroupRequest{
		DocID:   "dcVtxay44NP1VpJjHVPWUX1KTo-AANc5bUfNRRucCLcjgcJY4y89dvZRRKXKER6y1l5rlQeFNcn-F3vHpWrNXAxg",
		SheetID: "q979lj",
	})
	fmt.Printf("响应：%+v 错误：%+v\n", resp, err)
}

func TestUpdateFieldGroup(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	resp, err := cli.Wedoc.UpdateFieldGroup(context.Background(), &wedoc.UpdateFieldGroupRequest{
		DocID:        "test_doc_id",
		SheetID:      "test_sheet_id",
		FieldGroupID: "test_field_group_id",
		Name:         "更新后的编组名称",
		Children: []wedoc.FieldGroupChildren{
			{FieldID: "field1"},
			{FieldID: "field2"},
			{FieldID: "field3"},
		},
	})
	fmt.Printf("响应：%+v 错误：%+v\n", resp, err)
}

func TestDeleteFieldGroup(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	err = cli.Wedoc.DeleteFieldGroup(context.Background(), &wedoc.DeleteFieldGroupRequest{
		DocID:        "test_doc_id",
		SheetID:      "test_sheet_id",
		FieldGroupID: "test_field_group_id",
	})
	fmt.Printf("错误：%+v\n", err)
}

func TestCreateForm(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	resp, err := cli.Wedoc.CreateForm(context.Background(), &wedoc.CreateFormRequest{
		FormInfo: wedoc.FormInfo{
			FormTitle:  "测试收集表",
			FormDesc:   "这是一个测试收集表",
			FormHeader: "https://wework.qpic.cn/wwpic3az/291238_YYSv1K_mQn6oBN4_1767923045/0",
			FormQuestion: wedoc.FormQuestion{
				Items: []wedoc.QuestionItem{
					{
						QuestionID: 1,
						Title:      "您的姓名",
						Pos:        1,
						Status:     1,
						ReplyType:  1,
						MustReply:  true,
						QuestionExtendSetting: wedoc.QuestionExtendSetting{
							TextSetting: &wedoc.TextSetting{
								CharLen: 50,
							},
						},
					},
					{
						QuestionID: 2,
						Title:      "性别",
						Pos:        2,
						Status:     1,
						ReplyType:  2,
						MustReply:  true,
						OptionItem: []wedoc.OptionItem{
							{Key: 1, Value: "男", Status: 1},
							{Key: 2, Value: "女", Status: 1},
						},
					},
				},
			},
			FormSetting: wedoc.FormSetting{
				FillOutAuth:     1,
				AllowMultiFill:  false,
				CanAnonymous:    true,
				CanNotifySubmit: false,
				FillInRange: &wedoc.FillInRange{
					UserIDs: []string{"20170410022717"},
				},
				SettingManagerRange: &wedoc.SettingManagerRange{
					UserIDs: []string{"20170410022717"},
				},
			},
		},
	})
	fmt.Printf("响应：%+v 错误：%+v\n", resp, err)
}

func TestGetFormInfo(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	resp, err := cli.Wedoc.GetFormInfo(context.Background(), &wedoc.GetFormInfoRequest{
		FormID: "fm8M9G5UvsD9y3R3ybNBz48JEmEzE7WqMQ8yPvYIqBi_-h0pcQE-2iJW9tcv1fbDfI_Z2pTIHaeU-LiReIug7cuA",
	})
	fmt.Printf("响应：%+v 错误：%+v\n", resp, err)
}

func TestModifyForm(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	err = cli.Wedoc.ModifyForm(context.Background(), &wedoc.ModifyFormRequest{
		Oper:   1,
		FormID: "fm8M9G5UvsD9y3R3ybNBz48JEmEzE7WqMQ8yPvYIqBi_-h0pcQE-2iJW9tcv1fbDfI_Z2pTIHaeU-LiReIug7cuA",
		FormInfo: wedoc.FormInfo{
			FormTitle:  "修改后的收集表",
			FormDesc:   "修改后的描述",
			FormHeader: "https://example.com/new_header.jpg",
			FormQuestion: wedoc.FormQuestion{
				Items: []wedoc.QuestionItem{
					{
						QuestionID: 1,
						Title:      "您的姓名",
						Pos:        1,
						Status:     1,
						ReplyType:  1,
						MustReply:  true,
					},
					{
						QuestionID: 2,
						Title:      "性别",
						Pos:        2,
						Status:     1,
						ReplyType:  2,
						MustReply:  true,
						OptionItem: []wedoc.OptionItem{
							{Key: 1, Value: "男", Status: 1},
							{Key: 2, Value: "女", Status: 1},
							{Key: 3, Value: "其他", Status: 1},
						},
					},
				},
			},
		},
	})
	fmt.Printf("错误：%+v\n", err)
}

func TestGetFormAnswer(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}
	resp, err := cli.Wedoc.GetFormAnswer(context.Background(), &wedoc.GetFormAnswerRequest{
		RepeatedID: "fm8M9G5UvsD9y3R3ybNBz48JEmEzE7WqMQ8yPvYIqBi_-h0pcQE-2iJW9tcv1fbDfI_Z2pTIHaeU-LiReIug7cuA",
		AnswerIDs:  []uint64{1},
	})
	fmt.Printf("响应：%+v 错误：%+v\n", resp, err)
}

func TestGetFormStatistic(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		t.Errorf("创建默认客户端失败 %+v", err)
	}

	resp, err := cli.Wedoc.GetFormStatistic(context.Background(), &wedoc.GetFormStatisticRequest{
		RepeatedID: "test_repeated_id",
		ReqType:    1,
		Limit:      100,
	})
	fmt.Printf("统计信息响应：%+v 错误：%+v\n", resp, err)

	resp2, err := cli.Wedoc.GetFormStatistic(context.Background(), &wedoc.GetFormStatisticRequest{
		RepeatedID: "test_repeated_id",
		ReqType:    2,
		StartTime:  1700000000,
		EndTime:    1700086400,
		Limit:      100,
	})
	fmt.Printf("已提交列表响应：%+v 错误：%+v\n", resp2, err)

	resp3, err := cli.Wedoc.GetFormStatistic(context.Background(), &wedoc.GetFormStatisticRequest{
		RepeatedID: "test_repeated_id",
		ReqType:    3,
		Limit:      100,
	})
	fmt.Printf("未提交列表响应：%+v 错误：%+v\n", resp3, err)
}
