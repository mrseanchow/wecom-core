package hr

import "github.com/mrseanchow/wecom-core/types/common"

const (
	FieldValueTypeString = 1
	FieldValueTypeUint64 = 2
	FieldValueTypeUint32 = 3
	FieldValueTypeInt64  = 4
	FieldValueTypeMobile = 5
	FieldValueTypeFile   = 6
)

const (
	FieldTypeText   = 1
	FieldTypeOption = 2
	FieldTypeTime   = 3
	FieldTypeImage  = 4
	FieldTypeFile   = 5
	FieldTypeFiles  = 6
)

const (
	QueryResultSuccess        = 1
	QueryResultFailed         = 2
	QueryResultNotFound       = 3
	QueryResultNotSupported   = 5
	UpdateResultSuccess       = 1
	UpdateResultFailed        = 2
	UpdateResultNotFound      = 3
	UpdateResultRequiredEmpty = 4
	UpdateResultNotSupported  = 5
)

const (
	GroupTypeEducation = 1
	GroupTypeWork      = 2
	GroupTypeFamily    = 3
	GroupTypeEmergency = 4
	GroupTypeContract  = 5
)

type GetFieldsResponse struct {
	common.Response
	GroupList []FieldGroup `json:"group_list"`
}

type FieldGroup struct {
	GroupID   int     `json:"group_id"`
	GroupName string  `json:"group_name"`
	FieldList []Field `json:"field_list"`
}

type Field struct {
	FieldID    int      `json:"fieldid"`
	FieldName  string   `json:"field_name"`
	FieldType  int      `json:"field_type"`
	ValueType  int      `json:"value_type"`
	IsMust     bool     `json:"is_must"`
	OptionList []Option `json:"option_list,omitempty"`
}

type Option struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

type MobileValue struct {
	ValueCountryCode string `json:"value_country_code"`
	ValueMobile      string `json:"value_mobile"`
}

type FileValue struct {
	MediaID []string `json:"media_id"`
}

type FieldValue struct {
	FieldID     int          `json:"fieldid"`
	SubIdx      int          `json:"sub_idx"`
	ValueString *string      `json:"value_string,omitempty"`
	ValueUint64 *uint64      `json:"value_uint64,omitempty"`
	ValueUint32 *uint32      `json:"value_uint32,omitempty"`
	ValueInt64  *int64       `json:"value_int64,omitempty"`
	ValueMobile *MobileValue `json:"value_mobile,omitempty"`
	ValueFile   *FileValue   `json:"value_file,omitempty"`
}

type FieldInfo struct {
	FieldID     int          `json:"fieldid"`
	SubIdx      int          `json:"sub_idx"`
	Result      int          `json:"result"`
	ValueType   int          `json:"value_type"`
	ValueString *string      `json:"value_string,omitempty"`
	ValueUint64 *uint64      `json:"value_uint64,omitempty"`
	ValueUint32 *uint32      `json:"value_uint32,omitempty"`
	ValueInt64  *int64       `json:"value_int64,omitempty"`
	ValueMobile *MobileValue `json:"value_mobile,omitempty"`
	ValueFile   *FileValue   `json:"value_file,omitempty"`
}

