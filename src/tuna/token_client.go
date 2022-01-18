package tuna

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	newSessionPath        = "/api/Token/NewSession"
	validateSessionPath   = "/api/Token/ValidateSession"
	generateCardTokenPath = "/api/Token/Generate"
	listTokenPath         = "/api/Token/List"
	deleteCardTokenPath   = "/api/Token/Delete"
	bindCVVPath           = "/api/Token/Bind"
)

type TokenAPI interface {
	NewSession(request NewSessionRequest) (*NewSessionResponse, error)
	ValidateSession(request ValidadeSessionRequest) (*ValidadeSessionResponse, error)
	GenerateCardToken(request GenerateCardTokenRequest) (*GenerateCardTokenResponse, error)
	ListTokens(request ListTokensRequest) (*ListTokensResponse, error)
	DeleteCardToken(request DeleteCardTokenRequest) (*DeleteCardTokenResponse, error)
	BindCVV(request BindCVVRequest) (*BindCVVResponse, error)
}

type TokenClient struct {
	baseURL    *url.URL
	httpClient *http.Client
	userAgent  string
	appToken   string
}

func NewTokenClient(client *http.Client, conf Config) *TokenClient {
	parsedURL, _ := url.Parse(conf.BaseURL)

	return &TokenClient{
		httpClient: client,
		baseURL:    parsedURL,
		userAgent:  conf.UserAgent,
		appToken:   conf.AppToken,
	}
}

func (c *TokenClient) NewSession(request NewSessionRequest) (*NewSessionResponse, error) {
	var buf = new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(request)
	if err != nil {
		return nil, err
	}
	rel := &url.URL{Path: newSessionPath}
	u := c.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(http.MethodPost, u.String(), buf)
	if err != nil {
		return nil, err
	}

	setHeaders(req, c.userAgent, c.appToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("an error occurred, status code %v", resp.StatusCode)
	}

	var nsr NewSessionResponse
	err = json.NewDecoder(resp.Body).Decode(&nsr)

	return &nsr, err
}

func (c *TokenClient) ValidateSession(request ValidadeSessionRequest) (*ValidadeSessionResponse, error) {
	var buf = new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(request)
	if err != nil {
		return nil, err
	}
	rel := &url.URL{Path: validateSessionPath}
	u := c.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(http.MethodPost, u.String(), buf)
	if err != nil {
		return nil, err
	}

	setHeaders(req, c.userAgent, c.appToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("an error occurred, status code %v", resp.StatusCode)
	}

	var vsr ValidadeSessionResponse
	err = json.NewDecoder(resp.Body).Decode(&vsr)

	return &vsr, err
}

func (c *TokenClient) GenerateCardToken(request GenerateCardTokenRequest) (*GenerateCardTokenResponse, error) {
	var buf = new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(request)
	if err != nil {
		return nil, err
	}
	rel := &url.URL{Path: generateCardTokenPath}
	u := c.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(http.MethodPost, u.String(), buf)
	if err != nil {
		return nil, err
	}

	setHeaders(req, c.userAgent, c.appToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("an error occurred, status code %v", resp.StatusCode)
	}

	var gctr GenerateCardTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&gctr)

	return &gctr, err
}

func (c *TokenClient) ListTokens(request ListTokensRequest) (*ListTokensResponse, error) {
	var buf = new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(request)
	if err != nil {
		return nil, err
	}
	rel := &url.URL{Path: listTokenPath}
	u := c.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(http.MethodPost, u.String(), buf)
	if err != nil {
		return nil, err
	}

	setHeaders(req, c.userAgent, c.appToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("an error occurred, status code %v", resp.StatusCode)
	}

	var ltr ListTokensResponse
	err = json.NewDecoder(resp.Body).Decode(&ltr)

	return &ltr, err
}

func (c *TokenClient) DeleteCardToken(request DeleteCardTokenRequest) (*DeleteCardTokenResponse, error) {
	var buf = new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(request)
	if err != nil {
		return nil, err
	}
	rel := &url.URL{Path: deleteCardTokenPath}
	u := c.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(http.MethodDelete, u.String(), buf)
	if err != nil {
		return nil, err
	}

	setHeaders(req, c.userAgent, c.appToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("an error occurred, status code %v", resp.StatusCode)
	}

	var dctr DeleteCardTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&dctr)

	return &dctr, err
}

func (c *TokenClient) BindCVV(request BindCVVRequest) (*BindCVVResponse, error) {
	var buf = new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(request)
	if err != nil {
		return nil, err
	}
	rel := &url.URL{Path: bindCVVPath}
	u := c.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(http.MethodPost, u.String(), buf)
	if err != nil {
		return nil, err
	}

	setHeaders(req, c.userAgent, c.appToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("an error occurred, status code %v", resp.StatusCode)
	}

	var bcr BindCVVResponse
	err = json.NewDecoder(resp.Body).Decode(&bcr)

	return &bcr, err
}

type NewSessionRequest struct {
	Customer Customer `json:"customer"`
}

type NewSessionResponse struct {
	SessionID string `json:"sessionId"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}

type ValidadeSessionRequest struct {
	SessionID string `json:"sessionId"`
}

type ValidadeSessionResponse struct {
	PartnerID    int           `json:"partnerId"`
	CustomerID   int           `json:"customerId"`
	SessionID    string        `json:"sessionId"`
	CreationDate time.Time     `json:"creationDate"`
	Customer     Customer      `json:"customer"`
	SessionCards []SessionCard `json:"sessionCard"`
}

type GenerateCardTokenRequest struct {
	SessionID string   `json:"SessionId"`
	Card      CardData `json:"Card"`
}

type GenerateCardTokenResponse struct {
	Token     string `json:"token"`
	CardBrand string `json:"cardBrand"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}

type ListTokensRequest struct {
	SessionID string `json:"sessionId"`
}

type ListTokensResponse struct {
	Tokens  []TokenData `json:"tokens"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
}

type DeleteCardTokenRequest struct {
	Token     string `json:"token"`
	SessionID string `json:"sessionId"`
}

type DeleteCardTokenResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BindCVVRequest struct {
	Token     string `json:"token"`
	SessionID string `json:"sessionId"`
	CVV       string `json:"CVV"`
}

type BindCVVResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
