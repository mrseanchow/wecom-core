package callback

type DocMemberChangeEvent struct {
	BaseEvent
	DocId []string `xml:"DocId"`
}

type AddFieldEvent struct {
	BaseEvent
	DocId   []string `xml:"DocId"`
	SheetId string   `xml:"SheetId"`
	FieldId []string `xml:"FieldId"`
}

type UpdateFieldEvent struct {
	BaseEvent
	DocId   []string `xml:"DocId"`
	SheetId string   `xml:"SheetId"`
	FieldId []string `xml:"FieldId"`
}

type DeleteFieldEvent struct {
	BaseEvent
	DocId   []string `xml:"DocId"`
	SheetId string   `xml:"SheetId"`
	FieldId []string `xml:"FieldId"`
}

type DeleteDocEvent struct {
	BaseEvent
	DocId []string `xml:"DocId"`
}

type FormCompleteEvent struct {
	BaseEvent
	FormId []string `xml:"FormId"`
}

type DeleteFormEvent struct {
	BaseEvent
	FormId []string `xml:"FormId"`
}

type FormSettingsChangeEvent struct {
	BaseEvent
	FormId []string `xml:"FormId"`
}

type AddRecordEvent struct {
	BaseEvent
	DocId    []string `xml:"DocId"`
	SheetId  string   `xml:"SheetId"`
	RecordId []string `xml:"RecordId"`
}

type UpdateRecordEvent struct {
	BaseEvent
	DocId    []string `xml:"DocId"`
	SheetId  string   `xml:"SheetId"`
	RecordId []string `xml:"RecordId"`
}

type DeleteRecordEvent struct {
	BaseEvent
	DocId    []string `xml:"DocId"`
	SheetId  string   `xml:"SheetId"`
	RecordId []string `xml:"RecordId"`
}

func ParseDocMemberChange(data []byte) (*DocMemberChangeEvent, error) {
	return Parse[DocMemberChangeEvent](data)
}

func ParseAddField(data []byte) (*AddFieldEvent, error) {
	return Parse[AddFieldEvent](data)
}

func ParseUpdateField(data []byte) (*UpdateFieldEvent, error) {
	return Parse[UpdateFieldEvent](data)
}

func ParseDeleteField(data []byte) (*DeleteFieldEvent, error) {
	return Parse[DeleteFieldEvent](data)
}

func ParseDeleteDoc(data []byte) (*DeleteDocEvent, error) {
	return Parse[DeleteDocEvent](data)
}

func ParseFormComplete(data []byte) (*FormCompleteEvent, error) {
	return Parse[FormCompleteEvent](data)
}

func ParseDeleteForm(data []byte) (*DeleteFormEvent, error) {
	return Parse[DeleteFormEvent](data)
}

func ParseFormSettingsChange(data []byte) (*FormSettingsChangeEvent, error) {
	return Parse[FormSettingsChangeEvent](data)
}

func ParseAddRecord(data []byte) (*AddRecordEvent, error) {
	return Parse[AddRecordEvent](data)
}

func ParseUpdateRecord(data []byte) (*UpdateRecordEvent, error) {
	return Parse[UpdateRecordEvent](data)
}

func ParseDeleteRecord(data []byte) (*DeleteRecordEvent, error) {
	return Parse[DeleteRecordEvent](data)
}

