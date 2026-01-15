package miniapppay

// GetSignRequest 获取支付签名请求
type GetSignRequest struct {
	AppID     string `json:"appid"`
	PrepayID  string `json:"prepay_id"`
	SignType  string `json:"sign_type,omitempty"`
	Nonce     string `json:"nonce"`
	Timestamp int64  `json:"timestamp"`
}

// GetSignResponse 获取支付签名响应
type GetSignResponse struct {
	PaySign string `json:"pay_sign"`
}
