package requests

import (
    "encoding/json"
    "net/http"
    "gitlab.com/tokend/url-shortener-svc/internal/resources"
)

// creates a new request from the HTTP request
func NewShortenURLRequest(r *http.Request) (resources.ShortenURLRequest, error) {
    var request resources.ShortenURLRequest
    // Simulated decoding
    json.NewDecoder(r.Body).Decode(&request)
    return request, nil
}
