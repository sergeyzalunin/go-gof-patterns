package builder

import "testing"

var originEmail = email{
	from:	"test@gmail.com",
	to:		"friend@gmail.com",
	subject: "It is a test email",
	body:	"This is a body",
}

func TestSendEmail(t *testing.T) {
	var got email

	SendEmail(func(b *EmailBuild){
		got = b.From("test@gmail.com").
			To("friend@gmail.com").
			Subject("It is a test email").
			Body("This is a body").email
	})

	if got != originEmail {
		t.Errorf("by creating email via emailbuilder wants to see \n'%s' got \n'%s'", originEmail, got)
	}
}