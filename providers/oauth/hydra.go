package oauth

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/openpsd/auth-server/entities"
)

type Oauthclient interface {
	AcceptLoginRequest(challenge string, r entities.AcceptLoginRequest) (string, error)
}

type HydraClient struct {
	HydraURL string
}

func NewHydraClient(url string) *HydraClient {
	return &HydraClient{
		HydraURL: url,
	}
}

// AcceptLoginRequest sends the login request to the oauth server and expects the redirect link in return
func (h *HydraClient) AcceptLoginRequest(challenge string, r entities.AcceptLoginRequest) (string, error) {
	data, err := r.Marshal()
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, h.buildURL("login", "accept", challenge), bytes.NewBuffer(data))
	res, err := client.Do(req)
	return res.Header.Get("redirect_to"), err
}

func (h *HydraClient) buildURL(flow string, action string, challenge string) string {
	return fmt.Sprintf("%s/oauth2/auth/requests/%s/%s/%s", h.HydraURL, flow, challenge, action)
}
