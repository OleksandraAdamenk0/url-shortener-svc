package resources

type ShortenURLAttributes struct {
    ShortURL string `json:"short_url"`
}

type ShortenURL struct {
    Attributes ShortenURLAttributes `json:"attributes"`
}

type ShortenURLResponse struct {
    Data ShortenURL `json:"data"`
}

type ShortenURLRequest struct {
    URL string `json:"url"`
}
