package viewmodel

type HttpResponseViewModel struct {
	StatusCode int                `json:"statusCode"`
	Headers    []HeaderViewModel  `json:"headers"`
	Body       *BodyViewModel     `json:"body"`
	Cookies    []CookieViewModel  `json:"cookies"`
}
