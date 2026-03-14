package externalcontact

import "github.com/mrseanchow/wecom-core/types/common"

// TextContent 文本消息内容
type TextContent struct {
	Content string `json:"content,omitempty"`
}

// ImageAttachment 图片附件
type ImageAttachment struct {
	MediaID string `json:"media_id,omitempty"`
	PicURL  string `json:"pic_url,omitempty"`
}

// LinkAttachment 链接附件
type LinkAttachment struct {
	Title  string `json:"title"`
	PicURL string `json:"picurl,omitempty"`
	Desc   string `json:"desc,omitempty"`
	URL    string `json:"url"`
}

// MiniprogramAttachment 小程序附�?
type MiniprogramAttachment struct {
	Title      string `json:"title"`
	PicMediaID string `json:"pic_media_id"`
	AppID      string `json:"appid"`
	Page       string `json:"page"`
}

// VideoAttachment 视频附件
type VideoAttachment struct {
	MediaID string `json:"media_id"`
}

// FileAttachment 文件附件
type FileAttachment struct {
	MediaID string `json:"media_id"`
}

// Attachment 附件
type Attachment struct {
	MsgType     string                 `json:"msgtype"`
	Image       *ImageAttachment       `json:"image,omitempty"`
	Link        *LinkAttachment        `json:"link,omitempty"`
	Miniprogram *MiniprogramAttachment `json:"miniprogram,omitempty"`
	Video       *VideoAttachment       `json:"video,omitempty"`
	File        *FileAttachment        `json:"file,omitempty"`
}

// TagFilterGroup 标签过滤�?
type TagFilterGroup struct {
	TagList []string `json:"tag_list"`
}

// TagFilter 标签过滤�?
type TagFilter struct {
	GroupList []TagFilterGroup `json:"group_list,omitempty"`
}

// AddMsgTemplateRequest 创建企业群发请求
type AddMsgTemplateRequest struct {
	ChatType       string       `json:"chat_type,omitempty"`
	ExternalUserID []string     `json:"external_userid,omitempty"`
	ChatIDList     []string     `json:"chat_id_list,omitempty"`
	TagFilter      *TagFilter   `json:"tag_filter,omitempty"`
	Sender         string       `json:"sender,omitempty"`
	AllowSelect    bool         `json:"allow_select,omitempty"`
	Text           *TextContent `json:"text,omitempty"`
	Attachments    []Attachment `json:"attachments,omitempty"`
}

// AddMsgTemplateResponse 创建企业群发响应
type AddMsgTemplateResponse struct {
	common.Response
	FailList []string `json:"fail_list,omitempty"`
	MsgID    string   `json:"msgid"`
}

// GetGroupMsgListV2Request 获取群发记录列表请求
type GetGroupMsgListV2Request struct {
	ChatType   string `json:"chat_type"`
	StartTime  int64  `json:"start_time"`
	EndTime    int64  `json:"end_time"`
	Creator    string `json:"creator,omitempty"`
	FilterType int    `json:"filter_type,omitempty"`
	Limit      int    `json:"limit,omitempty"`
	Cursor     string `json:"cursor,omitempty"`
}

// GroupMsg 群发消息
type GroupMsg struct {
	MsgID       string       `json:"msgid"`
	Creator     string       `json:"creator,omitempty"`
	CreateTime  int64        `json:"create_time"`
	CreateType  int          `json:"create_type"`
	Text        *TextContent `json:"text,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

// GetGroupMsgListV2Response 获取群发记录列表响应
type GetGroupMsgListV2Response struct {
	common.Response
	NextCursor   string     `json:"next_cursor,omitempty"`
	GroupMsgList []GroupMsg `json:"group_msg_list"`
}

// GetGroupMsgTaskRequest 获取群发成员发送任务列表请�?
type GetGroupMsgTaskRequest struct {
	MsgID  string `json:"msgid"`
	Limit  int    `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}

// GroupMsgTask 群发成员发送任�?
type GroupMsgTask struct {
	UserID   string `json:"userid"`
	Status   int    `json:"status"`
	SendTime int64  `json:"send_time,omitempty"`
}

