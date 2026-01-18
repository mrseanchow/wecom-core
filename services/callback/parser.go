package callback

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

type Event string

const (
	MsgTypeEvent = "event"

	EventChangeContact         = "change_contact"
	EventChangeExternalContact = "change_external_contact"
	EventChangeExternalChat    = "change_external_chat"
	EventChangeExternalTag     = "change_external_tag"
	EventCustomerAcquisition   = "customer_acquisition"
	EventBatchJobResult        = "batch_job_result"
)

type ChangeType string

const (
	ChangeTypeCreateUser             ChangeType = "create_user"
	ChangeTypeUpdateUser             ChangeType = "update_user"
	ChangeTypeDeleteUser             ChangeType = "delete_user"
	ChangeTypeCreateParty            ChangeType = "create_party"
	ChangeTypeUpdateParty            ChangeType = "update_party"
	ChangeTypeDeleteParty            ChangeType = "delete_party"
	ChangeTypeUpdateTag              ChangeType = "update_tag"
	ChangeTypeAddExternalContact     ChangeType = "add_external_contact"
	ChangeTypeEditExternalContact    ChangeType = "edit_external_contact"
	ChangeTypeAddHalfExternalContact ChangeType = "add_half_external_contact"
	ChangeTypeDelExternalContact     ChangeType = "del_external_contact"
	ChangeTypeDelFollowUser          ChangeType = "del_follow_user"
	ChangeTypeTransferFail           ChangeType = "transfer_fail"
	ChangeTypeCreateChat             ChangeType = "create"
	ChangeTypeUpdateChat             ChangeType = "update"
	ChangeTypeDismissChat            ChangeType = "dismiss"
	ChangeTypeBalanceLow             ChangeType = "balance_low"
	ChangeTypeBalanceExhausted       ChangeType = "balance_exhausted"
	ChangeTypeBalanceIncreased       ChangeType = "balance_increased"
	ChangeTypeLinkUnavailable        ChangeType = "link_unavailable"
	ChangeTypeDeleteLink             ChangeType = "delete_link"
	ChangeTypeQuotaExpireSoon        ChangeType = "quota_expire_soon"
	ChangeTypeOpenProfile            ChangeType = "open_profile"
	ChangeTypeFriendRequest          ChangeType = "friend_request"
	ChangeTypeCustomerStartChat      ChangeType = "customer_start_chat"
	ChangeTypeMessageFromCustomer    ChangeType = "message_from_customer"
)

type BaseEvent struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Event        string   `xml:"Event"`
	ChangeType   string   `xml:"ChangeType"`
}

type ChangeExternalContactEvent struct {
	BaseEvent
	UserID         string `xml:"UserID"`
	ExternalUserID string `xml:"ExternalUserID"`
	State          string `xml:"State"`
	WelcomeCode    string `xml:"WelcomeCode"`
	Source         string `xml:"Source"`
	FailReason     string `xml:"FailReason"`
}

type ChangeExternalChatEvent struct {
	BaseEvent
	ChatId        string `xml:"ChatId"`
	UpdateDetail  string `xml:"UpdateDetail"`
	JoinScene     int    `xml:"JoinScene"`
	QuitScene     int    `xml:"QuitScene"`
	MemChangeCnt  int    `xml:"MemChangeCnt"`
	MemChangeList struct {
		Items []string `xml:"Item"`
	} `xml:"MemChangeList"`
	LastMemVer string `xml:"LastMemVer"`
	CurMemVer  string `xml:"CurMemVer"`
}

type ChangeExternalTagEvent struct {
	BaseEvent
	Id         string `xml:"Id"`
	TagType    string `xml:"TagType"`
	StrategyId string `xml:"StrategyId"`
}

type CustomerAcquisitionEvent struct {
	BaseEvent
}

type BalanceLowEvent struct {
	CustomerAcquisitionEvent
}

type BalanceExhaustedEvent struct {
	CustomerAcquisitionEvent
}

type BalanceIncreasedEvent struct {
	CustomerAcquisitionEvent
}

type LinkUnavailableEvent struct {
	CustomerAcquisitionEvent
	LinkId string `xml:"LinkId"`
}

type DeleteLinkEvent struct {
	CustomerAcquisitionEvent
	LinkId string `xml:"LinkId"`
}

type QuotaExpireSoonEvent struct {
	CustomerAcquisitionEvent
	ExpireTime     int64 `xml:"ExpireTime"`
	ExpireQuotaNum int   `xml:"ExpireQuotaNum"`
}

type OpenProfileEvent struct {
	CustomerAcquisitionEvent
	LinkId string `xml:"LinkId"`
	State  string `xml:"State"`
}

type FriendRequestEvent struct {
	CustomerAcquisitionEvent
	LinkId string `xml:"LinkId"`
	State  string `xml:"State"`
}

