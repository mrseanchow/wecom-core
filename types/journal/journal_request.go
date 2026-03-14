package journal

type DownloadWedriveFileRequest struct {
	JournalUUID string `json:"journaluuid"`
	FileID      string `json:"fileid"`
}

type GetRecordListRequest struct {
	StartTime uint32                    `json:"starttime"`
	EndTime   uint32                    `json:"endtime"`
	Cursor    uint32                    `json:"cursor"`
	Limit     uint32                    `json:"limit"`
	Filters   []GetRecordListFilterItem `json:"filters"`
}

type GetRecordListFilterItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type GetStatListRequest struct {
	TemplateID string `json:"template_id"`
	StartTime  uint64 `json:"starttime"`
	EndTime    uint64 `json:"endtime"`
}

type GetRecordDetailRequest struct {
	JournalUUID string `json:"journaluuid"`
}
