package builder

import "strings"

type email struct {
	from, to, subject, body string
}

// EmailBuild is builder
type EmailBuild struct {
	email email
}

// From email
func (b *EmailBuild) From(from string) *EmailBuild {
	if !strings.Contains(from, "@") {
		panic("email should contain @")
	}
	b.email.from = from
	return b
}

// To email
func (b *EmailBuild) To(to string) *EmailBuild {
	b.email.to = to
	return b
}

// Subject of email
func (b *EmailBuild) Subject(subject string) *EmailBuild {
	b.email.subject = subject
	return b
}

// Body of email
func (b *EmailBuild) Body(body string) *EmailBuild {
	b.email.body = body
	return b
}

func sendMailImpl(email *email) {
	// Some implementation of email sending
}

// build is action type
type build func(*EmailBuild)

// SendEmail is ....
func SendEmail(action build)  {
	builder := EmailBuild{}
	action(&builder)
	sendMailImpl(&builder.email)
}