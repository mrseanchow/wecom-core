package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/shuaidd/wecom-core"
	"github.com/shuaidd/wecom-core/types/message"
)

func TestSendTextMessage(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.Message.Send(context.Background(), &message.SendMessageRequest{
		ToUser:  "20170410022717",
		MsgType: "text",
		AgentID: 1000060,
		Text: &message.TextMessage{
			Content: "hello world",
		},
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestSendImageMessage(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.Message.Send(context.Background(), &message.SendMessageRequest{
		ToUser:  "20170410022717",
		MsgType: "image",
		AgentID: 1000060,
		Image: &message.MediaMessage{
			MediaID: "3GpqnD4o8ztI754ejNTtCPQRfjOgU19pmlJGZN_2wvfVJlGEATht_F2G4yD3Lk3ZI3PODYj99GiAipn2v-EgFhQ",
		},
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestSendVoiceMessage(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.Message.Send(context.Background(), &message.SendMessageRequest{
		ToUser:  "20170410022717",
		MsgType: "voice",
		AgentID: 1000060,
		Voice: &message.MediaMessage{
			MediaID: "MEDIA_ID",
		},
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestSendVideoMessage(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.Message.Send(context.Background(), &message.SendMessageRequest{
		ToUser:  "20170410022717",
		MsgType: "video",
		AgentID: 1000060,
		Video: &message.VideoMessage{
			MediaID:     "MEDIA_ID",
			Title:       "视频标题",
			Description: "视频描述",
		},
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestSendFileMessage(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.Message.Send(context.Background(), &message.SendMessageRequest{
		ToUser:  "20170410022717",
		MsgType: "file",
		AgentID: 1000060,
		File: &message.MediaMessage{
			MediaID: "3-uHsjTP0UegpkiB9Gx1RqkRhaMCarEmpInpwQMcovAaYF6v2Exg7ClCOGQ_1oF-L",
		},
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestSendTextCardMessage(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.Message.Send(context.Background(), &message.SendMessageRequest{
		ToUser:  "20170410022717",
		MsgType: "textcard",
		AgentID: 1000060,
		TextCard: &message.TextCardMessage{
			Title:       "通知标题",
			Description: "通知内容描述",
			URL:         "https://www.baidu.com",
			BtnTxt:      "查看详情",
		},
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestSendNewsMessage(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.Message.Send(context.Background(), &message.SendMessageRequest{
		ToUser:  "20170410022717",
		MsgType: "news",
		AgentID: 1000060,
		News: &message.NewsMessage{
			Articles: []message.NewsArticle{
				{
					Title:       "开发者中心",
					Description: "客户联系",
					URL:         "https://developer.work.weixin.qq.com/document/path/92228",
					PicURL:      "https://wework.qpic.cn/wwpic3az/291238_YYSv1K_mQn6oBN4_1767923045/0",
				},
			},
		},
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestSendMPNewsMessage(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.Message.Send(context.Background(), &message.SendMessageRequest{
		ToUser:  "20170410022717",
		MsgType: "mpnews",
		AgentID: 1000060,
		MPNews: &message.MPNewsMessage{
			Articles: []message.MPNewsArticle{
				{
					Title:            "图文标题",
					ThumbMediaID:     "3IN8NQFtxcnNLUQDV99RB0dHAp9009miO14W5qBb4B88lRHte5nscyuG3Umab1z1u",
					Author:           "作者",
					Content:          "图文内容",
					ContentSourceURL: "https://developer.work.weixin.qq.com/document/path/92228",
					Digest:           "图文摘要",
				},
			},
		},
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestSendMarkdownMessage(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.Message.Send(context.Background(), &message.SendMessageRequest{
		ToUser:  "20170410022717",
		MsgType: "markdown",
		AgentID: 1000060,
		Markdown: &message.MarkdownMessage{
			Content: "您的会议室已经预订，稍后会同步到`邮箱`\n> **收到** <font color='info'>[李四](https://work.weixin.qq.com)</font> 请注意查收",
		},
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestSendMiniProgramNoticeMessage(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.Message.Send(context.Background(), &message.SendMessageRequest{
		ToUser:  "20170410022717",
		MsgType: "miniprogram_notice",
		AgentID: 1000060,
		MiniProgramNotice: &message.MiniProgramNoticeMessage{
			AppID:             "wx1234567890abcdef",
			Page:              "pages/index/index",
			Title:             "审批通知",
			Description:       "您有一条新的审批申请",
			EmphasisFirstItem: true,
			ContentItem: []message.ContentItem{
				{
					Key:   "申请人",
					Value: "张三",
				},
				{
					Key:   "审批类型",
					Value: "报销审批",
				},
			},
		},
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}

func TestSendTemplateCardMessage(t *testing.T) {
	cli, err := wecom.Default()
	if err != nil {
		return
	}
	resp, err := cli.Message.Send(context.Background(), &message.SendMessageRequest{
		ToUser:  "20170410022717",
		MsgType: "template_card",
		AgentID: 1000060,
		TemplateCard: &message.TemplateCardMessage{
			CardType: "text_notice",
			Source: &message.CardSource{
				IconURL:   "https://example.com/icon.png",
				Desc:      "企业微信",
				DescColor: 0,
			},
			MainTitle: &message.CardMainTitle{
				Title: "欢迎使用",
				Desc:  "欢迎使用企业微信",
			},
			EmphasisContent: &message.CardEmphasisContent{
				Title: "100",
				Desc:  "数据处理完成",
			},
			SubTitleText: "这是副标题文本",
			HorizontalContentList: []message.CardHorizontalContent{
				{
					KeyName: "任务状态",
					Value:   "已完成",
				},
			},
			CardAction: &message.CardAction{
				Type: 1,
				URL:  "https://example.com",
			},
		},
	})
	fmt.Printf("响应：%+v 错误：%+v", resp, err)
}
