package callback

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

type Event string

const (
	MsgTypeEvent = "event"

	EventChangeContact               = "change_contact"
	EventChangeExternalContact       = "change_external_contact"
	EventChangeExternalChat          = "change_external_chat"
	EventChangeExternalTag           = "change_external_tag"
	EventCustomerAcquisition         = "customer_acquisition"
	EventBatchJobResult              = "batch_job_result"
	EventSecurity                    = "security"
	EventKFAccountAuthChange         = "kf_account_auth_change"
	EventDocChange                   = "doc_change"
	EventSmartSheetChange            = "smart_sheet_change"
	EventDeleteCalendar              = "delete_calendar"
	EventModifyCalendar              = "modify_calendar"
	EventModifySchedule              = "modify_schedule"
	EventDeleteSchedule              = "delete_schedule"
	EventRespondSchedule             = "respond_schedule"
	EventMeetingChange               = "meeting_change"
	EventMeetingStatistics           = "meeting_statistics"
	EventWeDriveSpaceChange          = "wedrive_space_change"
	EventWeDriveFileChange           = "wedrive_file_change"
	EventWeDriveInsufficientCapacity = "wedrive_insufficient_capacity"
)

type ChangeType string

const (
	ChangeTypeCreateUser                  ChangeType = "create_user"
	ChangeTypeUpdateUser                  ChangeType = "update_user"
	ChangeTypeDeleteUser                  ChangeType = "delete_user"
	ChangeTypeCreateParty                 ChangeType = "create_party"
	ChangeTypeUpdateParty                 ChangeType = "update_party"
	ChangeTypeDeleteParty                 ChangeType = "delete_party"
	ChangeTypeUpdateTag                   ChangeType = "update_tag"
	ChangeTypeAddExternalContact          ChangeType = "add_external_contact"
	ChangeTypeEditExternalContact         ChangeType = "edit_external_contact"
	ChangeTypeAddHalfExternalContact      ChangeType = "add_half_external_contact"
	ChangeTypeDelExternalContact          ChangeType = "del_external_contact"
	ChangeTypeDelFollowUser               ChangeType = "del_follow_user"
	ChangeTypeTransferFail                ChangeType = "transfer_fail"
	ChangeTypeCreateChat                  ChangeType = "create"
	ChangeTypeUpdateChat                  ChangeType = "update"
	ChangeTypeDismissChat                 ChangeType = "dismiss"
	ChangeTypeBalanceLow                  ChangeType = "balance_low"
	ChangeTypeBalanceExhausted            ChangeType = "balance_exhausted"
	ChangeTypeBalanceIncreased            ChangeType = "balance_increased"
	ChangeTypeLinkUnavailable             ChangeType = "link_unavailable"
	ChangeTypeDeleteLink                  ChangeType = "delete_link"
	ChangeTypeQuotaExpireSoon             ChangeType = "quota_expire_soon"
	ChangeTypeOpenProfile                 ChangeType = "open_profile"
	ChangeTypeFriendRequest               ChangeType = "friend_request"
	ChangeTypeCustomerStartChat           ChangeType = "customer_start_chat"
	ChangeTypeMessageFromCustomer         ChangeType = "message_from_customer"
	ChangeTypeChangeDomainIP              ChangeType = "change_domain_ip"
	ChangeTypeDocMemberChange             ChangeType = "doc_member_change"
	ChangeTypeAddField                    ChangeType = "add_field"
	ChangeTypeUpdateField                 ChangeType = "update_field"
	ChangeTypeDeleteField                 ChangeType = "delete_field"
	ChangeTypeDeleteDoc                   ChangeType = "delete_doc"
	ChangeTypeFormComplete                ChangeType = "form_complete"
	ChangeTypeDeleteForm                  ChangeType = "delete_form"
	ChangeTypeFormSettingsChange          ChangeType = "form_settings_change"
	ChangeTypeAddRecord                   ChangeType = "add_record"
	ChangeTypeUpdateRecord                ChangeType = "update_record"
	ChangeTypeDeleteRecord                ChangeType = "delete_record"
	ChangeTypeJoinMeeting                 ChangeType = "join_meeting"
	ChangeTypeQuitMeeting                 ChangeType = "quit_meeting"
	ChangeTypeJoinWaitingRoom             ChangeType = "join_waiting_room"
	ChangeTypeQuitWaitingRoom             ChangeType = "quit_waiting_room"
	ChangeTypeJoinFromMeetingRoom         ChangeType = "join_from_meeting_room"
	ChangeTypeMoveToWaitingRoom           ChangeType = "move_to_waiting_room"
	ChangeTypeJoinMeetingBeforeHost       ChangeType = "join_meeting_before_host"
	ChangeTypeRoleChange                  ChangeType = "role_change"
	ChangeTypeMeetingEnd                  ChangeType = "meeting_end"
	ChangeTypeOpenScreenShare             ChangeType = "open_screen_share"
	ChangeTypeCloseScreenShare            ChangeType = "close_screen_share"
	ChangeTypeResumeRecording             ChangeType = "resume_recording"
	ChangeTypeStartMeeting                ChangeType = "start_meeting"
	ChangeTypeMeetingMuteAll              ChangeType = "meeting_mute_all"
	ChangeTypeMeetingRoomResponse         ChangeType = "meeting_room_response"
	ChangeTypeMeetingStart                ChangeType = "meeting_start"
	ChangeTypeMeetingUnmuteAll            ChangeType = "meeting_unmute_all"
	ChangeTypeDeleteRecording             ChangeType = "delete_recording"
	ChangeTypeCancelMeeting               ChangeType = "cancel_meeting"
	ChangeTypeStartRecording              ChangeType = "start_recording"
	ChangeTypeMediumUpload                ChangeType = "medium_upload"
	ChangeTypePstnStatusUpdate            ChangeType = "pstn_status_update"
	ChangeTypeRecordingComplete           ChangeType = "recording_complete"
	ChangeTypeModifyMeeting               ChangeType = "modify_meeting"
	ChangeTypeStopRecording               ChangeType = "stop_recording"
	ChangeTypePauseRecording              ChangeType = "pause_recording"
	ChangeTypeCancelEnroll                ChangeType = "cancel_enroll"
	ChangeTypeEnroll                      ChangeType = "enroll"
	ChangeTypeWebinarRoleChange           ChangeType = "webinar_role_change"
	ChangeTypeWebinarWarmUpUpload         ChangeType = "webinar_warm_up_upload"
	ChangeTypeSpaceSecuritySettingsChange ChangeType = "space_security_settings_change"
	ChangeTypeSpaceMemberChange           ChangeType = "space_member_change"
	ChangeTypeDismissSpace                ChangeType = "dismiss_space"
	ChangeTypeCreateFile                  ChangeType = "create_file"
	ChangeTypeRenameFile                  ChangeType = "rename_file"
	ChangeTypeUpdateFile                  ChangeType = "update_file"
	ChangeTypeDeleteFile                  ChangeType = "delete_file"
	ChangeTypeMoveFile                    ChangeType = "move_file"
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
