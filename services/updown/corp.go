package updown

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/updown"
)
const (
	OrpgroupCorpGetChainCorpinfoListURL = "/cgi-bin/corpgroup/corp/get_chain_corpinfo_list"
	OrpgroupCorpGetChainCorpinfoURL = "/cgi-bin/corpgroup/corp/get_chain_corpinfo"
	OrpgroupCorpGetChainGroupURL = "/cgi-bin/corpgroup/corp/get_chain_group"
	OrpgroupCorpGetChainListURL = "/cgi-bin/corpgroup/corp/get_chain_list"
	OrpgroupCorpGetChainUserCustomIdURL = "/cgi-bin/corpgroup/corp/get_chain_user_custom_id"
	OrpgroupCorpRemoveCorpURL = "/cgi-bin/corpgroup/corp/remove_corp"
	OrpgroupGetCorpSharedChainListURL = "/cgi-bin/corpgroup/get_corp_shared_chain_list"
)

// GetChainList 获取上下游列表
// 文档: https://developer.work.weixin.qq.com/document/path/93355
func (s *Service) GetChainList(ctx context.Context) (*updown.GetChainListResponse, error) {
	return client.GetAndUnmarshal[updown.GetChainListResponse](s.client, ctx, OrpgroupCorpGetChainListURL, nil)
}

// GetChainGroup 获取上下游通讯录分组
// 文档: https://developer.work.weixin.qq.com/document/path/93355
func (s *Service) GetChainGroup(ctx context.Context, req *updown.GetChainGroupRequest) (*updown.GetChainGroupResponse, error) {
	return client.PostAndUnmarshal[updown.GetChainGroupResponse](s.client, ctx, OrpgroupCorpGetChainGroupURL, req)
}

// GetChainCorpInfoList 获取企业上下游通讯录分组下的企业详情列表
// 文档: https://developer.work.weixin.qq.com/document/path/93355
func (s *Service) GetChainCorpInfoList(ctx context.Context, req *updown.GetChainCorpInfoListRequest) (*updown.GetChainCorpInfoListResponse, error) {
	return client.PostAndUnmarshal[updown.GetChainCorpInfoListResponse](s.client, ctx, OrpgroupCorpGetChainCorpinfoListURL, req)
}

// GetChainCorpInfo 获取企业上下游通讯录下的企业信息
// 文档: https://developer.work.weixin.qq.com/document/path/93355
func (s *Service) GetChainCorpInfo(ctx context.Context, req *updown.GetChainCorpInfoRequest) (*updown.GetChainCorpInfoResponse, error) {
	return client.PostAndUnmarshal[updown.GetChainCorpInfoResponse](s.client, ctx, OrpgroupCorpGetChainCorpinfoURL, req)
}

// RemoveCorp 移除企业
// 文档: https://developer.work.weixin.qq.com/document/path/95820
func (s *Service) RemoveCorp(ctx context.Context, req *updown.RemoveCorpRequest) error {
	_, err := client.PostAndUnmarshal[updown.RemoveCorpResponse](s.client, ctx, OrpgroupCorpRemoveCorpURL, req)
	return err
}

// GetChainUserCustomID 查询成员自定义id
// 文档: https://developer.work.weixin.qq.com/document/path/95815
func (s *Service) GetChainUserCustomID(ctx context.Context, req *updown.GetChainUserCustomIDRequest) (string, error) {
	result, err := client.PostAndUnmarshal[updown.GetChainUserCustomIDResponse](s.client, ctx, OrpgroupCorpGetChainUserCustomIdURL, req)
	if err != nil {
		return "", err
	}
	return result.UserCustomID, nil
}

// GetCorpSharedChainList 获取下级企业加入的上下游
// 文档: https://developer.work.weixin.qq.com/document/path/95816
func (s *Service) GetCorpSharedChainList(ctx context.Context, req *updown.GetCorpSharedChainListRequest) (*updown.GetCorpSharedChainListResponse, error) {
	return client.PostAndUnmarshal[updown.GetCorpSharedChainListResponse](s.client, ctx, OrpgroupGetCorpSharedChainListURL, req)
}
