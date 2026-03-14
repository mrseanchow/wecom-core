package wedoc

import "github.com/mrseanchow/wecom-core/types/common"

// ==================== 记录相关类型 ====================

// AddRecordsRequest 添加记录请求
type AddRecordsRequest struct {
	DocID   string      `json:"docid"`
	SheetID string      `json:"sheet_id"`
	KeyType string      `json:"key_type,omitempty"` // CELL_VALUE_KEY_TYPE_FIELD_TITLE �?CELL_VALUE_KEY_TYPE_FIELD_ID
	Records []AddRecord `json:"records"`
}

// AddRecordsResponse 添加记录响应
type AddRecordsResponse struct {
	common.Response
	Records []CommonRecord `json:"records"`
}

// AddRecord 添加记录
type AddRecord struct {
	Values map[string]interface{} `json:"values"` // key为字段标题或ID, value为单元格�?
}

// CommonRecord 通用记录
type CommonRecord struct {
	RecordID string                 `json:"record_id"`
	Values   map[string]interface{} `json:"values"` // key为字段标题或ID, value为单元格�?
}

// GetRecordsRequest 查询记录请求
type GetRecordsRequest struct {
	DocID       string      `json:"docid"`
	SheetID     string      `json:"sheet_id"`
	ViewID      string      `json:"view_id,omitempty"`
	RecordIDs   []string    `json:"record_ids,omitempty"`
	KeyType     string      `json:"key_type,omitempty"`
	FieldTitles []string    `json:"field_titles,omitempty"`
	FieldIDs    []string    `json:"field_ids,omitempty"`
	Sort        []Sort      `json:"sort,omitempty"`
	Offset      uint32      `json:"offset,omitempty"`
	Limit       uint32      `json:"limit,omitempty"`
	Ver         uint32      `json:"ver,omitempty"`
	FilterSpec  *FilterSpec `json:"filter_spec,omitempty"`
}

// GetRecordsResponse 查询记录响应
type GetRecordsResponse struct {
	common.Response
	Total   uint32   `json:"total"`
	HasMore bool     `json:"has_more"`
	Next    uint32   `json:"next"`
	Records []Record `json:"records"`
	Ver     uint32   `json:"ver"`
}

// Sort 排序
type Sort struct {
	FieldTitle string `json:"field_title"`
	Desc       bool   `json:"desc,omitempty"`
}

// Record 记录
type Record struct {
	RecordID    string                 `json:"record_id"`
	CreateTime  string                 `json:"create_time"`
	UpdateTime  string                 `json:"update_time"`
	Values      map[string]interface{} `json:"values"`
	CreatorName string                 `json:"creator_name"`
	UpdaterName string                 `json:"updater_name"`
}

// UpdateRecordsRequest 更新记录请求
type UpdateRecordsRequest struct {
	DocID   string         `json:"docid"`
	SheetID string         `json:"sheet_id"`
	KeyType string         `json:"key_type,omitempty"`
	Records []UpdateRecord `json:"records"`
}

// UpdateRecordsResponse 更新记录响应
type UpdateRecordsResponse struct {
	common.Response
	Records []CommonRecord `json:"records"`
}

// UpdateRecord 更新记录
type UpdateRecord struct {
	RecordID string                 `json:"record_id"`
	Values   map[string]interface{} `json:"values"`
}

// DeleteRecordsRequest 删除记录请求
type DeleteRecordsRequest struct {
	DocID     string   `json:"docid"`
	SheetID   string   `json:"sheet_id"`
	RecordIDs []string `json:"record_ids"`
}

// ==================== 字段相关类型 ====================

// AddFieldsRequest 添加字段请求
type AddFieldsRequest struct {
	DocID   string     `json:"docid"`
	SheetID string     `json:"sheet_id"`
	Fields  []AddField `json:"fields"`
}

// AddFieldsResponse 添加字段响应
type AddFieldsResponse struct {
	common.Response
	Fields []Field `json:"fields"`
}

