package contact

import (
	"context"
	"fmt"
	"net/url"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/contact"
)
const (
	AgAddtagusersURL = "/cgi-bin/tag/addtagusers"
	AgCreateURL = "/cgi-bin/tag/create"
	AgDeleteURL = "/cgi-bin/tag/delete"
	AgDeltagusersURL = "/cgi-bin/tag/deltagusers"
	AgGetURL = "/cgi-bin/tag/get"
	AgListURL = "/cgi-bin/tag/list"
	AgUpdateURL = "/cgi-bin/tag/update"
)

// CreateTag 创建标签
// 文档: https://developer.work.weixin.qq.com/document/path/90210
func (s *Service) CreateTag(ctx context.Context, req *contact.CreateTagRequest) (int, error) {
	result, err := client.PostAndUnmarshal[contact.CreateTagResponse](s.client, ctx, AgCreateURL, req)
	if err != nil {
		return 0, err
	}

	return result.TagID, nil
}

// UpdateTag 更新标签名字
// 文档: https://developer.work.weixin.qq.com/document/path/90211
func (s *Service) UpdateTag(ctx context.Context, req *contact.UpdateTagRequest) error {
	_, err := client.PostAndUnmarshal[contact.UpdateTagResponse](s.client, ctx, AgUpdateURL, req)
	return err
}

// DeleteTag 删除标签
// 文档: https://developer.work.weixin.qq.com/document/path/90212
func (s *Service) DeleteTag(ctx context.Context, tagID int) error {
	query := url.Values{}
	query.Set("tagid", fmt.Sprintf("%d", tagID))

	_, err := client.GetAndUnmarshal[contact.DeleteTagResponse](s.client, ctx, AgDeleteURL, query)
	return err
}

// ListTags 获取标签列表
// 文档: https://developer.work.weixin.qq.com/document/path/90216
func (s *Service) ListTags(ctx context.Context) ([]contact.Tag, error) {
	result, err := client.GetAndUnmarshal[contact.ListTagsResponse](s.client, ctx, AgListURL, nil)
	if err != nil {
		return nil, err
	}

	return result.TagList, nil
}

// GetTag 获取标签成员
// 文档: https://developer.work.weixin.qq.com/document/path/90213
func (s *Service) GetTag(ctx context.Context, tagID int) (*contact.GetTagResponse, error) {
	query := url.Values{}
	query.Set("tagid", fmt.Sprintf("%d", tagID))

	return client.GetAndUnmarshal[contact.GetTagResponse](s.client, ctx, AgGetURL, query)
}

// AddTagUsers 增加标签成员
// 文档: https://developer.work.weixin.qq.com/document/path/90214
func (s *Service) AddTagUsers(ctx context.Context, req *contact.AddTagUsersRequest) error {
	_, err := client.PostAndUnmarshal[contact.AddTagUsersResponse](s.client, ctx, AgAddtagusersURL, req)
	return err
}

// DeleteTagUsers 删除标签成员
// 文档: https://developer.work.weixin.qq.com/document/path/90215
func (s *Service) DeleteTagUsers(ctx context.Context, req *contact.DeleteTagUsersRequest) error {
	_, err := client.PostAndUnmarshal[contact.DeleteTagUsersResponse](s.client, ctx, AgDeltagusersURL, req)
	return err
}
