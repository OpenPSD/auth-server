package oauth

import (
	"fmt"
	"net/http"

	"github.com/openpsd/auth-server/server/entities"
	"github.com/openpsd/auth-server/server/providers/log"
	"github.com/ory/hydra/sdk/go/hydra/swagger"

	"github.com/ory/hydra/sdk/go/hydra"
)

type Oauthclient interface {
	AcceptLoginRequest(challenge string, username string, remember bool) (string, error)
	GetLoginRequest(challenge string) (entities.ValidateLoginRequest, error)
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
	log.Info.Printf("user=%s challenge=%s remember=%t", username, challenge, remember)
	acceptLoginRequest := swagger.AcceptLoginRequest{
		Subject:     username,
		Remember:    remember,
		RememberFor: 3600,
	}
	req, res, err := h.HydraSDK.AcceptLoginRequest(challenge, acceptLoginRequest)

	if err != nil {
		log.Error.Printf("msg=%s", err)
		return "", err
	} else if res.StatusCode == http.StatusOK {
		log.Info.Printf("redirect_link=%s", req.RedirectTo)
		return req.RedirectTo, nil
	}
	log.Error.Printf("msg='not able to accept login request' http_status_code=%d", res.StatusCode)
	return "", fmt.Errorf("hydra returned status code %d", res.StatusCode)
}

func (h *HydraClient) GetLoginRequest(challenge string) (entities.ValidateLoginRequest, error) {
	validateLoginRequest := entities.ValidateLoginRequest{}
	req, res, err := h.HydraSDK.GetLoginRequest(challenge)
	if err != nil {
		log.Error.Printf("msg=%s", err)
		return validateLoginRequest, err
	}
	if res.StatusCode == http.StatusOK {
		validateLoginRequest.Skip = req.Skip
		validateLoginRequest.Subject = req.Subject
		log.Info.Printf("http_status_code=%d , validateLoginRequest=%s", http.StatusOK, validateLoginRequest.String())
		return validateLoginRequest, nil
	}
	log.Error.Printf("msg='not able to get login request' http_status_code=%d", res.StatusCode)
	return validateLoginRequest, fmt.Errorf("status code %d", res.StatusCode)
}
