package payment

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"sort"
	"strings"

	"github.com/mrseanchow/wecom-core/internal/client"
	"github.com/mrseanchow/wecom-core/pkg/logger"
	"github.com/mrseanchow/wecom-core/types/payment"
)

type Service struct {
	xmlClient *client.XMLClient
	config    *Config
	logger    logger.Logger
}

type Config struct {
	MchID     string
	APIKey    string
	APPSecret string
	CertPath  string
	KeyPath   string
}

func NewService(cfg *Config, log logger.Logger) (*Service, error) {
	xmlClient, err := client.NewXMLClient("https://api.mch.weixin.qq.com", cfg.CertPath, cfg.KeyPath, log)
	if err != nil {
		return nil, err
	}

	return &Service{
		xmlClient: xmlClient,
		config:    cfg,
		logger:    log,
	}, nil
}

func (s *Service) SendRedPack(ctx context.Context, req *payment.SendRedPackRequest) (*payment.SendRedPackResponse, error) {
	if err := s.generateRedPackSign(req); err != nil {
		return nil, fmt.Errorf("failed to generate sign: %w", err)
	}

	if err := s.generateWorkWXRedPackSign(req); err != nil {
		return nil, fmt.Errorf("failed to generate workwx sign: %w", err)
	}

	xmlData, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.xmlClient.PostXML(ctx, "/mmpaymkttransfers/sendworkwxredpack", xmlData)
	if err != nil {
		return nil, err
	}

	result := &payment.SendRedPackResponse{
		ReturnCode:          resp.ReturnCode,
		ReturnMsg:           resp.ReturnMsg,
		Sign:                resp.Sign,
		ResultCode:          resp.ResultCode,
		ErrCode:             resp.ErrCode,
		ErrCodeDes:          resp.ErrCodeDes,
		MchBillno:           resp.MchBillno,
		MchID:               resp.MchID,
		WxAppID:             resp.WxAppID,
		ReOpenID:            resp.ReOpenID,
		TotalAmount:         resp.TotalAmount,
		SendListID:          resp.SendListID,
		SenderName:          resp.SenderName,
		SenderHeaderMediaID: resp.SenderHeaderMediaID,
	}

	return result, nil
}

func (s *Service) QueryRedPack(ctx context.Context, req *payment.QueryRedPackRequest) (*payment.QueryRedPackResponse, error) {
	if err := s.generateQueryRedPackSign(req); err != nil {
		return nil, fmt.Errorf("failed to generate sign: %w", err)
	}

	xmlData, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.xmlClient.PostXML(ctx, "/mmpaymkttransfers/queryworkwxredpack", xmlData)
	if err != nil {
		return nil, err
	}

	result := &payment.QueryRedPackResponse{
		ReturnCode:          resp.ReturnCode,
		ReturnMsg:           resp.ReturnMsg,
		Sign:                resp.Sign,
		ResultCode:          resp.ResultCode,
		ErrCode:             resp.ErrCode,
		ErrCodeDes:          resp.ErrCodeDes,
		MchBillno:           resp.MchBillno,
		MchID:               resp.MchID,
		DetailID:            resp.DetailID,
		Status:              resp.Status,
		SendType:            resp.SendType,
		TotalAmount:         resp.TotalAmount,
		Reason:              resp.Reason,
		SendTime:            resp.SendTime,
		RefundTime:          resp.RefundTime,
		RefundAmount:        resp.RefundAmount,
		Wishing:             resp.Wishing,
		Remark:              resp.Remark,
		ActName:             resp.ActName,
		OpenID:              resp.OpenID,
		Amount:              resp.Amount,
		RcvTime:             resp.RcvTime,
		SenderName:          resp.SenderName,
		SenderHeaderMediaID: resp.SenderHeaderMediaID,
	}

	return result, nil
}

func (s *Service) SendTransfer(ctx context.Context, req *payment.SendTransferRequest) (*payment.SendTransferResponse, error) {
	if err := s.generateTransferSign(req); err != nil {
		return nil, fmt.Errorf("failed to generate sign: %w", err)
	}

	if err := s.generateWorkWXTransferSign(req); err != nil {
		return nil, fmt.Errorf("failed to generate workwx sign: %w", err)
	}

	xmlData, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.xmlClient.PostXML(ctx, "/mmpaymkttransfers/promotion/paywwsptrans2pocket", xmlData)
	if err != nil {
		return nil, err
	}

	result := &payment.SendTransferResponse{
		ReturnCode:     resp.ReturnCode,
		ReturnMsg:      resp.ReturnMsg,
		AppID:          resp.AppID,
		MchID:          resp.MchID,
		DeviceInfo:     resp.DeviceInfo,
		NonceStr:       resp.NonceStr,
		ResultCode:     resp.ResultCode,
		ErrCode:        resp.ErrCode,
		ErrCodeDes:     resp.ErrCodeDes,
		PartnerTradeNo: resp.PartnerTradeNo,
		PaymentNo:      resp.PaymentNo,
		PaymentTime:    resp.PaymentTime,
	}

	return result, nil
}

