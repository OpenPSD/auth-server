package entities

import "encoding/json"

type AcceptLoginRequest struct {
	subject      string
	remember     bool
	remember_for int
	acr          string
}

func NewAcceptLoginRequest(username string, remember bool) AcceptLoginRequest {
	return AcceptLoginRequest{
		subject:      username,
		remember:     remember,
		remember_for: 3600,
	}
}

// Marshal interface implementation
func (l *AcceptLoginRequest) Marshal() ([]byte, error) {
	if l == nil {
		return nil, nil
	}
	return json.Marshal(l)
}

// Unmarshal interface implementation
func (l *AcceptLoginRequest) Unmarshal(b []byte) error {
	var acceptLoginRequest AcceptLoginRequest
	if err := json.Unmarshal(b, &acceptLoginRequest); err != nil {
		return err
	}
	*l = acceptLoginRequest
	return nil
}
