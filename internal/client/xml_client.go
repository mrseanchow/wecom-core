package client

import (
	"context"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/mrseanchow/wecom-core/pkg/logger"
)

type XMLResponse struct {
	XMLName             xml.Name `json:"-" xml:"xml"`
	ReturnCode          string   `xml:"return_code"`
	ReturnMsg           string   `xml:"return_msg"`
	Sign                string   `xml:"sign,omitempty"`
	ResultCode          string   `xml:"result_code,omitempty"`
	ErrCode             string   `xml:"err_code,omitempty"`
	ErrCodeDes          string   `xml:"err_code_des,omitempty"`
	MchBillno           string   `xml:"mch_billno,omitempty"`
	MchID               string   `xml:"mch_id,omitempty"`
	AppID               string   `xml:"appid,omitempty"`
	WxAppID             string   `xml:"wxappid,omitempty"`
	ReOpenID            string   `xml:"re_openid,omitempty"`
	TotalAmount         int      `xml:"total_amount,omitempty"`
	SendListID          string   `xml:"send_listid,omitempty"`
	Status              string   `xml:"status,omitempty"`
	DetailID            string   `xml:"detail_id,omitempty"`
	SendType            string   `xml:"send_type,omitempty"`
	Reason              string   `xml:"reason,omitempty"`
	SendTime            string   `xml:"send_time,omitempty"`
	RefundTime          string   `xml:"refund_time,omitempty"`
	RefundAmount        int      `xml:"refund_amount,omitempty"`
	PartnerTradeNo      string   `xml:"partner_trade_no,omitempty"`
	PaymentNo           string   `xml:"payment_no,omitempty"`
	PaymentTime         string   `xml:"payment_time,omitempty"`
	PaymentAmount       int      `xml:"payment_amount,omitempty"`
	TransferTime        string   `xml:"transfer_time,omitempty"`
	TransferName        string   `xml:"transfer_name,omitempty"`
	Desc                string   `xml:"desc,omitempty"`
	SenderName          string   `xml:"sender_name,omitempty"`
	SenderHeaderMediaID string   `xml:"sender_header_media_id,omitempty"`
	Wishing             string   `xml:"wishing,omitempty"`
	Remark              string   `xml:"remark,omitempty"`
	ActName             string   `xml:"act_name,omitempty"`
	OpenID              string   `xml:"openid,omitempty"`
	Amount              int      `xml:"amount,omitempty"`
	RcvTime             string   `xml:"rcv_time,omitempty"`
	DeviceInfo          string   `xml:"device_info,omitempty"`
	NonceStr            string   `xml:"nonce_str,omitempty"`
}

type XMLClient struct {
	httpClient *http.Client
	baseURL    string
	logger     logger.Logger
}

func NewXMLClient(baseURL string, certPath, keyPath string, logger logger.Logger) (*XMLClient, error) {
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load certificate: %w", err)
	}

	return &XMLClient{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					Certificates: []tls.Certificate{cert},
				},
			},
		},
		baseURL: baseURL,
		logger:  logger,
	}, nil
}

func (c *XMLClient) PostXML(ctx context.Context, path string, xmlData []byte) (*XMLResponse, error) {
	url := c.baseURL + path

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(string(xmlData)))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/xml")

	c.logger.Debug("XML API Request", logger.F("url", url), logger.F("data", string(xmlData)))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.logger.Error("XML request failed", logger.F("error", err))
		return nil, fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("Failed to read response", logger.F("error", err))
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	c.logger.Debug("XML API Response", logger.F("status", resp.StatusCode), logger.F("data", string(body)))

	var result XMLResponse
	if err := xml.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.ReturnCode != "SUCCESS" {
		return &result, fmt.Errorf("api error: %s", result.ReturnMsg)
	}

	if result.ResultCode != "" && result.ResultCode != "SUCCESS" {
		return &result, fmt.Errorf("business error: %s - %s", result.ErrCode, result.ErrCodeDes)
	}

	return &result, nil
}
