package tuna

import "net/http"

func setHeaders(req *http.Request, userAgent, appToken string) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set(appTokenHeader, appToken)
}