// AddField 添加字段
type AddField struct {
	FieldTitle           string                     `json:"field_title"`
	FieldType            string                     `json:"field_type"`
	PropertyNumber       *NumberFieldProperty       `json:"property_number,omitempty"`
	PropertyCheckbox     *CheckboxFieldProperty     `json:"property_checkbox,omitempty"`
	PropertyDateTime     *DateTimeFieldProperty     `json:"property_date_time,omitempty"`
	PropertyAttachment   *AttachmentFieldProperty   `json:"property_attachment,omitempty"`
	PropertyUser         *UserFieldProperty         `json:"property_user,omitempty"`
	PropertyURL          *URLFieldProperty          `json:"property_url,omitempty"`
	PropertySelect       *SelectFieldProperty       `json:"property_select,omitempty"`
	PropertyCreatedTime  *CreatedTimeFieldProperty  `json:"property_created_time,omitempty"`
	PropertyModifiedTime *ModifiedTimeFieldProperty `json:"property_modified_time,omitempty"`
	PropertyProgress     *ProgressFieldProperty     `json:"property_progress,omitempty"`
	PropertySingleSelect *SingleSelectFieldProperty `json:"property_single_select,omitempty"`
	PropertyReference    *ReferenceFieldProperty    `json:"property_reference,omitempty"`
	PropertyLocation     *LocationFieldProperty     `json:"property_location,omitempty"`
	PropertyAutoNumber   *AutoNumberFieldProperty   `json:"property_auto_number,omitempty"`
	PropertyCurrency     *CurrencyFieldProperty     `json:"property_currency,omitempty"`
	PropertyWWGroup      *WWGroupFieldProperty      `json:"property_ww_group,omitempty"`
	PropertyPercentage   *PercentageFieldProperty   `json:"property_percentage,omitempty"`
	PropertyBarcode      *BarcodeFieldProperty      `json:"property_barcode,omitempty"`
}

// Field 字段
type Field struct {
	FieldID              string                     `json:"field_id"`
	FieldTitle           string                     `json:"field_title"`
	FieldType            string                     `json:"field_type"`
	PropertyNumber       *NumberFieldProperty       `json:"property_number,omitempty"`
	PropertyCheckbox     *CheckboxFieldProperty     `json:"property_checkbox,omitempty"`
	PropertyDateTime     *DateTimeFieldProperty     `json:"property_date_time,omitempty"`
	PropertyAttachment   *AttachmentFieldProperty   `json:"property_attachment,omitempty"`
	PropertyUser         *UserFieldProperty         `json:"property_user,omitempty"`
	PropertyURL          *URLFieldProperty          `json:"property_url,omitempty"`
	PropertySelect       *SelectFieldProperty       `json:"property_select,omitempty"`
	PropertyCreatedTime  *CreatedTimeFieldProperty  `json:"property_created_time,omitempty"`
	PropertyModifiedTime *ModifiedTimeFieldProperty `json:"property_modified_time,omitempty"`
	PropertyProgress     *ProgressFieldProperty     `json:"property_progress,omitempty"`
	PropertySingleSelect *SingleSelectFieldProperty `json:"property_single_select,omitempty"`
	PropertyReference    *ReferenceFieldProperty    `json:"property_reference,omitempty"`
	PropertyLocation     *LocationFieldProperty     `json:"property_location,omitempty"`
	PropertyAutoNumber   *AutoNumberFieldProperty   `json:"property_auto_number,omitempty"`
	PropertyCurrency     *CurrencyFieldProperty     `json:"property_currency,omitempty"`
	PropertyWWGroup      *WWGroupFieldProperty      `json:"property_ww_group,omitempty"`
	PropertyPercentage   *PercentageFieldProperty   `json:"property_percentage,omitempty"`
	PropertyBarcode      *BarcodeFieldProperty      `json:"property_barcode,omitempty"`
}

// 字段属性类�?

