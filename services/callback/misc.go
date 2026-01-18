package callback

type BatchJobEvent struct {
	BaseEvent
	BatchJob struct {
		JobId   string `xml:"JobId"`
		JobType string `xml:"JobType"`
		ErrCode int    `xml:"ErrCode"`
		ErrMsg  string `xml:"ErrMsg"`
	} `xml:"BatchJob"`
}

type SecurityEvent struct {
	BaseEvent
}

type KFAccountAuthChangeEvent struct {
	BaseEvent
	AuthAddOpenKfId []string `xml:"AuthAddOpenKfId"`
	AuthDelOpenKfId []string `xml:"AuthDelOpenKfId"`
}

func ParseBatchJobResult(data []byte) (*BatchJobEvent, error) {
	return Parse[BatchJobEvent](data)
}

func ParseSecurity(data []byte) (*SecurityEvent, error) {
	return Parse[SecurityEvent](data)
}

func ParseKFAccountAuthChange(data []byte) (*KFAccountAuthChangeEvent, error) {
	return Parse[KFAccountAuthChangeEvent](data)
}
