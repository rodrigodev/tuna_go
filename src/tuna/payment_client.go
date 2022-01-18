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
	initPaymentPath = "/api/Payment/Init"
	cancelPath      = "/api/Payment/Cancel"
	cancelItemPath  = "/api/Payment/CancelItem"
	capturePath     = "/api/Payment/Capture"
	continuePath    = "/api/Payment/Continue"
	statusPath      = "/api/Payment/Status"
	optionsPath     = "/api/Payment/Options"
	functionPath    = "/api/Payment/Function"
)

type PaymentAPI interface {
	Init(request InitRequest) (*InitResponse, error)
	Cancel(request CancelRequest) (*CancelResponse, error)
	CancelItem(request CancelItemRequest) (*CancelItemResponse, error)
	Capture(request CaptureRequest) (*CaptureResponse, error)
	Continue(request ContinueRequest) (*ContinueResponse, error)
	Status(request StatusRequest) (*StatusResponse, error)
	Options(request OptionsRequest) (*OptionsResponse, error)
	Function(request FunctionRequest) (*FunctionResponse, error)
}

type PaymentClient struct {
	baseURL    *url.URL
	httpClient *http.Client
	userAgent  string
	appToken   string
}

func NewPaymentClient(client *http.Client, conf Config) *PaymentClient {
	parsedURL, _ := url.Parse(conf.BaseURL)

	return &PaymentClient{
		httpClient: client,
		baseURL:    parsedURL,
		userAgent:  conf.UserAgent,
		appToken:   conf.AppToken,
	}
}

func (p PaymentClient) Init(request InitRequest) (*InitResponse, error) {
	var buf = new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(request)
	if err != nil {
		return nil, err
	}
	rel := &url.URL{Path: initPaymentPath}
	u := p.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(http.MethodPost, u.String(), buf)
	if err != nil {
		return nil, err
	}

	setHeaders(req, p.userAgent, p.appToken)

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("an error occurred, status code %v", resp.StatusCode)
	}

	var ir InitResponse
	err = json.NewDecoder(resp.Body).Decode(&ir)

	return &ir, err
}

func (p PaymentClient) Cancel(request CancelRequest) (*CancelResponse, error) {
	var buf = new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(request)
	if err != nil {
		return nil, err
	}
	rel := &url.URL{Path: cancelPath}
	u := p.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(http.MethodPost, u.String(), buf)
	if err != nil {
		return nil, err
	}

	setHeaders(req, p.userAgent, p.appToken)

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("an error occurred, status code %v", resp.StatusCode)
	}

	var cr CancelResponse
	err = json.NewDecoder(resp.Body).Decode(&cr)

	return &cr, err
}

func (p PaymentClient) CancelItem(request CancelItemRequest) (*CancelItemResponse, error) {
	var buf = new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(request)
	if err != nil {
		return nil, err
	}
	rel := &url.URL{Path: cancelItemPath}
	u := p.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(http.MethodPost, u.String(), buf)
	if err != nil {
		return nil, err
	}

	setHeaders(req, p.userAgent, p.appToken)

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("an error occurred, status code %v", resp.StatusCode)
	}

	var cir CancelItemResponse
	err = json.NewDecoder(resp.Body).Decode(&cir)

	return &cir, err
}

func (p PaymentClient) Capture(request CaptureRequest) (*CaptureResponse, error) {
	var buf = new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(request)
	if err != nil {
		return nil, err
	}
	rel := &url.URL{Path: capturePath}
	u := p.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(http.MethodPost, u.String(), buf)
	if err != nil {
		return nil, err
	}

	setHeaders(req, p.userAgent, p.appToken)

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("an error occurred, status code %v", resp.StatusCode)
	}

	var cr CaptureResponse
	err = json.NewDecoder(resp.Body).Decode(&cr)

	return &cr, err
}

func (p PaymentClient) Continue(request ContinueRequest) (*ContinueResponse, error) {
	var buf = new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(request)
	if err != nil {
		return nil, err
	}
	rel := &url.URL{Path: continuePath}
	u := p.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(http.MethodPost, u.String(), buf)
	if err != nil {
		return nil, err
	}

	setHeaders(req, p.userAgent, p.appToken)

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("an error occurred, status code %v", resp.StatusCode)
	}

	var cr ContinueResponse
	err = json.NewDecoder(resp.Body).Decode(&cr)

	return &cr, err
}

func (p PaymentClient) Status(request StatusRequest) (*StatusResponse, error) {
	var buf = new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(request)
	if err != nil {
		return nil, err
	}
	rel := &url.URL{Path: statusPath}
	u := p.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(http.MethodPost, u.String(), buf)
	if err != nil {
		return nil, err
	}

	setHeaders(req, p.userAgent, p.appToken)

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("an error occurred, status code %v", resp.StatusCode)
	}

	var sr StatusResponse
	err = json.NewDecoder(resp.Body).Decode(&sr)

	return &sr, err
}

