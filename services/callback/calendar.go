package callback

type DeleteCalendarEvent struct {
	BaseEvent
	CalId string `xml:"CalId"`
}

type ModifyCalendarEvent struct {
	BaseEvent
	CalId string `xml:"CalId"`
}

type ModifyScheduleEvent struct {
	BaseEvent
	CalId      string `xml:"CalId"`
	ScheduleId string `xml:"ScheduleId"`
}

type DeleteScheduleEvent struct {
	BaseEvent
	CalId      string `xml:"CalId"`
	ScheduleId string `xml:"ScheduleId"`
}

type RespondScheduleEvent struct {
	BaseEvent
	CalId      string `xml:"CalId"`
	ScheduleId string `xml:"ScheduleId"`
}

func ParseDeleteCalendar(data []byte) (*DeleteCalendarEvent, error) {
	return Parse[DeleteCalendarEvent](data)
}

func ParseModifyCalendar(data []byte) (*ModifyCalendarEvent, error) {
	return Parse[ModifyCalendarEvent](data)
}

func ParseModifySchedule(data []byte) (*ModifyScheduleEvent, error) {
	return Parse[ModifyScheduleEvent](data)
}

func ParseDeleteSchedule(data []byte) (*DeleteScheduleEvent, error) {
	return Parse[DeleteScheduleEvent](data)
}

func ParseRespondSchedule(data []byte) (*RespondScheduleEvent, error) {
	return Parse[RespondScheduleEvent](data)
}
