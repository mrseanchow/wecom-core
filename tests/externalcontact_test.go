package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/shuaidd/wecom-core"
	"github.com/shuaidd/wecom-core/types/externalcontact"
)

func Test_GetFollowUserList(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.GetFollowUserList(context.Background())
	if err != nil {
		t.Errorf("获取关注用户列表失败 %+v", err)
		return
	}

	fmt.Printf("%+v", resp)
}

func Test_GetCorpTagList(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.GetCorpTagList(context.Background(), &externalcontact.GetCorpTagListRequest{
		GroupID: []string{"etCRbQBwAAkyCwtvE_mvNKd0X-irpL-w"},
	})
	if err != nil {
		return
	}

	fmt.Printf("%+v", resp)
}

func TestListExternalContact(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.ListExternalContact(context.Background(), "20170410022717")
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestGetExternalContact(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.GetExternalContact(context.Background(), "woCRbQBwAAqNfflCX9uvtAnmQ9YpNJZQ", "")
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestBatchGetByUserRequest(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.BatchGetByUser(context.Background(), &externalcontact.BatchGetByUserRequest{
		UserIDList: []string{"20170410022717"},
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestUpdateRemark(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	err = cli.ExternalContact.UpdateRemark(context.Background(), &externalcontact.UpdateRemarkRequest{
		UserID:         "20170410022717",
		Remark:         "测试备注",
		ExternalUserID: "woCRbQBwAAqNfflCX9uvtAnmQ9YpNJZQ",
	})
	fmt.Printf("响应：%+v 错误：%+v", nil, err)
}

func TestAddCorpTag(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.AddCorpTag(context.Background(), &externalcontact.AddCorpTagRequest{
		GroupName: "测试标签组",
		Tag: []externalcontact.AddCorpTagItem{
			{Name: "测试标签1"},
			{Name: "测试标签2"},
		},
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestEditCorpTag(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	err = cli.ExternalContact.EditCorpTag(context.Background(), &externalcontact.EditCorpTagRequest{
		ID:    "etCRbQBwAAkyCwtvE_mvNKd0X-irpL-w",
		Name:  "修改后的标签组",
		Order: 1,
	})
	fmt.Printf("错误：%+v", err)
}

func TestDeleteCorpTag(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	err = cli.ExternalContact.DeleteCorpTag(context.Background(), &externalcontact.DeleteCorpTagRequest{
		TagID:   []string{"etCRbQBwAAkyCwtvE_mvNKd0X-irpL-w"},
		GroupID: []string{},
	})
	fmt.Printf("错误：%+v", err)
}

func TestMarkTag(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	err = cli.ExternalContact.MarkTag(context.Background(), &externalcontact.MarkTagRequest{
		UserID:         "20170410022717",
		ExternalUserID: "woCRbQBwAAqNfflCX9uvtAnmQ9YpNJZQ",
		AddTag:         []string{"etCRbQBwAAkyCwtvE_mvNKd0X-irpL-w"},
		RemoveTag:      []string{},
	})
	fmt.Printf("错误：%+v", err)
}

func TestGetStrategyTagList(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.GetStrategyTagList(context.Background(), &externalcontact.GetStrategyTagListRequest{
		StrategyID: 1,
		GroupID:    []string{},
		TagID:      []string{},
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestAddStrategyTag(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.AddStrategyTag(context.Background(), &externalcontact.AddStrategyTagRequest{
		StrategyID: 1,
		GroupName:  "测试策略标签组",
		Order:      1,
		Tag: []externalcontact.AddCorpTagItem{
			{Name: "策略标签1"},
		},
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestEditStrategyTag(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	err = cli.ExternalContact.EditStrategyTag(context.Background(), &externalcontact.EditStrategyTagRequest{
		ID:    "etCRbQBwAAkyCwtvE_mvNKd0X-irpL-w",
		Name:  "修改后的策略标签",
		Order: 1,
	})
	fmt.Printf("错误：%+v", err)
}

func TestDeleteStrategyTag(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	err = cli.ExternalContact.DeleteStrategyTag(context.Background(), &externalcontact.DeleteStrategyTagRequest{
		TagID:   []string{"etCRbQBwAAkyCwtvE_mvNKd0X-irpL-w"},
		GroupID: []string{},
	})
	fmt.Printf("错误：%+v", err)
}

func TestListGroupChat(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.ListGroupChat(context.Background(), &externalcontact.ListGroupChatRequest{
		Cursor: "",
		Limit:  100,
		OwnerFilter: &externalcontact.OwnerFilter{
			UserIDList: []string{"20170410022717"},
		},
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestGetGroupChat(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.GetGroupChat(context.Background(), &externalcontact.GetGroupChatRequest{
		ChatID:   "wrOgKjCwAAxJnNqT8qM8qM8qM8qM8qM8qM8qM8qM8qM8qM8qM8qM8qM8qM8qM8qM",
		NeedName: 1,
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestOpenGIDToChatID(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.OpenGIDToChatID(context.Background(), &externalcontact.OpenGIDToChatIDRequest{
		OpenGID: "wrgRbQBwAAk8vNqT8qM8qM8qM8qM8qM8qM8qM8qM8qM8qM8qM8qM8qM8qM8qM8qM",
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestTransferGroupChat(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.TransferGroupChat(context.Background(), &externalcontact.TransferGroupChatRequest{
		ChatIDList: []string{"wrOgKjCwAAxJnNqT8qM8qM8qM8qM8qM8qM8qM8qM8qM8qM8qM8qM8qM8qM"},
		NewOwner:   "20170410022717",
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestAddJoinWay(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.AddJoinWay(context.Background(), &externalcontact.AddJoinWayRequest{
		Scene:      1,
		ChatIDList: []string{},
		Remark:     "测试进群方式",
		State:      "test_state",
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestGetJoinWay(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.GetJoinWay(context.Background(), "config_id")
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestUpdateJoinWay(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	err = cli.ExternalContact.UpdateJoinWay(context.Background(), &externalcontact.UpdateJoinWayRequest{
		ConfigID:   "config_id",
		Scene:      1,
		ChatIDList: []string{},
		Remark:     "更新后的备注",
	})
	fmt.Printf("错误：%+v", err)
}

func TestDeleteJoinWay(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	err = cli.ExternalContact.DeleteJoinWay(context.Background(), "config_id")
	fmt.Printf("错误：%+v", err)
}

func TestOnJobTransferCustomer(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.OnJobTransferCustomer(context.Background(), &externalcontact.OnJobTransferCustomerRequest{
		HandoverUserID:     "20170410022717",
		TakeoverUserID:     "20170410022718",
		ExternalUserID:     []string{},
		TransferSuccessMsg: "客户已转移",
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestOnJobTransferGroupChat(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.OnJobTransferGroupChat(context.Background(), &externalcontact.OnJobTransferGroupChatRequest{
		ChatIDList: []string{},
		NewOwner:   "20170410022717",
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestGetTransferResult(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.GetTransferResult(context.Background(), &externalcontact.TransferResultRequest{
		HandoverUserID: "20170410022717",
		TakeoverUserID: "20170410022718",
		Cursor:         "",
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestGetUnassignedList(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.GetUnassignedList(context.Background(), &externalcontact.GetUnassignedListRequest{
		Cursor:   "",
		PageSize: 100,
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestTransferCustomer(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.TransferCustomer(context.Background(), &externalcontact.TransferCustomerRequest{
		HandoverUserID:  "20170410022717",
		TakeoverUserID:  "20170410022718",
		ExternalUserIDs: []string{},
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestListStrategy(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.ListStrategy(context.Background(), &externalcontact.ListStrategyRequest{
		Cursor: "",
		Limit:  100,
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestGetStrategy(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.GetStrategy(context.Background(), 1)
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestGetStrategyRange(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.GetStrategyRange(context.Background(), &externalcontact.GetStrategyRangeRequest{
		StrategyID: 1,
		Cursor:     "",
		Limit:      100,
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestCreateStrategy(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.CreateStrategy(context.Background(), &externalcontact.CreateStrategyRequest{
		ParentID:     0,
		StrategyName: "测试规则组",
		AdminList:    []string{"20170410022717"},
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestEditStrategy(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	err = cli.ExternalContact.EditStrategy(context.Background(), &externalcontact.EditStrategyRequest{
		StrategyID:   1,
		StrategyName: "修改后的规则组",
	})
	fmt.Printf("错误：%+v", err)
}

func TestDeleteStrategy(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	err = cli.ExternalContact.DeleteStrategy(context.Background(), 1)
	fmt.Printf("错误：%+v", err)
}

func TestGetGroupChatStatistic(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.GetGroupChatStatistic(context.Background(), &externalcontact.GroupChatStatisticRequest{
		DayBeginTime: 1704067200,
		DayEndTime:   1704153600,
		OwnerFilter: &externalcontact.OwnerFilter{
			UserIDList: []string{"20170410022717"},
		},
		OrderBy:  1,
		OrderAsc: 0,
		Offset:   0,
		Limit:    100,
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestGetGroupChatStatisticGroupByDay(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.GetGroupChatStatisticGroupByDay(context.Background(), &externalcontact.GroupChatStatisticGroupByDayRequest{
		DayBeginTime: 1704067200,
		DayEndTime:   1704153600,
		OwnerFilter: &externalcontact.OwnerFilter{
			UserIDList: []string{"20170410022717"},
		},
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestGetUserBehaviorData(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.GetUserBehaviorData(context.Background(), &externalcontact.GetUserBehaviorDataRequest{
		UserID:    []string{"20170410022717"},
		StartTime: 1704067200,
		EndTime:   1704153600,
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestAddMomentTask(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.AddMomentTask(context.Background(), &externalcontact.AddMomentTaskRequest{
		Text: &externalcontact.MomentText{
			Content: "测试朋友圈内容",
		},
		VisibleRange: &externalcontact.VisibleRange{
			SenderList: &externalcontact.SenderList{
				UserList: []string{"20170410022717"},
			},
		},
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestGetMomentTaskResult(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.GetMomentTaskResult(context.Background(), "job_id")
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestGetMomentList(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.GetMomentList(context.Background(), &externalcontact.GetMomentListRequest{
		StartTime: 1704067200,
		EndTime:   1704153600,
		Cursor:    "",
		Limit:     100,
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestGetMomentTask(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.GetMomentTask(context.Background(), &externalcontact.GetMomentTaskRequest{
		MomentID: "moment_id",
		Cursor:   "",
		Limit:    100,
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestGetMomentCustomerList(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.GetMomentCustomerList(context.Background(), &externalcontact.GetMomentCustomerListRequest{
		MomentID: "moment_id",
		UserID:   "20170410022717",
		Cursor:   "",
		Limit:    100,
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestGetMomentSendResult(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.GetMomentSendResult(context.Background(), &externalcontact.GetMomentSendResultRequest{
		MomentID: "moment_id",
		UserID:   "20170410022717",
		Cursor:   "",
		Limit:    100,
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestGetMomentComments(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.GetMomentComments(context.Background(), &externalcontact.GetMomentCommentsRequest{
		MomentID: "moment_id",
		UserID:   "20170410022717",
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestCancelMomentTask(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	err = cli.ExternalContact.CancelMomentTask(context.Background(), &externalcontact.CancelMomentTaskRequest{
		MomentID: "moment_id",
	})
	fmt.Printf("错误：%+v", err)
}

func TestListMomentStrategy(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.ListMomentStrategy(context.Background(), &externalcontact.ListMomentStrategyRequest{
		Cursor: "",
		Limit:  100,
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestGetMomentStrategy(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.GetMomentStrategy(context.Background(), &externalcontact.GetMomentStrategyRequest{
		StrategyID: 1,
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestGetMomentStrategyRange(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.GetMomentStrategyRange(context.Background(), &externalcontact.GetMomentStrategyRangeRequest{
		StrategyID: 1,
		Cursor:     "",
		Limit:      100,
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestCreateMomentStrategy(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.ExternalContact.CreateMomentStrategy(context.Background(), &externalcontact.CreateMomentStrategyRequest{
		StrategyName: "测试朋友圈规则组",
		AdminList:    []string{"20170410022717"},
		Range: []externalcontact.MomentStrategyRange{
			{Type: 1, UserID: "20170410022717"},
		},
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestEditMomentStrategy(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	err = cli.ExternalContact.EditMomentStrategy(context.Background(), &externalcontact.EditMomentStrategyRequest{
		StrategyID:   1,
		StrategyName: "修改后的朋友圈规则组",
		AdminList:    []string{"20170410022717"},
	})
	fmt.Printf("错误：%+v", err)
}

func TestDeleteMomentStrategy(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	err = cli.ExternalContact.DeleteMomentStrategy(context.Background(), &externalcontact.DeleteMomentStrategyRequest{
		StrategyID: 1,
	})
	fmt.Printf("错误：%+v", err)
}
