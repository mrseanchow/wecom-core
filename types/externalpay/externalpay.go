package externalpay

import "github.com/shuaidd/wecom-core/types/common"

type UploadImageResponse struct {
	common.Response
	OpenWxPayMediaID string `json:"open_wx_pay_media_id"`
}

type GetMerchantRequest struct {
	MchID string `json:"mch_id"`
}

type GetMerchantResponse struct {
	common.Response
	BindStatus    int            `json:"bind_status"`
	MchID         string         `json:"mch_id"`
	MerchantName  string         `json:"merchant_name"`
	AllowUseScope *AllowUseScope `json:"allow_use_scope,omitempty"`
}

type AllowUseScope struct {
	User    []string `json:"user,omitempty"`
	PartyID []int64  `json:"partyid,omitempty"`
	TagID   []int64  `json:"tagid,omitempty"`
}

type SetMchUseScopeRequest struct {
	MchID         string         `json:"mch_id"`
	AllowUseScope *AllowUseScope `json:"allow_use_scope"`
}

type SetMchUseScopeResponse struct {
	common.Response
}

type GetFundFlowRequest struct {
	BeginTime int64  `json:"begin_time"`
	EndTime   int64  `json:"end_time"`
	MchID     string `json:"mch_id,omitempty"`
	Cursor    string `json:"cursor,omitempty"`
	Limit     int    `json:"limit,omitempty"`
}

type GetFundFlowResponse struct {
	common.Response
	NextCursor   string     `json:"next_cursor,omitempty"`
	FundFlowList []FundFlow `json:"fund_flow_list"`
}

type FundFlow struct {
	Timestamp         int64   `json:"timestamp"`
	RequestNo         string  `json:"request_no"`
	TransactionType   int     `json:"transaction_type"`
	FundFlowType      int     `json:"fund_flow_type"`
	TransactionAmount int     `json:"transaction_amount"`
	AccountBalance    int     `json:"account_balance"`
	OutTradeNo        string  `json:"out_trade_no"`
	MchID             string  `json:"mch_id"`
	OperatorUserid    string  `json:"operator_userid"`
	GroupList         []Group `json:"group_list"`
	Remark            string  `json:"remark"`
}

type Group struct {
	GroupName string `json:"group_name"`
}

type GetBillListRequest struct {
	BeginTime   int64  `json:"begin_time"`
	EndTime     int64  `json:"end_time"`
	PayeeUserid string `json:"payee_userid,omitempty"`
	Cursor      string `json:"cursor,omitempty"`
	Limit       int    `json:"limit,omitempty"`
}

type GetBillListResponse struct {
	common.Response
	NextCursor string `json:"next_cursor,omitempty"`
	BillList   []Bill `json:"bill_list"`
}

type Bill struct {
	TransactionID   string           `json:"transaction_id"`
	BillType        int              `json:"bill_type"`
	TradeState      int              `json:"trade_state,omitempty"`
	PayTime         int64            `json:"pay_time"`
	OutTradeNo      string           `json:"out_trade_no"`
	OutRefundNo     string           `json:"out_refund_no,omitempty"`
	ExternalUserid  string           `json:"external_userid"`
	TotalFee        int              `json:"total_fee"`
	PayeeUserid     string           `json:"payee_userid"`
	PaymentType     int              `json:"payment_type"`
	MchID           string           `json:"mch_id"`
	Remark          string           `json:"remark"`
	CommodityList   []Commodity      `json:"commodity_list,omitempty"`
	TotalRefundFee  int              `json:"total_refund_fee,omitempty"`
	RefundList      []Refund         `json:"refund_list,omitempty"`
	ContactInfo     *ContactInfo     `json:"contact_info,omitempty"`
	MiniprogramInfo *MiniprogramInfo `json:"miniprogram_info,omitempty"`
}

type Commodity struct {
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}

type Refund struct {
	OutRefundNo   string `json:"out_refund_no"`
	RefundUserid  string `json:"refund_userid"`
	RefundComment string `json:"refund_comment"`
	RefundReqtime int64  `json:"refund_reqtime"`
	RefundStatus  int    `json:"refund_status"`
	RefundFee     int    `json:"refund_fee"`
}

type ContactInfo struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type MiniprogramInfo struct {
	AppID string `json:"appid"`
	Name  string `json:"name"`
}

type GetPaymentInfoRequest struct {
	PaymentID string `json:"payment_id"`
}

type GetPaymentInfoResponse struct {
	common.Response
	BillList []Trade `json:"bill_list"`
}

type Trade struct {
	OutTradeNo string `json:"out_trade_no"`
}

type BusinessLicenseInfo struct {
	CertType                            string `json:"cert_type,omitempty"`
	BusinessLicenseCopyOpenWxPayMediaID string `json:"business_license_copy_open_wx_pay_media_id"`
	BusinessLicenseNumber               string `json:"business_license_number"`
	MerchantName                        string `json:"merchant_name"`
	LegalPerson                         string `json:"legal_person"`
	CompanyAddress                      string `json:"company_address,omitempty"`
	BusinessTimeBeginTime               string `json:"business_time_begin_time,omitempty"`
	BusinessTimeEndTime                 string `json:"business_time_end_time,omitempty"`
}

type FinanceInstitutionInfo struct {
	FinanceType                        string   `json:"finance_type"`
	FinanceLicensePicsOpenWxPayMediaID []string `json:"finance_license_pics_open_wx_pay_media_id"`
}

