package callback

import (
	"fmt"
	"testing"
)

func TestParseBaseEvent(t *testing.T) {
	xmlData := []byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[sys]]></FromUserName> 
		<CreateTime>1403610513</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[change_external_contact]]></Event>
		<ChangeType><![CDATA[add_external_contact]]></ChangeType>
	</xml>`)

	event, err := ParseBase(xmlData)
	if err != nil {
		t.Fatalf("failed to parse: %v", err)
	}

	if event.ToUserName != "toUser" {
		t.Errorf("expected ToUserName to be 'toUser', got '%s'", event.ToUserName)
	}
	if event.FromUserName != "sys" {
		t.Errorf("expected FromUserName to be 'sys', got '%s'", event.FromUserName)
	}
	if event.CreateTime != 1403610513 {
		t.Errorf("expected CreateTime to be 1403610513, got %d", event.CreateTime)
	}
	if event.MsgType != "event" {
		t.Errorf("expected MsgType to be 'event', got '%s'", event.MsgType)
	}
	if event.Event != "change_external_contact" {
		t.Errorf("expected Event to be 'change_external_contact', got '%s'", event.Event)
	}
}

func TestParseAddExternalContact(t *testing.T) {
	xmlData := []byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[sys]]></FromUserName> 
		<CreateTime>1403610513</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[change_external_contact]]></Event>
		<ChangeType><![CDATA[add_external_contact]]></ChangeType>
		<UserID><![CDATA[zhangsan]]></UserID>
		<ExternalUserID><![CDATA[woAJ2GCAAAXtWyujaWJHDDGi0mAAAA]]></ExternalUserID>
		<State><![CDATA[teststate]]></State>
		<WelcomeCode><![CDATA[WELCOMECODE]]></WelcomeCode>
	</xml>`)

	event, err := ParseChangeExternalContact(xmlData)
	if err != nil {
		t.Fatalf("failed to parse: %v", err)
	}

	if event.UserID != "zhangsan" {
		t.Errorf("expected UserID to be 'zhangsan', got '%s'", event.UserID)
	}
	if event.ExternalUserID != "woAJ2GCAAAXtWyujaWJHDDGi0mAAAA" {
		t.Errorf("expected ExternalUserID to be 'woAJ2GCAAAXtWyujaWJHDDGi0mAAAA', got '%s'", event.ExternalUserID)
	}
	if event.State != "teststate" {
		t.Errorf("expected State to be 'teststate', got '%s'", event.State)
	}
	if event.WelcomeCode != "WELCOMECODE" {
		t.Errorf("expected WelcomeCode to be 'WELCOMECODE', got '%s'", event.WelcomeCode)
	}
}

func TestParseDelExternalContact(t *testing.T) {
	xmlData := []byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[sys]]></FromUserName> 
		<CreateTime>1403610513</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[change_external_contact]]></Event>
		<ChangeType><![CDATA[del_external_contact]]></ChangeType>
		<UserID><![CDATA[zhangsan]]></UserID>
		<ExternalUserID><![CDATA[woAJ2GCAAAXtWyujaWJHDDGi0mACAAAA]]></ExternalUserID>
		<Source><![CDATA[DELETE_BY_TRANSFER]]></Source>
	</xml>`)

	event, err := ParseChangeExternalContact(xmlData)
	if err != nil {
		t.Fatalf("failed to parse: %v", err)
	}

	if event.Source != "DELETE_BY_TRANSFER" {
		t.Errorf("expected Source to be 'DELETE_BY_TRANSFER', got '%s'", event.Source)
	}
}

func TestParseTransferFail(t *testing.T) {
	xmlData := []byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[sys]]></FromUserName> 
		<CreateTime>1403610513</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[change_external_contact]]></Event>
		<ChangeType><![CDATA[transfer_fail]]></ChangeType>
		<FailReason><![CDATA[customer_refused]]></FailReason>
		<UserID><![CDATA[zhangsan]]></UserID>
		<ExternalUserID><![CDATA[woAJ2GCAAAXtWyujaWJHDDGi0mAAAA]]></ExternalUserID>
	</xml>`)

	event, err := ParseChangeExternalContact(xmlData)
	if err != nil {
		t.Fatalf("failed to parse: %v", err)
	}

	if event.FailReason != "customer_refused" {
		t.Errorf("expected FailReason to be 'customer_refused', got '%s'", event.FailReason)
	}
}

