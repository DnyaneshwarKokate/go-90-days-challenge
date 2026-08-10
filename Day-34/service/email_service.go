package service

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"
	"sync"
	"time"

	"day-34/domain"

	"github.com/google/uuid"
)

type EmailService struct {
	smtpHost  string
	smtpPort  string
	sender    string
	password  string
	isMock    bool
	queueChan chan domain.EmailMessage
	logger    domain.Logger
	wg        sync.WaitGroup
}

func NewEmailService(host, port, sender, password string, isMock bool, logger domain.Logger) *EmailService {
	svc := &EmailService{
		smtpHost:  host,
		smtpPort:  port,
		sender:    sender,
		password:  password,
		isMock:    isMock,
		queueChan: make(chan domain.EmailMessage, 100),
		logger:    logger,
	}

	// Start background email queue worker
	svc.wg.Add(1)
	go svc.startWorker()

	return svc
}

func (s *EmailService) RenderTemplate(tplName string, name string) (string, error) {
	htmlTpl := `
	<!DOCTYPE html>
	<html>
	<head><style>body { font-family: Arial, sans-serif; }</style></head>
	<body>
		<h2>Hello {{.Name}},</h2>
		<p>Welcome to the <strong>90 Days Go Challenge</strong>!</p>
		<p>Action requested: <em>{{.Action}}</em></p>
		<hr/>
		<small>Automated Email System</small>
	</body>
	</html>`

	action := "Account Setup"
	if tplName == "PASSWORD_RESET" {
		action = "Password Reset Requested"
	} else if tplName == "ORDER_CONFIRMATION" {
		action = "Order Processed Successfully"
	}

	t, err := template.New("email").Parse(htmlTpl)
	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	data := map[string]string{
		"Name":   name,
		"Action": action,
	}
	if err := t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (s *EmailService) SendEmail(ctx context.Context, msg domain.EmailMessage) error {
	s.logger.Info(ctx, "Sending email via SMTP", "to", strings.Join(msg.To, ","), "subject", msg.Subject, "is_mock", s.isMock)

	if s.isMock {
		s.logger.Info(ctx, "MOCK SMTP: Email delivered successfully", "id", msg.ID, "recipients", len(msg.To))
		return nil
	}

	auth := smtp.PlainAuth("", s.sender, s.password, s.smtpHost)
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n%s\r\n%s", strings.Join(msg.To, ","), msg.Subject, mime, msg.BodyHTML))

	addr := fmt.Sprintf("%s:%s", s.smtpHost, s.smtpPort)
	if err := smtp.SendMail(addr, auth, s.sender, msg.To, body); err != nil {
		s.logger.Error(ctx, "Failed to send SMTP email", "error", err)
		return fmt.Errorf("SMTP error: %w", err)
	}

	return nil
}

func (s *EmailService) EnqueueEmail(ctx context.Context, msg domain.EmailMessage) {
	msg.ID = uuid.New().String()
	msg.Status = "QUEUED"
	msg.SentAt = time.Now()

	s.logger.Info(ctx, "Enqueuing email for background delivery", "job_id", msg.ID, "to", strings.Join(msg.To, ","))
	s.queueChan <- msg
}

func (s *EmailService) startWorker() {
	defer s.wg.Done()
	ctx := context.Background()

	for msg := range s.queueChan {
		if err := s.SendEmail(ctx, msg); err != nil {
			s.logger.Error(ctx, "Background worker email delivery failed", "job_id", msg.ID, "error", err)
		} else {
			s.logger.Info(ctx, "Background worker email delivered", "job_id", msg.ID)
		}
	}
}

func (s *EmailService) Close() {
	close(s.queueChan)
	s.wg.Wait()
}
