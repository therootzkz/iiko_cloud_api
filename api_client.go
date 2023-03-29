package iiko

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type ApiClient interface {
	GetMenu() (*GetMenuResponse, error)
}

func NewApiClient(apiKey, baseURL, organizationID string) (ApiClient, error) {
	orgID, err := uuid.Parse(organizationID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse organizationID: %w", err)
	}
	return &apiClient{
		apiKey:         apiKey,
		baseURL:        baseURL,
		organizationID: orgID,
	}, nil
}

type apiClient struct {
	apiKey         string
	baseURL        string
	organizationID uuid.UUID
}

func (a *apiClient) getAuthToken() (string, error) {
	url := a.baseURL + "/api/1/access_token"
	requestBytes, err := json.Marshal(struct {
		ApiLogin string `json:"apiLogin"`
	}{
		ApiLogin: a.apiKey,
	})
	if err != nil {
		return "", err
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(requestBytes))
	if err != nil {
		return "", err
	}
	var response struct {
		Token string `json:"token"`
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}
	return response.Token, nil
}

func (a *apiClient) GetMenu() (*GetMenuResponse, error) {
	authToken, err := a.getAuthToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get auth token: %w", err)
	}
	httpClient := http.DefaultClient
	requestBody := GetMenuRequest{
		OrganizationId: a.organizationID,
	}
	requestBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}
	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/api/1/nomenclature", a.baseURL),
		bytes.NewReader(requestBytes),
	)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authToken))
	request.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	response, err := httpClient.Do(request)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	bodyDecoder := json.NewDecoder(response.Body)
	if response.StatusCode != http.StatusOK {
		var errorDetails GetMenuErrorDetails
		err := bodyDecoder.Decode(&errorDetails)
		if err != nil {
			return nil, fmt.Errorf("invalid response status code %d", response.StatusCode)
		}
		return nil, fmt.Errorf("invalid response status code %d: %s", response.StatusCode, errorDetails.ErrorDescription)
	}

	var getMenuResponse GetMenuResponse
	err = bodyDecoder.Decode(&getMenuResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}
	return &getMenuResponse, nil
}
