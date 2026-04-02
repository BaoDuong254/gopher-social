package mailer

import "embed"

const (
	FromName                 = "Gopher Social"
	maxRetries               = 3
	UserWelcomeEmailTemplate = "user_invitation.tmpl"
)

//go:embed templates/*.tmpl
var FS embed.FS

type Client interface {
	Send(templateFile, username, email string, data any, isSandbox bool) error
}