func TestParseUpdateExternalChat(t *testing.T) {
	xmlData := []byte(`<xml>
		<ToUserName><![CDATA[ww55ca070cb9b7eb22]]></ToUserName>
		<FromUserName><![CDATA[sys]]></FromUserName>
		<CreateTime>1403610513</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[change_external_chat]]></Event>
		<ChatId><![CDATA[wrx7HUARsKwGRaQBVKPBTcEyzdHA4HrQ]]></ChatId>
		<ChangeType><![CDATA[update]]></ChangeType>
		<UpdateDetail><![CDATA[add_member]]></UpdateDetail>
		<JoinScene>1</JoinScene>
		<QuitScene>0</QuitScene>
		<MemChangeCnt>10</MemChangeCnt>
		<MemChangeList>
			<Item>Jack</Item>
			<Item>Rose</Item>
		</MemChangeList>
		<LastMemVer>9c3f97c2ada667dfb5f6d03308d963e1</LastMemVer>
		<CurMemVer>71217227bbd112ecfe3a49c482195cb4</CurMemVer>
	</xml>`)

	event, err := ParseChangeExternalChat(xmlData)
	if err != nil {
		t.Fatalf("failed to parse: %v", err)
	}

	if event.ChatId != "wrx7HUARsKwGRaQBVKPBTcEyzdHA4HrQ" {
		t.Errorf("expected ChatId to be 'wrx7HUARsKwGRaQBVKPBTcEyzdHA4HrQ', got '%s'", event.ChatId)
	}
	if event.UpdateDetail != "add_member" {
		t.Errorf("expected UpdateDetail to be 'add_member', got '%s'", event.UpdateDetail)
	}
	if event.JoinScene != 1 {
		t.Errorf("expected JoinScene to be 1, got %d", event.JoinScene)
	}
	if event.QuitScene != 0 {
		t.Errorf("expected QuitScene to be 0, got %d", event.QuitScene)
	}
	if event.MemChangeCnt != 10 {
		t.Errorf("expected MemChangeCnt to be 10, got %d", event.MemChangeCnt)
	}
	if len(event.MemChangeList.Items) != 2 {
		t.Errorf("expected 2 members, got %d", len(event.MemChangeList.Items))
	}
	if event.MemChangeList.Items[0] != "Jack" {
		t.Errorf("expected first member to be 'Jack', got '%s'", event.MemChangeList.Items[0])
	}
}

func TestParseCreateExternalTag(t *testing.T) {
	xmlData := []byte(`<xml>
		<ToUserName><![CDATA[CORPID]]></ToUserName>
		<FromUserName><![CDATA[sys]]></FromUserName> 
		<CreateTime>1403610513</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[change_external_tag]]></Event>
		<Id><![CDATA[TAG_ID]]></Id>
		<TagType><![CDATA[tag]]></TagType>
		<ChangeType><![CDATA[create]]></ChangeType>
		<StrategyId>1</StrategyId>
	</xml>`)

	event, err := ParseChangeExternalTag(xmlData)
	if err != nil {
		t.Fatalf("failed to parse: %v", err)
	}

	if event.Id != "TAG_ID" {
		t.Errorf("expected Id to be 'TAG_ID', got '%s'", event.Id)
	}
	if event.TagType != "tag" {
		t.Errorf("expected TagType to be 'tag', got '%s'", event.TagType)
	}
	if event.StrategyId != "1" {
		t.Errorf("expected StrategyId to be '1', got '%s'", event.StrategyId)
	}
}

func TestParseByEvent(t *testing.T) {
	testCases := []struct {
		name     string
		xmlData  []byte
		expected string
	}{
		{
			name: "change_external_contact event",
			xmlData: []byte(`<xml>
				<ToUserName><![CDATA[toUser]]></ToUserName>
				<FromUserName><![CDATA[sys]]></FromUserName> 
				<CreateTime>1403610513</CreateTime>
				<MsgType><![CDATA[event]]></MsgType>
				<Event><![CDATA[change_external_contact]]></Event>
				<ChangeType><![CDATA[add_external_contact]]></ChangeType>
				<UserID><![CDATA[zhangsan]]></UserID>
				<ExternalUserID><![CDATA[woAJ2GCAAAXtWyujaWJHDDGi0mAAAA]]></ExternalUserID>
			</xml>`),
			expected: "*callback.ChangeExternalContactEvent",
		},
		{
			name: "change_external_chat event",
			xmlData: []byte(`<xml>
				<ToUserName><![CDATA[toUser]]></ToUserName>
				<FromUserName><![CDATA[sys]]></FromUserName> 
				<CreateTime>1403610513</CreateTime>
				<MsgType><![CDATA[event]]></MsgType>
				<Event><![CDATA[change_external_chat]]></Event>
				<ChatId><![CDATA[CHAT_ID]]></ChatId>
				<ChangeType><![CDATA[create]]></ChangeType>
			</xml>`),
			expected: "*callback.ChangeExternalChatEvent",
		},
		{
			name: "change_external_tag event",
			xmlData: []byte(`<xml>
				<ToUserName><![CDATA[CORPID]]></ToUserName>
				<FromUserName><![CDATA[sys]]></FromUserName> 
				<CreateTime>1403610513</CreateTime>
				<MsgType><![CDATA[event]]></MsgType>
				<Event><![CDATA[change_external_tag]]></Event>
				<Id><![CDATA[TAG_ID]]></Id>
				<ChangeType><![CDATA[create]]></ChangeType>
			</xml>`),
			expected: "*callback.ChangeExternalTagEvent",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseByEvent(tc.xmlData)
			if err != nil {
				t.Fatalf("failed to parse: %v", err)
			}

			resultType := fmt.Sprintf("%T", result)
			if resultType != tc.expected {
				t.Errorf("expected type %s, got %s", tc.expected, resultType)
			}
		})
	}
}

func TestGenericParse(t *testing.T) {
	xmlData := []byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[sys]]></FromUserName> 
		<CreateTime>1403610513</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[change_external_contact]]></Event>
		<ChangeType><![CDATA[add_external_contact]]></ChangeType>
		<UserID><![CDATA[zhangsan]]></UserID>
		<ExternalUserID><![CDATA[woAJ2GCAAAXtWyujaWJHDDGi0mAAAA]]></ExternalUserID>
	</xml>`)

	event, err := Parse[ChangeExternalContactEvent](xmlData)
	if err != nil {
		t.Fatalf("failed to parse: %v", err)
	}

	if event.UserID != "zhangsan" {
		t.Errorf("expected UserID to be 'zhangsan', got '%s'", event.UserID)
	}
}
