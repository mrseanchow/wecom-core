package contact

import "github.com/mrseanchow/wecom-core/types/common"

// Tag 标签信息
type Tag struct {
	// TagID 标签id
	TagID int `json:"tagid"`
	// TagName 标签名称
	TagName string `json:"tagname"`
}

// TagUser 标签成员
type TagUser struct {
	// UserID 成员账号
	UserID string `json:"userid"`
	// Name 成员名称
	Name string `json:"name,omitempty"`
}

// CreateTagResponse 创建标签响应
type CreateTagResponse struct {
	common.Response
	// TagID 标签id
	TagID int `json:"tagid"`
}

// UpdateTagResponse 更新标签响应
type UpdateTagResponse struct {
	common.Response
}

// DeleteTagResponse 删除标签响应
type DeleteTagResponse struct {
	common.Response
}

// ListTagsResponse 获取标签列表响应
type ListTagsResponse struct {
	common.Response
	// TagList 标签列表
	TagList []Tag `json:"taglist"`
}

// GetTagResponse 获取标签成员响应
type GetTagResponse struct {
	common.Response
	// TagName 标签�?
	TagName string `json:"tagname"`
	// UserList 标签中包含的成员列表
	UserList []TagUser `json:"userlist,omitempty"`
	// PartyList 标签中包含的部门id列表
	PartyList []int `json:"partylist,omitempty"`
}

// AddTagUsersResponse 增加标签成员响应
type AddTagUsersResponse struct {
	common.Response
	// InvalidList 非法的成员帐号列�?
	InvalidList string `json:"invalidlist,omitempty"`
	// InvalidParty 非法的部门id列表
	InvalidParty []int `json:"invalidparty,omitempty"`
}

// DeleteTagUsersResponse 删除标签成员响应
type DeleteTagUsersResponse struct {
	common.Response
	// InvalidList 非法的成员帐号列�?
	InvalidList string `json:"invalidlist,omitempty"`
	// InvalidParty 非法的部门id列表
	InvalidParty []int `json:"invalidparty,omitempty"`
}

