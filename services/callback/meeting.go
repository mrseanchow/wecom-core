package callback

type MeetingUser struct {
	UserId    string `xml:"UserId"`
	TmpOpenId string `xml:"TmpOpenId"`
}

type JoinMeetingEvent struct {
	BaseEvent
	FromUserTmpOpenId string `xml:"FromUserTmpOpenId"`
	MeetingId         string `xml:"MeetingId"`
}

type QuitMeetingEvent struct {
	BaseEvent
	FromUserTmpOpenId string `xml:"FromUserTmpOpenId"`
	MeetingId         string `xml:"MeetingId"`
}

type JoinWaitingRoomEvent struct {
	BaseEvent
	FromUserTmpOpenId string `xml:"FromUserTmpOpenId"`
	MeetingId         string `xml:"MeetingId"`
}

type QuitWaitingRoomEvent struct {
	BaseEvent
	FromUserTmpOpenId string      `xml:"FromUserTmpOpenId"`
	OperatedUser      MeetingUser `xml:"OperatedUser"`
	MeetingId         string      `xml:"MeetingId"`
}

type JoinFromMeetingRoomEvent struct {
	BaseEvent
	FromUserTmpOpenId string      `xml:"FromUserTmpOpenId"`
	OperatedUser      MeetingUser `xml:"OperatedUser"`
	MeetingId         string      `xml:"MeetingId"`
}

type MoveToWaitingRoomEvent struct {
	BaseEvent
	FromUserTmpOpenId string      `xml:"FromUserTmpOpenId"`
	OperatedUser      MeetingUser `xml:"OperatedUser"`
	MeetingId         string      `xml:"MeetingId"`
}

type JoinMeetingBeforeHostEvent struct {
	BaseEvent
	FromUserTmpOpenId string `xml:"FromUserTmpOpenId"`
	MeetingId         string `xml:"MeetingId"`
}

type MeetingRoleUser struct {
	UserId    string `xml:"UserId"`
	TmpOpenId string `xml:"TmpOpenId"`
	UserRole  int    `xml:"UserRole"`
}

type RoleChangeEvent struct {
	BaseEvent
	FromUserTmpOpenId string          `xml:"FromUserTmpOpenId"`
	OperatedUser      MeetingRoleUser `xml:"OperatedUser"`
	MeetingId         string          `xml:"MeetingId"`
}

type MeetingEndEvent struct {
	BaseEvent
	FromUserTmpOpenId string `xml:"FromUserTmpOpenId"`
	MeetingId         string `xml:"MeetingId"`
}

type OpenScreenShareEvent struct {
	BaseEvent
	FromUserTmpOpenId string `xml:"FromUserTmpOpenId"`
	MeetingId         string `xml:"MeetingId"`
}

type CloseScreenShareEvent struct {
	BaseEvent
	FromUserTmpOpenId string `xml:"FromUserTmpOpenId"`
	MeetingId         string `xml:"MeetingId"`
}

type ResumeRecordingEvent struct {
	BaseEvent
	FromUserTmpOpenId string `xml:"FromUserTmpOpenId"`
	MeetingId         string `xml:"MeetingId"`
}

type StartMeetingEvent struct {
	BaseEvent
	Status int `xml:"Status"`
}

type MeetingMraAddress struct {
	Protocol   int    `xml:"Protocol"`
	DialString string `xml:"DialString"`
}

type MeetingMuteAllEvent struct {
	BaseEvent
	MeetingId string `xml:"MeetingId"`
}

type MeetingRoomResponseEvent struct {
	BaseEvent
	MeetingId          string            `xml:"MeetingId"`
	MeetingRoomId      string            `xml:"MeetingRoomId"`
	MraAddress         MeetingMraAddress `xml:"MraAddress"`
	RoomResponseStatus int               `xml:"RoomResponseStatus"`
}

type MeetingStartEvent struct {
	BaseEvent
	MeetingId string `xml:"MeetingId"`
}

type MeetingUnmuteAllEvent struct {
	BaseEvent
	MeetingId string `xml:"MeetingId"`
}

type DeleteRecordingEvent struct {
	BaseEvent
	MeetingId string `xml:"MeetingId"`
}

type CancelMeetingEvent struct {
	BaseEvent
	MeetingId string `xml:"MeetingId"`
}

type StartRecordingEvent struct {
	BaseEvent
	MeetingId string `xml:"MeetingId"`
}

type MediumUploadInfo struct {
	MediumUrl    string `xml:"MediumUrl"`
	MediumType   int    `xml:"MediumType"`
	UploadStatus int    `xml:"UploadStatus"`
	ErrorMsg     string `xml:"ErrorMsg"`
}

type MediumUploadEvent struct {
	BaseEvent
	MeetingId       string             `xml:"MeetingId"`
	AllUploadStatus bool               `xml:"AllUploadStatus"`
	UploadInfo      []MediumUploadInfo `xml:"UploadInfo"`
}

type PstnStatusUpdateEvent struct {
	BaseEvent
	MeetingId  string `xml:"MeetingId"`
	PstnStatus string `xml:"PstnStatus"`
}

type RecordingCompleteEvent struct {
	BaseEvent
	MeetingId string `xml:"MeetingId"`
}

type ModifyMeetingEvent struct {
	BaseEvent
	MeetingId string `xml:"MeetingId"`
}

type StopRecordingEvent struct {
	BaseEvent
	MeetingId string `xml:"MeetingId"`
}

type PauseRecordingEvent struct {
	BaseEvent
	MeetingId string `xml:"MeetingId"`
}

