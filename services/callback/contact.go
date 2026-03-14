package callback

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

