package callback

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
