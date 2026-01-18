# XML 解析器

用于解析企业微信回调 XML 事件的泛型解析器，支持灵活扩展。

## 特性

- 支持泛型解析任意结构体
- 自动提取公共字段到 BaseEvent
- 支持按事件类型自动路由
- 易于扩展新的事件类型
- 完整的测试覆盖

## 使用方法

### 1. 基础使用 - 解析特定事件类型

```go
import "github.com/shuaidd/wecom-core/services/callback"

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

event, err := callback.ParseChangeExternalContact(xmlData)
if err != nil {
    // 处理错误
}

fmt.Println(event.UserID)
fmt.Println(event.ExternalUserID)
```

### 2. 使用泛型解析

```go
type CustomEvent struct {
    callback.BaseEvent
    CustomField string `xml:"CustomField"`
}

custom, err := callback.Parse[CustomEvent](xmlData)
if err != nil {
    // 处理错误
}

fmt.Println(custom.CustomField)
```

### 3. 自动按事件类型路由

```go
event, err := callback.ParseByEvent(xmlData)
if err != nil {
    // 处理错误
}

switch v := event.(type) {
case *callback.ChangeExternalContactEvent:
    fmt.Printf("Contact Event: %s\n", v.UserID)
case *callback.ChangeExternalChatEvent:
    fmt.Printf("Chat Event: %s\n", v.ChatId)
case *callback.ChangeExternalTagEvent:
    fmt.Printf("Tag Event: %s\n", v.Id)
}
```

### 4. 只解析基础事件

```go
base, err := callback.ParseBase(xmlData)
if err != nil {
    // 处理错误
}

fmt.Printf("Event Type: %s\n", base.Event)
fmt.Printf("Change Type: %s\n", base.ChangeType)
```

## 支持的事件类型

### BaseEvent - 基础事件结构

所有事件都包含以下公共字段：

```go
type BaseEvent struct {
    XMLName      xml.Name
    ToUserName   string
    FromUserName string
    CreateTime   int64
    MsgType      string
    Event        string
    ChangeType   string
}
```

### ChangeExternalContactEvent - 外部联系人变更事件

包含字段：
- `UserID` - 企业服务人员的 UserID
- `ExternalUserID` - 外部联系人的 userid
- `State` - 联系我方式配置的 state 参数
- `WelcomeCode` - 欢迎语 code
- `Source` - 删除客户的操作来源
- `FailReason` - 接替失败的原因

支持的 ChangeType：
- `add_external_contact` - 添加企业客户
- `edit_external_contact` - 编辑企业客户
- `add_half_external_contact` - 外部联系人免验证添加成员
- `del_external_contact` - 删除企业客户
- `del_follow_user` - 删除跟进成员
- `transfer_fail` - 客户接替失败

### ChangeExternalChatEvent - 客户群变更事件

包含字段：
- `ChatId` - 群 ID
- `UpdateDetail` - 变更详情
- `JoinScene` - 成员入群方式
- `QuitScene` - 成员退群方式
- `MemChangeCnt` - 成员变更数量
- `MemChangeList` - 变更的成员列表
- `LastMemVer` - 变更前的群成员版本号
- `CurMemVer` - 变更后的群成员版本号

支持的 ChangeType：
- `create` - 客户群创建
- `update` - 客户群变更
- `dismiss` - 客户群解散

### ChangeExternalTagEvent - 企业客户标签事件

包含字段：
- `Id` - 标签或标签组的 ID
- `TagType` - `tag` 或 `tag_group`
- `StrategyId` - 规则组 id

支持的 ChangeType：
- `create` - 创建标签
- `update` - 变更标签
- `delete` - 删除标签
- `shuffle` - 标签重排

## 扩展新的事件类型

如果需要添加新的事件类型，只需：

1. 定义新的事件结构体，嵌入 BaseEvent：
```go
type NewEventType struct {
    callback.BaseEvent
    Field1 string `xml:"Field1"`
    Field2 int    `xml:"Field2"`
}
```

2. 创建对应的解析函数：
```go
func ParseNewEvent(data []byte) (*NewEventType, error) {
    return callback.Parse[NewEventType](data)
}
```

3. 在 `ParseByEvent` 函数中添加路由逻辑（如果需要）

## XML 标签说明

解析器使用 Go 标准库 `encoding/xml`，支持以下标签：

```go
type Example struct {
    Field1 string `xml:"FieldName"`           // 指定 XML 元素名称
    Field2 string `xml:",chardata"`           // 使用元素内容
    Field3 string `xml:"Field3,attr"`         // 属性
    Field4 string `xml:"field4,omitempty"`    // 可选字段
}
```

## 错误处理

解析函数返回的错误可能包括：
- XML 格式错误
- 字段类型不匹配
- 必填字段缺失

建议在调用解析函数后检查错误：

```go
event, err := callback.ParseByEvent(xmlData)
if err != nil {
    log.Printf("Failed to parse XML: %v", err)
    return
}
```

## 性能建议

1. 对于高频回调，建议使用 `ParseByEvent` 实现统一处理
2. 对于已知事件类型，使用专用解析函数（如 `ParseChangeExternalContact`）
3. 解析后的结构体可以缓存避免重复解析

## 示例

完整示例请参考 `services/callback/example/main.go`
