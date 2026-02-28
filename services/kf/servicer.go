package kf

import (
	"context"
	"net/url"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/kf"
)
const (
	FServicerAddURL = "/cgi-bin/kf/servicer/add"
	FServicerDelURL = "/cgi-bin/kf/servicer/del"
	FServicerListURL = "/cgi-bin/kf/servicer/list"
)

// AddServicer 添加接待人员
// 添加指定客服账号的接待人员,每个客服账号目前最多可添加2000个接待人员,20个部门
// 文档: https://developer.work.weixin.qq.com/document/path/94646
func (s *Service) AddServicer(ctx context.Context, req *kf.AddServicerRequest) (*kf.AddServicerResponse, error) {
	return client.PostAndUnmarshal[kf.AddServicerResponse](s.client, ctx, FServicerAddURL, req)
}

// DeleteServicer 删除接待人员
// 从客服账号删除接待人员
// 文档: https://developer.work.weixin.qq.com/document/path/94647
func (s *Service) DeleteServicer(ctx context.Context, req *kf.DeleteServicerRequest) (*kf.DeleteServicerResponse, error) {
	return client.PostAndUnmarshal[kf.DeleteServicerResponse](s.client, ctx, FServicerDelURL, req)
}

// ListServicer 获取接待人员列表
// 获取某个客服账号的接待人员列表
// 文档: https://developer.work.weixin.qq.com/document/path/94645
func (s *Service) ListServicer(ctx context.Context, req *kf.ListServicerRequest) (*kf.ListServicerResponse, error) {
	params := url.Values{}
	params.Set("open_kfid", req.OpenKfID)
	return client.GetAndUnmarshal[kf.ListServicerResponse](s.client, ctx, FServicerListURL, params)
}
