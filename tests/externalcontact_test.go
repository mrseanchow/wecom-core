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