// NumberFieldProperty 数字类型字段属�?
type NumberFieldProperty struct {
	DecimalPlaces int  `json:"decimal_places,omitempty"` // -1~4
	UseSeparate   bool `json:"use_separate,omitempty"`
}

// CheckboxFieldProperty 复选框类型字段属�?
type CheckboxFieldProperty struct {
	Checked bool `json:"checked,omitempty"`
}

// DateTimeFieldProperty 日期类型字段属�?
type DateTimeFieldProperty struct {
	Format   string `json:"format,omitempty"`
	AutoFill bool   `json:"auto_fill,omitempty"`
}

// AttachmentFieldProperty 文件类型字段属�?
type AttachmentFieldProperty struct {
	DisplayMode string `json:"display_mode,omitempty"` // DISPLAY_MODE_LIST �?DISPLAY_MODE_GRID
}

// UserFieldProperty 成员类型字段属�?
type UserFieldProperty struct {
	IsMultiple bool `json:"is_multiple,omitempty"`
	IsNotified bool `json:"is_notified,omitempty"`
}

// URLFieldProperty 超链接类型字段属�?
type URLFieldProperty struct {
	Type string `json:"type,omitempty"` // LINK_TYPE_PURE_TEXT �?LINK_TYPE_ICON_TEXT
}

// SelectFieldProperty 多选类型字段属�?
type SelectFieldProperty struct {
	IsQuickAdd bool     `json:"is_quick_add,omitempty"`
	Options    []Option `json:"options,omitempty"`
}

// CreatedTimeFieldProperty 创建时间类型字段属�?
type CreatedTimeFieldProperty struct {
	Format string `json:"format,omitempty"`
}

// ModifiedTimeFieldProperty 最后编辑时间类型字段属�?
type ModifiedTimeFieldProperty struct {
	Format string `json:"format,omitempty"`
}

// ProgressFieldProperty 进度类型字段属�?
type ProgressFieldProperty struct {
	DecimalPlaces int `json:"decimal_places,omitempty"`
}

// SingleSelectFieldProperty 单选类型字段属�?
type SingleSelectFieldProperty struct {
	IsQuickAdd bool     `json:"is_quick_add,omitempty"`
	Options    []Option `json:"options,omitempty"`
}

// ReferenceFieldProperty 关联字段属�?
type ReferenceFieldProperty struct {
	SubID      string `json:"sub_id,omitempty"`
	FieldID    string `json:"field_id,omitempty"`
	IsMultiple bool   `json:"is_multiple,omitempty"`
	ViewID     string `json:"view_id,omitempty"`
}

// LocationFieldProperty 地理位置字段属�?
type LocationFieldProperty struct {
	InputType string `json:"input_type,omitempty"` // LOCATION_INPUT_TYPE_MANUAL �?LOCATION_INPUT_TYPE_AUTO
}

// AutoNumberFieldProperty 自动编号字段属�?
type AutoNumberFieldProperty struct {
	Type                   string       `json:"type,omitempty"` // NUMBER_TYPE_INCR �?NUMBER_TYPE_CUSTOM
	Rules                  []NumberRule `json:"rules,omitempty"`
	ReformatExistingRecord bool         `json:"reformat_existing_record,omitempty"`
}

// NumberRule 自动编号规则
type NumberRule struct {
	Type  string `json:"type,omitempty"`  // NUMBER_RULE_TYPE_INCR, NUMBER_RULE_TYPE_FIXED_CHAR, NUMBER_RULE_TYPE_TIME
	Value string `json:"value,omitempty"` // 创建时间格式或固定字符或自增数字位数
}

// CurrencyFieldProperty 货币类型字段属�?
type CurrencyFieldProperty struct {
	CurrencyType  string `json:"currency_type,omitempty"` // CURRENCY_TYPE_CNY �?
	DecimalPlaces int    `json:"decimal_places,omitempty"`
	UseSeparate   bool   `json:"use_separate,omitempty"`
}

