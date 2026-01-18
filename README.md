# wecom-core

企业微信 Go SDK - 简洁、易用、功能完善的企业微信开发工具包

## 特性

- ✅ **统一日志监控**：支持自定义日志实现，完整的请求追踪
- ✅ **统一响应处理**：自动解析 JSON 响应，统一错误处理
- ✅ **统一重试逻辑**：智能重试机制，支持指数退避
- ✅ **Token 自动管理**：自动获取、缓存、刷新 access_token
- ✅ **并发安全**：所有操作都是并发安全的
- ✅ **接口化设计**：支持自定义 Logger、Cache 实现
- ✅ **易于扩展**：清晰的架构设计，易于添加新模块

## 安装

```bash
go get github.com/shuaidd/wecom-core
```

## 快速开始

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/shuaidd/wecom-core"
    "github.com/shuaidd/wecom-core/config"
    "github.com/shuaidd/wecom-core/pkg/logger"
)

func main() {
    // 创建企业微信客户端
    client, err := wecom.New(
        config.WithCorpID("your_corp_id"),
        config.WithAgent("customer", 100001, "agent_secret_1", "客户管理应用"),
		config.WithAgent("study-assistant", 100002, "agent_secret_2", "学习助手"),
        config.WithLogger(logger.NewStdLogger()),
    )
    if err != nil {
        log.Fatalf("Failed to create wecom client: %v", err)
    }

    ctx := context.Background()

    // 读取成员信息
    user, err := client.Contact.GetUser(ctx, "zhangsan")
    if err != nil {
        log.Fatalf("Failed to get user: %v", err)
    }

    fmt.Printf("成员信息: UserID=%s, Name=%s, Mobile=%s\n",
        user.UserID, user.Name, user.Mobile)
}
```

## 配置选项

```go
client, err := wecom.New(
    // 必填：企业ID和应用密钥
    config.WithCorpID("your_corp_id"),
    config.WithAgent("customer", 100001, "agent_secret_1", "客户管理应用"),
	config.WithAgent("study-assistant", 100002, "agent_secret_2", "学习助手"),

    // 可选：自定义日志
    config.WithLogger(logger.NewStdLogger()),

    // 可选：设置超时时间（默认 30 秒）
    config.WithTimeout(60 * time.Second),

    // 可选：设置重试次数（默认 3 次）
    config.WithRetry(5),

    // 可选：设置退避时间
    config.WithBackoff(1*time.Second, 30*time.Second),

    // 可选：自定义缓存（默认使用内存缓存）
    config.WithCache(yourCustomCache),
)
```

## 核心特性详解

### 自动 Token 管理

SDK 会自动处理 access_token 的获取、缓存和刷新：

- ✅ 首次调用时自动获取 token
- ✅ token 缓存在内存中（可自定义缓存实现）
- ✅ 提前 5 分钟自动刷新，避免过期
- ✅ 并发安全，防止重复获取
- ✅ token 失效时自动刷新并重试

### 智能重试机制

自动重试以下场景：

- ✅ Token 过期（errcode 40014, 42001）
- ✅ API 频率限制（errcode 45009）
- ✅ 系统繁忙（errcode 10001）
- ✅ 使用指数退避算法，避免频繁重试

### 统一日志记录

记录所有关键操作：

- 🔍 请求日志：URL、方法、耗时
- 🔍 响应日志：状态码、错误信息
- 🔍 Token 日志：获取、刷新、失效
- 🔍 重试日志：触发原因、次数

### 自定义 Logger

```go
// 实现 Logger 接口
type MyLogger struct{}

func (l *MyLogger) Debug(msg string, fields ...logger.Field) {}
func (l *MyLogger) Info(msg string, fields ...logger.Field)  {}
func (l *MyLogger) Warn(msg string, fields ...logger.Field)  {}
func (l *MyLogger) Error(msg string, fields ...logger.Field) {}

