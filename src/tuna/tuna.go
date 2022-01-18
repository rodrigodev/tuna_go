package tuna

type PaymentAdapter struct {
	tokenClient   *TokenClient
	paymentClient *PaymentClient
}

func NewTunaService(tokenClient *TokenClient, paymentClient *PaymentClient) *PaymentAdapter {
	return &PaymentAdapter{tokenClient: tokenClient, paymentClient: paymentClient}
}

func (s *PaymentAdapter) NewSession(userID string, email string) (string, error) {
	request := NewSessionRequest{
		Customer: Customer{
			ID:    userID,
			Email: email,
		},
	}

	session, err := s.tokenClient.NewSession(request)
	if err != nil {
		return "", err
	}

	return session.SessionID, nil
}