// WWGroupFieldProperty 群类型字段属�?
type WWGroupFieldProperty struct {
	AllowMultiple bool `json:"allow_multiple,omitempty"`
}

// PercentageFieldProperty 百分数类型字段属�?
type PercentageFieldProperty struct {
	DecimalPlaces int  `json:"decimal_places,omitempty"`
	UseSeparate   bool `json:"use_separate,omitempty"`
}

// BarcodeFieldProperty 条码类型字段属�?
type BarcodeFieldProperty struct {
	MobileScanOnly bool `json:"mobile_scan_only,omitempty"`
}

// Option 选项
type Option struct {
	ID    string `json:"id,omitempty"`
	Text  string `json:"text,omitempty"`
	Style int    `json:"style,omitempty"` // 1-27
}

// GetFieldsRequest 查询字段请求
type GetFieldsRequest struct {
	DocID       string   `json:"docid"`
	SheetID     string   `json:"sheet_id"`
	ViewID      string   `json:"view_id,omitempty"`
	FieldIDs    []string `json:"field_ids,omitempty"`
	FieldTitles []string `json:"field_titles,omitempty"`
	Offset      int      `json:"offset,omitempty"`
	Limit       int      `json:"limit,omitempty"`
}

// GetFieldsResponse 查询字段响应
type GetFieldsResponse struct {
	common.Response
	Total  int     `json:"total"`
	Fields []Field `json:"fields"`
}

// UpdateFieldsRequest 更新字段请求
type UpdateFieldsRequest struct {
	DocID   string        `json:"docid"`
	SheetID string        `json:"sheet_id"`
	Fields  []UpdateField `json:"fields"`
}

// UpdateFieldsResponse 更新字段响应
type UpdateFieldsResponse struct {
	common.Response
	Fields []Field `json:"fields"`
}

// UpdateField 更新字段
type UpdateField struct {
	FieldID              string                     `json:"field_id"`
	FieldTitle           string                     `json:"field_title,omitempty"`
	PropertyNumber       *NumberFieldProperty       `json:"property_number,omitempty"`
	PropertyCheckbox     *CheckboxFieldProperty     `json:"property_checkbox,omitempty"`
	PropertyDateTime     *DateTimeFieldProperty     `json:"property_date_time,omitempty"`
	PropertyAttachment   *AttachmentFieldProperty   `json:"property_attachment,omitempty"`
	PropertyUser         *UserFieldProperty         `json:"property_user,omitempty"`
	PropertyURL          *URLFieldProperty          `json:"property_url,omitempty"`
	PropertySelect       *SelectFieldProperty       `json:"property_select,omitempty"`
	PropertyCreatedTime  *CreatedTimeFieldProperty  `json:"property_created_time,omitempty"`
	PropertyModifiedTime *ModifiedTimeFieldProperty `json:"property_modified_time,omitempty"`
	PropertyProgress     *ProgressFieldProperty     `json:"property_progress,omitempty"`
	PropertySingleSelect *SingleSelectFieldProperty `json:"property_single_select,omitempty"`
	PropertyReference    *ReferenceFieldProperty    `json:"property_reference,omitempty"`
	PropertyLocation     *LocationFieldProperty     `json:"property_location,omitempty"`
	PropertyAutoNumber   *AutoNumberFieldProperty   `json:"property_auto_number,omitempty"`
	PropertyCurrency     *CurrencyFieldProperty     `json:"property_currency,omitempty"`
	PropertyWWGroup      *WWGroupFieldProperty      `json:"property_ww_group,omitempty"`
	PropertyPercentage   *PercentageFieldProperty   `json:"property_percentage,omitempty"`
	PropertyBarcode      *BarcodeFieldProperty      `json:"property_barcode,omitempty"`
}

// DeleteFieldsRequest 删除字段请求
type DeleteFieldsRequest struct {
	DocID    string   `json:"docid"`
	SheetID  string   `json:"sheet_id"`
	FieldIDs []string `json:"field_ids"`
}

