package mail

type Auth struct {
}

type Letter struct {
	Subject string
	Body    string
	Handler []string
}

// letter constructor
func NewLetter(subject string, body string, handler []string) *Letter {
	return &Letter{
		Subject: subject,
		Body:    body,
		Handler: handler,
	}
}

type Target struct {
	Target []string
}

//target constructor (target - mail_to_send)
func NewTarget(target []string) *Target {
	return &Target{
		Target: target,
	}
}
