package wedoc

import (
	"context"

	"github.com/mrseanchow/wecom-core/internal/client"
	"github.com/mrseanchow/wecom-core/types/common"
	"github.com/mrseanchow/wecom-core/types/wedoc"
)

const (
	createFormURL       = "/cgi-bin/wedoc/create_form"
	getFormInfoURL      = "/cgi-bin/wedoc/get_form_info"
	modifyFormURL       = "/cgi-bin/wedoc/modify_form"
	getFormAnswerURL    = "/cgi-bin/wedoc/get_form_answer"
	getFormStatisticURL = "/cgi-bin/wedoc/get_form_statistic"
)

// CreateForm 创建收集�?
// 该接口用于创建收集表
func (s *Service) CreateForm(ctx context.Context, req *wedoc.CreateFormRequest) (*wedoc.CreateFormResponse, error) {
	return client.PostAndUnmarshal[wedoc.CreateFormResponse](s.client, ctx, createFormURL, req)
}

// GetFormInfo 获取收集表信�?
// 该接口用于读取收集表的信�?
func (s *Service) GetFormInfo(ctx context.Context, req *wedoc.GetFormInfoRequest) (*wedoc.GetFormInfoResponse, error) {
	return client.PostAndUnmarshal[wedoc.GetFormInfoResponse](s.client, ctx, getFormInfoURL, req)
}

// ModifyForm 编辑收集�?
// 该接口用于编辑收集表
func (s *Service) ModifyForm(ctx context.Context, req *wedoc.ModifyFormRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, modifyFormURL, req)
	return err
}

// GetFormAnswer 读取收集表答�?
// 该接口用于读取收集表的答�?
func (s *Service) GetFormAnswer(ctx context.Context, req *wedoc.GetFormAnswerRequest) (*wedoc.GetFormAnswerResponse, error) {
	return client.PostAndUnmarshal[wedoc.GetFormAnswerResponse](s.client, ctx, getFormAnswerURL, req)
}

// GetFormStatistic 收集表的统计信息查询
// 该接口用于获取收集表的统计信息、已回答成员列表和未回答成员列表
func (s *Service) GetFormStatistic(ctx context.Context, req *wedoc.GetFormStatisticRequest) (*wedoc.GetFormStatisticResponse, error) {
	return client.PostAndUnmarshal[wedoc.GetFormStatisticResponse](s.client, ctx, getFormStatisticURL, req)
}

