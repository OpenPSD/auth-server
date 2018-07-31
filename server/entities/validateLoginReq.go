package entities

import "fmt"

type ValidateLoginRequest struct {
	Skip    bool
	Subject string
}

func (v ValidateLoginRequest) String() string {
	return fmt.Sprintf("{skip:%t , Subject: %s}", v.Skip, v.Subject)
}