// ==================== 视图相关类型 ====================

// AddViewRequest 添加视图请求
type AddViewRequest struct {
	DocID     string        `json:"docid"`
	SheetID   string        `json:"sheet_id"`
	ViewTitle string        `json:"view_title"`
	ViewType  string        `json:"view_type"` // VIEW_TYPE_GRID, VIEW_TYPE_KANBAN, VIEW_TYPE_GALLERY, VIEW_TYPE_GANTT
	Property  *ViewProperty `json:"property,omitempty"`
}

// AddViewResponse 添加视图响应
type AddViewResponse struct {
	common.Response
	View *View `json:"view"`
}

// View 视图
type View struct {
	ViewID    string        `json:"view_id"`
	ViewTitle string        `json:"view_title"`
	ViewType  string        `json:"view_type"`
	Property  *ViewProperty `json:"property,omitempty"`
}

// ViewProperty 视图属�?
type ViewProperty struct {
	AutoSort           bool             `json:"auto_sort,omitempty"`
	SortSpec           *SortSpec        `json:"sort_spec,omitempty"`
	GroupSpec          *GroupSpec       `json:"group_spec,omitempty"`
	FilterSpec         *FilterSpec      `json:"filter_spec,omitempty"`
	IsFieldStatEnabled bool             `json:"is_field_stat_enabled,omitempty"`
	FieldVisibility    map[string]bool  `json:"field_visibility,omitempty"` // key为字段ID, value为是否显�?
	FrozenFieldCount   int32            `json:"frozen_field_count,omitempty"`
	ColorConfig        *ViewColorConfig `json:"color_config,omitempty"`
}

// SortSpec 排序设置
type SortSpec struct {
	SortInfos []SortInfo `json:"sort_infos,omitempty"`
}

// SortInfo 排序信息
type SortInfo struct {
	FieldID string `json:"field_id"`
	Desc    bool   `json:"desc,omitempty"`
}

// GroupSpec 分组设置
type GroupSpec struct {
	Groups []GroupInfo `json:"groups,omitempty"`
}

// GroupInfo 分组信息
type GroupInfo struct {
	FieldID string `json:"field_id"`
	Desc    bool   `json:"desc,omitempty"`
}

// FilterSpec 过滤设置
type FilterSpec struct {
	Conjunction string      `json:"conjunction,omitempty"` // CONJUNCTION_AND �?CONJUNCTION_OR
	Conditions  []Condition `json:"conditions,omitempty"`
}

// Condition 过滤条件
type Condition struct {
	FieldID       string               `json:"field_id"`
	FieldType     string               `json:"field_type,omitempty"`
	Operator      string               `json:"operator"` // OPERATOR_IS, OPERATOR_IS_NOT, OPERATOR_CONTAINS �?
	StringValue   *FilterStringValue   `json:"string_value,omitempty"`
	NumberValue   *FilterNumberValue   `json:"number_value,omitempty"`
	BoolValue     *FilterBoolValue     `json:"bool_value,omitempty"`
	UserValue     *FilterUserValue     `json:"user_value,omitempty"`
	DateTimeValue *FilterDateTimeValue `json:"date_time_value,omitempty"`
}

// FilterStringValue 字符串过滤�?
type FilterStringValue struct {
	Value []string `json:"value"`
}

// FilterNumberValue 数字过滤�?
type FilterNumberValue struct {
	Value float64 `json:"value"`
}

// FilterBoolValue 布尔过滤�?
type FilterBoolValue struct {
	Value bool `json:"value"`
}

// FilterUserValue 用户过滤�?
type FilterUserValue struct {
	Value []string `json:"value"` // 成员ID列表
}

// FilterDateTimeValue 日期时间过滤�?
type FilterDateTimeValue struct {
	Type  string   `json:"type"`            // DATE_TIME_TYPE_DETAIL_DATE, DATE_TIME_TYPE_TODAY �?
	Value []string `json:"value,omitempty"` // 具体日期�?毫秒时间�?
}