// 使用自定义 Logger
client, err := wecom.New(
    config.WithCorpID("your_corp_id"),
    config.WithCorpSecret("your_corp_secret"),
    config.WithLogger(&MyLogger{}),
)
```

### 自定义 Cache

```go
// 实现 Cache 接口
type MyCache struct{}

func (c *MyCache) Get(ctx context.Context, key string) (token string, expireAt time.Time, err error) {
    // 从 Redis 获取
}

func (c *MyCache) Set(ctx context.Context, key string, token string, expireAt time.Time) error {
    // 存储到 Redis
}

func (c *MyCache) Delete(ctx context.Context, key string) error {
    // 从 Redis 删除
}

// 使用自定义 Cache
client, err := wecom.New(
    config.WithCorpID("your_corp_id"),
    config.WithCorpSecret("your_corp_secret"),
    config.WithCache(&MyCache{}),
)
```

## 错误处理

```go
user, err := client.Contact.GetUser(ctx, "zhangsan")
if err != nil {
    // 判断是否为企业微信错误
    if errors.IsWecomError(err) {
        errCode := errors.GetErrorCode(err)
        // 根据错误码处理
        switch errCode {
        case errors.ErrCodeUserNotFound:
            // 成员不存在
        case errors.ErrCodeInvalidParameter:
            // 参数错误
        default:
            // 其他错误
        }
    }
    return err
}
```

## 项目结构

```
wecom-core/
├── wecom.go                    # 主入口
├── config/                     # 配置管理
├── internal/                   # 内部包（不对外暴露）
│   ├── client/                # HTTP 客户端
│   ├── auth/                  # Token 管理
│   ├── retry/                 # 重试逻辑
│   └── errors/                # 错误处理
├── pkg/                        # 公共包（可被外部引用）
│   ├── logger/                # 日志接口
│   └── cache/                 # 缓存接口
├── types/                      # 数据类型定义
│   ├── common/                # 通用类型
│   └── contact/               # 通讯录相关
└── services/                   # 业务服务
    └── contact/               # 通讯录服务
