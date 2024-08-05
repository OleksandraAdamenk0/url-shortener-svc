package resources

type ResolveURLAttributes struct {
    OriginalURL string `json:"original_url"`
}

type ResolveURL struct {
    Attributes ResolveURLAttributes `json:"attributes"`
}

type ResolveURLResponse struct {
    Data ResolveURL `json:"data"`
}