type CustomerStartChatEvent struct {
	CustomerAcquisitionEvent
	UserID         string `xml:"UserID"`
	ExternalUserID string `xml:"ExternalUserID"`
	LinkId         string `xml:"LinkId"`
}

type MessageFromCustomerEvent struct {
	CustomerAcquisitionEvent
	UserID         string `xml:"UserID"`
	ExternalUserID string `xml:"ExternalUserID"`
	LinkId         string `xml:"LinkId"`
	ChatSeq        int    `xml:"ChatSeq"`
	ChatKey        string `xml:"ChatKey"`
}

type ExtAttrItem struct {
	Name string `xml:"Name"`
	Type int    `xml:"Type"`
	Text struct {
		Value string `xml:"Value"`
	} `xml:"Text"`
	Web struct {
		Title string `xml:"Title"`
		Url   string `xml:"Url"`
	} `xml:"Web"`
}

type CreateUserEvent struct {
	BaseEvent
	UserID         string `xml:"UserID"`
	Name           string `xml:"Name"`
	Department     string `xml:"Department"`
	MainDepartment string `xml:"MainDepartment"`
	IsLeaderInDept string `xml:"IsLeaderInDept"`
	DirectLeader   string `xml:"DirectLeader"`
	Position       string `xml:"Position"`
	Mobile         string `xml:"Mobile"`
	Gender         int    `xml:"Gender"`
	Email          string `xml:"Email"`
	BizMail        string `xml:"BizMail"`
	Status         int    `xml:"Status"`
	Avatar         string `xml:"Avatar"`
	Alias          string `xml:"Alias"`
	Telephone      string `xml:"Telephone"`
	Address        string `xml:"Address"`
	ExtAttr        struct {
		Items []ExtAttrItem `xml:"Item"`
	} `xml:"ExtAttr"`
}

type UpdateUserEvent struct {
	BaseEvent
	UserID         string `xml:"UserID"`
	NewUserID      string `xml:"NewUserID"`
	Name           string `xml:"Name"`
	Department     string `xml:"Department"`
	MainDepartment string `xml:"MainDepartment"`
	IsLeaderInDept string `xml:"IsLeaderInDept"`
	DirectLeader   string `xml:"DirectLeader"`
	Position       string `xml:"Position"`
	Mobile         string `xml:"Mobile"`
	Gender         int    `xml:"Gender"`
	Email          string `xml:"Email"`
	BizMail        string `xml:"BizMail"`
	Status         int    `xml:"Status"`
	Avatar         string `xml:"Avatar"`
	Alias          string `xml:"Alias"`
	Telephone      string `xml:"Telephone"`
	Address        string `xml:"Address"`
	ExtAttr        struct {
		Items []ExtAttrItem `xml:"Item"`
	} `xml:"ExtAttr"`
}

type DeleteUserEvent struct {
	BaseEvent
	UserID string `xml:"UserID"`
}

type CreatePartyEvent struct {
	BaseEvent
	Id       int    `xml:"Id"`
	Name     string `xml:"Name"`
	ParentId string `xml:"ParentId"`
	Order    int    `xml:"Order"`
}

type UpdatePartyEvent struct {
	BaseEvent
	Id       int    `xml:"Id"`
	Name     string `xml:"Name"`
	ParentId string `xml:"ParentId"`
}

type DeletePartyEvent struct {
	BaseEvent
	Id int `xml:"Id"`
}

type UpdateTagEvent struct {
	BaseEvent
	TagId         string `xml:"TagId"`
	AddUserItems  string `xml:"AddUserItems"`
	DelUserItems  string `xml:"DelUserItems"`
	AddPartyItems string `xml:"AddPartyItems"`
	DelPartyItems string `xml:"DelPartyItems"`
}

type BatchJobEvent struct {
	BaseEvent
	BatchJob struct {
		JobId   string `xml:"JobId"`
		JobType string `xml:"JobType"`
		ErrCode int    `xml:"ErrCode"`
		ErrMsg  string `xml:"ErrMsg"`
	} `xml:"BatchJob"`
}

func Parse[T any](data []byte) (*T, error) {
	var result T
	if err := xml.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse XML: %w", err)
	}
	return &result, nil
}

func ParseBase(data []byte) (*BaseEvent, error) {
	return Parse[BaseEvent](data)
}

func ParseChangeExternalContact(data []byte) (*ChangeExternalContactEvent, error) {
	return Parse[ChangeExternalContactEvent](data)
}

func ParseChangeExternalChat(data []byte) (*ChangeExternalChatEvent, error) {
	return Parse[ChangeExternalChatEvent](data)
}

func ParseChangeExternalTag(data []byte) (*ChangeExternalTagEvent, error) {
	return Parse[ChangeExternalTagEvent](data)
}

func ParseBalanceLow(data []byte) (*BalanceLowEvent, error) {
	return Parse[BalanceLowEvent](data)
}

func ParseBalanceExhausted(data []byte) (*BalanceExhaustedEvent, error) {
	return Parse[BalanceExhaustedEvent](data)
}

