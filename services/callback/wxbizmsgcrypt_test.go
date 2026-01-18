package callback

import (
	"fmt"
	"testing"
)

func TestNewWXBizMsgCrypt(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		token           string
		encoding_aeskey string
		receiver_id     string
		protocol_type   ProtocolType
		want            *WXBizMsgCrypt
	}{
		// TODO: Add test cases.
		{
			token: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token := "EwqzXfUhCSV"
			receiverId := "ww36e0a51aab349a7d"
			encodingAeskey := "N9MfTG76WmKEwqzXfUhCSVlNKW9YG2raDvz1TdMPanw"
			wxcpt := NewWXBizMsgCrypt(token, encodingAeskey, receiverId, XmlType)
			/*
			   	------------使用示例一：验证回调URL---------------
			   	*企业开启回调模式时，企业微信会向验证url发送一个get请求
			   	假设点击验证时，企业收到类似请求：
			   	* GET /cgi-bin/wxpush?msg_signature=5c45ff5e21c57e6ad56bac8758b79b1d9ac89fd3&timestamp=1409659589&nonce=263014780&echostr=P9nAzCzyDtyTWESHep1vC5X9xho%2FqYX3Zpb4yKa9SKld1DsH3Iyt3tP3zNdtp%2B4RPcs8TgAE7OaBO%2BFZXvnaqQ%3D%3D
			   	* HTTP/1.1 Host: qy.weixin.qq.com

			   	接收到该请求时，企业应
			        1.解析出Get请求的参数，包括消息体签名(msg_signature)，时间戳(timestamp)，随机数字串(nonce)以及企业微信推送过来的随机加密字符串(echostr),
			        这一步注意作URL解码。
			        2.验证消息体签名的正确性
			        3. 解密出echostr原文，将原文当作Get请求的response，返回给企业微信
			        第2，3步可以用企业微信提供的库函数VerifyURL来实现。

			*/
			// 解析出url上的参数值如下：
			// verifyMsgSign := HttpUtils.ParseUrl("msg_signature")
			verifyMsgSign := "f903b456678c5b042fb3ac8c13215d362be20015"
			// verifyTimestamp := HttpUtils.ParseUrl("timestamp")
			verifyTimestamp := "1763695192"
			// verifyNonce := HttpUtils.ParseUrl("nonce")
			verifyNonce := "1763358121"
			// verifyEchoStr := HttpUtils.ParseUrl("echoStr")
			verifyEchoStr := "LHx6PhMuAWVUZz6+kTVC0E58HZPwgaVbPbJh2oBTV75RRYDg4oUcabSkq0bmsle9WC3oslxV4AAPAKJIkV/ijA=="
			echoStr, cryptErr := wxcpt.VerifyURL(verifyMsgSign, verifyTimestamp, verifyNonce, verifyEchoStr)
			if nil != cryptErr {
				fmt.Println("verifyUrl fail", cryptErr)
			}
			fmt.Println("verifyUrl success echoStr", string(echoStr))
			
			if true {
				t.Errorf("NewWXBizMsgCrypt() = %v, want %v", echoStr, tt.want)
			}
		})
	}
}
