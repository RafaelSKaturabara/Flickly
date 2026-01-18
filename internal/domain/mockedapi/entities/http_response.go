package entities

type HttpResponse struct {
	StatusCode int
	Headers    []Header
	Body       *Body
	Cookies    []Cookie
}