func (p PaymentClient) Options(request OptionsRequest) (*OptionsResponse, error) {
	var buf = new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(request)
	if err != nil {
		return nil, err
	}
	rel := &url.URL{Path: optionsPath}
	u := p.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(http.MethodPost, u.String(), buf)
	if err != nil {
		return nil, err
	}

	setHeaders(req, p.userAgent, p.appToken)

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("an error occurred, status code %v", resp.StatusCode)
	}

	var or OptionsResponse
	err = json.NewDecoder(resp.Body).Decode(&or)

	return &or, err
}

func (p PaymentClient) Function(request FunctionRequest) (*FunctionResponse, error) {
	var buf = new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(request)
	if err != nil {
		return nil, err
	}
	rel := &url.URL{Path: functionPath}
	u := p.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(http.MethodPost, u.String(), buf)
	if err != nil {
		return nil, err
	}

	setHeaders(req, p.userAgent, p.appToken)

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("an error occurred, status code %v", resp.StatusCode)
	}

	var fr FunctionResponse
	err = json.NewDecoder(resp.Body).Decode(&fr)

	return &fr, err
}

type InitRequest struct {
	PartnerUniqueID string        `json:"PartnerUniqueID"`
	Customer        Customer      `json:"Customer"`
	PaymentItems    []PaymentItem `json:"PaymentItems"`
	PaymentData     PaymentData   `json:"PaymentData"`
	FrontData       FrontData     `json:"FrontData"`
}

type InitResponse struct {
	Status          string       `json:"status"`
	Methods         []Method     `json:"methods"`
	PaymentKey      string       `json:"paymentKey"`
	PartnerUniqueId string       `json:"partnerUniqueId"`
	Message         Message      `json:"message"`
	RedirectInfo    RedirectInfo `json:"redirectInfo"`
}

type CancelRequest struct {
	PartnerUniqueID string       `json:"PartnerUniqueID"`
	PaymentDate     string       `json:"PaymentDate"`
	CancelAll       bool         `json:"CancelAll"`
	CardsDetail     []CardDetail `json:"CardsDetail"`
}

type CancelResponse struct {
	Status  string   `json:"status"`
	Methods []Method `json:"methods"`
	Message Message  `json:"message"`
}

type CancelItemRequest struct {
	PartnerUniqueID string       `json:"PartnerUniqueID"`
	PaymentDate     string       `json:"PaymentDate"`
	ItemsDetail     []ItemDetail `json:"ItemsDetail"`
}

type CancelItemResponse struct {
	Status  string  `json:"status"`
	Items   []Item  `json:"Items"`
	Message Message `json:"message"`
}

type CaptureRequest struct {
	Amount          int          `json:"amount"`
	CardsDetail     []CardDetail `json:"cardsDetail"`
	PaymentKey      string       `json:"paymentKey"`
	PartnerUniqueID string       `json:"partnerUniqueID"`
	PaymentDate     time.Time    `json:"paymentDate"`
	AppToken        string       `json:"appToken"`
	Account         string       `json:"account"`
	ExtraInfo       string       `json:"extraInfo"`
}

type CaptureResponse struct {
	Status  string   `json:"status"`
	Methods []Method `json:"methods"`
	Message Message  `json:"message"`
}

type ContinueRequest struct {
	Amount          int            `json:"amount"`
	AdditionalInfo  AdditionalInfo `json:"additionalInfo"`
	PaymentKey      string         `json:"paymentKey"`
	PartnerUniqueID string         `json:"partnerUniqueID"`
	PaymentDate     time.Time      `json:"paymentDate"`
	AppToken        string         `json:"appToken"`
	Account         string         `json:"account"`
	ExtraInfo       string         `json:"extraInfo"`
}

type ContinueResponse struct {
	Message Message `json:"message"`
}

type StatusRequest struct {
	PartnerUniqueID string    `json:"partnerUniqueID"`
	PaymentDate     time.Time `json:"paymentDate"`
	PaymentKey      string    `json:"paymentKey"`
	PartnerID       int       `json:"partnerID"`
	AppToken        string    `json:"appToken"`
	Account         string    `json:"account"`
	ExtraInfo       string    `json:"extraInfo"`
}

type StatusResponse struct{}

type OptionsRequest struct {
	PartnerID int    `json:"partnerID"`
	AppToken  string `json:"appToken"`
	Account   string `json:"account"`
	ExtraInfo string `json:"extraInfo"`
}

type OptionsResponse struct {
	PaymentOptions []PaymentOption `json:"paymentOptions"`
	Message        Message         `json:"message"`
}

type FunctionRequest struct {
	PartnerUniqueID string    `json:"PartnerUniqueID"`
	FunctionName    string    `json:"FunctionName"`
	Arguments       Arguments `json:"Arguments"`
}

type FunctionResponse struct {
	Response interface{} `json:"response"`
	Message  Message     `json:"message"`
}
