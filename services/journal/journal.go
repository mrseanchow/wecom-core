package journal

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/journal"
)
const (
	AJournalDownloadWedriveFileURL = "/cgi-bin/oa/journal/download_wedrive_file"
	AJournalGetRecordDetailURL = "/cgi-bin/oa/journal/get_record_detail"
	AJournalGetRecordListURL = "/cgi-bin/oa/journal/get_record_list"
	AJournalGetStatListURL = "/cgi-bin/oa/journal/get_stat_list"
)

type Service struct {
	client *client.Client
}

func NewService(client *client.Client) *Service {
	return &Service{client: client}
}

func (s *Service) DownloadWedriveFile(ctx context.Context, req *journal.DownloadWedriveFileRequest) (*journal.DownloadWedriveFileResponse, error) {
	return client.PostAndUnmarshal[journal.DownloadWedriveFileResponse](s.client, ctx, AJournalDownloadWedriveFileURL, req)
}

func (s *Service) GetRecordList(ctx context.Context, req *journal.GetRecordListRequest) (*journal.GetRecordListResponse, error) {
	return client.PostAndUnmarshal[journal.GetRecordListResponse](s.client, ctx, AJournalGetRecordListURL, req)
}

func (s *Service) GetStatList(ctx context.Context, req *journal.GetStatListRequest) (*journal.GetStatListResponse, error) {
	return client.PostAndUnmarshal[journal.GetStatListResponse](s.client, ctx, AJournalGetStatListURL, req)
}

func (s *Service) GetRecordDetail(ctx context.Context, req *journal.GetRecordDetailRequest) (*journal.GetRecordDetailResponse, error) {
	return client.PostAndUnmarshal[journal.GetRecordDetailResponse](s.client, ctx, AJournalGetRecordDetailURL, req)
}