func ParseBalanceIncreased(data []byte) (*BalanceIncreasedEvent, error) {
	return Parse[BalanceIncreasedEvent](data)
}

func ParseLinkUnavailable(data []byte) (*LinkUnavailableEvent, error) {
	return Parse[LinkUnavailableEvent](data)
}

func ParseDeleteLink(data []byte) (*DeleteLinkEvent, error) {
	return Parse[DeleteLinkEvent](data)
}

func ParseQuotaExpireSoon(data []byte) (*QuotaExpireSoonEvent, error) {
	return Parse[QuotaExpireSoonEvent](data)
}

func ParseOpenProfile(data []byte) (*OpenProfileEvent, error) {
	return Parse[OpenProfileEvent](data)
}

func ParseFriendRequest(data []byte) (*FriendRequestEvent, error) {
	return Parse[FriendRequestEvent](data)
}

func ParseCustomerStartChat(data []byte) (*CustomerStartChatEvent, error) {
	return Parse[CustomerStartChatEvent](data)
}

func ParseMessageFromCustomer(data []byte) (*MessageFromCustomerEvent, error) {
	return Parse[MessageFromCustomerEvent](data)
}

func ParseCreateUser(data []byte) (*CreateUserEvent, error) {
	return Parse[CreateUserEvent](data)
}

func ParseUpdateUser(data []byte) (*UpdateUserEvent, error) {
	return Parse[UpdateUserEvent](data)
}

func ParseDeleteUser(data []byte) (*DeleteUserEvent, error) {
	return Parse[DeleteUserEvent](data)
}

func ParseCreateParty(data []byte) (*CreatePartyEvent, error) {
	return Parse[CreatePartyEvent](data)
}

func ParseUpdateParty(data []byte) (*UpdatePartyEvent, error) {
	return Parse[UpdatePartyEvent](data)
}

func ParseDeleteParty(data []byte) (*DeletePartyEvent, error) {
	return Parse[DeletePartyEvent](data)
}

func ParseUpdateTag(data []byte) (*UpdateTagEvent, error) {
	return Parse[UpdateTagEvent](data)
}

func ParseBatchJobResult(data []byte) (*BatchJobEvent, error) {
	return Parse[BatchJobEvent](data)
}

func ParseByEvent(data []byte) (any, error) {
	base, err := ParseBase(data)
	if err != nil {
		return nil, err
	}

	switch Event(base.Event) {
	case EventChangeContact:
		switch ChangeType(base.ChangeType) {
		case ChangeTypeCreateUser:
			return ParseCreateUser(data)
		case ChangeTypeUpdateUser:
			return ParseUpdateUser(data)
		case ChangeTypeDeleteUser:
			return ParseDeleteUser(data)
		case ChangeTypeCreateParty:
			return ParseCreateParty(data)
		case ChangeTypeUpdateParty:
			return ParseUpdateParty(data)
		case ChangeTypeDeleteParty:
			return ParseDeleteParty(data)
		case ChangeTypeUpdateTag:
			return ParseUpdateTag(data)
		default:
			return base, nil
		}
	case EventChangeExternalContact:
		return ParseChangeExternalContact(data)
	case EventChangeExternalChat:
		return ParseChangeExternalChat(data)
	case EventChangeExternalTag:
		return ParseChangeExternalTag(data)
	case EventCustomerAcquisition:
		switch ChangeType(base.ChangeType) {
		case ChangeTypeBalanceLow:
			return ParseBalanceLow(data)
		case ChangeTypeBalanceExhausted:
			return ParseBalanceExhausted(data)
		case ChangeTypeBalanceIncreased:
			return ParseBalanceIncreased(data)
		case ChangeTypeLinkUnavailable:
			return ParseLinkUnavailable(data)
		case ChangeTypeDeleteLink:
			return ParseDeleteLink(data)
		case ChangeTypeQuotaExpireSoon:
			return ParseQuotaExpireSoon(data)
		case ChangeTypeOpenProfile:
			return ParseOpenProfile(data)
		case ChangeTypeFriendRequest:
			return ParseFriendRequest(data)
		case ChangeTypeCustomerStartChat:
			return ParseCustomerStartChat(data)
		case ChangeTypeMessageFromCustomer:
			return ParseMessageFromCustomer(data)
		default:
			return Parse[CustomerAcquisitionEvent](data)
		}
	case EventBatchJobResult:
		return ParseBatchJobResult(data)
	default:
		return base, nil
	}
}

func FormatXML(data []byte) string {
	decoder := xml.NewDecoder(bytes.NewReader(data))
	var buf bytes.Buffer
	buf.WriteString("```xml\n")
	encoder := xml.NewEncoder(&buf)
	encoder.Indent("", "  ")
	for {
		token, err := decoder.Token()
		if err != nil {
			break
		}
		encoder.EncodeToken(token)
	}
	encoder.Flush()
	buf.WriteString("\n```")
	return buf.String()
}
