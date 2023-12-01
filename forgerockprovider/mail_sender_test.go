//go:build unit_tests
// +build unit_tests

package forgerockprovider

import (
	"context"
	"fmt"
	"net/smtp"
	"strings"
	"terraform-provider-forgerock/sdk"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockSMTPMailer struct {
	mock.Mock
}

// PlainAuth implements SMTPMailer.
func (m *MockSMTPMailer) PlainAuth(identity string, username string, password string, host string) smtp.Auth {
	args := m.Called(identity, username, password, host)
	return args.Get(0).(smtp.Auth)
}

// SendMail implements SMTPMailer.
func (m *MockSMTPMailer) SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	args := m.Called(addr, a, from, to, msg)
	return args.Error(0)
}

type MockSmtpAuth struct {
	mock.Mock
}

func (m *MockSmtpAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	args := m.Called(server)
	return args.String(0), args.Get(1).([]byte), args.Error(2)
}

func (m *MockSmtpAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	args := m.Called(fromServer, more)
	return args.Get(0).([]byte), args.Error(1)
}

// Method to send an email on new client creation
func TestSendClientInfo(t *testing.T) {
	ctx := context.Background()
	clientId := "testClientId"
	clientSecret := "testClientSecret"
	clientType := "randomClientType"
	status := "randomStatus"
	scopes := []string{"randomScope1", "randomScope2"}
	redirectionUris := []string{"randomRedirectUri1", "randomRedirectUri2"}
	postLogoutRedirectUris := []string{"randomPostLogoutRedirectUri1", "randomPostLogoutRedirectUri2"}
	wellknownEndpoint := "testWellknownEndpoint"
	recipient := "testRecipient@example.com"

	oAuth2Client := sdk.OAuth2Client{
		CoreOAuth2ClientConfig: sdk.CoreOAuth2ClientConfig{
			ClientType:      NewWrappedValue(&clientType),
			Status:          NewWrappedValue(&status),
			Scopes:          NewWrappedValue(&scopes),
			RedirectionUris: NewWrappedValue(&redirectionUris),
		},
		CoreOpenIDClientConfig: sdk.CoreOpenIDClientConfig{
			PostLogoutRedirectUri: NewWrappedValue(&postLogoutRedirectUris),
		},
	}

	smtpAuthMock := new(MockSmtpAuth)
	smtpMailerMock := new(MockSMTPMailer)

	email := &MailSender{
		SenderEmail:    "test@example.com",
		SenderUsername: "username",
		SenderPassword: "password",
		SMTPServer:     "smtp.example.com",
		SMTPPort:       587,
	}
	email.smtpMailer = smtpMailerMock
	body := email.generateBody(clientId, clientSecret, oAuth2Client, wellknownEndpoint, recipient)

	smtpMailerMock.On("PlainAuth", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(smtpAuthMock, []byte{}, nil)
	smtpMailerMock.On("SendMail", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, []byte{}, nil)

	err := email.SendClientInfo(ctx, oAuth2Client, clientId, clientSecret, recipient, wellknownEndpoint)

	smtpMailerMock.AssertCalled(t, "PlainAuth", "", email.SenderUsername, email.SenderPassword, email.SMTPServer)
	smtpMailerMock.AssertCalled(t, "SendMail", email.SMTPServer+":"+fmt.Sprint(email.SMTPPort), smtpAuthMock, email.SenderEmail, []string{recipient}, []byte(body))

	if err != nil {
		t.Errorf("Failed to send email : %v", err)
	}
}

func TestGenerateBody(t *testing.T) {
	clientId := "testClientId"
	clientSecret := "testClientSecret"
	clientType := "randomClientType"
	status := "randomStatus"
	scopes := []string{"randomScope1", "randomScope2"}
	redirectionUris := []string{"randomRedirectUri1", "randomRedirectUri2"}
	postLogoutRedirectUris := []string{"randomPostLogoutRedirectUri1", "randomPostLogoutRedirectUri2"}
	wellknownEndpoint := "testWellknownEndpoint"
	recipient := "testRecipient@example.com"

	oAuth2Client := sdk.OAuth2Client{
		CoreOAuth2ClientConfig: sdk.CoreOAuth2ClientConfig{
			ClientType:      NewWrappedValue(&clientType),
			Status:          NewWrappedValue(&status),
			Scopes:          NewWrappedValue(&scopes),
			RedirectionUris: NewWrappedValue(&redirectionUris),
		},
		CoreOpenIDClientConfig: sdk.CoreOpenIDClientConfig{
			PostLogoutRedirectUri: NewWrappedValue(&postLogoutRedirectUris),
		},
	}

	emailSender := &MailSender{SenderEmail: "sender@example.com"}

	body := emailSender.generateBody(clientId, clientSecret, oAuth2Client, wellknownEndpoint, recipient)

	expectedBody := "To: " + recipient + "\r\n" +
		"From: " + emailSender.SenderEmail + "\r\n" +
		"Subject: Your ForgeRock client info\r\n" +
		"Content-Type: text/plain; charset=UTF-8\r\n\r\n" +
		"Your ForgeRock client has been created/updated. Please find below the client information:\n\n" +
		"ID: " + clientId + "\n" +
		"Secret: " + clientSecret + "\n" +
		"Type: " + clientType + "\n" +
		"Status: " + status + "\n" +
		"Scopes: " + strings.Join(scopes, " ") + "\n" +
		"Redirection URIs: " + strings.Join(redirectionUris, " ") + "\n" +
		"Post Logout Redirect URIs: " + strings.Join(postLogoutRedirectUris, " ") + "\n" +
		"Wellknown endpoint: " + wellknownEndpoint + "\n"

	if body != expectedBody {
		t.Errorf("Expected body to be '%s', but got '%s'", expectedBody, body)
	}
}

func TestSendEmail(t *testing.T) {

	smtpAuthMock := new(MockSmtpAuth)
	smtpMailerMock := new(MockSMTPMailer)

	recipient := "recipient@example.com"
	body := "Ceci est le corps de l'e-mail de test."
	email := &MailSender{
		SenderEmail:    "test@example.com",
		SenderUsername: "username",
		SenderPassword: "password",
		SMTPServer:     "smtp.example.com",
		SMTPPort:       587,
	}
	email.smtpMailer = smtpMailerMock

	smtpMailerMock.On("PlainAuth", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(smtpAuthMock, []byte{}, nil)
	smtpMailerMock.On("SendMail", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, []byte{}, nil)

	err := email.Send(context.Background(), recipient, body)

	smtpMailerMock.AssertCalled(t, "PlainAuth", "", email.SenderUsername, email.SenderPassword, email.SMTPServer)
	smtpMailerMock.AssertCalled(t, "SendMail", email.SMTPServer+":"+fmt.Sprint(email.SMTPPort), smtpAuthMock, email.SenderEmail, []string{recipient}, []byte(body))

	if err != nil {
		t.Errorf("Erreur lors de l'envoi de l'e-mail : %v", err)
	}

}
