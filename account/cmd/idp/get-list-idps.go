package idp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
)

type IDPHandler struct {
	baseURL  string
	username string
	password string
	realm    string
}

func NewIDPHandler(baseURL, username, password, realm string) *IDPHandler {
	return &IDPHandler{
		baseURL:  baseURL,
		username: username,
		password: password,
		realm:    realm,
	}
}

func (h *IDPHandler) getToken() (string, error) {
	data := url.Values{}
	data.Set("client_id", "admin-cli")
	data.Set("username", h.username)
	data.Set("password", h.password)
	data.Set("grant_type", "password")

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/realms/master/protocol/openid-connect/token", h.baseURL), strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create token request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to get token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get token: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read token response body: %w", err)
	}

	var tokenResponse struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return "", fmt.Errorf("failed to decode token response: %w", err)
	}

	return tokenResponse.AccessToken, nil
}

func (h *IDPHandler) ListIDPs() ([]interface{}, error) {
	token, err := h.getToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/admin/realms/%s/identity-provider/instances", h.baseURL, h.realm)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get IDPs: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get IDPs: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var idps []interface{}
	if err := json.Unmarshal(body, &idps); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return idps, nil
}

var ListIDPsCmd = &cobra.Command{
	Use:   "list-idps",
	Short: "List identity providers",
	Run: func(cmd *cobra.Command, args []string) {
		baseURL, _ := cmd.Flags().GetString("base-url")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		realm, _ := cmd.Flags().GetString("realm")

		handler := NewIDPHandler(baseURL, username, password, realm)
		idps, err := handler.ListIDPs()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		idpsJSON, err := json.MarshalIndent(idps, "", "  ")
		if err != nil {
			fmt.Println("Failed to encode response:", err)
			return
		}

		fmt.Println(string(idpsJSON))
	},
}

func init() {
	ListIDPsCmd.Flags().String("base-url", "", "Base URL of the IDP service")
	ListIDPsCmd.Flags().String("username", "", "Admin username")
	ListIDPsCmd.Flags().String("password", "", "Admin password")
	ListIDPsCmd.Flags().String("realm", "", "Realm name")
	ListIDPsCmd.MarkFlagRequired("base-url")
	ListIDPsCmd.MarkFlagRequired("username")
	ListIDPsCmd.MarkFlagRequired("password")
	ListIDPsCmd.MarkFlagRequired("realm")
}