// ViewColorConfig 填色设置
type ViewColorConfig struct {
	Conditions []ViewColorCondition `json:"conditions,omitempty"`
}

// ViewColorCondition 填色条件
type ViewColorCondition struct {
	ID        string     `json:"id,omitempty"` // 填色id
	Type      string     `json:"type"`         // VIEW_COLOR_CONDITION_TYPE_ROW, VIEW_COLOR_CONDITION_TYPE_COLUMN, VIEW_COLOR_CONDITION_TYPE_CELL
	Color     string     `json:"color"`        // fillColorGray_5, accentBlueLighten_5 �?
	Condition *Condition `json:"condition"`
}

// GetViewsRequest 查询视图请求
type GetViewsRequest struct {
	DocID   string   `json:"docid"`
	SheetID string   `json:"sheet_id"`
	ViewIDs []string `json:"view_ids,omitempty"`
	Offset  uint32   `json:"offset,omitempty"`
	Limit   uint32   `json:"limit,omitempty"`
}

// GetViewsResponse 查询视图响应
type GetViewsResponse struct {
	common.Response
	Total   uint32 `json:"total"`
	HasMore bool   `json:"has_more"`
	Next    uint32 `json:"next"`
	Views   []View `json:"views"`
}

// UpdateViewRequest 更新视图请求
type UpdateViewRequest struct {
	DocID     string        `json:"docid"`
	SheetID   string        `json:"sheet_id"`
	ViewID    string        `json:"view_id"`
	ViewTitle string        `json:"view_title,omitempty"`
	Property  *ViewProperty `json:"property,omitempty"`
}

// UpdateViewResponse 更新视图响应
type UpdateViewResponse struct {
	common.Response
	View *View `json:"view"`
}

// DeleteViewRequest 删除视图请求
type DeleteViewRequest struct {
	DocID   string `json:"docid"`
	SheetID string `json:"sheet_id"`
	ViewID  string `json:"view_id"`
}

// ==================== 子表相关类型 ====================

// AddSheetRequest 添加子表请求
type AddSheetRequest struct {
	DocID      string         `json:"docid"`
	Properties *SheetProperty `json:"properties,omitempty"`
}

// AddSheetResponse 添加子表响应
type AddSheetResponse struct {
	common.Response
	Properties *SheetProperty `json:"properties"`
}

// SheetProperty 子表属�?
type SheetProperty struct {
	SheetID string `json:"sheet_id,omitempty"`
	Title   string `json:"title,omitempty"`
	Index   int32  `json:"index,omitempty"`
}

// GetSheetRequest 查询子表请求
type GetSheetRequest struct {
	DocID            string `json:"docid"`
	SheetID          string `json:"sheet_id,omitempty"`
	NeedAllTypeSheet bool   `json:"need_all_type_sheet,omitempty"`
}

// GetSheetResponse 查询子表响应
type GetSheetResponse struct {
	common.Response
	SheetList []SheetInfo `json:"sheet_list"`
}

// SheetInfo 子表信息
type SheetInfo struct {
	SheetID   string `json:"sheet_id"`
	Title     string `json:"title"`
	IsVisible bool   `json:"is_visible"`
	Type      string `json:"type"` // smartsheet, dashboard, external
}

// UpdateSheetRequest 更新子表请求
type UpdateSheetRequest struct {
	DocID      string         `json:"docid"`
	SheetID    string         `json:"sheet_id"`
	Properties *SheetProperty `json:"properties"`
}

// UpdateSheetResponse 更新子表响应
type UpdateSheetResponse struct {
	common.Response
	Properties *SheetProperty `json:"properties"`
}

// DeleteSheetRequest 删除子表请求
type DeleteSheetRequest struct {
	DocID   string `json:"docid"`
	SheetID string `json:"sheet_id"`
}

