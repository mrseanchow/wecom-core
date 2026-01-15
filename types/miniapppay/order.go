package miniapppay

// CreateOrderRequest 小程序下单请求
type CreateOrderRequest struct {
	AppID       string       `json:"appid"`
	MchID       string       `json:"mchid"`
	OutTradeNo  string       `json:"out_trade_no"`
	Description string       `json:"description"`
	SceneKey    string       `json:"scenekey,omitempty"`
	Amount      *Amount      `json:"amount"`
	Payer       *Payer       `json:"payer"`
	TimeExpire  string       `json:"time_expire,omitempty"`
	Attach      string       `json:"attach,omitempty"`
	GoodsTag    string       `json:"goods_tag,omitempty"`
	SceneInfo   *SceneInfo   `json:"scene_info,omitempty"`
	Detail      *OrderDetail `json:"detail,omitempty"`
}

// OrderDetail 订单详情
type OrderDetail struct {
	CostPrice   int           `json:"cost_price,omitempty"`
	InvoiceID   string        `json:"invoice_id,omitempty"`
	GoodsDetail []GoodsDetail `json:"goods_detail,omitempty"`
}

// CreateOrderResponse 小程序下单响应
type CreateOrderResponse struct {
	PrepayID string `json:"prepay_id"`
}

// GetOrderRequest 查询订单请求
type GetOrderRequest struct {
	MchID      string `json:"mchid"`
	OutTradeNo string `json:"out_trade_no"`
}

// GetOrderResponse 查询订单响应
type GetOrderResponse struct {
	MchID           string             `json:"mchid"`
	OutTradeNo      string             `json:"out_trade_no"`
	TradeType       string             `json:"trade_type,omitempty"`
	TradeState      string             `json:"trade_state"`
	TradeStateDesc  string             `json:"trade_state_desc"`
	Payer           *Payer             `json:"payer,omitempty"`
	TransactionID   string             `json:"transaction_id"`
	BankType        string             `json:"bank_type,omitempty"`
	Attach          string             `json:"attach,omitempty"`
	SuccessTime     string             `json:"success_time,omitempty"`
	Amount          *Amount            `json:"amount,omitempty"`
	SceneInfo       *OrderSceneInfo    `json:"scene_info,omitempty"`
	PromotionDetail []*PromotionDetail `json:"promotion_detail,omitempty"`
}

// OrderSceneInfo 订单场景信息
type OrderSceneInfo struct {
	DeviceID string `json:"device_id,omitempty"`
}

// CloseOrderRequest 关闭订单请求
type CloseOrderRequest struct {
	MchID      string `json:"mchid"`
	OutTradeNo string `json:"out_trade_no"`
}

// PaymentNotification 支付通知回调结构
type PaymentNotification struct {
	AppID          string  `json:"appid"`
	MchID          string  `json:"mchid"`
	OutTradeNo     string  `json:"out_trade_no"`
	TradeState     string  `json:"trade_state"`
	TradeStateDesc string  `json:"trade_state_desc"`
	TradeType      string  `json:"trade_type,omitempty"`
	Attach         string  `json:"attach,omitempty"`
	SuccessTime    string  `json:"success_time,omitempty"`
	Payer          *Payer  `json:"payer"`
	Amount         *Amount `json:"amount"`
}
