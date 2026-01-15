package journal

type DownloadWedriveFileResponse struct {
	DownloadURL string `json:"download_url"`
	CookieName  string `json:"cookie_name"`
	CookieValue string `json:"cookie_value"`
}

type GetRecordListResponse struct {
	JournalUUIDList []string `json:"journaluuid_list"`
	NextCursor      uint32   `json:"next_cursor"`
	EndFlag         uint32   `json:"endflag"`
}

type GetStatListResponse struct {
	StatList []StatInfo `json:"stat_list"`
}

type StatInfo struct {
	TemplateID     string           `json:"template_id"`
	TemplateName   string           `json:"template_name"`
	ReportRange    *RangeInfo       `json:"report_range"`
	WhiteRange     *RangeInfo       `json:"white_range"`
	Receivers      *ReceiverInfo    `json:"receivers"`
	CycleBeginTime uint64           `json:"cycle_begin_time"`
	CycleEndTime   uint64           `json:"cycle_end_time"`
	StatBeginTime  uint64           `json:"stat_begin_time"`
	StatEndTime    uint64           `json:"stat_end_time"`
	ReportList     []UserReportInfo `json:"report_list"`
	UnreportList   []UserReportInfo `json:"unreport_list"`
	ReportType     uint32           `json:"report_type"`
}

type RangeInfo struct {
	UserList  []UserIDInfo  `json:"user_list"`
	PartyList []PartyIDInfo `json:"party_list"`
	TagList   []TagIDInfo   `json:"tag_list"`
}

type UserIDInfo struct {
	UserID string `json:"userid"`
}

type PartyIDInfo struct {
	OpenPartyID string `json:"open_partyid"`
}

type TagIDInfo struct {
	OpenTagID string `json:"open_tagid"`
}

type ReceiverInfo struct {
	UserList   []UserIDInfo `json:"user_list"`
	TagList    []TagIDInfo  `json:"tag_list"`
	LeaderList []LeaderInfo `json:"leader_list"`
}

type LeaderInfo struct {
	Level uint64 `json:"level"`
}

type UserReportInfo struct {
	User     UserIDInfo   `json:"user"`
	ItemList []ReportItem `json:"itemlist"`
}

type ReportItem struct {
	JournalUUID string `json:"journaluuid"`
	ReportTime  uint32 `json:"reporttime"`
	Flag        uint32 `json:"flag"`
}

type GetRecordDetailResponse struct {
	Info JournalInfo `json:"info"`
}

type JournalInfo struct {
	JournalUUID     string     `json:"journal_uuid"`
	TemplateName    string     `json:"template_name"`
	TemplateID      string     `json:"template_id"`
	ReportTime      int32      `json:"report_time"`
	Submitter       UserInfo   `json:"submitter"`
	Receivers       []UserInfo `json:"receivers"`
	ReadedReceivers []UserInfo `json:"readed_receivers"`
	ApplyData       ApplyData  `json:"apply_data"`
	SysJournalData  string     `json:"sys_journal_data"`
	Comments        []Comment  `json:"comments"`
}

type UserInfo struct {
	UserID string `json:"userid"`
}

type ApplyData struct {
	Contents []Content `json:"contents"`
}

type Content struct {
	Control string       `json:"control"`
	ID      string       `json:"id"`
	Title   []TextInfo   `json:"title"`
	Value   ContentValue `json:"value"`
}

type TextInfo struct {
	Text string `json:"text"`
}

type ContentValue struct {
	Text            string            `json:"text"`
	Tips            []string          `json:"tips"`
	Members         []UserIDInfo      `json:"members"`
	Departments     []DepartmentInfo  `json:"departments"`
	Files           []FileInfo        `json:"files"`
	Children        []ChildContent    `json:"children"`
	StatField       []interface{}     `json:"stat_field"`
	SumField        []interface{}     `json:"sum_field"`
	RelatedApproval []interface{}     `json:"related_approval"`
	Students        []StudentInfo     `json:"students"`
	Classes         []ClassInfo       `json:"classes"`
	NewNumber       string            `json:"new_number"`
	NewMoney        string            `json:"new_money"`
	Date            *DateInfo         `json:"date"`
	Selector        *SelectorInfo     `json:"selector"`
	DateRange       *DateRangeInfo    `json:"date_range"`
	Location        *LocationInfo     `json:"location"`
	Formula         *FormulaInfo      `json:"formula"`
	Docs            []DocInfo         `json:"docs"`
	WedriveFiles    []WedriveFileInfo `json:"wedrive_files"`
}

type DepartmentInfo struct {
	OpenAPIID string `json:"openapi_id"`
}

type FileInfo struct {
	FileID string `json:"file_id"`
}

type ChildContent struct {
	List []Content `json:"list"`
}

type StudentInfo struct {
	Name string `json:"name"`
}

type ClassInfo struct {
	Name string `json:"name"`
}

type DateInfo struct {
	Type       string `json:"type"`
	STimestamp string `json:"s_timestamp"`
}

type SelectorInfo struct {
	Type    string           `json:"type"`
	Options []SelectorOption `json:"options"`
}

type SelectorOption struct {
	Key   string     `json:"key"`
	Value []TextInfo `json:"value"`
}

type DateRangeInfo struct {
	Type        string `json:"type"`
	NewBegin    uint64 `json:"new_begin"`
	NewEnd      uint64 `json:"new_end"`
	NewDuration uint64 `json:"new_duration"`
}

type LocationInfo struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Title     string `json:"title"`
	Address   string `json:"address"`
	Time      uint64 `json:"time"`
}

type FormulaInfo struct {
	Value string `json:"value"`
}

type DocInfo struct {
	DocID  string `json:"docid"`
	DocURL string `json:"doc_url"`
}

type WedriveFileInfo struct {
	FileID string `json:"fileid"`
}

type Comment struct {
	CommentID       uint64   `json:"commentid"`
	ToCommentID     uint64   `json:"tocommentid"`
	CommentUserInfo UserInfo `json:"comment_userinfo"`
	Content         string   `json:"content"`
	CommentTime     uint32   `json:"comment_time"`
}
