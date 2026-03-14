package payment

import "encoding/xml"

type SendRedPackRequest struct {
	XMLName             xml.Name `xml:"xml"`
	NonceStr            string   `xml:"nonce_str"`
	Sign                string   `xml:"sign"`
	MchBillno           string   `xml:"mch_billno"`
	MchID               string   `xml:"mch_id"`
	WxAppID             string   `xml:"wxappid"`
	SenderName          string   `xml:"sender_name,omitempty"`
	AgentID             int      `xml:"agentid,omitempty"`
	SenderHeaderMediaID string   `xml:"sender_header_media_id,omitempty"`
	ReOpenID            string   `xml:"re_openid"`
	TotalAmount         int      `xml:"total_amount"`
	Wishing             string   `xml:"wishing"`
	ActName             string   `xml:"act_name"`
	Remark              string   `xml:"remark"`
	SceneID             string   `xml:"scene_id,omitempty"`
	WorkWXSign          string   `xml:"workwx_sign"`
}

type SendRedPackResponse struct {
	ReturnCode          string `xml:"return_code"`
	ReturnMsg           string `xml:"return_msg"`
	Sign                string `xml:"sign,omitempty"`
	ResultCode          string `xml:"result_code,omitempty"`
	ErrCode             string `xml:"err_code,omitempty"`
	ErrCodeDes          string `xml:"err_code_des,omitempty"`
	MchBillno           string `xml:"mch_billno,omitempty"`
	MchID               string `xml:"mch_id,omitempty"`
	WxAppID             string `xml:"wxappid,omitempty"`
	ReOpenID            string `xml:"re_openid,omitempty"`
	TotalAmount         int    `xml:"total_amount,omitempty"`
	SendListID          string `xml:"send_listid,omitempty"`
	SenderName          string `xml:"sender_name,omitempty"`
	SenderHeaderMediaID string `xml:"sender_header_media_id,omitempty"`
}

type QueryRedPackRequest struct {
	XMLName   xml.Name `xml:"xml"`
	NonceStr  string   `xml:"nonce_str"`
	Sign      string   `xml:"sign"`
	MchBillno string   `xml:"mch_billno"`
	MchID     string   `xml:"mch_id"`
	AppID     string   `xml:"appid"`
}

type QueryRedPackResponse struct {
	ReturnCode          string `xml:"return_code"`
	ReturnMsg           string `xml:"return_msg"`
	Sign                string `xml:"sign,omitempty"`
	ResultCode          string `xml:"result_code,omitempty"`
	ErrCode             string `xml:"err_code,omitempty"`
	ErrCodeDes          string `xml:"err_code_des,omitempty"`
	MchBillno           string `xml:"mch_billno,omitempty"`
	MchID               string `xml:"mch_id,omitempty"`
	DetailID            string `xml:"detail_id,omitempty"`
	Status              string `xml:"status,omitempty"`
	SendType            string `xml:"send_type,omitempty"`
	TotalAmount         int    `xml:"total_amount,omitempty"`
	Reason              string `xml:"reason,omitempty"`
	SendTime            string `xml:"send_time,omitempty"`
	RefundTime          string `xml:"refund_time,omitempty"`
	RefundAmount        int    `xml:"refund_amount,omitempty"`
	Wishing             string `xml:"wishing,omitempty"`
	Remark              string `xml:"remark,omitempty"`
	ActName             string `xml:"act_name,omitempty"`
	OpenID              string `xml:"openid,omitempty"`
	Amount              int    `xml:"amount,omitempty"`
	RcvTime             string `xml:"rcv_time,omitempty"`
	SenderName          string `xml:"sender_name,omitempty"`
	SenderHeaderMediaID string `xml:"sender_header_media_id,omitempty"`
}

