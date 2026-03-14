package hr

import "github.com/mrseanchow/wecom-core/types/common"

type GetStaffInfoRequest struct {
	UserID   string           `json:"userid"`
	GetAll   bool             `json:"get_all,omitempty"`
	FieldIDs []FieldIDWithIdx `json:"fieldids,omitempty"`
}

type FieldIDWithIdx struct {
	FieldID int `json:"fieldid"`
	SubIdx  int `json:"sub_idx,omitempty"`
}

type GetStaffInfoResponse struct {
	common.Response
	FieldInfo []FieldInfo `json:"field_info"`
}

type UpdateStaffInfoRequest struct {
	UserID      string       `json:"userid"`
	UpdateItems []UpdateItem `json:"update_items,omitempty"`
	RemoveItems []RemoveItem `json:"remove_items,omitempty"`
	InsertItems []InsertItem `json:"insert_items,omitempty"`
}

type UpdateItem struct {
	FieldID     int          `json:"fieldid"`
	SubIdx      int          `json:"sub_idx,omitempty"`
	ValueString *string      `json:"value_string,omitempty"`
	ValueUint64 *uint64      `json:"value_uint64,omitempty"`
	ValueUint32 *uint32      `json:"value_uint32,omitempty"`
	ValueInt64  *int64       `json:"value_int64,omitempty"`
	ValueMobile *MobileValue `json:"value_mobile,omitempty"`
}

type RemoveItem struct {
	GroupType int `json:"group_type"`
	SubIdx    int `json:"sub_idx"`
}

type InsertItem struct {
	GroupType int          `json:"group_type"`
	Item      []UpdateItem `json:"item,omitempty"`
}

type UpdateStaffInfoResponse struct {
	common.Response
	UpdateResults []UpdateResult `json:"update_results,omitempty"`
	RemoveResults []RemoveResult `json:"remove_results,omitempty"`
	InsertResults []InsertResult `json:"insert_result,omitempty"`
}

type UpdateResult struct {
	FieldID int `json:"fieldid"`
	SubIdx  int `json:"sub_idx"`
	Result  int `json:"result"`
}

type RemoveResult struct {
	GroupType int `json:"group_type"`
	SubIdx    int `json:"sub_idx"`
	Result    int `json:"result"`
}

type InsertResult struct {
	GroupType int `json:"group_type"`
	Idx       int `json:"idx"`
	Result    int `json:"result"`
}

