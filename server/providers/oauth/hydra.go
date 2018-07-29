package oauth

import (
	"fmt"
	"net/http"

	"github.com/ory/hydra/sdk/go/hydra/swagger"

	"github.com/ory/hydra/sdk/go/hydra"
)

type Oauthclient interface {
	AcceptLoginRequest(challenge string, username string, remember bool) (string, error)
}

type HydraClient struct {
	HydraURL string
	HydraSDK *hydra.CodeGenSDK
}

func NewHydraClient(url string) (*HydraClient, error) {
	sdk, err := hydra.NewSDK(&hydra.Configuration{
		EndpointURL: url,
	})
	if err != nil {
		return nil, err
	}
	return &HydraClient{
		HydraURL: url,
		HydraSDK: sdk,
	}, nil
}

// AcceptLoginRequest sends the login request to the oauth server and expects the redirect link in return
func (h *HydraClient) AcceptLoginRequest(challenge string, username string, remember bool) (string, error) {
	acceptLoginRequest := swagger.AcceptLoginRequest{
		Subject:     username,
		Remember:    remember,
		RememberFor: 3600,
	}
	req, res, err := h.HydraSDK.AcceptLoginRequest(challenge, acceptLoginRequest)

	if err != nil {
		return "", err
	} else if res.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("hydra returned status code %d", res.StatusCode)
	}
	return req.RedirectTo, nil
}
