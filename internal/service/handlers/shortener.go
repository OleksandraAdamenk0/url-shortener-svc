package handlers

import (
    "net/http"
    "gitlab.com/distributed_lab/ape"
    "gitlab.com/tokend/url-shortener-svc/internal/resources"
)

// handles requests for shortening URLs
func ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
    // Simulated response
    response := resources.ShortenURLResponse{
        Data: resources.ShortenURL{
            Attributes: resources.ShortenURLAttributes{
                ShortURL: "http://short.url/abcd",
            },
        },
    }

    ape.Render(w, response)
}