type CancelEnrollEvent struct {
	BaseEvent
	MeetingId string `xml:"MeetingId"`
	EnrollId  string `xml:"EnrollId"`
}

type EnrollEvent struct {
	BaseEvent
	MeetingId string `xml:"MeetingId"`
	EnrollId  string `xml:"EnrollId"`
}

type WebinarRoleChangeEvent struct {
	BaseEvent
	FromUserTmpOpenId string          `xml:"FromUserTmpOpenId"`
	OperatedUser      MeetingRoleUser `xml:"OperatedUser"`
	MeetingId         string          `xml:"MeetingId"`
}

type WarmUpInfo struct {
	WarmUpPicture string `xml:"WarmUpPicture"`
	WarmUpVideo   string `xml:"WarmUpVideo"`
	UploadStatus  int    `xml:"UploadStatus"`
	ErrorMsg      string `xml:"ErrorMsg"`
}

type WebinarWarmUpUploadEvent struct {
	BaseEvent
	MeetingId  string     `xml:"MeetingId"`
	WarmUpInfo WarmUpInfo `xml:"WarmUpInfo"`
}

func ParseJoinMeeting(data []byte) (*JoinMeetingEvent, error) {
	return Parse[JoinMeetingEvent](data)
}

func ParseQuitMeeting(data []byte) (*QuitMeetingEvent, error) {
	return Parse[QuitMeetingEvent](data)
}

func ParseJoinWaitingRoom(data []byte) (*JoinWaitingRoomEvent, error) {
	return Parse[JoinWaitingRoomEvent](data)
}

func ParseQuitWaitingRoom(data []byte) (*QuitWaitingRoomEvent, error) {
	return Parse[QuitWaitingRoomEvent](data)
}

func ParseJoinFromMeetingRoom(data []byte) (*JoinFromMeetingRoomEvent, error) {
	return Parse[JoinFromMeetingRoomEvent](data)
}

func ParseMoveToWaitingRoom(data []byte) (*MoveToWaitingRoomEvent, error) {
	return Parse[MoveToWaitingRoomEvent](data)
}

func ParseJoinMeetingBeforeHost(data []byte) (*JoinMeetingBeforeHostEvent, error) {
	return Parse[JoinMeetingBeforeHostEvent](data)
}

func ParseRoleChange(data []byte) (*RoleChangeEvent, error) {
	return Parse[RoleChangeEvent](data)
}

func ParseMeetingEnd(data []byte) (*MeetingEndEvent, error) {
	return Parse[MeetingEndEvent](data)
}

func ParseOpenScreenShare(data []byte) (*OpenScreenShareEvent, error) {
	return Parse[OpenScreenShareEvent](data)
}

func ParseCloseScreenShare(data []byte) (*CloseScreenShareEvent, error) {
	return Parse[CloseScreenShareEvent](data)
}

func ParseResumeRecording(data []byte) (*ResumeRecordingEvent, error) {
	return Parse[ResumeRecordingEvent](data)
}

func ParseStartMeeting(data []byte) (*StartMeetingEvent, error) {
	return Parse[StartMeetingEvent](data)
}

func ParseMeetingMuteAll(data []byte) (*MeetingMuteAllEvent, error) {
	return Parse[MeetingMuteAllEvent](data)
}

func ParseMeetingRoomResponse(data []byte) (*MeetingRoomResponseEvent, error) {
	return Parse[MeetingRoomResponseEvent](data)
}

func ParseMeetingStart(data []byte) (*MeetingStartEvent, error) {
	return Parse[MeetingStartEvent](data)
}

func ParseMeetingUnmuteAll(data []byte) (*MeetingUnmuteAllEvent, error) {
	return Parse[MeetingUnmuteAllEvent](data)
}

func ParseDeleteRecording(data []byte) (*DeleteRecordingEvent, error) {
	return Parse[DeleteRecordingEvent](data)
}

func ParseCancelMeeting(data []byte) (*CancelMeetingEvent, error) {
	return Parse[CancelMeetingEvent](data)
}

func ParseStartRecording(data []byte) (*StartRecordingEvent, error) {
	return Parse[StartRecordingEvent](data)
}

func ParseMediumUpload(data []byte) (*MediumUploadEvent, error) {
	return Parse[MediumUploadEvent](data)
}

func ParsePstnStatusUpdate(data []byte) (*PstnStatusUpdateEvent, error) {
	return Parse[PstnStatusUpdateEvent](data)
}

func ParseRecordingComplete(data []byte) (*RecordingCompleteEvent, error) {
	return Parse[RecordingCompleteEvent](data)
}

func ParseModifyMeeting(data []byte) (*ModifyMeetingEvent, error) {
	return Parse[ModifyMeetingEvent](data)
}

func ParseStopRecording(data []byte) (*StopRecordingEvent, error) {
	return Parse[StopRecordingEvent](data)
}

func ParsePauseRecording(data []byte) (*PauseRecordingEvent, error) {
	return Parse[PauseRecordingEvent](data)
}

func ParseCancelEnroll(data []byte) (*CancelEnrollEvent, error) {
	return Parse[CancelEnrollEvent](data)
}

func ParseEnroll(data []byte) (*EnrollEvent, error) {
	return Parse[EnrollEvent](data)
}

func ParseWebinarRoleChange(data []byte) (*WebinarRoleChangeEvent, error) {
	return Parse[WebinarRoleChangeEvent](data)
}

func ParseWebinarWarmUpUpload(data []byte) (*WebinarWarmUpUploadEvent, error) {
	return Parse[WebinarWarmUpUploadEvent](data)
}

