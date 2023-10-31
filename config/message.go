package config

type EmailConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

var MailConfig EmailConfig
var AlertEmailList = []string{
	"pythonup@hotmail.com",
}

func init() {
	MailConfig = EmailConfig{
		"smtp.163.com",
		465,
		"username",
		"password",
	}
}
