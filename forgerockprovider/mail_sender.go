package forgerockprovider

import (
	"context"
	"fmt"
	"net/smtp"
	"strings"

	"terraform-provider-forgerock/sdk"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Type to represent email information
type MailSender struct {
	SMTPServer     string
	SMTPPort       int64
	SenderEmail    string
	SenderUsername string
	SenderPassword string
	smtpMailer     SMTPMailer
}

type SMTPMailer interface {
	SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error
	PlainAuth(identity, username, password, host string) smtp.Auth
}

type DefaultSMTPMailer struct{}

func (d *DefaultSMTPMailer) SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	return smtp.SendMail(addr, a, from, to, msg)
}

func (d *DefaultSMTPMailer) PlainAuth(identity, username, password, host string) smtp.Auth {
	return smtp.PlainAuth(identity, username, password, host)
}

func NewDefaultSMTPMailer() SMTPMailer {
	return &DefaultSMTPMailer{}
}

// Method to send an email on new client creation
func (email *MailSender) SendClientInfo(ctx context.Context, oAuth2Client sdk.OAuth2Client, clientId string, clientSecret string, recipient string, wellknownEndpoint string) error {
	body := email.generateBody(clientId, clientSecret, oAuth2Client, wellknownEndpoint, recipient)
	return email.Send(ctx, recipient, body)
}

func (email *MailSender) generateBody(clientId string, clientSecret string, oAuth2Client sdk.OAuth2Client, wellknownEndpoint string, recipient string) string {
	subject := "Your ForgeRock client info"
	message := "Your ForgeRock client has been created/updated. Please find below the client information:\n\n" +
		"ID: " + clientId + "\n" +
		"Secret: " + clientSecret + "\n" +
		"Type: " + oAuth2Client.CoreOAuth2ClientConfig.ClientType.Value + "\n" +
		"Status: " + oAuth2Client.CoreOAuth2ClientConfig.Status.Value + "\n" +
		"Scopes: " + strings.Join(oAuth2Client.CoreOAuth2ClientConfig.Scopes.Value, " ") + "\n" +
		"Redirection URIs: " + strings.Join(oAuth2Client.CoreOAuth2ClientConfig.RedirectionUris.Value, " ") + "\n" +
		"Post Logout Redirect URIs: " + strings.Join(oAuth2Client.CoreOpenIDClientConfig.PostLogoutRedirectUri.Value, " ") + "\n" +
		"Wellknown endpoint: " + wellknownEndpoint + "\n"

	body := "To: " + recipient + "\r\n" +
		"From: " + email.SenderEmail + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/plain; charset=UTF-8\r\n\r\n" +
		message

	return body
}

// Method to send an email
func (email *MailSender) Send(ctx context.Context, recipient string, body string) error {

	tflog.Info(ctx, "Sending email to "+recipient+" from "+email.SenderEmail)

	// Authenticate using credentials
	auth := email.smtpMailer.PlainAuth("", email.SenderUsername, email.SenderPassword, email.SMTPServer)

	// Send the email
	err := email.smtpMailer.SendMail(email.SMTPServer+":"+fmt.Sprint(email.SMTPPort), auth, email.SenderEmail, []string{recipient}, []byte(body))
	tflog.Info(ctx, "Email sent to "+recipient)
	return err
}