```

## 开发计划

- ✅ **阶段一：基础框架**（已完成）
  - 统一 HTTP 客户端
  - Token 自动管理
  - 智能重试机制
  - 错误处理

- ✅ **阶段二：核心业务模块**（已完成）
  - ✅ 通讯录管理 (Contact)
  - ✅ 身份验证 (OAuth)
  - ✅ 企业二维码 (QRCode)
  - ✅ IP 管理 (IP)
  - ✅ 上下游服务 (UpDown)
  - ✅ 企业互联 (CorpGroup)
  - ✅ 安全管理 (Security)
  - ✅ 消息管理 (Message)
  - ✅ 应用管理 (Agent)
    - ✅ 应用信息管理（获取应用详情、获取应用列表、设置应用）
    - ✅ 菜单管理（创建菜单、获取菜单、删除菜单）
    - ✅ 工作台自定义展示（设置/获取模板、设置/批量设置/获取用户数据）
  - ✅ 外部联系人 (ExternalContact)
    - ✅ 客户管理（获取客户列表、获取客户详情、修改备注、批量获取）
    - ✅ 客户标签管理（企业标签、规则组标签、客户打标）
    - ✅ 客户联系规则组管理（规则组CRUD、管理范围）
    - ✅ 获客助手（获客链接、额度管理、使用统计）
    - ✅ 客户群管理（获取群列表、获取群详情、opengid转换）
    - ✅ 客户朋友圈（发表任务、获取列表、互动数据、规则组管理）
    - ✅ 联系我与客户入群方式（「联系我」配置、客户群进群方式管理）
    - ✅ 企业服务人员管理（获取配置了客户联系功能的成员列表）
    - ✅ 统计管理（群聊数据统计、联系客户统计）
    - ✅ 消息推送（创建企业群发、获取群发记录、发送新客户欢迎语、入群欢迎语素材管理）
    - ✅ 在职继承（分配在职成员的客户、分配在职成员的客户群、查询客户接替状态）
    - ✅ 商品图册管理（创建、获取、列表、编辑、删除）
    - ✅ 聊天敏感词管理（新建、获取列表、获取详情、修改、删除）
    - ✅ 获取已服务的外部联系人
    - ⏳ 离职继承（部分完成）
    - ⏳ 上传附件资源（需要文件上传功能支持，待实现）

- ⏳ **阶段三：更多业务模块**（规划中）
  - ✅ 素材管理 (Media)
    - ✅ 上传图片（永久有效）
    - ✅ 上传临时素材（图片、语音、视频、文件）
    - ✅ 获取临时素材（支持Range分块下载）
    - ✅ 获取高清语音素材
    - ✅ 异步上传临时素材（支持200M大文件）
    - ✅ 查询异步上传任务结果
  - ✅ 电子发票 (Invoice)
    - ✅ 查询电子发票
    - ✅ 批量查询电子发票
    - ✅ 更新发票状态（锁定、解锁、核销）
    - ✅ 批量更新发票状态
  - ✅ 微信客服 (KF)
    - ✅ 客服账号管理（添加、删除、修改、获取列表）
    - ✅ 获取客服账号链接
    - ✅ 接待人员管理（添加、删除、获取列表）
    - ✅ 会话分配与消息收发（获取会话状态、变更会话状态）
    - ✅ 消息发送（文本、图片、语音、视频、文件、图文链接、小程序、菜单、地理位置、获客链接）
    - ✅ 事件响应消息（发送欢迎语等场景化消息）
    - ✅ 客户基础信息管理（批量获取客户基础信息）
    - ✅ 升级服务配置（获取配置的专员与客户群、为客户升级服务、取消推荐）
  - ✅ 邮件服务 (Email)
    - ✅ 发送邮件（普通邮件、日程邮件、会议邮件）
    - ✅ 公共邮箱管理（创建、更新、删除、获取、搜索）
    - ✅ 客户端专用密码管理（获取列表、删除）
    - ✅ 应用邮箱账号管理（查询、更新）
    - ✅ 邮件群组管理（创建、获取、更新、搜索、删除）
  - ✅ 微文档 (Wedoc)
    - ✅ 文档管理（新建文档/表格/智能表格、获取基础信息、重命名、分享、删除）
    - ✅ 文档内容管理（获取文档数据、批量编辑文档内容、上传文档图片）
    - ✅ 表格内容管理（获取表格行列信息、获取表格数据、批量编辑表格内容）
    - ✅ 收集表管理（创建收集表、获取信息、编辑收集表）
    - ✅ 收集表数据（读取答案、获取统计信息、已提交/未提交列表）
    - ✅ 智能表格内容管理（记录、字段、视图、子表、编组的完整CRUD操作）
    - ✅ 文档权限管理（获取权限信息、修改安全设置、修改查看规则、修改通知范围及权限）
    - ✅ 智能表格内容权限（查询和更新子表权限、管理成员额外权限、字段和记录级别权限）
    - ✅ 高级功能账号管理（分配、取消、获取高级功能账号列表）
  - ✅ 日历管理 (Calendar)
    - ✅ 日历管理（创建、获取、更新、删除日历）
    - ✅ 日程管理（创建、获取、更新、取消日程）
    - ✅ 日程参与者管理（新增、删除参与者）
    - ✅ 获取日历下的日程列表
    - ✅ 重复日程管理（支持不同操作模式）
  - ✅ 会议管理 (Meeting)
    - ✅ 创建预约会议
    - ✅ 修改预约会议
    - ✅ 取消预约会议
    - ✅ 获取会议详情
    - ✅ 获取成员会议ID列表
    - ✅ 会议统计管理（获取会议发起记录）
  - ✅ OA 审批

## 示例

查看 [examples](./examples) 目录获取更多示例：

- [基础示例](./examples/basic/main.go)
- [通讯录示例](./examples/contact/main.go)
- [应用管理示例](./examples/agent/main.go)
- [微文档示例](./examples/wedoc/main.go)

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License
