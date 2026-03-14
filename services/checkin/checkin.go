package checkin

import (
	"context"

	"github.com/mrseanchow/wecom-core/internal/client"
	"github.com/mrseanchow/wecom-core/types/checkin"
	"github.com/mrseanchow/wecom-core/types/common"
)

// Service 打卡相关接口服务
type Service struct {
	client *client.Client
}

// NewService 创建打卡服务
func NewService(c *client.Client) *Service {
	return &Service{client: c}
}

// Request / Response types 已移动到 types/checkin 包，服务中仅引用 types�?
// 详见: types/checkin/checkin.go

//
// Methods: 调用内部 client 并返回对应类�?
//

// SetCheckinScheduleList 为打卡人员排�?
func (s *Service) SetCheckinScheduleList(ctx context.Context, req *checkin.SetCheckinScheduleListRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/checkin/setcheckinschedulist", req)
	return err
}

// PunchCorrection 为打卡人员补�?
func (s *Service) PunchCorrection(ctx context.Context, req *checkin.PunchCorrectionRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/checkin/punch_correction", req)
	return err
}

// AddCheckinUserFace 录入人脸信息
func (s *Service) AddCheckinUserFace(ctx context.Context, req *checkin.AddCheckinUserFaceRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/checkin/addcheckinuserface", req)
	return err
}

// AddCheckinRecord 添加打卡记录
func (s *Service) AddCheckinRecord(ctx context.Context, req *checkin.AddCheckinRecordRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/checkin/add_checkin_record", req)
	return err
}

// AddCheckinOption 创建打卡规则
func (s *Service) AddCheckinOption(ctx context.Context, req *checkin.AddCheckinOptionRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/checkin/add_checkin_option", req)
	return err
}

// UpdateCheckinOption 修改打卡规则
func (s *Service) UpdateCheckinOption(ctx context.Context, req *checkin.UpdateCheckinOptionRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/checkin/update_checkin_option", req)
	return err
}

// ClearCheckinOptionArrayField 清空打卡规则数组元素
func (s *Service) ClearCheckinOptionArrayField(ctx context.Context, req *checkin.ClearCheckinOptionArrayFieldRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/checkin/clear_checkin_option_array_field", req)
	return err
}

// DeleteCheckinOption 删除打卡规则
func (s *Service) DeleteCheckinOption(ctx context.Context, req *checkin.DeleteCheckinOptionRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/checkin/del_checkin_option", req)
	return err
}

// GetCorpCheckinOption 获取企业所有打卡规�?
func (s *Service) GetCorpCheckinOption(ctx context.Context) (*checkin.GetCorpCheckinOptionResponse, error) {
	return client.PostAndUnmarshal[checkin.GetCorpCheckinOptionResponse](s.client, ctx, "/cgi-bin/checkin/getcorpcheckinoption", nil)
}

// GetCheckinOption 获取员工打卡规则
func (s *Service) GetCheckinOption(ctx context.Context, req *checkin.GetCheckinOptionRequest) (*checkin.GetCheckinOptionResponse, error) {
	return client.PostAndUnmarshal[checkin.GetCheckinOptionResponse](s.client, ctx, "/cgi-bin/checkin/getcheckinoption", req)
}

// GetCheckinScheduleList 获取打卡人员排班信息
func (s *Service) GetCheckinScheduleList(ctx context.Context, req *checkin.GetCheckinScheduleListRequest) (*checkin.GetCheckinScheduleListResponse, error) {
	return client.PostAndUnmarshal[checkin.GetCheckinScheduleListResponse](s.client, ctx, "/cgi-bin/checkin/getcheckinschedulist", req)
}

// GetCheckinDayData 获取打卡日报数据
func (s *Service) GetCheckinDayData(ctx context.Context, req *checkin.GetCheckinDayDataRequest) (*checkin.GetCheckinDayDataResponse, error) {
	return client.PostAndUnmarshal[checkin.GetCheckinDayDataResponse](s.client, ctx, "/cgi-bin/checkin/getcheckin_daydata", req)
}

// GetCheckinMonthData 获取打卡月报数据
func (s *Service) GetCheckinMonthData(ctx context.Context, req *checkin.GetCheckinMonthDataRequest) (*checkin.GetCheckinMonthDataResponse, error) {
	return client.PostAndUnmarshal[checkin.GetCheckinMonthDataResponse](s.client, ctx, "/cgi-bin/checkin/getcheckin_monthdata", req)
}

// GetCheckinData 获取打卡记录数据
func (s *Service) GetCheckinData(ctx context.Context, req *checkin.GetCheckinDataRequest) (*checkin.GetCheckinDataResponse, error) {
	return client.PostAndUnmarshal[checkin.GetCheckinDataResponse](s.client, ctx, "/cgi-bin/checkin/getcheckindata", req)
}

// GetHardwareCheckinData 获取设备打卡数据
func (s *Service) GetHardwareCheckinData(ctx context.Context, req *checkin.GetHardwareCheckinDataRequest) (*checkin.GetHardwareCheckinDataResponse, error) {
	return client.PostAndUnmarshal[checkin.GetHardwareCheckinDataResponse](s.client, ctx, "/cgi-bin/hardware/get_hardware_checkin_data", req)
}

