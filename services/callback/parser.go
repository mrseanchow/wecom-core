package callback

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
	case EventSecurity:
		return ParseSecurity(data)
	case EventKFAccountAuthChange:
		return ParseKFAccountAuthChange(data)
	case EventDocChange:
		switch ChangeType(base.ChangeType) {
		case ChangeTypeDocMemberChange:
			return ParseDocMemberChange(data)
		case ChangeTypeDeleteDoc:
			return ParseDeleteDoc(data)
		case ChangeTypeFormComplete:
			return ParseFormComplete(data)
		case ChangeTypeDeleteForm:
			return ParseDeleteForm(data)
		case ChangeTypeFormSettingsChange:
			return ParseFormSettingsChange(data)
		default:
			return base, nil
		}
	case EventSmartSheetChange:
		switch ChangeType(base.ChangeType) {
		case ChangeTypeAddField:
			return ParseAddField(data)
		case ChangeTypeUpdateField:
			return ParseUpdateField(data)
		case ChangeTypeDeleteField:
			return ParseDeleteField(data)
		case ChangeTypeAddRecord:
			return ParseAddRecord(data)
		case ChangeTypeUpdateRecord:
			return ParseUpdateRecord(data)
		case ChangeTypeDeleteRecord:
			return ParseDeleteRecord(data)
		default:
			return base, nil
		}
	case EventDeleteCalendar:
		return ParseDeleteCalendar(data)
	case EventModifyCalendar:
		return ParseModifyCalendar(data)
	case EventModifySchedule:
		return ParseModifySchedule(data)
	case EventDeleteSchedule:
		return ParseDeleteSchedule(data)
	case EventRespondSchedule:
		return ParseRespondSchedule(data)
	case EventMeetingChange:
		switch ChangeType(base.ChangeType) {
		case ChangeTypeJoinMeeting:
			return ParseJoinMeeting(data)
		case ChangeTypeQuitMeeting:
			return ParseQuitMeeting(data)
		case ChangeTypeJoinWaitingRoom:
			return ParseJoinWaitingRoom(data)
		case ChangeTypeQuitWaitingRoom:
			return ParseQuitWaitingRoom(data)
		case ChangeTypeJoinFromMeetingRoom:
			return ParseJoinFromMeetingRoom(data)
		case ChangeTypeMoveToWaitingRoom:
			return ParseMoveToWaitingRoom(data)
		case ChangeTypeJoinMeetingBeforeHost:
			return ParseJoinMeetingBeforeHost(data)
		case ChangeTypeRoleChange:
			return ParseRoleChange(data)
		case ChangeTypeMeetingEnd:
			return ParseMeetingEnd(data)
		case ChangeTypeOpenScreenShare:
			return ParseOpenScreenShare(data)
		case ChangeTypeCloseScreenShare:
			return ParseCloseScreenShare(data)
		case ChangeTypeResumeRecording:
			return ParseResumeRecording(data)
		case ChangeTypeMeetingMuteAll:
			return ParseMeetingMuteAll(data)
		case ChangeTypeMeetingRoomResponse:
			return ParseMeetingRoomResponse(data)
		case ChangeTypeMeetingStart:
			return ParseMeetingStart(data)
		case ChangeTypeMeetingUnmuteAll:
			return ParseMeetingUnmuteAll(data)
		case ChangeTypeDeleteRecording:
			return ParseDeleteRecording(data)
		case ChangeTypeCancelMeeting:
			return ParseCancelMeeting(data)
		case ChangeTypeStartRecording:
			return ParseStartRecording(data)
		case ChangeTypeMediumUpload:
			return ParseMediumUpload(data)
		case ChangeTypePstnStatusUpdate:
			return ParsePstnStatusUpdate(data)
		case ChangeTypeRecordingComplete:
			return ParseRecordingComplete(data)
		case ChangeTypeModifyMeeting:
			return ParseModifyMeeting(data)
		case ChangeTypeStopRecording:
			return ParseStopRecording(data)
		case ChangeTypePauseRecording:
			return ParsePauseRecording(data)
		case ChangeTypeCancelEnroll:
			return ParseCancelEnroll(data)
		case ChangeTypeEnroll:
			return ParseEnroll(data)
		case ChangeTypeWebinarRoleChange:
			return ParseWebinarRoleChange(data)
		case ChangeTypeWebinarWarmUpUpload:
			return ParseWebinarWarmUpUpload(data)
		default:
			return base, nil
		}
	case EventMeetingStatistics:
		switch ChangeType(base.ChangeType) {
		case ChangeTypeStartMeeting:
			return ParseStartMeeting(data)
		default:
			return base, nil
		}
	case EventWeDriveSpaceChange:
		switch ChangeType(base.ChangeType) {
		case ChangeTypeSpaceSecuritySettingsChange:
			return ParseSpaceSecuritySettingsChange(data)
		case ChangeTypeSpaceMemberChange:
			return ParseSpaceMemberChange(data)
		case ChangeTypeDismissSpace:
			return ParseDismissSpace(data)
		default:
			return base, nil
		}
	case EventWeDriveFileChange:
		switch ChangeType(base.ChangeType) {
		case ChangeTypeCreateFile:
			return ParseCreateFile(data)
		case ChangeTypeRenameFile:
			return ParseRenameFile(data)
		case ChangeTypeUpdateFile:
			return ParseUpdateFile(data)
		case ChangeTypeDeleteFile:
			return ParseDeleteFile(data)
		case ChangeTypeMoveFile:
			return ParseMoveFile(data)
		default:
			return base, nil
		}
	case EventWeDriveInsufficientCapacity:
		return ParseWeDriveInsufficientCapacity(data)
	default:
		return base, nil
	}
}
