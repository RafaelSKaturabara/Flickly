package entities

type HttpRequest struct {
	Method      HttpMethod
	Path        string
	Headers     []Header
	QueryParams []QueryParam
	Body        *Body
	Route       *Route
}
