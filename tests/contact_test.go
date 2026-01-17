package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/shuaidd/wecom-core"
	"github.com/shuaidd/wecom-core/types/contact"
)

func TestContactCreateUser(t *testing.T) {
	t.Run("default config", func(t *testing.T) {
		cli, err := wecom.Default()
		if err != nil {
			t.Errorf("获取cli失败 %+v", err)
			return
		}
		resp, err := cli.Contact.CreateUser(context.Background(), &contact.CreateUserRequest{})
		if err != nil {
			t.Errorf("创建成员失败 %+v", err)
		}
		if resp == nil {
			t.Error("创建成员返回为空")
		}
	})
}

func TestContactGetUser(t *testing.T) {
	t.Run("default config", func(t *testing.T) {
		cli, err := wecom.Default()
		if err != nil {
			t.Errorf("获取cli失败 %+v", err)
			return
		}
		user, err := cli.Contact.GetUser(context.Background(), "test_userid")
		if err != nil {
			t.Errorf("读取成员失败 %+v", err)
		}
		if user == nil {
			t.Error("读取成员返回为空")
		}
	})
}

func TestContactUpdateUser(t *testing.T) {
	t.Run("default config", func(t *testing.T) {
		cli, err := wecom.Default()
		if err != nil {
			t.Errorf("获取cli失败 %+v", err)
			return
		}
		err = cli.Contact.UpdateUser(context.Background(), &contact.UpdateUserRequest{})
		if err != nil {
			t.Errorf("更新成员失败 %+v", err)
		}
	})
}

func TestContactDeleteUser(t *testing.T) {
	t.Run("default config", func(t *testing.T) {
		cli, err := wecom.Default()
		if err != nil {
			t.Errorf("获取cli失败 %+v", err)
			return
		}
		err = cli.Contact.DeleteUser(context.Background(), "test_userid")
		if err != nil {
			t.Errorf("删除成员失败 %+v", err)
		}
	})
}

func TestContactListUsers(t *testing.T) {
	t.Run("default config", func(t *testing.T) {
		cli, err := wecom.Default()
		if err != nil {
			t.Errorf("获取cli失败 %+v", err)
			return
		}
		users, err := cli.Contact.ListUsers(context.Background(), 1, false)
		if err != nil {
			t.Errorf("获取部门成员失败 %+v", err)
		}
		if users == nil {
			t.Error("获取部门成员返回为空")
		}
	})
}

func TestContactListUsersWithChild(t *testing.T) {
	t.Run("default config", func(t *testing.T) {
		cli, err := wecom.Default()
		if err != nil {
			t.Errorf("获取cli失败 %+v", err)
			return
		}
		users, err := cli.Contact.ListUsers(context.Background(), 1, true)
		if err != nil {
			t.Errorf("获取部门成员失败 %+v", err)
		}
		if users == nil {
			t.Error("获取部门成员返回为空")
		}
	})
}

func TestContactListUsersDetail(t *testing.T) {
	t.Run("default config", func(t *testing.T) {
		cli, err := wecom.Default()
		if err != nil {
			t.Errorf("获取cli失败 %+v", err)
			return
		}
		users, err := cli.Contact.ListUsersDetail(context.Background(), 1, false)
		if err != nil {
			t.Errorf("获取部门成员详情失败 %+v", err)
		}
		if users == nil {
			t.Error("获取部门成员详情返回为空")
		}
	})
}

func TestContactListUsersDetailWithChild(t *testing.T) {
	t.Run("default config", func(t *testing.T) {
		cli, err := wecom.Default()
		if err != nil {
			t.Errorf("获取cli失败 %+v", err)
			return
		}
		users, err := cli.Contact.ListUsersDetail(context.Background(), 1, true)
		if err != nil {
			t.Errorf("获取部门成员详情失败 %+v", err)
		}
		if users == nil {
			t.Error("获取部门成员详情返回为空")
		}
	})
}

func TestContactListUserIDs(t *testing.T) {
	t.Run("default config", func(t *testing.T) {
		cli, err := wecom.Default()
		if err != nil {
			t.Errorf("获取cli失败 %+v", err)
			return
		}
		ctx := wecom.WithAgentName(context.Background(), "address-book")
		resp, err := cli.Contact.ListUserIDs(ctx, &contact.ListUserIDsRequest{})
		fmt.Printf("响应：%+v 错误：%+v", resp, err)
	})
}

func TestContactAuthSuccess(t *testing.T) {
	t.Run("default config", func(t *testing.T) {
		cli, err := wecom.Default()
		if err != nil {
			t.Errorf("获取cli失败 %+v", err)
			return
		}
		err = cli.Contact.AuthSuccess(context.Background(), "test_userid")
		if err != nil {
			t.Errorf("二次验证失败 %+v", err)
		}
	})
}

func TestContactConvertToOpenID(t *testing.T) {
	t.Run("default config", func(t *testing.T) {
		cli, err := wecom.Default()
		if err != nil {
			t.Errorf("获取cli失败 %+v", err)
			return
		}
		openid, err := cli.Contact.ConvertToOpenID(context.Background(), "test_userid")
		if err != nil {
			t.Errorf("userid转openid失败 %+v", err)
		}
		if openid == "" {
			t.Error("userid转openid返回为空")
		}
	})
}

func TestContactConvertToUserID(t *testing.T) {
	t.Run("default config", func(t *testing.T) {
		cli, err := wecom.Default()
		if err != nil {
			t.Errorf("获取cli失败 %+v", err)
			return
		}
		userid, err := cli.Contact.ConvertToUserID(context.Background(), "test_openid")
		if err != nil {
			t.Errorf("openid转userid失败 %+v", err)
		}
		if userid == "" {
			t.Error("openid转userid返回为空")
		}
	})
}

func TestContactGetUserIDByEmail(t *testing.T) {
	t.Run("default config", func(t *testing.T) {
		cli, err := wecom.Default()
		if err != nil {
			t.Errorf("获取cli失败 %+v", err)
			return
		}
		userid, err := cli.Contact.GetUserIDByEmail(context.Background(), "test@example.com", 1)
		if err != nil {
			t.Errorf("通过邮箱获取userid失败 %+v", err)
		}
		if userid == "" {
			t.Error("通过邮箱获取userid返回为空")
		}
	})
}

func TestContactGetUserIDByMobile(t *testing.T) {
	t.Run("default config", func(t *testing.T) {
		cli, err := wecom.Default()
		if err != nil {
			t.Errorf("获取cli失败 %+v", err)
			return
		}
		userid, err := cli.Contact.GetUserIDByMobile(context.Background(), "13800000000")
		if err != nil {
			t.Errorf("通过手机号获取userid失败 %+v", err)
		}
		if userid == "" {
			t.Error("通过手机号获取userid返回为空")
		}
	})
}

func TestContactBatchDeleteUsers(t *testing.T) {
	t.Run("default config", func(t *testing.T) {
		cli, err := wecom.Default()
		if err != nil {
			t.Errorf("获取cli失败 %+v", err)
			return
		}
		err = cli.Contact.BatchDeleteUsers(context.Background(), []string{"user1", "user2"})
		if err != nil {
			t.Errorf("批量删除成员失败 %+v", err)
		}
	})
}
