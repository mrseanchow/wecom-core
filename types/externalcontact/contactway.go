package externalcontact

import "github.com/mrseanchow/wecom-core/types/common"

// Conclusions 结束语定�?
type Conclusions struct {
	Text        *ConclusionText        `json:"text,omitempty"`
	Image       *ConclusionImage       `json:"image,omitempty"`
	Link        *ConclusionLink        `json:"link,omitempty"`
	Miniprogram *ConclusionMiniprogram `json:"miniprogram,omitempty"`
}

// ConclusionText 文本消息
type ConclusionText struct {
	Content string `json:"content"`
}

// ConclusionImage 图片消息
type ConclusionImage struct {
	MediaID string `json:"media_id,omitempty"`
	PicURL  string `json:"pic_url,omitempty"`
}

// ConclusionLink 图文消息
type ConclusionLink struct {
	Title  string `json:"title"`
	PicURL string `json:"picurl"`
	Desc   string `json:"desc"`
	URL    string `json:"url"`
}

// ConclusionMiniprogram 小程序消�?
type ConclusionMiniprogram struct {
	Title      string `json:"title"`
	PicMediaID string `json:"pic_media_id"`
	AppID      string `json:"appid"`
	Page       string `json:"page"`
}

// AddContactWayRequest 配置客户联系「联系我」方式请�?
type AddContactWayRequest struct {
	Type          int          `json:"type"`
	Scene         int          `json:"scene"`
	Style         int          `json:"style,omitempty"`
	Remark        string       `json:"remark,omitempty"`
	SkipVerify    bool         `json:"skip_verify,omitempty"`
	State         string       `json:"state,omitempty"`
	User          []string     `json:"user,omitempty"`
	Party         []int        `json:"party,omitempty"`
	IsTemp        bool         `json:"is_temp,omitempty"`
	ExpiresIn     int          `json:"expires_in,omitempty"`
	ChatExpiresIn int          `json:"chat_expires_in,omitempty"`
	UnionID       string       `json:"unionid,omitempty"`
	IsExclusive   bool         `json:"is_exclusive,omitempty"`
	MarkSource    bool         `json:"mark_source,omitempty"`
	Conclusions   *Conclusions `json:"conclusions,omitempty"`
}

// AddContactWayResponse 配置客户联系「联系我」方式响�?
type AddContactWayResponse struct {
	common.Response
	ConfigID string `json:"config_id"`
	QRCode   string `json:"qr_code,omitempty"`
}

// GetContactWayRequest 获取企业已配置的「联系我」方式请�?
type GetContactWayRequest struct {
	ConfigID string `json:"config_id"`
}

// ContactWay 联系方式配置详情
type ContactWay struct {
	ConfigID      string       `json:"config_id"`
	Type          int          `json:"type"`
	Scene         int          `json:"scene"`
	Style         int          `json:"style,omitempty"`
	Remark        string       `json:"remark,omitempty"`
	SkipVerify    bool         `json:"skip_verify,omitempty"`
	State         string       `json:"state,omitempty"`
	QRCode        string       `json:"qr_code,omitempty"`
	User          []string     `json:"user,omitempty"`
	Party         []int        `json:"party,omitempty"`
	IsTemp        bool         `json:"is_temp,omitempty"`
	ExpiresIn     int          `json:"expires_in,omitempty"`
	ChatExpiresIn int          `json:"chat_expires_in,omitempty"`
	UnionID       string       `json:"unionid,omitempty"`
	MarkSource    bool         `json:"mark_source,omitempty"`
	Conclusions   *Conclusions `json:"conclusions,omitempty"`
}

// GetContactWayResponse 获取企业已配置的「联系我」方式响�?
type GetContactWayResponse struct {
	common.Response
	ContactWay ContactWay `json:"contact_way"`
}

// ListContactWayRequest 获取企业已配置的「联系我」列表请�?
type ListContactWayRequest struct {
	StartTime int64  `json:"start_time,omitempty"`
	EndTime   int64  `json:"end_time,omitempty"`
	Cursor    string `json:"cursor,omitempty"`
	Limit     int    `json:"limit,omitempty"`
}

// ContactWayItem 联系方式配置�?
type ContactWayItem struct {
	ConfigID string `json:"config_id"`
}

// ListContactWayResponse 获取企业已配置的「联系我」列表响�?
type ListContactWayResponse struct {
	common.Response
	ContactWay []ContactWayItem `json:"contact_way"`
	NextCursor string           `json:"next_cursor,omitempty"`
}

// UpdateContactWayRequest 更新企业已配置的「联系我」方式请�?
type UpdateContactWayRequest struct {
	ConfigID      string       `json:"config_id"`
	Remark        string       `json:"remark,omitempty"`
	SkipVerify    bool         `json:"skip_verify,omitempty"`
	Style         int          `json:"style,omitempty"`
	State         string       `json:"state,omitempty"`
	User          []string     `json:"user,omitempty"`
	Party         []int        `json:"party,omitempty"`
	ExpiresIn     int          `json:"expires_in,omitempty"`
	ChatExpiresIn int          `json:"chat_expires_in,omitempty"`
	UnionID       string       `json:"unionid,omitempty"`
	MarkSource    bool         `json:"mark_source,omitempty"`
	Conclusions   *Conclusions `json:"conclusions,omitempty"`
}

// DeleteContactWayRequest 删除企业已配置的「联系我」方式请�?
type DeleteContactWayRequest struct {
	ConfigID string `json:"config_id"`
}

// CloseTempChatRequest 结束临时会话请求
type CloseTempChatRequest struct {
	UserID         string `json:"userid"`
	ExternalUserID string `json:"external_userid"`
}

