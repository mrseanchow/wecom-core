package payment

import "encoding/xml"

type SendTransferRequest struct {
	XMLName        xml.Name `xml:"xml"`
	AppID          string   `xml:"appid"`
	MchID          string   `xml:"mch_id"`
	DeviceInfo     string   `xml:"device_info,omitempty"`
	NonceStr       string   `xml:"nonce_str"`
	Sign           string   `xml:"sign"`
	PartnerTradeNo string   `xml:"partner_trade_no"`
	OpenID         string   `xml:"openid"`
	CheckName      string   `xml:"check_name"`
	ReUserName     string   `xml:"re_user_name,omitempty"`
	Amount         int      `xml:"amount"`
	Desc           string   `xml:"desc"`
	SpbillCreateIP string   `xml:"spbill_create_ip"`
	WorkWXSign     string   `xml:"workwx_sign"`
	WWMsgType      string   `xml:"ww_msg_type"`
	ApprovalNumber string   `xml:"approval_number,omitempty"`
	ApprovalType   int      `xml:"approval_type,omitempty"`
	ActName        string   `xml:"act_name"`
	AgentID        int      `xml:"agentid,omitempty"`
}

type SendTransferResponse struct {
	ReturnCode     string `xml:"return_code"`
	ReturnMsg      string `xml:"return_msg"`
	AppID          string `xml:"appid,omitempty"`
	MchID          string `xml:"mch_id,omitempty"`
	DeviceInfo     string `xml:"device_info,omitempty"`
	NonceStr       string `xml:"nonce_str,omitempty"`
	ResultCode     string `xml:"result_code,omitempty"`
	ErrCode        string `xml:"err_code,omitempty"`
	ErrCodeDes     string `xml:"err_code_des,omitempty"`
	PartnerTradeNo string `xml:"partner_trade_no,omitempty"`
	PaymentNo      string `xml:"payment_no,omitempty"`
	PaymentTime    string `xml:"payment_time,omitempty"`
}

type QueryTransferRequest struct {
	XMLName        xml.Name `xml:"xml"`
	NonceStr       string   `xml:"nonce_str"`
	Sign           string   `xml:"sign"`
	PartnerTradeNo string   `xml:"partner_trade_no"`
	MchID          string   `xml:"mch_id"`
	AppID          string   `xml:"appid"`
}

type QueryTransferResponse struct {
	ReturnCode     string `xml:"return_code"`
	ReturnMsg      string `xml:"return_msg"`
	ResultCode     string `xml:"result_code,omitempty"`
	ErrCode        string `xml:"err_code,omitempty"`
	ErrCodeDes     string `xml:"err_code_des,omitempty"`
	MchID          string `xml:"mch_id,omitempty"`
	AppID          string `xml:"appid,omitempty"`
	DetailID       string `xml:"detail_id,omitempty"`
	PartnerTradeNo string `xml:"partner_trade_no,omitempty"`
	Status         string `xml:"status,omitempty"`
	Reason         string `xml:"reason,omitempty"`
	OpenID         string `xml:"openid,omitempty"`
	TransferName   string `xml:"transfer_name,omitempty"`
	PaymentAmount  int    `xml:"payment_amount,omitempty"`
	TransferTime   string `xml:"transfer_time,omitempty"`
	Desc           string `xml:"desc,omitempty"`
}

