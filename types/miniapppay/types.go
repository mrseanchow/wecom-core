package miniapppay

// Amount 金额信息
type Amount struct {
	Total          int    `json:"total"`
	PayerTotal     int    `json:"payer_total,omitempty"`
	Currency       string `json:"currency,omitempty"`
	PayerCurrency  string `json:"payer_currency,omitempty"`
	Refund         int    `json:"refund,omitempty"`
	PayerRefund    int    `json:"payer_refund,omitempty"`
	DiscountRefund int    `json:"discount_refund,omitempty"`
}

// Payer 支付者信息
type Payer struct {
	OpenID string `json:"openid"`
}

// SceneInfo 场景信息
type SceneInfo struct {
	PayerClientIP string     `json:"payer_client_ip"`
	DeviceID      string     `json:"device_id,omitempty"`
	StoreInfo     *StoreInfo `json:"store_info,omitempty"`
}

// StoreInfo 门店信息
type StoreInfo struct {
	ID       string `json:"id"`
	Name     string `json:"name,omitempty"`
	AreaCode string `json:"area_code,omitempty"`
	Address  string `json:"address,omitempty"`
}

// GoodsDetail 商品详情
type GoodsDetail struct {
	MerchantGoodsID  string `json:"merchant_goods_id"`
	WeChatPayGoodsID string `json:"wechatpay_goods_id,omitempty"`
	GoodsName        string `json:"goods_name,omitempty"`
	Quantity         int    `json:"quantity"`
	UnitPrice        int    `json:"unit_price"`
	DiscountAmount   int    `json:"discount_amount,omitempty"`
	GoodsRemark      string `json:"goods_remark,omitempty"`
}

// PromotionDetail 优惠详情
type PromotionDetail struct {
	PromotionID         string            `json:"promotion_id,omitempty"`
	Scope               string            `json:"scope,omitempty"`
	Type                string            `json:"type,omitempty"`
	Amount              int               `json:"amount,omitempty"`
	RefundAmount        int               `json:"refund_amount,omitempty"`
	WeChatPayContribute int               `json:"wechatpay_contribute,omitempty"`
	MerchantContribute  int               `json:"merchant_contribute,omitempty"`
	OtherContribute     int               `json:"other_contribute,omitempty"`
	Currency            string            `json:"currency,omitempty"`
	GoodsDetail         []GoodsDetailItem `json:"goods_detail,omitempty"`
}

// GoodsDetailItem 商品优惠详情
type GoodsDetailItem struct {
	GoodsID        string `json:"goods_id"`
	Quantity       int    `json:"quantity"`
	UnitPrice      int    `json:"unit_price"`
	DiscountAmount int    `json:"discount_amount"`
}

// Resource 加密资源
type Resource struct {
	Algorithm      string `json:"algorithm"`
	Ciphertext     string `json:"ciphertext"`
	Nonce          string `json:"nonce"`
	AssociatedData string `json:"associated_data,omitempty"`
}

// NotifyBase 通知基础结构
type NotifyBase struct {
	NotifyType   string    `json:"notify_type"`
	ID           string    `json:"id"`
	CreateTime   string    `json:"create_time"`
	EventType    string    `json:"event_type"`
	ResourceType string    `json:"resource_type"`
	Resource     *Resource `json:"resource"`
	Summary      string    `json:"summary,omitempty"`
}
