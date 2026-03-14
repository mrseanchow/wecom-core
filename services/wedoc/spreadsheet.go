package wedoc

import (
	"context"

	"github.com/mrseanchow/wecom-core/internal/client"
	"github.com/mrseanchow/wecom-core/types/wedoc"
)

const (
	getSheetPropertiesURL     = "/cgi-bin/wedoc/spreadsheet/get_sheet_properties"
	getSheetRangeDataURL      = "/cgi-bin/wedoc/spreadsheet/get_sheet_range_data"
	batchUpdateSpreadsheetURL = "/cgi-bin/wedoc/spreadsheet/batch_update"
)

// GetSheetProperties 获取表格行列信息
// 该接口用于获取在线表格的工作表、行数、列数等
func (s *Service) GetSheetProperties(ctx context.Context, req *wedoc.GetSheetPropertiesRequest) (*wedoc.GetSheetPropertiesResponse, error) {
	return client.PostAndUnmarshal[wedoc.GetSheetPropertiesResponse](s.client, ctx, getSheetPropertiesURL, req)
}

// GetSheetRangeData 获取表格数据
// 本接口用于获取指定范围内的在线表格信�?
func (s *Service) GetSheetRangeData(ctx context.Context, req *wedoc.GetSheetRangeDataRequest) (*wedoc.GetSheetRangeDataResponse, error) {
	return client.PostAndUnmarshal[wedoc.GetSheetRangeDataResponse](s.client, ctx, getSheetRangeDataURL, req)
}

// BatchUpdateSpreadsheet 批量编辑表格内容
// 该接口可以对一个在线表格批量执行多个更新操�?
func (s *Service) BatchUpdateSpreadsheet(ctx context.Context, req *wedoc.BatchUpdateSpreadsheetRequest) (*wedoc.BatchUpdateSpreadsheetResponse, error) {
	return client.PostAndUnmarshal[wedoc.BatchUpdateSpreadsheetResponse](s.client, ctx, batchUpdateSpreadsheetURL, req)
}

