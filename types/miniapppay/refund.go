package miniapppay

import "github.com/mrseanchow/wecom-core/types/common"

// RefundRequest 申请退款请�?
type RefundRequest struct {
	MchID        string        `json:"mchid"`
	AppID        string        `json:"appid"`
	OutTradeNo   string        `json:"out_trade_no"`
	OutRefundNo  string        `json:"out_refund_no"`
	Reason       string        `json:"reason,omitempty"`
	FundsAccount string        `json:"funds_account,omitempty"`
	Amount       *RefundAmount `json:"amount"`
}

// RefundAmount 退款金额信�?
type RefundAmount struct {
	Refund   int    `json:"refund"`
	Total    int    `json:"total"`
	Currency string `json:"currency"`
}

// RefundResponse 申请退款响�?
type RefundResponse struct {
	common.Response
	OutRefundNo     string                `json:"out_refund_no"`
	Amount          *RefundResponseAmount `json:"amount,omitempty"`
	PromotionDetail []*PromotionDetail    `json:"promotion_detail,omitempty"`
}

// RefundResponseAmount 退款响应金额信�?
type RefundResponseAmount struct {
	Refund         int    `json:"refund"`
	PayerRefund    int    `json:"payer_refund"`
	DiscountRefund int    `json:"discount_refund"`
	Currency       string `json:"currency,omitempty"`
}

// GetRefundRequest 查询退款请�?
type GetRefundRequest struct {
	MchID       string `json:"mchid"`
	OutRefundNo string `json:"out_refund_no"`
}

// GetRefundResponse 查询退款响�?
type GetRefundResponse struct {
	common.Response
	RefundID            string                `json:"refund_id,omitempty"`
	OutRefundNo         string                `json:"out_refund_no"`
	TransactionID       string                `json:"transaction_id,omitempty"`
	OutTradeNo          string                `json:"out_trade_no"`
	Channel             string                `json:"channel,omitempty"`
	UserReceivedAccount string                `json:"user_received_account,omitempty"`
	SuccessTime         string                `json:"success_time,omitempty"`
	CreateTime          string                `json:"create_time,omitempty"`
	Status              string                `json:"status"`
	Amount              *RefundResponseAmount `json:"amount"`
	PromotionDetail     []*PromotionDetail    `json:"promotion_detail,omitempty"`
}

// RefundNotification 退款通知回调结构
type RefundNotification struct {
	MchID               string                    `json:"mchid"`
	OutTradeNo          string                    `json:"out_trade_no"`
	RefundID            string                    `json:"refund_id"`
	OutRefundNo         string                    `json:"out_refund_no"`
	RefundStatus        string                    `json:"refund_status"`
	SuccessTime         string                    `json:"success_time,omitempty"`
	UserReceivedAccount string                    `json:"user_received_account,omitempty"`
	Amount              *RefundNotificationAmount `json:"amount"`
}

// RefundNotificationAmount 退款通知金额信息
type RefundNotificationAmount struct {
	Total       int `json:"total"`
	Refund      int `json:"refund"`
	PayerTotal  int `json:"payer_total"`
	PayerRefund int `json:"payer_refund"`
}