// ==================== 编组相关类型 ====================

// AddFieldGroupRequest 添加编组请求
type AddFieldGroupRequest struct {
	DocID    string               `json:"docid"`
	SheetID  string               `json:"sheet_id"`
	Name     string               `json:"name"`
	Children []FieldGroupChildren `json:"children,omitempty"`
}

// AddFieldGroupResponse 添加编组响应
type AddFieldGroupResponse struct {
	common.Response
	FieldGroup *FieldGroup `json:"field_group"`
}

// FieldGroup 编组
type FieldGroup struct {
	FieldGroupID string               `json:"field_group_id"`
	Name         string               `json:"name"`
	Children     []FieldGroupChildren `json:"children,omitempty"`
}

// FieldGroupChildren 编组内容
type FieldGroupChildren struct {
	FieldID string `json:"field_id"`
}

// GetFieldGroupRequest 获取编组请求
type GetFieldGroupRequest struct {
	DocID   string `json:"docid"`
	SheetID string `json:"sheet_id"`
	Offset  int    `json:"offset,omitempty"`
	Limit   int    `json:"limit,omitempty"`
}

// GetFieldGroupResponse 获取编组响应
type GetFieldGroupResponse struct {
	common.Response
	Total       int          `json:"total"`
	FieldGroups []FieldGroup `json:"field_groups"`
}

// UpdateFieldGroupRequest 更新编组请求
type UpdateFieldGroupRequest struct {
	DocID        string               `json:"docid"`
	SheetID      string               `json:"sheet_id"`
	FieldGroupID string               `json:"field_group_id"`
	Name         string               `json:"name,omitempty"`
	Children     []FieldGroupChildren `json:"children,omitempty"`
}

// UpdateFieldGroupResponse 更新编组响应
type UpdateFieldGroupResponse struct {
	common.Response
	FieldGroup *FieldGroup `json:"field_group"`
}

// DeleteFieldGroupRequest 删除编组请求
type DeleteFieldGroupRequest struct {
	DocID        string `json:"docid"`
	SheetID      string `json:"sheet_id"`
	FieldGroupID string `json:"field_group_id"`
}

// ==================== 内容权限相关类型 ====================

// GetSheetPrivRequest 查询智能表格子表权限请求
type GetSheetPrivRequest struct {
	DocID      string   `json:"docid"`                  // 智能表ID
	Type       uint32   `json:"type"`                   // 权限规则类型�?-全员权限�?-额外权限
	RuleIDList []uint32 `json:"rule_id_list,omitempty"` // 需要查询的规则id列表，查询额外权限时填写
}

// GetSheetPrivResponse 查询智能表格子表权限响应
type GetSheetPrivResponse struct {
	common.Response
	RuleList []SheetPrivRule `json:"rule_list"` // 权限列表
}

// SheetPrivRule 子表权限规则
type SheetPrivRule struct {
	RuleID   uint32      `json:"rule_id"`   // 规则id
	Type     uint32      `json:"type"`      // 权限规则类型�?-全员权限�?-额外权限
	Name     string      `json:"name"`      // 权限名称，仅当type�?时有�?
	PrivList []SheetPriv `json:"priv_list"` // 针对不同子表设置内容权限
}

// SheetPriv 子表权限配置
type SheetPriv struct {
	SheetID                   string      `json:"sheet_id"`                      // 子表ID
	Priv                      uint32      `json:"priv"`                          // 子表权限: 1-全部权限�?-可编辑；3-仅浏览；4-无权�?
	CanInsertRecord           bool        `json:"can_insert_record"`             // 是否可以新增记录
	CanDeleteRecord           bool        `json:"can_delete_record"`             // 是否可以删除记录
	CanCreateModifyDeleteView bool        `json:"can_create_modify_delete_view"` // 是否可以增、删、改视图
	FieldPriv                 *FieldPriv  `json:"field_priv,omitempty"`          // 按字段配置权�?
	RecordPriv                *RecordPriv `json:"record_priv,omitempty"`         // 按记录配置权�?
	Clear                     bool        `json:"clear"`                         // 清除子表的设置，恢复默认权限
}

