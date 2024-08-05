package handlers

import (
    "net/http"
    "gitlab.com/distributed_lab/ape"
    "gitlab.com/tokend/url-shortener-svc/internal/resources"
)

// handles requests to resolve short URLs
func ResolveURLHandler(w http.ResponseWriter, r *http.Request) {
    // Simulated response
    response := resources.ResolveURLResponse{
        Data: resources.ResolveURL{
            Attributes: resources.ResolveURLAttributes{
                OriginalURL: "http://example.com/original-url",
            },
        },
    }

    ape.Render(w, response)
}
