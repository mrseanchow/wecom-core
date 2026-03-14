package email

import (
	"context"

	"github.com/mrseanchow/wecom-core/internal/client"
	"github.com/mrseanchow/wecom-core/types/email"
)

const (
	sendEmailURL = "/cgi-bin/exmail/app/compose_send"
)

// SendEmail 发送邮�?支持普通邮件、日程邮件、会议邮�?
// 普通邮�? 只需填写基本字段(to, subject, content)
// 日程邮件: 需要填写schedule字段
// 会议邮件: 需要同时填写schedule和meeting字段
func (s *Service) SendEmail(ctx context.Context, req *email.SendEmailRequest) (*email.SendEmailResponse, error) {
	return client.PostAndUnmarshal[email.SendEmailResponse](s.client, ctx, sendEmailURL, req)
}

// SendNormalEmail 发送普通邮�?
func (s *Service) SendNormalEmail(ctx context.Context, to *email.EmailRecipient, subject, content string, opts ...SendEmailOption) (*email.SendEmailResponse, error) {
	req := &email.SendEmailRequest{
		To:      to,
		Subject: subject,
		Content: content,
	}

	for _, opt := range opts {
		opt(req)
	}

	return s.SendEmail(ctx, req)
}

// SendScheduleEmail 发送日程邮�?
func (s *Service) SendScheduleEmail(ctx context.Context, to *email.EmailRecipient, subject, content string, schedule *email.Schedule, opts ...SendEmailOption) (*email.SendEmailResponse, error) {
	req := &email.SendEmailRequest{
		To:       to,
		Subject:  subject,
		Content:  content,
		Schedule: schedule,
	}

	for _, opt := range opts {
		opt(req)
	}

	return s.SendEmail(ctx, req)
}

// SendMeetingEmail 发送会议邮�?
func (s *Service) SendMeetingEmail(ctx context.Context, to *email.EmailRecipient, subject, content string, schedule *email.Schedule, meeting *email.Meeting, opts ...SendEmailOption) (*email.SendEmailResponse, error) {
	req := &email.SendEmailRequest{
		To:       to,
		Subject:  subject,
		Content:  content,
		Schedule: schedule,
		Meeting:  meeting,
	}

	for _, opt := range opts {
		opt(req)
	}

	return s.SendEmail(ctx, req)
}

// SendEmailOption 发送邮件选项
type SendEmailOption func(*email.SendEmailRequest)

// WithCC 设置抄�?
func WithCC(cc *email.EmailRecipient) SendEmailOption {
	return func(req *email.SendEmailRequest) {
		req.CC = cc
	}
}

// WithBCC 设置密�?
func WithBCC(bcc *email.EmailRecipient) SendEmailOption {
	return func(req *email.SendEmailRequest) {
		req.BCC = bcc
	}
}

// WithAttachments 设置附件
func WithAttachments(attachments []*email.Attachment) SendEmailOption {
	return func(req *email.SendEmailRequest) {
		req.AttachmentList = attachments
	}
}

// WithContentType 设置内容类型
func WithContentType(contentType string) SendEmailOption {
	return func(req *email.SendEmailRequest) {
		req.ContentType = contentType
	}
}

// WithEnableIDTrans 设置是否开启id转译
func WithEnableIDTrans(enable bool) SendEmailOption {
	return func(req *email.SendEmailRequest) {
		if enable {
			req.EnableIDTrans = 1
		} else {
			req.EnableIDTrans = 0
		}
	}
}