// FieldPriv 字段权限
type FieldPriv struct {
	FieldRangeType   uint32      `json:"field_range_type"`             // 子表权限对所有字段生效还是部分字段生效：1-所有字段；2-部分字段
	FieldRuleList    []FieldRule `json:"field_rule_list,omitempty"`    // 按字段分别配置权�?
	FieldDefaultRule *FieldRule  `json:"field_default_rule,omitempty"` // 未指定字段和后续新增字段的默认配�?
}

// FieldRule 字段权限规则
type FieldRule struct {
	FieldID   string `json:"field_id"`   // 字段id
	FieldType string `json:"field_type"` // 字段类型
	CanEdit   bool   `json:"can_edit"`   // 可编�?
	CanInsert bool   `json:"can_insert"` // 可首次提�?
	CanView   bool   `json:"can_view"`   // 可查�?
}

// RecordPriv 记录权限
type RecordPriv struct {
	RecordRangeType uint32       `json:"record_range_type"`          // 子表权限对记录生效范围：1-全部记录�?-满足任意条件的记录；3-满足全部条件的记�?
	RecordRuleList  []RecordRule `json:"record_rule_list,omitempty"` // 记录的条件列�?
	OtherPriv       uint32       `json:"other_priv"`                 // 当记录不满足条件的时的权限类型：1-不可编辑 2-不可查看
}

// RecordRule 记录权限规则
type RecordRule struct {
	FieldID   string   `json:"field_id"`             // 字段id，当field_id为CREATED_USER时表示记录创建�?
	FieldType string   `json:"field_type,omitempty"` // 字段类型
	OperType  uint32   `json:"oper_type"`            // 逻辑判断类型�?-包含自己�?-包含value�?-不包含value�?-等于value�?-不等于value�?-为空�?-非空
	Value     []string `json:"value,omitempty"`      // 用于单选、多选字段的option_id
}

// UpdateSheetPrivRequest 更新智能表格子表权限请求
type UpdateSheetPrivRequest struct {
	DocID    string      `json:"docid"`               // 智能表ID
	Type     uint32      `json:"type"`                // 权限规则类型�?-全员权限�?-额外权限
	RuleID   uint32      `json:"rule_id,omitempty"`   // 当type�?时必�?
	Name     string      `json:"name,omitempty"`      // 更新权限名称，仅当type�?时有�?
	PrivList []SheetPriv `json:"priv_list,omitempty"` // 针对不同子表设置内容权限
}

// CreateRuleRequest 新增智能表格指定成员额外权限请求
type CreateRuleRequest struct {
	DocID string `json:"docid"` // 智能表ID
	Name  string `json:"name"`  // 权限规则名称，不可重�?
}

// CreateRuleResponse 新增智能表格指定成员额外权限响应
type CreateRuleResponse struct {
	common.Response
	RuleID uint32 `json:"rule_id"` // 成员权限规则id
}

// ModRuleMemberRequest 更新智能表格指定成员额外权限请求
type ModRuleMemberRequest struct {
	DocID          string       `json:"docid"`                      // 智能表ID
	RuleID         uint32       `json:"rule_id"`                    // 需要更新的id
	AddMemberRange *MemberRange `json:"add_member_range,omitempty"` // 新增成员
	DelMemberRange *MemberRange `json:"del_member_range,omitempty"` // 删除成员
}

// MemberRange 成员范围
type MemberRange struct {
	UserIDList []string `json:"userid_list,omitempty"` // 成员的userid列表
}

// DeleteRuleRequest 删除智能表格指定成员额外权限请求
type DeleteRuleRequest struct {
	DocID      string   `json:"docid"`        // 智能表ID
	RuleIDList []uint32 `json:"rule_id_list"` // 需要删除的规则id列表
}