func (s *Service) QueryTransfer(ctx context.Context, req *payment.QueryTransferRequest) (*payment.QueryTransferResponse, error) {
	if err := s.generateQueryTransferSign(req); err != nil {
		return nil, fmt.Errorf("failed to generate sign: %w", err)
	}

	xmlData, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.xmlClient.PostXML(ctx, "/mmpaymkttransfers/promotion/querywwsptrans2pocket", xmlData)
	if err != nil {
		return nil, err
	}

	result := &payment.QueryTransferResponse{
		ReturnCode:     resp.ReturnCode,
		ReturnMsg:      resp.ReturnMsg,
		ResultCode:     resp.ResultCode,
		ErrCode:        resp.ErrCode,
		ErrCodeDes:     resp.ErrCodeDes,
		MchID:          resp.MchID,
		AppID:          resp.AppID,
		DetailID:       resp.DetailID,
		PartnerTradeNo: resp.PartnerTradeNo,
		Status:         resp.Status,
		Reason:         resp.Reason,
		OpenID:         resp.OpenID,
		TransferName:   resp.TransferName,
		PaymentAmount:  resp.PaymentAmount,
		TransferTime:   resp.TransferTime,
		Desc:           resp.Desc,
	}

	return result, nil
}

func (s *Service) generateRedPackSign(req *payment.SendRedPackRequest) error {
	params := map[string]string{
		"act_name":     req.ActName,
		"mch_billno":   req.MchBillno,
		"mch_id":       req.MchID,
		"nonce_str":    req.NonceStr,
		"re_openid":    req.ReOpenID,
		"total_amount": fmt.Sprintf("%d", req.TotalAmount),
		"wxappid":      req.WxAppID,
	}

	sign, err := s.md5Sign(params, s.config.APIKey)
	if err != nil {
		return err
	}

	req.Sign = sign
	return nil
}

func (s *Service) generateWorkWXRedPackSign(req *payment.SendRedPackRequest) error {
	params := map[string]string{
		"act_name":     req.ActName,
		"mch_billno":   req.MchBillno,
		"mch_id":       req.MchID,
		"nonce_str":    req.NonceStr,
		"re_openid":    req.ReOpenID,
		"total_amount": fmt.Sprintf("%d", req.TotalAmount),
		"wxappid":      req.WxAppID,
	}

	sign, err := s.md5Sign(params, s.config.APPSecret)
	if err != nil {
		return err
	}

	req.WorkWXSign = sign
	return nil
}

func (s *Service) generateQueryRedPackSign(req *payment.QueryRedPackRequest) error {
	params := map[string]string{
		"appid":      req.AppID,
		"mch_billno": req.MchBillno,
		"mch_id":     req.MchID,
		"nonce_str":  req.NonceStr,
	}

	sign, err := s.md5Sign(params, s.config.APIKey)
	if err != nil {
		return err
	}

	req.Sign = sign
	return nil
}

func (s *Service) generateTransferSign(req *payment.SendTransferRequest) error {
	params := map[string]string{
		"amount":           fmt.Sprintf("%d", req.Amount),
		"appid":            req.AppID,
		"desc":             req.Desc,
		"mch_id":           req.MchID,
		"nonce_str":        req.NonceStr,
		"openid":           req.OpenID,
		"partner_trade_no": req.PartnerTradeNo,
		"ww_msg_type":      req.WWMsgType,
	}

	sign, err := s.md5Sign(params, s.config.APIKey)
	if err != nil {
		return err
	}

	req.Sign = sign
	return nil
}

func (s *Service) generateWorkWXTransferSign(req *payment.SendTransferRequest) error {
	params := map[string]string{
		"amount":           fmt.Sprintf("%d", req.Amount),
		"appid":            req.AppID,
		"desc":             req.Desc,
		"mch_id":           req.MchID,
		"nonce_str":        req.NonceStr,
		"openid":           req.OpenID,
		"partner_trade_no": req.PartnerTradeNo,
		"ww_msg_type":      req.WWMsgType,
	}

	sign, err := s.md5Sign(params, s.config.APPSecret)
	if err != nil {
		return err
	}

	req.WorkWXSign = sign
	return nil
}

func (s *Service) generateQueryTransferSign(req *payment.QueryTransferRequest) error {
	params := map[string]string{
		"appid":            req.AppID,
		"mch_id":           req.MchID,
		"nonce_str":        req.NonceStr,
		"partner_trade_no": req.PartnerTradeNo,
	}

	sign, err := s.md5Sign(params, s.config.APIKey)
	if err != nil {
		return err
	}

	req.Sign = sign
	return nil
}

func (s *Service) md5Sign(params map[string]string, key string) (string, error) {
	keys := make([]string, 0, len(params))
	for k := range params {
		if params[k] == "" {
			continue
		}
		keys = append(keys, k)
	}

	sort.Strings(keys)

	var buf strings.Builder
	for i, k := range keys {
		if i > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(params[k])
	}

	buf.WriteByte('&')
	buf.WriteString(key)

	hash := md5.Sum([]byte(buf.String()))
	return strings.ToUpper(hex.EncodeToString(hash[:])), nil
}

