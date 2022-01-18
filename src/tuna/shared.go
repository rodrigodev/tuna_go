package tuna

const appTokenHeader = "x-tuna-apptoken"

type Config struct {
	BaseURL   string
	UserAgent string
	AppToken  string
}

type Customer struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type SessionCard struct {
	PartnerID  int    `json:"partnerId"`
	CustomerID int    `json:"customerId"`
	SessionID  string `json:"sessionId"`
	CardID     int    `json:"cardId"`
}

type CardData struct {
	CardHolderName  string `json:"cardHolderName"`
	ExpirationMonth int64  `json:"expirationMonth"`
	ExpirationYear  int64  `json:"expirationYear"`
	SingleUse       bool   `json:"singleUse"`
}

type TokenData struct {
	Token           string `json:"token"`
	Brand           string `json:"brand"`
	CardHolderName  string `json:"cardHolderName"`
	ExpirationMonth int    `json:"expirationMonth"`
	ExpirationYear  int    `json:"expirationYear"`
	MaskedNumber    string `json:"maskedNumber"`
}

type FrontData struct {
	SessionID       string `json:"SessionID"`
	Origin          string `json:"Origin"`
	IPAddress       string `json:"IpAddress"`
	CookiesAccepted bool   `json:"CookiesAccepted"`
}

type Address struct {
	Street       string `json:"Street"`
	Number       string `json:"Number"`
	Complement   string `json:"Complement"`
	Neighborhood string `json:"Neighborhood"`
	City         string `json:"City"`
	State        string `json:"State"`
	Country      string `json:"Country"`
	PostalCode   string `json:"PostalCode"`
	Phone        string `json:"Phone"`
}

type BillingInfo struct {
	Document     string  `json:"Document"`
	DocumentType string  `json:"DocumentType"`
	Address      Address `json:"Address"`
}

type CardInfo struct {
	CardNumber      interface{} `json:"CardNumber"`
	CardHolderName  string      `json:"CardHolderName"`
	BrandName       string      `json:"BrandName"`
	ExpirationMonth interface{} `json:"ExpirationMonth"`
	ExpirationYear  interface{} `json:"ExpirationYear"`
	Token           string      `json:"Token"`
	TokenSingleUse  int         `json:"TokenSingleUse"`
	SaveCard        bool        `json:"SaveCard"`
	BillingInfo     BillingInfo `json:"BillingInfo"`
}

type DeliveryAddress struct {
	Street       string `json:"Street"`
	Number       string `json:"Number"`
	Complement   string `json:"Complement"`
	Neighborhood string `json:"Neighborhood"`
	City         string `json:"City"`
	State        string `json:"State"`
	Country      string `json:"Country"`
	PostalCode   string `json:"PostalCode"`
	Phone        string `json:"Phone"`
}

type PaymentMethods struct {
	PaymentMethodType string   `json:"PaymentMethodType"`
	Amount            int      `json:"Amount"`
	Installments      int      `json:"Installments"`
	CardInfo          CardInfo `json:"CardInfo"`
}

type PaymentData struct {
	PaymentMethods  []PaymentMethods `json:"PaymentMethods"`
	CountryCode     string           `json:"Countrycode"`
	AntiFraud       AntiFraud        `json:"AntiFraud"`
	DeliveryAddress DeliveryAddress  `json:"DeliveryAddress"`
}

type AntiFraud struct {
	DeliveryAddressee string `json:"DeliveryAddressee"`
}

type PaymentItem struct {
	Amount             int       `json:"Amount"`
	ProductDescription string    `json:"ProductDescription"`
	ItemQuantity       int       `json:"ItemQuantity"`
	CategoryName       string    `json:"CategoryName"`
	Ean                string    `json:"Ean"`
	AntiFraud          AntiFraud `json:"AntiFraud"`
}

type RedirectInfo struct {
	Url string `json:"url"`
}

type Message struct {
	Source  int    `json:"source"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Info    string `json:"info"`
}

type Method struct {
	Message        Message        `json:"message"`
	AdditionalInfo AdditionalInfo `json:"additionalInfo"`
	MethodType     string         `json:"methodType"`
	Status         string         `json:"status"`
	MethodId       int            `json:"methodId"`
}

type AdditionalInfo struct {
	Property1 string `json:"property1"`
	Property2 string `json:"property2"`
}

type CardDetail struct {
	MethodId int `json:"MethodId"`
	Amount   int `json:"Amount"`
	Data     struct {
		CardNumber string `json:"cardNumber"`
	} `json:"Data,omitempty"`
}

type ItemDetail struct {
	ItemQuantity   int    `json:"ItemQuantity"`
	DetailUniqueID string `json:"DetailUniqueID"`
}

type Item struct {
	Message         Message `json:"message"`
	Status          string  `json:"Status"`
	PartnerUniqueId string  `json:"PartnerUniqueId"`
	MethodType      string  `json:"MethodType"`
}

type PaymentOption struct {
	Name            string   `json:"name"`
	DisplayName     string   `json:"displayName"`
	AcceptedBrands  []string `json:"acceptedBrands"`
	PaymentBehavior string   `json:"paymentBehavior"`
}

type GiftCard struct {
	CardNumber   string `json:"CardNumber"`
	Organization string `json:"Organization"`
}

type Arguments struct {
	GiftCard GiftCard `json:"GiftCard"`
}
