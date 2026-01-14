package living

import (
	"context"
	"net/url"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/common"
	"github.com/shuaidd/wecom-core/types/living"
)

// Service 直播服务
type Service struct {
	client *client.Client
}

// NewService 创建直播服务实例
func NewService(c *client.Client) *Service {
	return &Service{client: c}
}

// Create 创建预约直播
// 文档: docs/创建预约直播.md
func (s *Service) Create(ctx context.Context, req *living.CreateRequest) (*living.CreateResponse, error) {
	return client.PostAndUnmarshal[living.CreateResponse](s.client, ctx, "/cgi-bin/living/create", req)
}

// Modify 修改预约直播
// 文档: docs/修改预约直播.md
func (s *Service) Modify(ctx context.Context, req *living.ModifyRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/living/modify", req)
	return err
}

// Cancel 取消预约直播
// 文档: docs/取消预约直播.md
func (s *Service) Cancel(ctx context.Context, req *living.CancelRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/living/cancel", req)
	return err
}

// DeleteReplay 删除直播回放
// 文档: docs/删除直播回放.md
func (s *Service) DeleteReplay(ctx context.Context, req *living.DeleteReplayRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/living/delete_replay_data", req)
	return err
}

// GetLivingCode 获取微信观看直播凭证
// 文档: docs/在微信中观看直播或直播回放.md
func (s *Service) GetLivingCode(ctx context.Context, req *living.GetLivingCodeRequest) (*living.GetLivingCodeResponse, error) {
	return client.PostAndUnmarshal[living.GetLivingCodeResponse](s.client, ctx, "/cgi-bin/living/get_living_code", req)
}

// GetUserAllLivingID 获取成员直播ID列表
// 文档: docs/获取成员直播ID列表.md
func (s *Service) GetUserAllLivingID(ctx context.Context, req *living.GetUserAllLivingIDRequest) (*living.GetUserAllLivingIDResponse, error) {
	return client.PostAndUnmarshal[living.GetUserAllLivingIDResponse](s.client, ctx, "/cgi-bin/living/get_user_all_livingid", req)
}

// GetWatchStat 获取直播观看明细
// 文档: docs/获取直播观看明细.md
func (s *Service) GetWatchStat(ctx context.Context, req *living.GetWatchStatRequest) (*living.GetWatchStatResponse, error) {
	return client.PostAndUnmarshal[living.GetWatchStatResponse](s.client, ctx, "/cgi-bin/living/get_watch_stat", req)
}

// GetLivingInfo 获取直播详情（GET 接口，使用 query 参数）
// 文档: docs/获取直播详情.md
func (s *Service) GetLivingInfo(ctx context.Context, livingID string) (*living.GetLivingInfoResponse, error) {
	q := url.Values{}
	q.Set("livingid", livingID)
	return client.GetAndUnmarshal[living.GetLivingInfoResponse](s.client, ctx, "/cgi-bin/living/get_living_info", q)
}

// GetLivingShareInfo 获取跳转小程序商城的直播观众信息
// 文档: docs/获取跳转小程序商城的直播观众信息.md
func (s *Service) GetLivingShareInfo(ctx context.Context, req *living.GetLivingShareInfoRequest) (*living.GetLivingShareInfoResponse, error) {
	return client.PostAndUnmarshal[living.GetLivingShareInfoResponse](s.client, ctx, "/cgi-bin/living/get_living_share_info", req)
}