type IDCardInfo struct {
	IDCardCopyOpenWxPayMediaID     string `json:"id_card_copy_open_wx_pay_media_id"`
	IDCardNationalOpenWxPayMediaID string `json:"id_card_national_open_wx_pay_media_id"`
	IDCardName                     string `json:"id_card_name"`
	IDCardNumber                   string `json:"id_card_number"`
	IDCardAddress                  string `json:"id_card_address,omitempty"`
	IDCardValidTimeBegin           string `json:"id_card_valid_time_begin"`
	IDCardValidTime                string `json:"id_card_valid_time"`
	IDDocType                      int    `json:"id_doc_type"`
}

type ContactInfoReq struct {
	ContactType                                 string      `json:"contact_type"`
	ContactInfo                                 *IDCardInfo `json:"contact_info"`
	BusinessAuthorizationLetterOpenWxPayMediaID string      `json:"business_authorization_letter_open_wx_pay_media_id,omitempty"`
	MobilePhone                                 string      `json:"mobile_phone"`
	ContactEmail                                string      `json:"contact_email"`
}

type AccountInfo struct {
	BankAccountType    int                 `json:"bank_account_type,omitempty"`
	AccountBank        string              `json:"account_bank"`
	AccountName        string              `json:"account_name"`
	BankAddressCode    string              `json:"bank_address_code"`
	BankName           string              `json:"bank_name,omitempty"`
	AccountNumber      string              `json:"account_number"`
	BankCardSupplement *BankCardSupplement `json:"bank_card_supplement,omitempty"`
}

type BankCardSupplement struct {
	SettlementCertificateOpenWxPayMediaID   string   `json:"settlement_certificate_open_wx_pay_media_id,omitempty"`
	RelationshipCertificateOpenWxPayMediaID string   `json:"relationship_certificate_open_wx_pay_media_id,omitempty"`
	OtherCertificateOpenWxPayMediaID        []string `json:"other_certificate_open_wx_pay_media_id,omitempty"`
}

type SalesSceneInfo struct {
	Type                        int    `json:"type"`
	StoreURL                    string `json:"store_url,omitempty"`
	StorePicOpenWxPayMediaID    string `json:"store_pic_open_wx_pay_media_id,omitempty"`
	AddressCode                 string `json:"address_code,omitempty"`
	OfflineAddress              string `json:"offline_address,omitempty"`
	EntrancePicOpenWxPayMediaID string `json:"entrance_pic_open_wx_pay_media_id,omitempty"`
	IndoorPicOpenWxPayMediaID   string `json:"indoor_pic_open_wx_pay_media_id,omitempty"`
}

type Qualifications struct {
	ID []string `json:"id"`
}

type BusinessAdditionPics struct {
	ID []string `json:"id"`
}

type ApplyMchRequest struct {
	OutRequestNo           string                  `json:"out_request_no"`
	OrganizationType       int                     `json:"organization_type"`
	BusinessLicenseInfo    *BusinessLicenseInfo    `json:"business_license_info"`
	FinanceInstitutionInfo *FinanceInstitutionInfo `json:"finance_institution_info,omitempty"`
	MerchantShortName      string                  `json:"merchant_short_name"`
	IDCardInfo             *IDCardInfo             `json:"id_card_info"`
	Owner                  bool                    `json:"owner,omitempty"`
	UboInfo                *IDCardInfo             `json:"ubo_info,omitempty"`
	ContactInfo            *ContactInfoReq         `json:"contact_info"`
	AccountInfo            *AccountInfo            `json:"account_info"`
	SalesSceneInfo         *SalesSceneInfo         `json:"sales_scene_info"`
	BusinessID             int                     `json:"business_id"`
	Qualifications         *Qualifications         `json:"qualifications,omitempty"`
	BusinessAdditionPics   *BusinessAdditionPics   `json:"business_addition_pics,omitempty"`
	UserID                 string                  `json:"userid"`
}

type ApplyMchResponse struct {
	common.Response
	ErrorDescription string `json:"error_description,omitempty"`
}

type GetApplymentStatusRequest struct {
	OutRequestNo string `json:"out_request_no"`
}

type Status struct {
	ApplymentState     string             `json:"applyment_state"`
	ApplymentStateDesc string             `json:"applyment_state_desc"`
	SignState          string             `json:"sign_state"`
	SignURL            string             `json:"sign_url,omitempty"`
	SubMchid           string             `json:"sub_mchid,omitempty"`
	AuditDetail        []AuditDetail      `json:"audit_detail,omitempty"`
	AccountValidation  *AccountValidation `json:"account_validation,omitempty"`
	LegalValidationURL string             `json:"legal_validation_url,omitempty"`
}

type AuditDetail struct {
	ParamName    string `json:"param_name"`
	RejectReason string `json:"reject_reason"`
}

type AccountValidation struct {
	AccountName              string `json:"account_name"`
	AccountNo                string `json:"account_no"`
	PayAmount                int    `json:"pay_amount"`
	DestinationAccountNumber string `json:"destination_account_number"`
	DestinationAccountName   string `json:"destination_account_name"`
	DestinationAccountBank   string `json:"destination_account_bank"`
	City                     string `json:"city"`
	Remark                   string `json:"remark"`
	Deadline                 string `json:"deadline"`
}

type GetApplymentStatusResponse struct {
	common.Response
	Status        *Status `json:"status,omitempty"`
	ApplyState    int     `json:"apply_state"`
	RealSignState int     `json:"real_sign_state"`
	RejectReason  string  `json:"reject_reason,omitempty"`
}
