package entities

import "encoding/json"

type LoginRequest struct {
	Username  string
	Password  string
	Challenge string
}

// Marshal interface implementation
func (l *LoginRequest) Marshal() ([]byte, error) {
	if l == nil {
		return nil, nil
	}
	return json.Marshal(l)
}

// Unmarshal interface implementation
func (l *LoginRequest) Unmarshal(b []byte) error {
	var loginRequest LoginRequest
	if err := json.Unmarshal(b, &loginRequest); err != nil {
		return err
	}
	*l = loginRequest
	return nil
}
