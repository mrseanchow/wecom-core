package miniapppay

// GetBillRequest 交易账单申请请求
type GetBillRequest struct {
	BillDate string `json:"bill_date"`
	MchID    string `json:"mchid"`
	BillType string `json:"bill_type,omitempty"`
	TarType  string `json:"tar_type,omitempty"`
}

// GetBillResponse 交易账单申请响应
type GetBillResponse struct {
	DownloadURL string `json:"download_url"`
	HashType    string `json:"hash_type"`
	HashValue   string `json:"hash_value"`
	Auth        string `json:"auth,omitempty"`
}
