package handler

type CreateShortUrlRequest struct {
	Url string `json:"url"`
}

func (request *CreateShortUrlRequest) Validate() error {
	if request.Url == "" {
		panic("Url is required")
	}
	return nil
}