// GetGroupMsgTaskResponse 获取群发成员发送任务列表响�?
type GetGroupMsgTaskResponse struct {
	common.Response
	NextCursor string         `json:"next_cursor,omitempty"`
	TaskList   []GroupMsgTask `json:"task_list"`
}

// GetGroupMsgSendResultRequest 获取企业群发成员执行结果请求
type GetGroupMsgSendResultRequest struct {
	MsgID  string `json:"msgid"`
	UserID string `json:"userid"`
	Limit  int    `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}

// GroupMsgSendResult 群发成员执行结果
type GroupMsgSendResult struct {
	ExternalUserID string `json:"external_userid,omitempty"`
	ChatID         string `json:"chat_id,omitempty"`
	UserID         string `json:"userid"`
	Status         int    `json:"status"`
	SendTime       int64  `json:"send_time,omitempty"`
}

// GetGroupMsgSendResultResponse 获取企业群发成员执行结果响应
type GetGroupMsgSendResultResponse struct {
	common.Response
	NextCursor string               `json:"next_cursor,omitempty"`
	SendList   []GroupMsgSendResult `json:"send_list"`
}

// SendWelcomeMsgRequest 发送新客户欢迎语请�?
type SendWelcomeMsgRequest struct {
	WelcomeCode string       `json:"welcome_code"`
	Text        *TextContent `json:"text,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

// CancelGroupMsgSendRequest 停止企业群发请求
type CancelGroupMsgSendRequest struct {
	MsgID string `json:"msgid"`
}

// RemindGroupMsgSendRequest 提醒成员群发请求
type RemindGroupMsgSendRequest struct {
	MsgID string `json:"msgid"`
}

// GroupWelcomeTemplate 入群欢迎语素�?
type GroupWelcomeTemplate struct {
	Text        *TextContent           `json:"text,omitempty"`
	Image       *ImageAttachment       `json:"image,omitempty"`
	Link        *LinkAttachment        `json:"link,omitempty"`
	Miniprogram *MiniprogramAttachment `json:"miniprogram,omitempty"`
	File        *FileAttachment        `json:"file,omitempty"`
	Video       *VideoAttachment       `json:"video,omitempty"`
	AgentID     int                    `json:"agentid,omitempty"`
	Notify      int                    `json:"notify,omitempty"`
}

// AddGroupWelcomeTemplateRequest 添加入群欢迎语素材请�?
type AddGroupWelcomeTemplateRequest GroupWelcomeTemplate

// AddGroupWelcomeTemplateResponse 添加入群欢迎语素材响�?
type AddGroupWelcomeTemplateResponse struct {
	common.Response
	TemplateID string `json:"template_id"`
}

// EditGroupWelcomeTemplateRequest 编辑入群欢迎语素材请�?
type EditGroupWelcomeTemplateRequest struct {
	TemplateID  string                 `json:"template_id"`
	Text        *TextContent           `json:"text,omitempty"`
	Image       *ImageAttachment       `json:"image,omitempty"`
	Link        *LinkAttachment        `json:"link,omitempty"`
	Miniprogram *MiniprogramAttachment `json:"miniprogram,omitempty"`
	File        *FileAttachment        `json:"file,omitempty"`
	Video       *VideoAttachment       `json:"video,omitempty"`
	AgentID     int                    `json:"agentid,omitempty"`
}

// GetGroupWelcomeTemplateRequest 获取入群欢迎语素材请�?
type GetGroupWelcomeTemplateRequest struct {
	TemplateID string `json:"template_id"`
}

// GetGroupWelcomeTemplateResponse 获取入群欢迎语素材响�?
type GetGroupWelcomeTemplateResponse struct {
	common.Response
	Text        *TextContent           `json:"text,omitempty"`
	Image       *ImageAttachment       `json:"image,omitempty"`
	Link        *LinkAttachment        `json:"link,omitempty"`
	Miniprogram *MiniprogramAttachment `json:"miniprogram,omitempty"`
	File        *FileAttachment        `json:"file,omitempty"`
	Video       *VideoAttachment       `json:"video,omitempty"`
}

// DelGroupWelcomeTemplateRequest 删除入群欢迎语素材请�?
type DelGroupWelcomeTemplateRequest struct {
	TemplateID string `json:"template_id"`
	AgentID    int    `json:"agentid,omitempty"`
}

