package updown

import (
	"context"

	"github.com/mrseanchow/wecom-core/internal/client"
	"github.com/mrseanchow/wecom-core/types/updown"
)

// GetChainList 获取上下游列�?
// 文档: https://developer.work.weixin.qq.com/document/path/93355
func (s *Service) GetChainList(ctx context.Context) (*updown.GetChainListResponse, error) {
	return client.GetAndUnmarshal[updown.GetChainListResponse](s.client, ctx, "/cgi-bin/corpgroup/corp/get_chain_list", nil)
}

// GetChainGroup 获取上下游通讯录分�?
// 文档: https://developer.work.weixin.qq.com/document/path/93355
func (s *Service) GetChainGroup(ctx context.Context, req *updown.GetChainGroupRequest) (*updown.GetChainGroupResponse, error) {
	return client.PostAndUnmarshal[updown.GetChainGroupResponse](s.client, ctx, "/cgi-bin/corpgroup/corp/get_chain_group", req)
}

// GetChainCorpInfoList 获取企业上下游通讯录分组下的企业详情列�?
// 文档: https://developer.work.weixin.qq.com/document/path/93355
func (s *Service) GetChainCorpInfoList(ctx context.Context, req *updown.GetChainCorpInfoListRequest) (*updown.GetChainCorpInfoListResponse, error) {
	return client.PostAndUnmarshal[updown.GetChainCorpInfoListResponse](s.client, ctx, "/cgi-bin/corpgroup/corp/get_chain_corpinfo_list", req)
}

// GetChainCorpInfo 获取企业上下游通讯录下的企业信�?
// 文档: https://developer.work.weixin.qq.com/document/path/93355
func (s *Service) GetChainCorpInfo(ctx context.Context, req *updown.GetChainCorpInfoRequest) (*updown.GetChainCorpInfoResponse, error) {
	return client.PostAndUnmarshal[updown.GetChainCorpInfoResponse](s.client, ctx, "/cgi-bin/corpgroup/corp/get_chain_corpinfo", req)
}

// RemoveCorp 移除企业
// 文档: https://developer.work.weixin.qq.com/document/path/95820
func (s *Service) RemoveCorp(ctx context.Context, req *updown.RemoveCorpRequest) error {
	_, err := client.PostAndUnmarshal[updown.RemoveCorpResponse](s.client, ctx, "/cgi-bin/corpgroup/corp/remove_corp", req)
	return err
}

// GetChainUserCustomID 查询成员自定义id
// 文档: https://developer.work.weixin.qq.com/document/path/95815
func (s *Service) GetChainUserCustomID(ctx context.Context, req *updown.GetChainUserCustomIDRequest) (string, error) {
	result, err := client.PostAndUnmarshal[updown.GetChainUserCustomIDResponse](s.client, ctx, "/cgi-bin/corpgroup/corp/get_chain_user_custom_id", req)
	if err != nil {
		return "", err
	}
	return result.UserCustomID, nil
}

// GetCorpSharedChainList 获取下级企业加入的上下游
// 文档: https://developer.work.weixin.qq.com/document/path/95816
func (s *Service) GetCorpSharedChainList(ctx context.Context, req *updown.GetCorpSharedChainListRequest) (*updown.GetCorpSharedChainListResponse, error) {
	return client.PostAndUnmarshal[updown.GetCorpSharedChainListResponse](s.client, ctx, "/cgi-bin/corpgroup/get_corp_shared_chain_list", req)
}

